-- SQL statements for creating accounts and contestant history for ?-Party.
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
-- github:kevindamm/q-party/sql/create_4_basedata.sql

--
-- ENUM TABLES
--

INSERT INTO DataQuality
    ("dqID", "quality",            "summary")
  VALUES
         (0, "Needs Review",       "Some votes are needed to determine the accuracy of this answer.")
       , (1, "Entirely Incorrect", "The answer is incorrect and was never correct.")
       , (2, "Recently Incorrect", "This answer was correct when it was made but ")
       , (3, "Suspected Outdated", "The challenge was written a long time ago or the field has seen recent changes, but accuracy has not been checked.")
       , (4, "Needs Minor Change", "The answer needs clarification, it is too broad to be compared fairly to an accurate or inaccurate answer.")
       , (5, "Disagreement",       "Some votes claim correctness while other votes claim incorrectness.")
       , (6, "Correct",            "This answer is acceptable and correct.")
       , (7, "Confirmed Correct",  "Multiple 'correct' votes and no supported 'incorrect' votes.")
       ;

-- The tie-breakers are extremely rare, perhaps only a handful on record.
-- We could use other final questions instead.
INSERT INTO RoundEnum
    ("round", "title",       "notes")
  VALUES
         ( 0, "UNKNOWN",     NULL)
       , ( 1, "Single",      "The first of three rounds")
       , ( 2, "Double",      "Second round, double values")
       , ( 3, "Final",       "Third and final round, single question with bidding")
       , ( 4, "Tie-Breaker", "To resolve any ties at the end of the Final round (format is same as final)")
       ;

-- These difficulty values are approximately ordered but there is considerable overlap.
-- The ordering and variance could be estimated more specifically from matches.
INSERT INTO MatchDifficultyEnum
    ("match_difficulty", "title",                         "season", "notes")
  VALUES
         (            0, "UNKNOWN",                       "unk",           NULL)
       , (            1, "Trebek Pilots",                 "pilot",         "")
       , (            2, "Teen Tournament",               "teen",          "younger players")
       , (            3, "Celebrity Match",               "celeb",         "")
       , (            4, "National College Championship", "ncc",           "")
       , (            5, "Seniors Tournament",            "seniors",       "")
       , (            6, "Standard Competition",          "s%02d",         "assume approximately equal difficulty")
       , (            7, "Battle of the Bay Area Brains", "bbab",          "")
       , (            8, "Tournament of Champions",       "champ",         "returning champions")
       , (            9, "Masters Tournament",            "masters",       "")
       , (           10, "Super-Jeopardy",                "super",         "")
       , (           11, "Watson vs Humans",              "watson",        "")
       , (           12, "Greatest of All Time",          "goat",          "")
       ;

-- These values were calculated from aggregate correct-response measurements
-- for challenges at each value level.  There is a slight difference between
-- single & double, they've been combined here because the difference is small.
INSERT INTO ChallengeDifficultyEnum
    ( "difficulty", "base_value", "success_rate")
  VALUES
         (       0,            0,      "UNKNOWN")
       , (       1,          100,          "70%")
       , (       2,          200,          "60%")
       , (       3,          300,          "50%")
       , (       4,          400,          "41%")
       , (       5,          500,          "34%")
       ;

--
-- UNKNOWNS
--

INSERT INTO Qs
     ("qID", "challenge") VALUES (0, "UNKNOWN");

INSERT INTO Categories
    ("catID", "title") VALUES ( 0, "UNKNOWN");

INSERT INTO Matches
    ("matchID",  "season", "jeid", "jaid")
  VALUES (   0, "UNKNOWN",      0,      0)
  ;
