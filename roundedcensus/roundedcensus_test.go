package roundedcensus

import (
	"testing"

	qt "github.com/frankban/quicktest"
)

func Test_GroupAndRoundCensus(t *testing.T)     {}
func Test_zScore(t *testing.T)                  {}
func Test_groupAndRoundCensus(t *testing.T)     {}
func Test_roundGroups(t *testing.T)             {}
func Test_roundToFirstCommonDigit(t *testing.T) {}

func Test_roundGap(t *testing.T) {
	c := qt.New(t)
	c.Assert(roundGap(1), qt.Equals, int64(1))
	c.Assert(roundGap(10), qt.Equals, int64(1))
	c.Assert(roundGap(20), qt.Equals, int64(1))
	c.Assert(roundGap(40), qt.Equals, int64(2))
	c.Assert(roundGap(100), qt.Equals, int64(5))
	c.Assert(roundGap(110), qt.Equals, int64(6))
}

func Test_calculateAccuracy(t *testing.T) {}
