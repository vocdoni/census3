package gitcoin

import (
	"math"
	"math/big"
	"strconv"
	"time"

	"github.com/ethereum/go-ethereum/common"
)

const dateLayout = "2006-01-02T15:04:05.999Z"

// GitcoinScore represents a Gitcoin score result object from the JSONL API.
type GitcoinScore struct {
	Passport struct {
		Address string `json:"address"`
	} `json:"passport"`
	LastScore          string         `json:"score"`
	StampScores        map[string]any `json:"stamp_scores"`
	LastScoreTimestamp string         `json:"last_score_timestamp"`
	Evidence           *struct {
		Score   string `json:"rawScore"`
		Success bool   `json:"success"`
	} `json:"evidence"`
}

// Address method returns the address of the score object as a common.Address.
func (score *GitcoinScore) Address() common.Address {
	return common.HexToAddress(score.Passport.Address)
}

// Timestamp method returns the last score timestamp of the score object as a
// time.Time and an error if something fails parsing the timestamp.
func (score *GitcoinScore) Timestamp() (time.Time, error) {
	return time.Parse(dateLayout, score.LastScoreTimestamp)
}

// Valid method returns true if the score object is valid, false otherwise. A
// score object is valid if the last score and evidence are not empty nor zero
// and the evidence is successful.
func (score *GitcoinScore) Valid() bool {
	validLastScore := score.LastScore != "" && score.LastScore != "0E-9"
	validEvidence := score.Evidence != nil && score.Evidence.Success
	validEvidenceScore := score.Evidence != nil && score.Evidence.Score != "" && score.Evidence.Score != "0E-9"
	return validLastScore && validEvidence && validEvidenceScore
}

// Score method returns the score of the score object as a *big.Int or nil if
// the score object is not valid. It uses the parseScore function to parse the
// score from the evidence.
func (score *GitcoinScore) Score() *big.Int {
	if !score.Valid() {
		return nil
	}
	return parseScore(score.Evidence.Score)
}

// Stamps method returns the stamps of the score object as a map[string]*big.Int
// or an error if something fails parsing the stamps. It uses the parseStamps
// function to parse the stamps from the stamp scores grouped by stamp.
func (score *GitcoinScore) Stamps() (map[string]*big.Int, error) {
	return parseStamps(score.StampScores)
}

// parseScore function returns the score of the input as a *big.Int or nil if
// the input is not a valid score. It uses the strconv.ParseFloat function to
// parse the score from a string and the math.Round function to round the score
// to the nearest integer. If the score is between 0 and 1, it sets it to 1 to
// avoid rounding errors. If the score is 0, it omits it returning nil.
func parseScore(input any) *big.Int {
	var fScore float64
	switch value := input.(type) {
	case string:
		var err error
		fScore, err = strconv.ParseFloat(value, 64)
		if err != nil {
			return nil
		}
	case float64:
		fScore = value
	default:
		return nil
	}
	// truncate the score to the nearest integer and return it as a big.Int
	if biScore := big.NewInt(int64(math.Trunc(fScore))); biScore.Cmp(big.NewInt(0)) == 1 {
		return biScore
	}
	return nil
}
