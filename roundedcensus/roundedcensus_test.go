package roundedcensus

import (
	"math/big"
	"testing"

	qt "github.com/frankban/quicktest"
)

func Test_GroupAndRoundCensus(t *testing.T) {
	c := qt.New(t)

	participants := []*Participant{
		{Address: "0x1", Balance: big.NewInt(1)},
		{Address: "0x2", Balance: big.NewInt(2)},
		{Address: "0x3", Balance: big.NewInt(3)},
	}

	result, accuracy, err := GroupAndRoundCensus(participants, DefaultGroupsConfig)
	c.Assert(err, qt.IsNil)
	c.Assert(len(result), qt.Equals, 3)
	c.Assert(accuracy, qt.Equals, 50.0)
	for _, p := range result {
		c.Assert(p.Balance.String(), qt.Equals, "1")
	}

	participants[1].Balance = big.NewInt(1000)
	participants[2].Balance = big.NewInt(1000000)
	result, accuracy, err = GroupAndRoundCensus(participants, DefaultGroupsConfig)
	c.Assert(err, qt.IsNotNil)
	c.Assert(len(result), qt.Equals, 3)
	c.Assert(accuracy < 0.01, qt.IsTrue, qt.Commentf("accuracy should be near to 0"))
	for _, p := range result {
		c.Assert(p.Balance.String(), qt.Equals, "1")
	}

	participants = []*Participant{
		{Address: "0x1", Balance: big.NewInt(1)},
		{Address: "0x2", Balance: big.NewInt(2)},
		{Address: "0x3", Balance: big.NewInt(3)},
		{Address: "0x4", Balance: big.NewInt(1)},
		{Address: "0x5", Balance: big.NewInt(2)},
		{Address: "0x6", Balance: big.NewInt(3)},
		{Address: "0x7", Balance: big.NewInt(1)},
		{Address: "0x8", Balance: big.NewInt(2)},
		{Address: "0x9", Balance: big.NewInt(3)},
		{Address: "0x10", Balance: big.NewInt(1)},
		{Address: "0x11", Balance: big.NewInt(2)},
		{Address: "0x12", Balance: big.NewInt(3)},
	}
	result, accuracy, err = GroupAndRoundCensus(participants, DefaultGroupsConfig)
	c.Assert(err, qt.IsNil)
	c.Assert(len(result), qt.Equals, 12)
	c.Assert(accuracy, qt.Equals, 50.0)
	for _, p := range result {
		c.Assert(p.Balance.String(), qt.Equals, "1")
	}
}

func Test_zScore(t *testing.T) {
	c := qt.New(t)
	// zScore method calculates the outliers for a given dataset. Outliers are
	// every value that is more than n standard deviations away from the mean.
	// The value of n is required to use the zScore method.

	// initial participants
	participants := []*Participant{
		{Address: "0x1", Balance: big.NewInt(1)},
		{Address: "0x2", Balance: big.NewInt(50)},
		{Address: "0x3", Balance: big.NewInt(51)},
		{Address: "0x4", Balance: big.NewInt(52)},
		{Address: "0x5", Balance: big.NewInt(53)},
		{Address: "0x6", Balance: big.NewInt(52)},
		{Address: "0x7", Balance: big.NewInt(51)},
		{Address: "0x8", Balance: big.NewInt(50)},
		{Address: "0x9", Balance: big.NewInt(100)},
	}
	// expected participants after removing outliers
	expectedParticipants := []*Participant{
		{Address: "0x2", Balance: big.NewInt(50)},
		{Address: "0x3", Balance: big.NewInt(51)},
		{Address: "0x4", Balance: big.NewInt(52)},
		{Address: "0x5", Balance: big.NewInt(53)},
		{Address: "0x6", Balance: big.NewInt(52)},
		{Address: "0x7", Balance: big.NewInt(51)},
		{Address: "0x8", Balance: big.NewInt(50)},
	}
	// expected outliers
	expectedOutliers := []*Participant{
		{Address: "0x1", Balance: big.NewInt(1)},
		{Address: "0x9", Balance: big.NewInt(100)},
	}
	// results
	cleanedParticipants, outliers := zScore(participants, DefaultGroupsConfig.OutliersThreshold)
	c.Assert(len(cleanedParticipants), qt.Equals, len(expectedParticipants))
	for i, p := range cleanedParticipants {
		c.Assert(p.Address, qt.Equals, expectedParticipants[i].Address)
	}
	for i, p := range outliers {
		c.Assert(p.Address, qt.Equals, expectedOutliers[i].Address)
	}
}

