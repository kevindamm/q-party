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
-- Define enum tables (constants).
--

INSERT INTO DataQuality
    ("dqID", "quality",                 "summary")
  VALUES (0, "Needs Review",       "Some votes are needed to determine the accuracy of this answer.")
       , (1, "Entirely Incorrect", "The answer is incorrect and was never correct.")
       , (2, "Recently Incorrect", "This answer was correct when it was made but ")
       , (3, "Suspected Outdated", "The challenge was written a long time ago or the field has seen recent changes, but accuracy has not been checked.")
       , (4, "Needs Minor Change", "The answer needs clarification, it is too broad to be compared fairly to an accurate or inaccurate answer.")
       , (5, "Disagreement",       "Some votes claim correctness while other votes claim incorrectness.")
       , (6, "Correct",            "This answer is acceptable and correct.")
       , (7, "Confirmed Correct",  "Multiple 'correct' votes and no supported 'incorrect' votes.")
       ;

INSERT INTO RoundEnum
    ("round", "title",       "notes")
  VALUES ( 0, "UNKNOWN",     NULL)
       , ( 1, "Single",      "The first of three rounds")
       , ( 2, "Double",      "Second round, double values")
       , ( 3, "Final",       "Third and final round, single question with bidding")
       , ( 4, "Tie-Breaker", "To resolve any ties at the end of the Final round (format is same as final)")
       ;
                  
INSERT INTO DifficultyEnum
    ("difficulty", "title",                   "notes")
  VALUES (      0, "UNKNOWN",                 NULL)
       , (      1, "Teen Tournament",         "")
       , (      2, "Celebrity Match",         "")
       , (      3, "College Championship",    "")
       , (      4, "Seniors Tournament",      "")
       , (      5, "Standard Competition",    "")
       , (      6, "Tournament of Champions", "")
       , (      7, "Masters Tournament",      "")
       , (      8, "Greatest of All Time",    "")
       ;

--
-- Define unknown records (sentinels).
--

INSERT INTO Qs
     ("qID", "challenge")
  VALUES (0, "UNKNOWN")
  ;

INSERT INTO Categories
    ("catID", "title")
  VALUES ( 0, "UNKNOWN")
  ;

INSERT INTO Matches
    ("matchID",  "season", "jeid", "jaid")
  VALUES (   0, "UNKNOWN",      0,      0)
  ;
