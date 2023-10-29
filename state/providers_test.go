package state

import (
	"testing"

	qt "github.com/frankban/quicktest"
)

func TestCheckWeb3Providers(t *testing.T) {
	c := qt.New(t)

	providers := []string{DefaultWeb3Provider.URI}
	w3p, err := CheckWeb3Providers(providers)
	c.Assert(err, qt.IsNil)
	c.Assert(w3p[DefaultWeb3Provider.ChainID].URI, qt.Equals, DefaultWeb3Provider.URI)
	c.Assert(w3p[DefaultWeb3Provider.ChainID].ShortName, qt.Equals, DefaultWeb3Provider.ShortName)
	c.Assert(w3p[DefaultWeb3Provider.ChainID].Name, qt.Equals, DefaultWeb3Provider.Name)

	t.Run("URIByChainID", func(t *testing.T) {
		_, ok := w3p.URIByChainID(DefaultWeb3Provider.ChainID + 1)
		c.Assert(ok, qt.IsFalse)
		uri, ok := w3p.URIByChainID(DefaultWeb3Provider.ChainID)
		c.Assert(ok, qt.Equals, true)
		c.Assert(uri, qt.Equals, DefaultWeb3Provider.URI)
	})
	t.Run("ChainIDByShortName", func(t *testing.T) {
		_, ok := w3p.ChainIDByShortName("UNKNOWN")
		c.Assert(ok, qt.IsFalse)
		chainID, ok := w3p.ChainIDByShortName(DefaultWeb3Provider.ShortName)
		c.Assert(ok, qt.Equals, true)
		c.Assert(chainID, qt.Equals, DefaultWeb3Provider.ChainID)
	})
	t.Run("ChainAddress", func(t *testing.T) {
		_, ok := w3p.ChainAddress(DefaultWeb3Provider.ChainID+1, "0x1234567890")
		c.Assert(ok, qt.IsFalse)
		prefix, ok := w3p.ChainAddress(DefaultWeb3Provider.ChainID, "0x1234567890")
		c.Assert(ok, qt.Equals, true)
		c.Assert(prefix, qt.Equals, "gor:0x1234567890")
	})
}