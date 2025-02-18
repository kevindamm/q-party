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
// github:kevindamm/q-party/cmd/jarchive/write_json.go

package json

import (
	"encoding/json"
	"fmt"
	"os"
	"path"

	qparty "github.com/kevindamm/q-party"
	"github.com/kevindamm/q-party/service"
)

func WriteSeasonIndex(data_path string) func(*service.JArchiveIndex) error {
	return func(jarchive *service.JArchiveIndex) error {
		filepath := path.Join(data_path, "jarchive.json")
		writer, err := os.Create(filepath)
		if err != nil {
			return fmt.Errorf("failed to create file %s\n%s", filepath, err)
		}
		defer writer.Close()

		bytes, err := json.MarshalIndent(jarchive, "", "  ")
		if err != nil {
			return fmt.Errorf("failed to marshal index into json\n%s", err)
		}

		_, err = writer.Write(bytes)
		if err != nil {
			return fmt.Errorf("failed to write bytes\n%s", err)
		}
		return nil
	}
}

func WriteEpisode(json_path string) func(qparty.FullEpisode) error {
	return func(episode qparty.FullEpisode) error {
		filepath := path.Join(json_path, episode.Show.JSON())
		writer, err := os.Create(filepath)
		if err != nil {
			return fmt.Errorf("failed to create file %s\n%s", filepath, err)
		}
		defer writer.Close()

		bytes, err := json.MarshalIndent(episode, "", "  ")
		if err != nil {
			return fmt.Errorf("failed to marshal index into json\n%s", err)
		}

		_, err = writer.Write(bytes)
		if err != nil {
			return fmt.Errorf("failed to write bytes\n%s", err)
		}
		return nil
	}
}

func WriteEpisodeMetadata(json_path string, episode qparty.EpisodeMetadata) error {
	filepath := path.Join(json_path, episode.Show.JSON())
	writer, err := os.Create(filepath)
	if err != nil {
		return err
	}
	bytes, err := json.MarshalIndent(episode, "", "  ")
	if err != nil {
		return err
	}

	_, err = writer.Write(bytes)
	return err
}