func Test_groupAndRoundCensus(t *testing.T) {
	c := qt.New(t)

	participants := []*Participant{
		{Address: "0x1", Balance: big.NewInt(1)},
		{Address: "0x2", Balance: big.NewInt(2)},
		{Address: "0x3", Balance: big.NewInt(3)},
		{Address: "0x4", Balance: big.NewInt(118)},
		{Address: "0x5", Balance: big.NewInt(119)},
		{Address: "0x6", Balance: big.NewInt(120)},
		{Address: "0x7", Balance: big.NewInt(1200)},
		{Address: "0x8", Balance: big.NewInt(1290)},
		{Address: "0x9", Balance: big.NewInt(1299)},
		{Address: "0x10", Balance: big.NewInt(1400)},
		{Address: "0x11", Balance: big.NewInt(1460)},
		{Address: "0x12", Balance: big.NewInt(1560)},
	}
	expectedCensus := []*Participant{
		// group1
		{Address: "0x1", Balance: big.NewInt(1)},
		{Address: "0x2", Balance: big.NewInt(1)},
		{Address: "0x3", Balance: big.NewInt(1)},
		// group2
		{Address: "0x4", Balance: big.NewInt(118)},
		{Address: "0x5", Balance: big.NewInt(118)},
		{Address: "0x6", Balance: big.NewInt(118)},
		// group3
		{Address: "0x7", Balance: big.NewInt(1200)},
		{Address: "0x8", Balance: big.NewInt(1200)},
		{Address: "0x9", Balance: big.NewInt(1200)},
		// group4
		{Address: "0x10", Balance: big.NewInt(1400)},
		{Address: "0x11", Balance: big.NewInt(1400)},
		{Address: "0x12", Balance: big.NewInt(1400)},
	}
	result := groupAndRoundCensus(participants, 3, DefaultGroupsConfig.GroupBalanceDiff)
	c.Assert(len(result), qt.Equals, len(expectedCensus))
	for i, p := range result {
		c.Assert(p.Address, qt.Equals, expectedCensus[i].Address)
		c.Assert(p.Balance.String(), qt.Equals, expectedCensus[i].Balance.String())
	}
}

func Test_flatAndRoundGroups(t *testing.T) {
	c := qt.New(t)
	// flatAndRoundGroups method rounds a given list of groups to the first
	// common digit of their balances. It also flattens the groups into a single
	// list of participants.
	groups := [][]*Participant{
		{
			{Address: "0x1", Balance: big.NewInt(1)},
			{Address: "0x2", Balance: big.NewInt(2)},
			{Address: "0x3", Balance: big.NewInt(3)},
		},
		{
			{Address: "0x4", Balance: big.NewInt(118)},
			{Address: "0x5", Balance: big.NewInt(119)},
			{Address: "0x6", Balance: big.NewInt(120)},
		},
		{
			{Address: "0x7", Balance: big.NewInt(1200)},
			{Address: "0x8", Balance: big.NewInt(1290)},
			{Address: "0x9", Balance: big.NewInt(1299)},
			{Address: "0x10", Balance: big.NewInt(1400)},
			{Address: "0x11", Balance: big.NewInt(1560)},
			{Address: "0x12", Balance: big.NewInt(1560)},
		},
	}
	expectedRoundedGroups := []*Participant{
		// group1
		{Address: "0x1", Balance: big.NewInt(1)},
		{Address: "0x2", Balance: big.NewInt(1)},
		{Address: "0x3", Balance: big.NewInt(1)},
		// group2
		{Address: "0x4", Balance: big.NewInt(118)},
		{Address: "0x5", Balance: big.NewInt(118)},
		{Address: "0x6", Balance: big.NewInt(118)},
		// group3
		{Address: "0x7", Balance: big.NewInt(1200)},
		{Address: "0x8", Balance: big.NewInt(1200)},
		{Address: "0x9", Balance: big.NewInt(1200)},
		{Address: "0x10", Balance: big.NewInt(1200)},
		{Address: "0x11", Balance: big.NewInt(1200)},
		{Address: "0x12", Balance: big.NewInt(1200)},
	}
	roundedGroups := flatAndRoundGroups(groups)
	c.Assert(len(roundedGroups), qt.Equals, len(expectedRoundedGroups))
	for i, g := range roundedGroups {
		c.Assert(g.Address, qt.Equals, expectedRoundedGroups[i].Address)
		c.Assert(g.Balance.String(), qt.Equals, expectedRoundedGroups[i].Balance.String())
	}
}

func Test_roundGap(t *testing.T) {
	c := qt.New(t)
	c.Assert(roundGap(1), qt.Equals, int64(1))
	c.Assert(roundGap(10), qt.Equals, int64(1))
	c.Assert(roundGap(20), qt.Equals, int64(1))
	c.Assert(roundGap(40), qt.Equals, int64(2))
	c.Assert(roundGap(100), qt.Equals, int64(5))
	c.Assert(roundGap(110), qt.Equals, int64(6))
}

func Test_calculateAccuracy(t *testing.T) {
	c := qt.New(t)
	// calculateAccuracy method calculates the accuracy of a given list of
	// participants. The accuracy is calculated comparing the sum of the
	// participants balances with the sum of the rounded balances.
	originalParticipants := []*Participant{
		{Address: "0x1", Balance: big.NewInt(11)},
		{Address: "0x1", Balance: big.NewInt(12)},
		{Address: "0x1", Balance: big.NewInt(9)},
		{Address: "0x1", Balance: big.NewInt(8)},
	}
	roundedParticipants := []*Participant{
		{Address: "0x1", Balance: big.NewInt(10)},
		{Address: "0x1", Balance: big.NewInt(11)},
		{Address: "0x1", Balance: big.NewInt(8)},
		{Address: "0x1", Balance: big.NewInt(7)},
	}
	c.Assert(calculateAccuracy(originalParticipants, roundedParticipants), qt.Equals, 90.0)
	roundedParticipants = []*Participant{
		{Address: "0x1", Balance: big.NewInt(2)},
		{Address: "0x1", Balance: big.NewInt(2)},
		{Address: "0x1", Balance: big.NewInt(0)},
		{Address: "0x1", Balance: big.NewInt(0)},
	}
	c.Assert(calculateAccuracy(originalParticipants, roundedParticipants) > 9.9, qt.IsTrue,
		qt.Commentf("accuracy should be greater than 19.9"))
	c.Assert(calculateAccuracy(originalParticipants, roundedParticipants) <= 10, qt.IsTrue,
		qt.Commentf("and also should be less or equal than 20"))
}
