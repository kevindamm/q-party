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
-- github:kevindamm/q-party/sql/create_1_accounts.sql

-------------------------------------------------------------------------------
-- UserAccounts and related properties
--
--    [--------------]
--    | UserAccounts |---[ Account__Username ]
--    [--------------]
--         |    |        [---------------] + fullname
--         |    '--------| User_Profiles | + occupation
--         |             [---------------] + residence
--         |
--         |         [-------------]
--         |---------| User_Banned |
--         |         [-------------]
--         |       [--------------]
--         |-------| User_Avatars |
--         |       [--------------]
--         |     [-------------]
--         '-----| User_Tokens |
--               [-------------]

-- Sometimes we know the contestant's name and background but they aren't
-- registered as an account here; sometimes there is an account here that has
-- only played games on the server and has never been a contestant.  I will be
-- very excited if the intersection of those two sets ever becomes non-null :)
CREATE TABLE IF NOT EXISTS "UserAccounts" (
    "accountID"    INTEGER
      PRIMARY KEY
      REFERENCES     UserAccounts (accountID)
  , "username"     TEXT
      NOT NULL       CHECK (username <> "")
      UNIQUE         ON CONFLICT ROLLBACK

  -- The email address for this account, only used for game invites and updates.
  -- (optional, may be NULL)
  , "email"        TEXT  -- valid email format 'user@example.com'

  -- If NULL, a verification code has not been emailed yet.
  , "email_sent"   TEXT  -- YYYY/MM/DD
      CHECK (email_sent <> "")

  -- If NULL, the email address has not yet been validated.
  , "email_valid"  TEXT  -- YYYY/MM/DD
      CHECK (email_valid <> "")
);

CREATE UNIQUE INDEX IF NOT EXISTS "User__Active"
  ON UserAccounts (username)
  WHERE (email_valid IS NOT NULL)
  ;

CREATE TABLE IF NOT EXISTS "User_Banned" (
    "accountID"    INTEGER
      PRIMARY KEY
      REFERENCES     UserAccounts (accountID)

  , "date_banned"  TEXT  -- YYYY/MM/DD or NULL if the user is not banned
      NOT NULL
  , "notes"        TEXT  -- useful for recording a reason for later reference
);

CREATE TABLE IF NOT EXISTS "User_Profiles" (
    "accountID"    INTEGER
      PRIMARY KEY
      REFERENCES     UserAccounts (accountID)

  , "fullname"     TEXT
      CHECK (fullname <> "")
  , "occupation"   TEXT
  , "residence"    TEXT

  , "notes"        TEXT
);

CREATE TABLE IF NOT EXISTS "User_Avatars" (
    "accountID"    INTEGER
      PRIMARY KEY
      REFERENCES     UserAccounts (accountID)

  , "obj_path"     TEXT
      NOT NULL
      CHECK (obj_path <> "")
  , "filetype"     TEXT
  , "width"        INTEGER
      CHECK (width IS NULL OR width > 0)
  , "height"       INTEGER
      CHECK (height IS NULL OR height > 0)
);

CREATE TABLE IF NOT EXISTS "User_Tokens" (
    "accountID"    INTEGER
      PRIMARY KEY
      REFERENCES     UserAccounts (accountID)
  , "token"        TEXT
      NOT NULL
  , "refresh"      TEXT
);
