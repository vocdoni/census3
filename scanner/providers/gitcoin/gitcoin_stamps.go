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

var stampsIcons = map[string]string{
	// include and empty string to match the Gitcoin base token with the same icon
	// than the Gitcoin stamp
	"":                      "https://ipfs.io/ipfs/bafybeieuv2qxisdhpyye7nsklzf7dlircabomj632su7klhbpg5sv56wl4/gitcoin.svg",
	"BrightID":              "https://ipfs.io/ipfs/bafybeieuv2qxisdhpyye7nsklzf7dlircabomj632su7klhbpg5sv56wl4/brightid.svg",
	"Civic":                 "https://ipfs.io/ipfs/bafybeieuv2qxisdhpyye7nsklzf7dlircabomj632su7klhbpg5sv56wl4/civic.svg",
	"Coinbase":              "https://ipfs.io/ipfs/bafybeieuv2qxisdhpyye7nsklzf7dlircabomj632su7klhbpg5sv56wl4/coinbase.svg",
	"GTCStaking":            "https://ipfs.io/ipfs/bafybeieuv2qxisdhpyye7nsklzf7dlircabomj632su7klhbpg5sv56wl4/gtcStaking.svg",
	"Discord":               "https://ipfs.io/ipfs/bafybeieuv2qxisdhpyye7nsklzf7dlircabomj632su7klhbpg5sv56wl4/discord.svg",
	"Ens":                   "https://ipfs.io/ipfs/bafybeieuv2qxisdhpyye7nsklzf7dlircabomj632su7klhbpg5sv56wl4/ens.svg",
	"Ethereum":              "https://ipfs.io/ipfs/bafybeieuv2qxisdhpyye7nsklzf7dlircabomj632su7klhbpg5sv56wl4/ethereum.svg",
	"Gitcoin":               "https://ipfs.io/ipfs/bafybeieuv2qxisdhpyye7nsklzf7dlircabomj632su7klhbpg5sv56wl4/gitcoin.svg",
	"Github":                "https://ipfs.io/ipfs/bafybeieuv2qxisdhpyye7nsklzf7dlircabomj632su7klhbpg5sv56wl4/github.svg",
	"GnosisSafe":            "https://ipfs.io/ipfs/bafybeieuv2qxisdhpyye7nsklzf7dlircabomj632su7klhbpg5sv56wl4/gnosisSafe.svg",
	"Google":                "https://ipfs.io/ipfs/bafybeieuv2qxisdhpyye7nsklzf7dlircabomj632su7klhbpg5sv56wl4/google.svg",
	"GuildMembership&Roles": "https://ipfs.io/ipfs/bafybeieuv2qxisdhpyye7nsklzf7dlircabomj632su7klhbpg5sv56wl4/guild.svg",
	"Hololym":               "https://ipfs.io/ipfs/bafybeieuv2qxisdhpyye7nsklzf7dlircabomj632su7klhbpg5sv56wl4/holonym.svg",
	"Idena":                 "https://ipfs.io/ipfs/bafybeieuv2qxisdhpyye7nsklzf7dlircabomj632su7klhbpg5sv56wl4/idena.svg",
	"Lens":                  "https://ipfs.io/ipfs/bafybeieuv2qxisdhpyye7nsklzf7dlircabomj632su7klhbpg5sv56wl4/lens.svg",
	"LinkedIn":              "https://ipfs.io/ipfs/bafybeieuv2qxisdhpyye7nsklzf7dlircabomj632su7klhbpg5sv56wl4/linkedin.svg",
	"NFTHolder":             "https://ipfs.io/ipfs/bafybeieuv2qxisdhpyye7nsklzf7dlircabomj632su7klhbpg5sv56wl4/nft.svg",
	"PHI":                   "https://ipfs.io/ipfs/bafybeieuv2qxisdhpyye7nsklzf7dlircabomj632su7klhbpg5sv56wl4/phi.svg",
	"ProofOfHumanity":       "https://ipfs.io/ipfs/bafybeieuv2qxisdhpyye7nsklzf7dlircabomj632su7klhbpg5sv56wl4/poh.svg",
	"Snapshot":              "https://ipfs.io/ipfs/bafybeieuv2qxisdhpyye7nsklzf7dlircabomj632su7klhbpg5sv56wl4/snapshot.svg",
	"TrustaLabs":            "https://ipfs.io/ipfs/bafybeieuv2qxisdhpyye7nsklzf7dlircabomj632su7klhbpg5sv56wl4/trustaLabs.svg",
	"Twitter":               "https://ipfs.io/ipfs/bafybeieuv2qxisdhpyye7nsklzf7dlircabomj632su7klhbpg5sv56wl4/twitter.svg",
	"ZkSync":                "https://ipfs.io/ipfs/bafybeieuv2qxisdhpyye7nsklzf7dlircabomj632su7klhbpg5sv56wl4/zksync.svg",
}

var stampsNames = map[string]string{
	"BrightID":              "BrightID",
	"Civic":                 "Civic",
	"Coinbase":              "Coinbase",
	"GTCStaking":            "GTC Staking",
	"Discord":               "Discord",
	"Ens":                   "ENS",
	"Ethereum":              "Ethereum",
	"Gitcoin":               "Gitcoin",
	"Github":                "Github",
	"GnosisSafe":            "Gnosis Safe",
	"Google":                "Google",
	"GuildMembership&Roles": "Guild Membership & Roles",
	"Hololym":               "Holonym",
	"Idena":                 "Idena",
	"Lens":                  "Lens",
	"LinkedIn":              "LinkedIn",
	"NFTHolder":             "NFT Holder",
	"PHI":                   "PHI",
	"ProofOfHumanity":       "Proof of Humanity",
	"Snapshot":              "Snapshot",
	"TrustaLabs":            "Trusta Labs",
	"Twitter":               "Twitter",
	"ZkSync":                "ZkSync",
}

const noIconURI = "https://ipfs.io/ipfs/bafybeiehrtssiqivcxq3af2fnyeys5n75cka62irkks66ueq2hgllq43ji/no-image.svg"

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

func stampIcon(stamp string) string {
	if icon, ok := stampsIcons[stamp]; ok {
		return icon
	}
	return noIconURI
}
