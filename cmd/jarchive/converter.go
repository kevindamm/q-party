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
// github:kevindamm/q-party/cmd/jarchive/converter.go

package main

// The de-normed representation as found in some datasets, e.g. on Kaggle.
type JArchiveChallenge struct {
	Category string            `json:"category"`
	AirDate  `json:"air_date"` // YYYY-MM-DD

	Value    DollarValue  `json:"value"` // '$' (\d+)
	Question string       `json:"question"`
	Answer   string       `json:"answer"` // excluding "what is..." preface
	Round    EpisodeRound `json:"round"`
}

func ConvertGamesInDir(dir_path string) <-chan *JArchiveEpisode {
	channel := make(chan *JArchiveEpisode)

	go func() {
		// for file, _, _ in os.WalkDir(path)
		//   convert and write to channel
		channel <- ParseEpisode("...") // parse(filepath)
		// ...
		close(channel)
	}()

	return channel
}

func ParseEpisode(file_path string) *JArchiveEpisode {
	episode := new(JArchiveEpisode)

	return episode
}
