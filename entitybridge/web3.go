package entitybridge

import (
	ens "gitlab.com/vocdoni/go-dvote/chain/contracts"
)

// ENS wraps the ENS Registry and the ENS Resolver contracts
type ENS struct {
	*ens.EntityResolver
	*ens.EnsRegistryWithFallback
}
