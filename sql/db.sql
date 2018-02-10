CREATE TABLE IF NOT EXISTS Checksums(
	Key varchar(25) NOT NULL PRIMARY KEY,
	Value char(64)
);
CREATE TABLE IF NOT EXISTS Factions(
	Name varchar(11) NOT NULL PRIMARY KEY
);
INSERT INTO Factions(Name) VALUES ON DUPLICATE KEY UPDATE
	('Troika'),
	('Nightmares');
CREATE TABLE IF NOT EXISTS Leaders(
	Name varchar(11) PRIMARY KEY,
	Faction varchar(11) FOREIGN KEY REFERENCES Factions(Name)
);
INSERT INTO Leader(Name, Faction) VALUES ON DUPLICATE KEY UPDATE
('Bast', 'Troika'),
('Igrath', 'Troika'),
('Lilith', 'Troika'),
('Scinter', 'Troika'),
('Vi', 'Troika'),
('Ravat', 'Nightmares'),
('Scuttler', 'Nightmares'),
('Tendril', 'Nightmares'),
('Tinsel', 'Nightmares'),
('Wisp', 'Nightmares');