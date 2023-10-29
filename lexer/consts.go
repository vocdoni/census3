package lexer

const (
	// byte constants characters
	bScape      = byte(92)
	bSpace      = byte(32)
	bStartGroup = byte(40)
	bEndGroup   = byte(41)
	// str constants characters
	scape      = string(bScape)
	space      = string(bSpace)
	startGroup = string(bStartGroup)
	endGroup   = string(bEndGroup)
	// others
	maxGroupID = 99999999
)
