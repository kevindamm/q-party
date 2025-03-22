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
-- github:kevindamm/q-party/sql/create_3_matches.sql

-------------------------------------------------------------------------------
-- Matches and MatchRounds
--
--   [---------] + (season, episode)
--   | Matches | + (jeid) episode ID
--   [---------] + (jaid) jarchive ID
--         A
--         |    [-------------]
--         '----| MatchRounds |
--              [-------------]
--                  A  A
--                  |  |    [----------------------]
--                  '--+----| MatchRound_Positions |
--                     |    [----------------------]
--                     |       [------------------------]
--                     '-------| MatchRound_Contestants |
--                             [------------------------]

-- Enum table for consistent tracking of the mechanics of different round types
-- (single, double, final, tiebreaker) having different questions and Q values.
CREATE TABLE IF NOT EXISTS "RoundEnum" (
    "round"       INTEGER
      PRIMARY KEY

  , "title"       TEXT
      NOT NULL      CHECK (title <> "")
  , "notes"       TEXT
  -- (optional, may be NULL)
);

CREATE TABLE IF NOT EXISTS "MatchDifficultyEnum" (
    "match_difficulty"   INTEGER
      PRIMARY KEY

  , "title"        TEXT
      NOT NULL       CHECK (title <> "")
  , "notes"        TEXT
  -- (optional, may be NULL)
);

-- A match is a series of rounds with each player seeing the same questions.
-- Matches are traditionally played synchronously but may be played async.
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

CREATE INDEX IF NOT EXISTS "Match__Season"
  ON Matches ("season")
  WHERE (season IS NOT NULL)
  ;
CREATE UNIQUE INDEX IF NOT EXISTS "Match__SeasonEpisode"
  ON Matches ("season", "episode")
  WHERE (season IS NOT NULL AND episode IS NOT NULL)
  ;
CREATE UNIQUE INDEX IF NOT EXISTS "Match__JEID"
  ON Matches ("jeid")
  WHERE (jeid IS NOT NULL)
  ;
CREATE UNIQUE INDEX IF NOT EXISTS "Match__JAID"
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
      REFERENCES    MatchDifficultyEnum (match_difficulty)

  , PRIMARY KEY ("matchID", "round")
) WITHOUT ROWID;

CREATE INDEX IF NOT EXISTS "MatchRound__Difficulty"
  ON MatchRounds (difficulty)
  WHERE (difficulty <> 0)
  ;

-- There are usually three contestants per (match, round), but there may be any.
CREATE TABLE IF NOT EXISTS "MatchRound_Contestants" (
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

-- There are as many positions as the round type can have (usually max 30 or 1).
CREATE TABLE IF NOT EXISTS "MatchRound_Positions" (
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
) WITHOUT ROWID;
