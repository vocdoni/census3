package web3

import (
	"testing"

	qt "github.com/frankban/quicktest"
)

func TestCheckWeb3Providers(t *testing.T) {
	c := qt.New(t)

	testNetwork, err := InitNetworkEndpoints(DefaultNetworkEndpoint.URIs)
	c.Assert(err, qt.IsNil)
	c.Assert(testNetwork[DefaultNetworkEndpoint.ChainID].URIs, qt.ContentEquals, DefaultNetworkEndpoint.URIs)
	c.Assert(testNetwork[DefaultNetworkEndpoint.ChainID].ShortName, qt.Equals, DefaultNetworkEndpoint.ShortName)
	c.Assert(testNetwork[DefaultNetworkEndpoint.ChainID].Name, qt.Equals, DefaultNetworkEndpoint.Name)

	t.Run("URIByChainID", func(t *testing.T) {
		_, ok := testNetwork.URIsByChainID(DefaultNetworkEndpoint.ChainID + 1)
		c.Assert(ok, qt.IsFalse)
		uri, ok := testNetwork.URIsByChainID(DefaultNetworkEndpoint.ChainID)
		c.Assert(ok, qt.Equals, true)
		c.Assert(uri, qt.ContentEquals, DefaultNetworkEndpoint.URIs)
	})
	t.Run("ChainIDByShortName", func(t *testing.T) {
		_, ok := testNetwork.ChainIDByShortName("UNKNOWN")
		c.Assert(ok, qt.IsFalse)
		chainID, ok := testNetwork.ChainIDByShortName(DefaultNetworkEndpoint.ShortName)
		c.Assert(ok, qt.Equals, true)
		c.Assert(chainID, qt.ContentEquals, DefaultNetworkEndpoint.ChainID)
	})
	t.Run("ChainAddress", func(t *testing.T) {
		_, ok := testNetwork.ChainAddress(DefaultNetworkEndpoint.ChainID+1, "0x1234567890")
		c.Assert(ok, qt.IsFalse)
		prefix, ok := testNetwork.ChainAddress(DefaultNetworkEndpoint.ChainID, "0x1234567890")
		c.Assert(ok, qt.Equals, true)
		c.Assert(prefix, qt.Equals, "gor:0x1234567890")
	})
}
