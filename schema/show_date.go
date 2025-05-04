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
// github:kevindamm/q-party/schema/show_date.go

package schema

import (
	"fmt"
	"regexp"
	"strconv"
)

type ShowDate struct {
	Year  int `json:"year"`
	Month int `json:"month"`
	Day   int `json:"day"`
}

// formatted as YYYY-MM-DD
func ParseShowDate(image string) *ShowDate {
	reShowDate := regexp.MustCompile(`(\d{4})-(\d{2})-(\d{2})`)
	matches := reShowDate.FindStringSubmatch(image)
	// We can ignore the error here because we know these are all digit patterns.
	year, _ := strconv.Atoi(matches[1])
	month, _ := strconv.Atoi(matches[2])
	day, _ := strconv.Atoi(matches[3])

	return &ShowDate{
		Year:  year,
		Month: month,
		Day:   day,
	}
}

func (sd ShowDate) String() string {
	if sd.Year+sd.Month+sd.Day == 0 {
		return ""
	}
	return fmt.Sprintf("%04d/%02d/%02d", sd.Year, sd.Month, sd.Day)
}

// Returns 0 if `this` and `other` are equal;
// +1 if this is later than other, and -1 if before.
func (this ShowDate) Compare(other *ShowDate) int {
	if this.Year < other.Year {
		return -1
	}
	if this.Year > other.Year {
		return +1
	}
	// (this.Year == other.Year)
	if this.Month < other.Month {
		return -1
	}
	if this.Month > other.Month {
		return +1
	}
	// (this.Month == other.Month)
	if this.Day < other.Day {
		return -1
	}
	if this.Day > other.Day {
		return +1
	}
	// The dates are equal.
	return 0
}

type ShowDateRange struct {
	From  *ShowDate `json:"from,omitempty"`
	Until *ShowDate `json:"until,omitempty"`
}

func (scope ShowDateRange) Contains(date ShowDate) bool {
	return (                         // including endpoints,
	date.Compare(scope.From) >= 0 && // after beginning and
		date.Compare(scope.Until) <= 0) // before ending
}
