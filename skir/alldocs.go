// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// DO NOT EDIT THIS FILE. GENERATED BY mkalldocs.sh.
// Edit the documentation in other files and rerun mkalldocs.sh to generate this one.

// Skir is a tool for developing the Skirmish card game.
//
// Usage:
//
// 	skir command [arguments]
//
// The commands are:
//
// 	card        show information about a specific card
// 	dump        save the current database to disk
// 	export      compile cards from the database to a specific format
// 	ps          fill out Photoshop templates
// 	recover     reload the database from disk
// 	sql         query the database
// 	version     print skir version
//
// Use "skir help [command]" for more information about a command.
//
//
// Show information about a specific card
//
// Usage:
//
// 	skir card [-fmt=[format]] [card name]
//
// Card prints data for the given card to standard output.
//
// The -fmt flag can be used to alter the output format. The valid formats are: "string", "ue", and "xml".
//
//
// Save the current database to disk
//
// Usage:
//
// 	skir dump
//
// Dump saves the current state of the database to the "skirmish_db.sql"
// file in the dreamkeepers-dat repository.
//
//
// Compile cards from the database to a specific format
//
// Usage:
//
// 	skir export [format]
//
// 'Skir export' pulls information for all cards from the database and compiles them into the given format.
//
// The valid formats are:
// 	csv	csv formatted files to use as datasets in Photoshop.
// 		One file is generated for Deck Cards, and another is generated for Non-Deck Cards.
// 		The files are generated in the top level of the "dreamkeepers-psd" repository.
// 	ue	a collection of JSON files for importing into Unreal Engine.
// 		Deck cards are grouped by deck, Non-Deck Cards are grouped together.
// 		The files can be found in the "Unreal_JSONs" folder in the skirmish repository.
//
//
// Fill out Photoshop templates
//
// Usage:
//
// 	skir ps [deck] [all] [card name]
//
// 'skir ps' fills out a Photoshop template file with information from the database.
//
//
// See the skirmish/ps package for instructions on setting up this tool.
//
// To start generating files, invoke the tool with the cards you want. Available options are:
// 	- 'skir ps all' to generate all cards.
// 	- 'skir ps deck [leader name]' to generate all cards lead by the named character.
// 	- 'skir ps [card id]' to generate one card.
//
// Running the tool will open Photoshop and the necessary template .psd,
// after which it will pause and ask you to load a dataset file.
// Dataset files are csv formatted files that correspond to fields in the Photoshop template.
// The tool generates them for you and puts them in the dreamkeepers data folder as:
// 	- 'deckcards.csv' for deck cards.
// 	- 'nondeckcards.csv' for non-deckcards
//
// To load a dataset file, open Photoshop and navigate to 'Image > Variables > Data Sets...'.
// Make sure Encoding is set to "Automatic", and "Use First Column For Data Set Names" and
// "Replace Existing Data Sets" are selected, then click 'Import' on the right side of the pop-up menu.
// It will take a minute to load, but once it does,
// hit 'OK' and then return to the terminal where you ran the tool and hit enter to continue.
// After this, the program should not require are further user interaction.
//
// The dataset file will only need to be reloaded when the Template is opened or the data is changed.
//
// Deck cards will be output to "[dreamkeepers-psd]/Decks/[leader name]/[card id].png", Nondeck cards will
// be output to "[dreamkeepers-psd]/Decks/Heroes/[card_id].png".
//
// Photoshop is very slow, generating every card could take 15+ minutes, so be ready to wait.
//
//
// Reload the database from disk
//
// Usage:
//
// 	skir recover
//
// Recover runs the skirmish_db.sql file from the dreamkeepers-dat repository
// on the database, effectively resetting it to the most recently saved state. To overwrite this
// file, see 'skir dump'.
//
//
// Query the database
//
// Usage:
//
// 	skir sql [PSQL query]
//
// 'Skir sql' queries the database for any desired information.
//
//
// Print skir version
//
// Usage:
//
// 	skir version
//
// Version prints the installed version of skir.
//
//
package main
