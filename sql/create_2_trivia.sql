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
-- github:kevindamm/q-party/sql/create_2_trivia.sql

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

--
-- Qs and Answers and Categories
--
--   [----]
--   | Qs |--- +index[aired_date, updated, data_quality]
--   [----]
--      A 
--      |             [---------]
--      |             | Answers |
--      |             [---------]
--      |                A
--      |     UNIQUE     |
--      '----[ Q_A ]-----/
--

-- Qs contain just the challenge part of the challenge/answer pair.
-- Category inclusion is normed out to its own table, and the answers
-- are in a separate table because there may be more than one valid answer.
CREATE TABLE IF NOT EXISTS "Qs" (
    "qID"           INTEGER
      PRIMARY KEY

  , "challenge"     TEXT
      NOT NULL        CHECK (challenge <> "")
  
  , "aired_date"    TEXT -- YYYY/MM/DD
  --  (optional, may be NULL)
);

CREATE INDEX IF NOT EXISTS "Q__Aired"
  ON Qs (aired_date)
  WHERE (aired_date IS NOT NULL)
  ;

-- There may be many Answers for a Qs.qID, all equally acceptable.
-- No distinction is made between Answers with the same text answer,
-- and in the interest of de-duplicating the same answer given on
-- different occasions (such that random selection can be uniform,
-- knowing only how many associated rows there are, via Q_A).
CREATE TABLE IF NOT EXISTS "Answers" (
    "aID"           INTEGER
      PRIMARY KEY

  , "answer"        TEXT
      NOT NULL        CHECK (answer <> "")
  , "data_quality"  INTEGER
      NOT NULL        DEFAULT 0
      REFERENCES      DataQuality (dqID)
      ON DELETE       RESTRICT
      ON UPDATE       RESTRICT
  , "updated_date"  TEXT -- YYYY/MM/DD, NULL indicates unknown
      CHECK (updated_date <> "")
);

CREATE INDEX IF NOT EXISTS "Answer__DataQuality"
  ON Answers (data_quality)
  ;
CREATE INDEX IF NOT EXISTS "Answer__Updated"
  ON Answers (updated_date)
  WHERE (updated_date IS NOT NULL)
  ;

-- Many-to-many relation binding Qs and Answers, with an index on each
-- independently and unique index on the pair of foreign keys (the primary key).
CREATE TABLE IF NOT EXISTS "Q_Answer" (
    "qID"         INTEGER
      REFERENCES    Qs (qID)
  , "aID"         INTEGER
      REFERENCES    Answers (aID)

  , PRIMARY KEY ("qID", "aID")
) WITHOUT ROWID;

CREATE INDEX IF NOT EXISTS "Q__Answer"
  ON Q_Answer (aID)
  ;
CREATE INDEX IF NOT EXISTS "Answer__Q"
  ON Q_Answer (qID)
  ;


--
-- Category Titles
--
--   [----]                              [------------]
--   | Qs |         - - -(N..N)- - -     | Categories | +index[Category__Title]
--   [----]                              [------------]                    
--      A                                       A
--      |        [--------------------]         |
--      '--------| CategoryMembership |---------/
--               [--------------------]
--     

-- Categories are useful for grouping related Qs.
CREATE TABLE IF NOT EXISTS "Categories" (
    "catID"        INTEGER
      PRIMARY KEY
  , "title"        TEXT
      NOT NULL       CHECK (title <> "")

  , "notes"        TEXT
);

-- Restrict duplicate category names and facilitate ordered lookup.
CREATE UNIQUE INDEX IF NOT EXISTS "Category__Title"
  ON Categories (title)
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

  , PRIMARY KEY (qID, catID)
) WITHOUT ROWID;

CREATE INDEX IF NOT EXISTS "Category__Q"
  ON CategoryMembership (qID)
  ;
CREATE INDEX IF NOT EXISTS "Q__Category"
  ON CategoryMembership (catID)
  ;

--
-- Matches and MatchRounds
--
--   [---------]
--   | Matches |
--   [---------]
-- 

-- Enum table for consistent tracking of the mechanics of different round types
-- (single, double, final, tiebreaker) having different questions and Q values.
CREATE TABLE IF NOT EXISTS "RoundEnum" (
    "round"       INTEGER
      PRIMARY KEY

  , "title"       TEXT
      NOT NULL      CHECK (title <> "")
  , "describe"    TEXT
  -- (optional, may be NULL)
);

-- TODO populate RoundEnum values

CREATE TABLE IF NOT EXISTS "DifficultyEnum" (
    "difficulty"   INTEGER
      PRIMARY KEY

  , "title"        TEXT
      NOT NULL       CHECK (title <> "")
  , "describe"     TEXT
  -- (optional, may be NULL)
);

