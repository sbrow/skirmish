CREATE SCHEMA IF NOT EXISTS Skirmish;

CREATE TABLE IF NOT EXISTS Skirmish.Factions(
	Name varchar(11) NOT NULL PRIMARY KEY UNIQUE
);

INSERT IGNORE INTO Skirmish.Factions(Name) VALUES
	('Troika')
	,('Nightmares')
	,('Neutral');

CREATE TABLE IF NOT EXISTS Skirmish.NonDeckCards(
	Name varchar(11) PRIMARY KEY
    ,Faction varchar(11)
    ,Banner char(6) DEFAULT '000000'
    ,Indicator char(6) DEFAULT '808080'
	,FOREIGN KEY (Faction) REFERENCES Skirmish.Factions(Name)
);

REPLACE INTO Skirmish.NonDeckCards(Name, Faction) VALUES
('Bast', (SELECT Name from Skirmish.Factions Where Name='Troika'))
,('Igrath', (SELECT Name from Skirmish.Factions Where Name='Troika'))
,('Lilith', (SELECT Name from Skirmish.Factions Where Name='Troika'))
,('Scinter', (SELECT Name from Skirmish.Factions Where Name='Troika'))
,('Vi', (SELECT Name from Skirmish.Factions Where Name='Troika'))
,('Ravat', (SELECT Name from Skirmish.Factions Where Name='Nightmares'))
,('Scuttler', (SELECT Name from Skirmish.Factions Where Name='Nightmares'))
,('Tendril', (SELECT Name from Skirmish.Factions Where Name='Nightmares'))
,('Tinsel', (SELECT Name from Skirmish.Factions Where Name='Nightmares'))
,('Wisp', (SELECT Name from Skirmish.Factions Where Name='Nightmares'));

CREATE TABLE IF NOT EXISTS Skirmish.Keywords(
	Name varchar(30) NOT NULL PRIMARY KEY UNIQUE
	,Description varchar(255)
);

CREATE TABLE IF NOT EXISTS Skirmish.Bold(
	Regex varchar(255) NOT NULL PRIMARY KEY UNIQUE
	,CaseSensitive boolean NOT NULL DEFAULT TRUE
	,Plural varchar(2) DEFAULT ''
	,PastTense varchar(3) DEFAULT ''
);

REPLACE INTO Skirmish.Bold(Regex)
SELECT Name FROM Skirmish.NonDeckCards