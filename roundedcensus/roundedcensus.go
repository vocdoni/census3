package roundedcensus

/*
roundedcensus package provides an algorithm to anonymize participant balances in
a voting system while maintaining a certain level of accuracy. It sorts participants
by balance, groups them based on a privacy threshold and balance differences,
rounds their balances, and calculates lost balance for accuracy measurement.

The main steps of the algorithm are:

1. Sort Participants by Balance:
   - Participants are sorted in ascending order based on their balances.

2. Group Participants:
   - Participants are initially grouped with a size equal to the privacy threshold.
   - The group can extend if consecutive participants have the same balance or if
     the difference in balances between consecutive participants is less than or
     equal to the groupBalanceDiff threshold.

3. Round Group Balances:
   - Each group's balances are rounded down to the lowest common value within
     that group.

4. (optional) Accuracy loop:
   - The algorithm tries to find the highest accuracy possible while maintaining
     a minimum privacy threshold. It starts with the minimum privacy threshold
	 and increases it by a small amount until the accuracy is maximized.

5. Output Rounded Balances and Accuracy:
   - The algorithm provides the new list of participants with their rounded
     balances and the calculated accuracy to quantify the balance preservation.
*/

import (
	"fmt"
	"math"
	"math/big"
	"sort"
)

// GroupsConfig represents the configuration for the grouping and rounding process.
type GroupsConfig struct {
	// GroupBalanceDiff is the maximum difference between consecutive balances
	GroupBalanceDiff *big.Int
	// MinPrivacyThreshold is the minimum number of participants in a group.
	MinPrivacyThreshold int64
	// MinAccuracy is the minimum accuracy required for the rounding process.
	MinAccuracy float64
	// OutliersThreshold is the z-score threshold for identifying outliers.
	OutliersThreshold float64
}

var DefaultGroupsConfig GroupsConfig = GroupsConfig{
	GroupBalanceDiff:    big.NewInt(10),
	MinPrivacyThreshold: 3,
	MinAccuracy:         50.0,
	OutliersThreshold:   2.0,
}

// Participant represents a participant with an Ethereum address and balance.
type Participant struct {
	Address string
	Balance *big.Int
}

// ByBalance implements sort.Interface for []Participant based on the Balance field.
type ByBalance []*Participant

func (a ByBalance) Len() int           { return len(a) }
func (a ByBalance) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByBalance) Less(i, j int) bool { return a[i].Balance.Cmp(a[j].Balance) < 0 }

// GroupAndRoundCensus groups the participants and rounds their balances. It
// rounds the balances of the participants with the highest accuracy possible
// while maintaining a minimum privacy threshold. It discards outliers from the
// rounding process but returns them in the final list of participants.
func GroupAndRoundCensus(participants []*Participant, config GroupsConfig) ([]*Participant, float64, error) {
	cleanedParticipants, outliers := zScore(participants, config.OutliersThreshold)

	maxPrivacyThreshold := int64(len(participants)) / config.MinPrivacyThreshold
	currentPrivacyThreshold := config.MinPrivacyThreshold
	maxAccuracy := 0.0
	maxAccuracyPrivacyThreshold := currentPrivacyThreshold
	for currentPrivacyThreshold <= maxPrivacyThreshold {
		groupedParticipants := groupAndRoundCensus(cleanedParticipants, currentPrivacyThreshold, config.GroupBalanceDiff)
		lastAccuracy := calculateAccuracy(cleanedParticipants, groupedParticipants)
		if lastAccuracy > maxAccuracy {
			maxAccuracy = lastAccuracy
			maxAccuracyPrivacyThreshold = currentPrivacyThreshold
		}
		currentPrivacyThreshold += roundGap(currentPrivacyThreshold)
	}
	roundedCensus := groupAndRoundCensus(cleanedParticipants, maxAccuracyPrivacyThreshold, config.GroupBalanceDiff)
	outliersCensus := groupAndRoundCensus(outliers, maxAccuracyPrivacyThreshold, config.GroupBalanceDiff)
	roundedCensus = append(roundedCensus, outliersCensus...)
	accuracy := calculateAccuracy(participants, roundedCensus)
	if accuracy < config.MinAccuracy {
		return roundedCensus, accuracy, fmt.Errorf("could not find a privacy threshold that satisfies the minimum accuracy")
	}
	return roundedCensus, accuracy, nil
}

// roundGap calculate the gap between privacy thresholds for the accuracy loop.
// It returns the gap as a percentage of the current privacy threshold. The gap
// is calculated as 5% of the current privacy threshold, rounded to the nearest
// integer. If the gap is less than 1, it returns 1. The privacy threshold is
// increased by the gap in each iteration of the accuracy loop.
func roundGap(x int64) int64 {
	gap := float64(x) * 0.05
	if gap < 1 {
		return 1
	}
	return int64(math.Round(gap))
}

// roundGroups rounds the balances within each group to the lowest value in the group.
func roundGroups(groups [][]*Participant) []*Participant {
	roundedCensus := []*Participant{}
	for _, group := range groups {
		if len(group) == 0 {
			continue
		}
		lowestBalance := roundToFirstCommonDigit(group)
		for _, participant := range group {
			roundedCensus = append(roundedCensus, &Participant{Address: participant.Address, Balance: lowestBalance})
		}
	}
	return roundedCensus
}

