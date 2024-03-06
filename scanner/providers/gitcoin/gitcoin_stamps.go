package gitcoin

import (
	"math/big"
	"regexp"
)

var stamps = map[string]*regexp.Regexp{
	"BrightID":              regexp.MustCompile(`Brightid`),
	"Civic":                 regexp.MustCompile(`^Civic[Captcha|Liveness|Uniqueness]+Pass`),
	"Coinbase":              regexp.MustCompile(`^Coinbase[DualVerification]*`),
	"GTCStaking":            regexp.MustCompile(`^[Beginner|Experienced]+CommunityStaker|^TrustedCitizen|^CommunityStaking[Bronze|Gold|Silver]+`),
	"Discord":               regexp.MustCompile(`^Discord`),
	"Ens":                   regexp.MustCompile(`^Ens`),
	"Ethereum":              regexp.MustCompile(`^ETH[Advocate|Enthusiast|Maxi|Pioneer]+`),
	"Gitcoin":               regexp.MustCompile(`^Gitcoin.+#numGrantsContributeToGte#[1|10|25|100]|^Gitcoin.+#totalContributionAmountGte#[10|100|1000]`),
	"Github":                regexp.MustCompile(`^githubAccountCreationGte#[90|180|365].+|^githubContributionActivityGte#[30|60|120]`),
	"GnosisSafe":            regexp.MustCompile(`^GnosisSafe`),
	"Google":                regexp.MustCompile(`^Google`),
	"GuildMembership&Roles": regexp.MustCompile(`^Guild[Admin|Member|PassportMember]`),
	"Hololym":               regexp.MustCompile(`^HolonymGovIdProvider`),
	"Idena":                 regexp.MustCompile(`^IdenaAge#[5|10]|^IdenaStake#[1k|10k|100k]|^IdenaState#[Human|Newbie|Verified]`),
	"Lens":                  regexp.MustCompile(`^Lens`),
	"LinkedIn":              regexp.MustCompile(`^Linkedin`),
	"NFTHolder":             regexp.MustCompile(`^NFT`),
	"PHI":                   regexp.MustCompile(`^PHIActivity[Gold|Silver]`),
	"ProofOfHumanity":       regexp.MustCompile(`^Poh`),
	"Snapshot":              regexp.MustCompile(`^Snapshot[Proposals|Votes]+Provider`),
	"TrustaLabs":            regexp.MustCompile(`^TrustaLabs`),
	"Twitter":               regexp.MustCompile(`^twitterAccountAgeGte#[180|365|730]`),
	"ZkSync":                regexp.MustCompile(`^ZkSync[Era]*`),
}

func findStamp(alias string) (string, bool) {
	for stamp, re := range stamps {
		if re.MatchString(alias) {
			return stamp, true
		}
	}
	return "", false
}

func parseStamps(stamps map[string]any) (map[string]*big.Int, error) {
	results := make(map[string]*big.Int)
	for alias, score := range stamps {
		stamp, ok := findStamp(alias)
		if !ok {
			continue
		}
		if stampScore := parseScore(score); stampScore != nil {
			if currentScore, ok := results[stamp]; ok {
				results[stamp] = new(big.Int).Add(currentScore, stampScore)
			} else {
				results[stamp] = stampScore
			}
		}
	}
	return results, nil
}
