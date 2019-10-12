package snapshot

import "flag"

var shouldUpdate *bool

// init is called, only once, when the package is imported
func init() {
	// @TODO: Support more flags e.g. custom snapshot directory, custom snapshot extension
	// Flags can be passed to the `go test` command
	// e.g.
	// 	go test ./... -u # updates snapshots
	shouldUpdate = flag.Bool("u", false, "Updates a snapshot when not matched correctly")
	flag.Parse()
}
