// Package skirmish contains code for production of the
// Dreamkeepers: Skirmish battle card game.
//
// More specifically, it provides an interface between the SQL database
// that contains card data, Photoshop, and the user (via CLI).
//
// Photoshop
//
// This package selects cards from SQL and creates .csv files to be read into
// Photoshop as datasets.
package skirmish
