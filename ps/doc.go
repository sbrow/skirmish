/*
Package ps fills in Photoshop templates with Card information.DeckTemplate

ps acts as a bridge between github.com/sbrow/skirmish and github.com/sbrow/ps,
allowing the user to fill out Photoshop templates with Card data pulled from
an SQL database.

Setup

For the ps tool to work, you will need to have a supported version of Photoshop installed
(see github.com/sbrow/skirmish for details),
and you will need a local instance of the 'Dreamkeepers Data' repository on your computer.
The location of said repository must match database.dir in skirmish/config.yml.

Usage

Most users will want to use the 'skir ps' command to automate card production,
refer to the documentation for that package (github.com/sbrow/skirmish/skir).
*/
package ps
