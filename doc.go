//go:generate sh -c "godoc2md -template ./.doc.template github.com/sbrow/skirmish > README.md"

// Package skirmish contains code for production of the
// Dreamkeepers: Skirmish battle card game.
//
// More specifically, it provides an interface between the SQL database
// that contains card data, Photoshop, and the user (via CLI).
//
// TODO(sbrow): Cameo card flavor text.
package skirmish
