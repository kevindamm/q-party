// Copyright (c) 2024 Kevin Damm
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.
//
// github:kevindamm/q-party/db/db.go

package db

import (
	"sync"

	qparty "github.com/kevindamm/q-party"
	"github.com/kevindamm/q-party/ent"
)

type DatabaseConnection interface {
	// Get all season metadata.
	Seasons() []qparty.SeasonMetadata

	// Get all episode metadata for a season.
	SeasonEpisodes(season_id qparty.SeasonID) []qparty.EpisodeMetadata

	// Get specifics of a single episode.
	EpisodeStats(show qparty.ShowNumber) qparty.EpisodeStats

	// Initialize the database structure.
	Create() error

	HasError() bool
	Error() string
}

func JArchive() DatabaseConnection {
	client, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		return db_error{""}
	}

	return &db{client: client}
}

/*
 * DB (sum type of DatabaseConnection, indicates success)
 */

type db struct {
	client *ent.Client
	mutex  sync.Mutex
}

// Always evaluates to SUCCESS outcome when Open() results in type [db].
func (*db) HasError() bool { return false }
func (*db) Error() string  { return "" }

func (db *db) Seasons() []qparty.SeasonMetadata {
	seasons := make([]qparty.SeasonMetadata, 0)
	// TODO

	return seasons
}

func (db *db) SeasonEpisodes(season_id qparty.SeasonID) []qparty.EpisodeMetadata {
	episodes := make([]qparty.EpisodeMetadata, 0)
	db.mutex.Lock()
	defer db.mutex.Unlock()
	// TODO

	return episodes
}

func (db *db) EpisodeStats(show qparty.ShowNumber) qparty.EpisodeStats {
	episode_stats := qparty.EpisodeStats{}
	db.mutex.Lock()
	defer db.mutex.Unlock()
	// TODO

	return episode_stats
}

func (db *db) Create() error {
	db.mutex.Lock()
	defer db.mutex.Unlock()
	// TODO

	return nil
}

/*
 * DB Error (sum type of DatabaseConnection, indicates error in opening)
 */

type db_error struct {
	string
}

func (db_error) HasError() bool       { return true }
func (db_err db_error) Error() string { return db_err.string }

func (db_error) Seasons() []qparty.SeasonMetadata                        { return nil }
func (db_error) SeasonEpisodes(qparty.SeasonID) []qparty.EpisodeMetadata { return nil }
func (db_error) EpisodeStats(qparty.ShowNumber) qparty.EpisodeStats      { return qparty.EpisodeStats{} }
func (db_error) Create() error                                           { return nil }