-- TODO populate DifficultyEnum values


CREATE TABLE IF NOT EXISTS "Matches" (
    "matchID"     INTEGER
      PRIMARY KEY

  , "season"      TEXT
  -- (optional, may be NULL)
  , "episode"     INTEGER
  -- (optional, may be NULL, unique per season)

  , "jeid"        INTEGER  -- enumerated episode ID
  -- (optional, may be NULL, must be UNIQUE)
  , "jaid"        INTEGER  -- jarchive unique ID
  -- (optional, may be NULL, must be UNIQUE)
);

CREATE UNIQUE INDEX IF NOT EXISTS "Matches__Season"
  ON Matches ("season")
  WHERE (season IS NOT NULL)
  ;
CREATE UNIQUE INDEX IF NOT EXISTS "Matches__SeasonEpisode"
  ON Matches ("season", "episode")
  WHERE (season IS NOT NULL AND episode IS NOT NULL)
  ;
CREATE UNIQUE INDEX IF NOT EXISTS "Matches__JEID"
  ON Matches ("jeid")
  WHERE (jeid IS NOT NULL)
  ;
CREATE UNIQUE INDEX IF NOT EXISTS "Matches__JAID"
  ON Matches ("jaid")
  WHERE (jaid IS NOT NULL)
  ;

-- The segment of a single board (a collection of one or more Qs)
-- where Contestants are consistent (note: contestants may change
-- within a single episode, though that has been exceptionally rare).
CREATE TABLE IF NOT EXISTS "MatchRounds" (
    "matchID"     INTEGER
      NOT NULL      CHECK (matchID > 0)
  , "round"       INTEGER
      NOT NULL      DEFAULT 0
      REFERENCES    RoundEnum (round)

  , "difficulty"  INTEGER
      NOT NULL      DEFAULT 0
      REFERENCES    DifficultyEnum (difficulty)

  , PRIMARY KEY ("matchID", "round")
) WITHOUT ROWID;

CREATE INDEX IF NOT EXISTS "MatchRound__Difficulty"
  ON MatchRounds (difficulty)
  WHERE (difficulty <> 0)
  ;

CREATE TABLE IF NOT EXISTS "MatchRound_Contestant" (
    "matchID"       INTEGER
      NOT NULL        CHECK (matchID > 0)
  , "round"         INTEGER
      NOT NULL        DEFAULT 0
      REFERENCES      RoundEnum (round)
      ON DELETE       RESTRICT
  , "contestant"    INTEGER
      NOT NULL        CHECK (contestant <> 0)
      REFERENCES      UserAccounts (accountID)
      ON DELETE       CASCADE

  , "is_returning"  BOOLEAN
      NOT NULL        DEFAULT FALSE

  , FOREIGN KEY            ("matchID", "round")
    REFERENCES MatchRounds ("matchID", "round")

  , PRIMARY KEY ("matchID", "round", "contestant")
) WITHOUT ROWID;

CREATE TABLE IF NOT EXISTS "MatchRound_Position" (
    "matchID"      INTEGER
      NOT NULL
      REFERENCES     Matches (matchID)
  , "round"        INTEGER
      NOT NULL
      REFERENCES     RoundEnum (round)

  , "across"       INTEGER
      NOT NULL       CHECK (across > 0 AND across <= 6)
  , "down"         INTEGER
      NOT NULL       CHECK (down > 0 AND down <= 5)
  , "qID"          INTEGER
      NOT NULL
      REFERENCES     Qs (qID)
      ON DELETE      CASCADE
  , "special"      BOOLEAN
      NOT NULL       DEFAULT FALSE

  , FOREIGN KEY            ("matchID", "round")
    REFERENCES MatchRounds ("matchID", "round")
  , PRIMARY KEY            ("matchID", "round")
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
      NOT NULL     CHECK (media_url <> "")
  , "notes"      TEXT
);

CREATE TABLE IF NOT EXISTS "MediaLinks" (
    "qID"       INTEGER
      NOT NULL
      REFERENCES  BoardPosition (qID)
      ON DELETE   RESTRICT
  , "mediaID"   INTEGER
      NOT NULL
      REFERENCES  MediaClue (mediaID)
      ON DELETE   CASCADE
  
  , PRIMARY KEY ("qID", "mediaID")
) WITHOUT ROWID;

CREATE INDEX IF NOT EXISTS "Media__Q"
  ON MediaLinks (qID)
  ;
CREATE INDEX IF NOT EXISTS "Q__Media"
  ON MediaLinks (mediaID)
  ;