// calculateAccuracy computes the accuracy of the rounding process.
func calculateAccuracy(original, rounded []*Participant) float64 {
	var totalOriginal, totalRounded big.Int
	for i := range original {
		totalOriginal.Add(&totalOriginal, original[i].Balance)
		totalRounded.Add(&totalRounded, rounded[i].Balance)
	}
	totalOriginalFloat := new(big.Float).SetInt(&totalOriginal)
	if totalOriginalFloat.Cmp(big.NewFloat(0)) == 0 {
		return 0
	}
	lostWeight := new(big.Float).Sub(totalOriginalFloat, new(big.Float).SetInt(&totalRounded))
	accuracy, _ := new(big.Float).Quo(lostWeight, totalOriginalFloat).Float64()
	return 100 - (accuracy * 100)
}

// groupAndRoundCensus groups the cleanedParticipants and rounds their balances.
func groupAndRoundCensus(participants []*Participant, privacyThreshold int64, groupBalanceDiff *big.Int) []*Participant {
	sort.Sort(ByBalance(participants))
	var groups [][]*Participant
	var currentGroup []*Participant
	for i, participant := range participants {
		if len(currentGroup) == 0 {
			currentGroup = append(currentGroup, participant)
		} else {
			lastParticipant := currentGroup[len(currentGroup)-1]
			balanceDiff := new(big.Int).Abs(new(big.Int).Sub(participant.Balance, lastParticipant.Balance))

			if int64(len(currentGroup)) < privacyThreshold || balanceDiff.Cmp(groupBalanceDiff) <= 0 {
				currentGroup = append(currentGroup, participant)
			} else {
				groups = append(groups, currentGroup)
				currentGroup = []*Participant{participant}
			}
		}
		// Ensure the last group is added
		if i == len(participants)-1 {
			groups = append(groups, currentGroup)
		}
	}
	roundedCensus := roundGroups(groups)
	return roundedCensus
}

// zScore identifies and returns outliers based on a specified z-score
// threshold. The z-score is the number of standard deviations from the mean a
// data point is. For example, a z-score of 2 means the data point is 2 standard
// deviations above the mean. A z-score of -2 means it is 2 standard deviations
// below the mean. A number is considered an outlier if its z-score is greater
// than the threshold.
func zScore(participants []*Participant, threshold float64) ([]*Participant, []*Participant) {
	// calculate mean and standard deviation
	// mean = sum of all values / number of values
	// standard deviation = sqrt(sum of all (value - mean)^2 / number of values)
	mean := new(big.Float)
	stdDev := new(big.Float)
	n := new(big.Float).SetInt64(int64(len(participants)))
	fBalances := make([]*big.Float, len(participants))
	for i, p := range participants {
		fBalance := new(big.Float).SetInt(p.Balance)
		fBalances[i] = fBalance
		mean = new(big.Float).Add(mean, fBalance)
	}
	mean = new(big.Float).Quo(mean, n)
	for _, balance := range fBalances {
		diff := new(big.Float).Sub(balance, mean)
		stdDev = new(big.Float).Add(stdDev, new(big.Float).Mul(diff, diff))
	}
	stdDev = new(big.Float).Quo(stdDev, n)
	stdDev = new(big.Float).Sqrt(stdDev)
	// calculate z-score for each value to determine outliers
	outliers := make([]*Participant, 0)
	newParticipants := make([]*Participant, 0)
	for _, p := range participants {
		// z-score = (value - mean) / standard deviation
		fBalance := new(big.Float).SetInt(p.Balance)
		diff := new(big.Float).Sub(fBalance, mean)
		if stdDev.Cmp(big.NewFloat(0)) == 0 {
			newParticipants = append(newParticipants, p)
			continue
		}
		zScore := new(big.Float).Quo(diff, stdDev)
		// if z-score is greater than threshold, it is an outlier
		if zScore.Abs(zScore).Cmp(big.NewFloat(threshold)) > 0 {
			outliers = append(outliers, p)
		} else {
			newParticipants = append(newParticipants, p)
		}
	}
	return newParticipants, outliers
}

func roundToFirstCommonDigit(participants []*Participant) *big.Int {
	// check if at least two numbers is provided
	if len(participants) == 0 {
		return big.NewInt(0)
	}
	if len(participants) == 1 {
		return participants[0].Balance
	}
	// get the minimun length of any number
	sBalances := []string{}
	minBalance := participants[0].Balance
	minLenght := int64(len(participants[0].Balance.String()))
	for _, n := range participants {
		sBalances = append(sBalances, n.Balance.String())
		if l := int64(len(n.Balance.String())); l < minLenght {
			minLenght = l
			minBalance = n.Balance
		}
	}

	firstCommonByte := int64(minLenght - 1)
	for firstCommonByte >= 0 {
		commonNumber := true
		currentNumber := sBalances[0][firstCommonByte]
		for _, n := range sBalances[1:] {
			if n[firstCommonByte] != currentNumber {
				commonNumber = false
				break
			}
		}
		if commonNumber {
			firstCommonByte++
			padding := new(big.Int).Exp(big.NewInt(10), big.NewInt(minLenght-firstCommonByte), nil)
			rounded := new(big.Int)
			rounded.SetString(sBalances[0][:firstCommonByte], 10)
			return rounded.Mul(rounded, padding)
		}
		firstCommonByte--
	}
	return minBalance
}
