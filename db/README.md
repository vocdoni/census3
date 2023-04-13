# Known issues

## Bindings generation with sqlc

- `Censuses` table generated bindings are wrongly named `Censuse`. For now change this manually when regenerating the bindings.
- In `sqlc.yaml` some tables with composed names like `TokenHolders` are changed to lower-case `tokenholders` for sqlc compatibility.