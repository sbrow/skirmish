// TODO: config file for non-programmers
// TODO: command - "card" display info for a card
// TODO: separate commands for each ps operation.
package skirmish

import (
	"os"
	"path/filepath"
)

var Template = filepath.Join(os.Getenv("SK_DIR"), "Template009.1.psd")
var DataDir = filepath.Join(os.Getenv("SK_DIR"), "card_jsons")
var ImageDir = filepath.Join(os.Getenv("SK_DIR"), "Images")
