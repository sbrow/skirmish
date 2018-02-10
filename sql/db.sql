IF OBJECT_ID('Factions', 'U') IS NULL
BEGIN
CREATE TABLE Factions(
	Name varchar(255)
	PRIMARY KEY (Name)
);
INSERT INTO Factions VALUES
('Troika'), ('Nightmares');
END
IF OBJECT_ID('Leaders', 'U') IS NULL
BEGIN
CREATE TABLE Leaders(
	Name varchar(255),
	Faction varchar(255)
	PRIMARY KEY (Name)
	FOREIGN KEY (Faction) REFERENCES Factions(Name)
);
INSERT INTO Factions VALUES
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
END