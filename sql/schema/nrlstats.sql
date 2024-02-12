


CREATE TABLE CompYear (
    year_ INT,
    PRIMARY KEY (year_)
);


CREATE TABLE Round (
    round INT,
    year_ INT,
    PRIMARY KEY (round, year_),
    FOREIGN KEY (year_) REFERENCES CompYear(year_)
);

CREATE TABLE Player (
    ID SERIAL,
    name VARCHAR(100),
    PRIMARY KEY (ID)
);

CREATE TABLE PlayerPosition (
    playerID INT,
    year_ INT,
    position_ VARCHAR(100),
    PRIMARY KEY (playerID, year_),
    FOREIGN KEY (playerID) REFERENCES Player(ID),
    FOREIGN KEY (year_) REFERENCES CompYear(year_)
);

CREATE TABLE Team (
    ID Serial,
    name VARCHAR,
    PRIMARY KEY (ID)
);

CREATE TABLE PlayerScore (
    round INT,
    year_ INT,
    playerID INT,
    teamID INT,
    opponentID INT,
    price INT,
    score INT,
    minutes INT,
    weather VARCHAR(100),
    jersey INT,
    base INT,
    PRIMARY KEY (round, year_, playerID),
    FOREIGN KEY (round, year_) REFERENCES Round(round, year_),
    FOREIGN KEY (playerID) REFERENCES Player(ID),
    FOREIGN KEY (teamID) REFERENCES Team(ID),
    FOREIGN KEY (opponentID) REFERENCES Team(ID)
);

CREATE TABLE MatchOdds (
    round INT,
    year_ INT,
    teamID INT,
    odds FLOAT,
    FOREIGN KEY (round, year_) REFERENCES Round(round, year_),
    FOREIGN KEY (teamID) REFERENCES Team(ID)
    );