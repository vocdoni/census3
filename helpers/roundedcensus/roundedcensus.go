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
	"runtime"
	"sort"
	"sync"
	"sync/atomic"
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
	GroupBalanceDiff:    big.NewInt(0),
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
func GroupAndRoundCensus(participants []*Participant, config GroupsConfig,
	progressCh chan float64,
) ([]*Participant, float64, error) {
	// create vars to parallelize the accuracy optimization loop, create a
	// bestResult struct to store the best accuracy and privacy threshold used
	// to achieve it, and the best struct to store the best result and a mutex
	// to update it safely
	type bestResult struct {
		accuracy         float64
		privacyThreshold int64
	}
	best := struct {
		sync.Mutex
		result bestResult
	}{result: bestResult{accuracy: 0.0}}
	// create a wait group and a semaphore to limit the number of goroutines to
	// the number of CPU cores, calculate the maximum privacy
	var wg sync.WaitGroup
	numWorkers := runtime.NumCPU() // Limit the number of goroutines to the number of CPU cores
	goroutineSem := make(chan struct{}, numWorkers)
	// calculate bounds for the accuracy optimization loop, if the number of
	// participants divided by the minimum privacy threshold is less than the
	// minimum privacy threshold, use the minimum privacy threshold as the
	// maximum privacy threshold to iterate at least once
	maxPrivacyThreshold := int64(len(participants)) / config.MinPrivacyThreshold
	if maxPrivacyThreshold < config.MinPrivacyThreshold {
		maxPrivacyThreshold = config.MinPrivacyThreshold
	}
	// calculate outliers and cleaned participants using z-score and sort the
	// cleaned ones by balance
	cleanedParticipants, outliers := zScore(participants, config.OutliersThreshold)
	sortedParticipants := append([]*Participant{}, cleanedParticipants...)
	sort.Sort(ByBalance(sortedParticipants))
	// create some vars to track the progress of the accuracy optimization loop,
	// incrase the number of setups to process by one to include the last
	// iteration to calculate the final rounded census for the chosen setup
	setupsToProcess := maxPrivacyThreshold - config.MinPrivacyThreshold + 1
	proccessedSetups := atomic.Int64{}
	// iterate over privacy thresholds using the roundGap function to calculate
	// the gap between them, and calculate the accuracy for each privacy
	// threshold using a goroutine for each iteration
	for current := config.MinPrivacyThreshold; current <= maxPrivacyThreshold; current += roundGap(current) {
		goroutineSem <- struct{}{} // acquire goroutines semaphore
		wg.Add(1)
		go func(privacyThreshold int64) {
			defer wg.Done()
			// group and round participants and calculate accuracy
			groupedParticipants := groupAndRoundCensus(sortedParticipants, privacyThreshold, config.GroupBalanceDiff)
			accuracy := calculateAccuracy(sortedParticipants, groupedParticipants)
			// lock best result before checking and updating it
			best.Lock()
			// update best result if the current accuracy is higher or if it is
			// the same and the current privacy threshold is higher
			if accuracy > best.result.accuracy || (accuracy == best.result.accuracy && privacyThreshold > best.result.privacyThreshold) {
				best.result = bestResult{accuracy, privacyThreshold}
			}
			// unlock best result
			best.Unlock()
			lastProcessed := current - config.MinPrivacyThreshold
			if alreadyProcessed := proccessedSetups.Load(); lastProcessed > alreadyProcessed {
				proccessedSetups.Store(lastProcessed)
				if progressCh != nil {
					progressCh <- float64(alreadyProcessed+1) / float64(setupsToProcess) * 100
				}
			}
			proccessedSetups.CompareAndSwap(lastProcessed, lastProcessed+1)
			<-goroutineSem // release goroutines semaphore
		}(current)
	}
	// close iteration channel when all goroutines are done
	wg.Wait()
	// calculate the final rounded census with the highest calculated accuracy
	finalPrivacyThreshold := best.result.privacyThreshold
	roundedCensus := groupAndRoundCensus(sortedParticipants, finalPrivacyThreshold, config.GroupBalanceDiff)
	outliersCensus := groupAndRoundCensus(outliers, finalPrivacyThreshold, config.GroupBalanceDiff)
	roundedCensus = append(roundedCensus, outliersCensus...)
	// send 100% progress to the channel to indicate the process is done
	if progressCh != nil {
		progressCh <- 100
	}
	// return the final rounded census and the highest accuracy, if it does not
	// satisfy the minimum accuracy requirement for the rounding process, return
	// an error, if not, just return the final rounded census and the highest
	// accuracy
	if best.result.accuracy < config.MinAccuracy {
		return roundedCensus, best.result.accuracy, fmt.Errorf("could not find a privacy threshold that satisfies the minimum accuracy")
	}
	return roundedCensus, best.result.accuracy, nil
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
	if stdDev.Cmp(big.NewFloat(0)) == 0 {
		return participants, []*Participant{}
	}
	// calculate z-score for each value to determine outliers
	outliers := make([]*Participant, 0)
	newParticipants := make([]*Participant, 0)
	for _, p := range participants {
		// z-score = (value - mean) / standard deviation
		fBalance := new(big.Float).SetInt(p.Balance)
		diff := new(big.Float).Sub(fBalance, mean)
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

// groupAndRoundCensus groups the cleanedParticipants and rounds their balances.
func groupAndRoundCensus(participants []*Participant, privacyThreshold int64, groupBalanceDiff *big.Int) []*Participant {
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
	roundedCensus := flatAndRoundGroups(groups)
	return roundedCensus
}

// flatAndRoundGroups function iterate over formed groups and round their
// balances using the lowest group balance. The function returns the list of
// participants with rounded balances flattened.
func flatAndRoundGroups(groups [][]*Participant) []*Participant {
	roundedCensus := []*Participant{}
	for _, group := range groups {
		if len(group) == 0 {
			continue
		}
		for _, participant := range group {
			roundedCensus = append(roundedCensus, &Participant{Address: participant.Address, Balance: group[0].Balance})
		}
	}
	return roundedCensus
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

// calculateAccuracy computes the accuracy of the rounding process.
func calculateAccuracy(original, rounded []*Participant) float64 {
	totalOriginal, totalRounded := new(big.Int), new(big.Int)
	for i := range original {
		totalOriginal.Add(totalOriginal, original[i].Balance)
		totalRounded.Add(totalRounded, rounded[i].Balance)
	}
	totalOriginalFloat := new(big.Float).SetInt(totalOriginal)
	if totalOriginalFloat.Cmp(big.NewFloat(0)) == 0 {
		return 0
	}
	lostWeight := new(big.Float).Sub(totalOriginalFloat, new(big.Float).SetInt(totalRounded))
	accuracy, _ := new(big.Float).Abs(new(big.Float).Quo(lostWeight, totalOriginalFloat)).Float64()
	return 100 - (accuracy * 100)
}
