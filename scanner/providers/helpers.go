package providers

import (
	"context"
	"fmt"
	"math/big"
	"net/http"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"go.vocdoni.io/dvote/log"
)

// CalcPartialHolders calculates the partial holders from the current and new holders
// maps. It returns a map with the partial holders and their balances. The final
// holders are calculated as:
//  1. Holders that are in the new holders map but not in the current holders
//     map. They are created with the balance of the new holders.
//  2. Holders that are in the new and the current holders maps and have a
//     different balance. They are updated with difference between the new
//     and the current balances.
//  3. Holders that are in the current holders map but not in the new holders
//     map. They are updated with the balance of the current holders negated.
func CalcPartialHolders(currentHolders, newHolders map[common.Address]*big.Int) map[common.Address]*big.Int {
	partialHolders := make(map[common.Address]*big.Int)
	// calculate holders of type 1 and 2
	for addr, newBalance := range newHolders {
		// if the address is not in the current holders, it is a holder of type 1
		// so we add it to the partial holders with the new balance
		currentBalance, alreadyExists := currentHolders[addr]
		if !alreadyExists {
			partialHolders[addr] = newBalance
			continue
		}
		// if the address is in the current holders, it is a holder of type 2
		// so we add it to the partial holders with the difference between the
		// new and the current balances, if the difference is not zero (if it
		// is zero, it has not changed it balance)
		if diff := new(big.Int).Sub(newBalance, currentBalance); diff.Cmp(big.NewInt(0)) != 0 {
			partialHolders[addr] = diff
		}
	}
	// calculate holders of type 3
	for addr, currentBalance := range currentHolders {
		// if the address is not in the new holders, it is a holder of type 3
		// so we add it to the partial holders with the current balance negated
		if _, exists := newHolders[addr]; !exists {
			partialHolders[addr] = new(big.Int).Neg(currentBalance)
		}
	}
	return partialHolders
}

// ServeTestStaticFiles serves the given files in the given port and returns a
// map with the URIs of the served files. The server is shutdown when the given
// context is done. It is useful for testing external providers to emulate
// external API services.
func ServeTestStaticFiles(ctx context.Context, port int, files map[string]string) map[string]string {
	srv := &http.Server{Addr: fmt.Sprintf(":%d", port)}
	// iterate over the files and serve them storing the URIs in a map
	uris := map[string]string{}
	for path := range files {
		uris[path] = fmt.Sprintf("http://localhost:%d%s", port, path)
	}
	// serve the files
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		for path, file := range files {
			currentFile := file
			if strings.HasPrefix(r.URL.Path, path) {
				http.ServeFile(w, r, currentFile)
				return
			}
		}
	})
	// run the server in a goroutine
	go func() {
		if err := srv.ListenAndServe(); err != http.ErrServerClosed {
			log.Errorf("HTTP server error: %v", err)
		}
	}()
	// shutdown the server when the context is done in a goroutine
	go func() {
		<-ctx.Done()
		shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := srv.Shutdown(shutdownCtx); err != nil {
			log.Errorf("HTTP server error: %v", err)
		}
	}()
	// return the URIs
	return uris
}
