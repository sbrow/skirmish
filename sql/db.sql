CREATE SCHEMA IF NOT EXISTS Skirmish;

DROP TABLE IF EXISTS Skirmish.NonDeckCards;
DROP TABLE IF EXISTS Skirmish.Factions;
DROP TABLE IF EXISTS Skirmish.keywords;

CREATE TABLE Skirmish.Factions(
	Name varchar(11) NOT NULL PRIMARY KEY UNIQUE
);

INSERT IGNORE INTO Skirmish.Factions(Name) VALUES
	('Troika')
	,('Nightmares')
	,('Neutral');

CREATE TABLE Skirmish.NonDeckCards(
	Name varchar(11) PRIMARY KEY
    ,Faction varchar(11)
    ,Banner char(6) DEFAULT '000000'
    ,Indicator char(6) DEFAULT '808080'
    ,ShortText varchar(255)
	,FOREIGN KEY (Faction) REFERENCES Skirmish.Factions(Name)
);

REPLACE INTO Skirmish.NonDeckCards(Name, Faction, ShortText) VALUES
('Bast', (SELECT Name from Skirmish.Factions Where Name='Troika'), '')
,('Igrath', (SELECT Name from Skirmish.Factions Where Name='Troika'), '')
,('Lilith', (SELECT Name from Skirmish.Factions Where Name='Troika'), '')
,('Scinter', (SELECT Name from Skirmish.Factions Where Name='Troika'), '')
,('Vi', (SELECT Name from Skirmish.Factions Where Name='Troika'), '')
,('Ravat', (SELECT Name from Skirmish.Factions Where Name='Nightmares'), 'Speed damage')
,('Scuttler', (SELECT Name from Skirmish.Factions Where Name='Nightmares'), '')
,('Tendril', (SELECT Name from Skirmish.Factions Where Name='Nightmares'), '')
,('Tinsel', (SELECT Name from Skirmish.Factions Where Name='Nightmares'), '')
,('Wisp', (SELECT Name from Skirmish.Factions Where Name='Nightmares'), '');

CREATE TABLE Skirmish.Keywords(
	Word varchar(30) NOT NULL PRIMARY KEY UNIQUE 
	,Description varchar(255)
    ,Regex varchar(30)
	,FULLTEXT (Regex)
);

REPLACE INTO Skirmish.Keywords(Word, Description, Regex) VALUES
('Speed', 'Spend speed on things', 'Speed')
,('Damage', 'Damage hurts.', 'Damage(s|ed)');

CREATE TABLE IF NOT EXISTS Skirmish.Bold(
	Regex varchar(255) NOT NULL PRIMARY KEY UNIQUE
	,CaseSensitive boolean NOT NULL DEFAULT TRUE
	,Plural varchar(2) DEFAULT ''
	,PastTense varchar(3) DEFAULT ''
    ,FULLTEXT (Regex)
);

REPLACE INTO Skirmish.Bold(Regex)
SELECT Name FROM Skirmish.NonDeckCards;

DROP FUNCTION IF EXISTS Skirmish.BoldWords;
DELIMITER //
CREATE FUNCTION Skirmish.BoldWords(ShortText varchar(255)) RETURNS varchar(255)
BEGIN
	RETURN (SELECT GROUP_CONCAT(Regex SEPARATOR ',')
    FROM (
    SELECT Regex FROM Skirmish.Keywords AS t0
    WHERE Match(Regex) AGAINST (ShortText)
    UNION
    SELECT Regex FROM Skirmish.Bold AS t1
    WHERE MATCH(Regex) AGAINST (ShortText)) as t);
END;//

CREATE OR REPLACE VIEW Skirmish.Leaders AS 
SELECT *, BoldWords(ShortText) as BoldWords
FROM Skirmish.nondeckcards 
WHERE NOT Faction='Neutral';//

CREATE OR REPLACE VIEW Skirmish.Guests AS
SELECT *, BoldWords(ShortText) as BoldWords
FROM Skirmish.NonDeckCards
WHERE Faction='Neutral';//