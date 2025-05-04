// Copyright (c) 2025 Kevin Damm
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
// github:kevindamm/q-party/schema/show_date_test.go

package schema_test

import (
	_ "embed"
	"testing"

	"github.com/kevindamm/q-party/schema"
)

func TestDateCompare(t *testing.T) {
	tests := []struct {
		name  string
		this  schema.ShowDate
		other schema.ShowDate
		want  int
	}{
		{"less-than",
			schema.ShowDate{Year: 2000, Month: 9, Day: 1},
			schema.ShowDate{Year: 2001, Month: 9, Day: 1},
			-1},
		{"greater-than",
			schema.ShowDate{Year: 2002, Month: 9, Day: 2},
			schema.ShowDate{Year: 2001, Month: 9, Day: 2},
			1},
		{"less-than",
			schema.ShowDate{Year: 1999, Month: 7, Day: 2},
			schema.ShowDate{Year: 1999, Month: 9, Day: 1},
			-1},
		{"greater-than",
			schema.ShowDate{Year: 2025, Month: 10, Day: 2},
			schema.ShowDate{Year: 2025, Month: 9, Day: 3},
			1},
		{"less-than",
			schema.ShowDate{Year: 2001, Month: 9, Day: 1},
			schema.ShowDate{Year: 2001, Month: 9, Day: 2},
			-1},
		{"greater-than",
			schema.ShowDate{Year: 2001, Month: 9, Day: 3},
			schema.ShowDate{Year: 2001, Month: 9, Day: 2},
			1},
		{"equal",
			schema.ShowDate{Year: 1999, Month: 12, Day: 19},
			schema.ShowDate{Year: 1999, Month: 12, Day: 19},
			0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.this.Compare(&tt.other); got != tt.want {
				t.Errorf("ShowDate.Compare() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDateRangeContains(t *testing.T) {
	scope := schema.ShowDateRange{
		From:  &schema.ShowDate{Year: 1995, Month: 3, Day: 25},
		Until: &schema.ShowDate{Year: 1997, Month: 6, Day: 2},
	}
	tests := []struct {
		name  string
		scope schema.ShowDateRange
		date  schema.ShowDate
		want  bool
	}{
		{"before", scope,
			schema.ShowDate{Year: 1979, Month: 1, Day: 13},
			false},
		{"after", scope,
			schema.ShowDate{Year: 2000, Month: 1, Day: 1},
			false},
		{"within", scope,
			schema.ShowDate{Year: 1996, Month: 10, Day: 31},
			true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.scope.Contains(tt.date); got != tt.want {
				t.Errorf("ShowDateRange.Contains() = %v, want %v", got, tt.want)
			}
		})
	}
}
