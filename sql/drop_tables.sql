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
-- github:kevindamm/q-party/sql/drop_tables.sql

-- matches
DROP TABLE IF EXISTS "MatchRound_Positions";
DROP TABLE IF EXISTS "MatchRound_Contestants";
DROP TABLE IF EXISTS "MatchRounds";
DROP TABLE IF EXISTS "Matches";
DROP TABLE IF EXISTS "MatchDifficultyEnum";
DROP TABLE IF EXISTS "RoundEnum";

-- trivia
DROP TABLE IF EXISTS "Category_Qs";
DROP TABLE IF EXISTS "Categories";
DROP TABLE IF EXISTS "Q_Media";
DROP TABLE IF EXISTS "MediaClue";
DROP TABLE IF EXISTS "Q_Answer";
DROP TABLE IF EXISTS "Answers";
DROP TABLE IF EXISTS "ChallengeDifficultyEnum";
DROP TABLE IF EXISTS "Qs";
DROP TABLE IF EXISTS "DataQuality";

-- accounts
DROP TABLE IF EXISTS "User_Tokens";
DROP TABLE IF EXISTS "User_Avatars";
DROP TABLE IF EXISTS "User_Profiles";
DROP TABLE IF EXISTS "User_Banned";
DROP TABLE IF EXISTS "UserAccounts";
