-- SQL statements for creating ?-Party database tables.
-- Copyright (c) 2025, Kevin Damm
-- All rights reserved.
-- MIT License:
--
-- Permission is hereby granted, free of charge, to any person obtaining a copy
-- of this software and associated documentation files (the "Software"), to deal
-- in the Software without restriction, including without limitation the rights
-- to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
-- copies of the Software, and to permit persons to whom the Software is
-- furnished to do so, subject to the following conditions:
--
-- The above copyright notice and this permission notice shall be included in
-- all copies or substantial portions of the Software.
--
-- THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
-- IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
-- FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
-- AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
-- LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
-- OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
-- SOFTWARE.
--
-- github:kevindamm/q-party/sql/create_tables.sql

-- Enum Table representing the correctness of a Q entry.
CREATE TABLE IF NOT EXISTS "DataQuality" (
    "dqID"     INTEGER
      PRIMARY KEY

  , "quality"  TEXT
      CHECK (quality IN ("Unknown"
                       , "Incorrect"
                       , "Needs Minor Change"
                       , "Disagreement"
                       , "Correct"
            ))
  , "summary"  TEXT
      NOT NULL
);

-- Questions/Qlues
CREATE TABLE IF NOT EXISTS "Qs" (
    "qID"           INTEGER
      PRIMARY KEY

  , "challenge"     TEXT
      NOT NULL        CHECK (challenge <> "")
  , "aired_date"    TEXT -- YYYY/MM/DD
);

CRAETE INDEX IF NOT EXISTS "Q__Aired"
  ON Qs (aired_date)
  ;
CREATE INDEX IF NOT EXISTS "Q__Updated"
  ON Qs (updated)
  WHERE (updated IS NOT NULL)
  ;
CREATE INDEX IF NOT EXISTS "Q__DataQuality"
  ON Qs (data_quality)
  ;

CREATE TABLE IF NOT EXISTS "As" (
    "qID"           INTEGER
      NOT NULL        CHECK (qID <> 0)
      REFERENCES      Qs (qID)
      ON DELETE       CASCADE
      ON UPDATE       RESTRICT

  , "answer"        TEXT
      NOT NULL        CHECK (answer <> "")
  , "data_quality"  INTEGER
      NOT NULL        DEFAULT 0
      REFERENCES      DataQuality (dqID)
      ON DELETE       RESTRICT
      ON UPDATE       RESTRICT
  , "updated"       TEXT -- YYYY/MM/DD
);

CREATE INDEX IF NOT EXISTS "A__Q"
  ON "As" (qID)

-- Category Names
CREATE TABLE IF NOT EXISTS "Categories" (
    "catID"     INTEGER
      PRIMARY KEY
  , "title"  TEXT
      NOT NULL    CHECK (cat_name <> "")

  , "notes"  TEXT
);

CREATE INDEX IF NOT EXISTS "Category__Title"
  ON Category (title)
  ;

-- Category that a Q can be found in (many-to-many relation).
CREATE TABLE IF NOT EXISTS "CategoryMembership" (
    "qID"        INTEGER
      NOT NULL
      REFERENCES   Qs (qID)
      ON DELETE    CASCADE
  , "catID"      INTEGER
      NOT NULL
      REFERENCES   Category (catID)
      ON DELETE    CASCADE

  PRIMARY KEY (qID, catID)
) WITHOUT ROWID;

CREATE INDEX IF NOT EXISTS "Category__Q"
  ON CategoryMembership (qID)
  ;
CREATE INDEX IF NOT EXISTS "Q__Category"
  ON CategoryMembership (catID)
  ;


CREATE TABLE IF NOT EXISTS "Seasons" (
    "seasonID"  INTEGER
      PRIMARY KEY
  , "title"     TEXT
      NOT NULL    CHECK (title <> "")

  , "air_date"  TEXT  -- YYYY/MM/DD, or NULL if this is a private season
  , "notes"     TEXT
);

CREATE TABLE IF NOT EXISTS "RoundType" (
    "roundID"     INTEGER
      PRIMARY KEY
  , "round"       TEXT
      NOT NULL      CHECK (round <> "")
  , "describe"    TEXT
);

CREATE TABLE IF NOT EXISTS "Contests" (
    "season"      INTEGER
      NOT NULL      CHECK (season > 0)
      REFERENCES    Seasons (seasonID)
  , "episode"     INTEGER
      NOT NULL      CHECK (episode > 0)
  , "round"       INTEGER
      NOT NULL
      REFERENCES    RoundType (roundID)

  PRIMARY KEY ("season", "episode", "round")
) WITHOUT ROWID;

ALTER TABLE Contestant
  ADD CONSTRAINT "contestant_fk_season"
  FOREIGN KEY        (seasonID)
  REFERENCES Seasons (seasonID)
  ;

ALTER TABLE Contestant
  ADD CONSTRAINT "contestant_fk_round"
  FOREIGN KEY         (round)
  REFERENCES Contests (round)
  ;

ALTER TABLE Contestant
  ADD CONSTRAINT "contestant_fk_contest"
  FOREIGN KEY
    (seasonID, episode, round)
  REFERENCES Contests
    (seasonID, episode, round)
  ;

CREATE TABLE IF NOT EXISTS "BoardPosition" (
    "qID"          INTEGER
      PRIMARY KEY
      REFERENCES     Qs (qID)

  , "season"     INTEGER
      NOT NULL
      REFERENCES   Seasons (seasonID)
  , "episode"    INTEGER
      NOT NULL
  , "round"      INTEGER
      NOT NULL
      REFERENCES   RoundType (roundID)

  , "across"     INTEGER
      NOT NULL     CHECK (across > 0 AND across <= 6)
  , "down"       INTEGER
      NOT NULL     CHECK (down > 0 AND down <= 5)

  , FOREIGN KEY
      ("season", "episode", "round")
    REFERENCES Contests
      ("season", "episode", "round")
);

CREATE TABLE IF NOT EXISTS "MediaClue" (
    "mediaID"    INTEGER
      PRIMARY KEY
  , "filetype"   TEXT
      NOT NULL     CHECK (filetype IN ("JPEG"
                                     , "PNG"
                                     , "SVG"
                                     , "MP3"
                                     , "MP4"
                                     , "MOV"
                         ))
  , "media_url"  TEXT  -- URL encoded
      NOT NULL     CHECK (url <> "")
  , "notes"      TEXT
);

CREATE TABLE IF NOT EXISTS "MediaLinks" (
    "qID"       INTEGER
      NOT NULL
      REFERENCES  BoardPosition (qID)
  , "mediaID"   INTEGER
      NOT NULL
      REFERENCES  MediaClue (mediaID)
  
  PRIMARY KEY ("qID", "mediaID")
) WITHOUT ROWID;

CREATE INDEX IF NOT EXISTS "Media__Q"
  ON MediaLinks (qID)
  ;
CREATE INDEX IF NOT EXISTS "Q__Media"
  ON MediaLinks (mediaID)
  ;
