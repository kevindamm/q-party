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
-- github:kevindamm/q-party/sql/drop_indices.sql

DROP INDEX IF EXISTS "MatchRound__Difficulty";

DROP INDEX IF EXISTS "Match__JAID";
DROP INDEX IF EXISTS "Match__JEID";
DROP INDEX IF EXISTS "Match__SeasonEpisode";
DROP INDEX IF EXISTS "Match__Season";

DROP INDEX IF EXISTS "Category__Q";
DROP INDEX IF EXISTS "Q__Category";
DROP INDEX IF EXISTS "Category__Title";

DROP INDEX IF EXISTS "Q__Media";
DROP INDEX IF EXISTS "Media__Q";

DROP INDEX IF EXISTS "Q__Answer";
DROP INDEX IF EXISTS "Answer__Q";

DROP INDEX IF EXISTS "Answer__DataQuality";
DROP INDEX IF EXISTS "Answer__Updated";
DROP INDEX IF EXISTS "Q__Aired";
DROP INDEX IF EXISTS "User__Active";
