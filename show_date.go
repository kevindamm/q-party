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
// github:kevindamm/q-party/show_date.go

package qparty

import (
	"fmt"
	"strconv"
	"time"
)

type ShowDateRange struct {
	From  ShowDate `json:"from"`
	Until ShowDate `json:"until,omitempty"`
}

type ShowDate struct {
	Year  int `json:"year"`
	Month int `json:"month"`
	Day   int `json:"day"`
}

func (date ShowDate) IsZero() bool {
	return date.Year == 0 && date.Month == 0 && date.Day == 0
}

func (date ShowDate) Equal(other ShowDate) bool {
	return other.Year == date.Year &&
		other.Month == date.Month &&
		other.Day == date.Day
}

func (date ShowDate) ToTime() time.Time {
	return time.Date(
		date.Year, time.Month(date.Month), date.Day,
		23, 7, 42, 0, time.UTC)
}

func (date ShowDate) String() string {
	if date.IsZero() {
		return ""
	}
	return fmt.Sprintf("%d/%02d/%02d",
		date.Year, date.Month, date.Day)
}

func (date ShowDate) MarshalText() ([]byte, error) {
	return []byte(date.String()), nil
}

func (date *ShowDate) UnmarshalText(text []byte) error {
	if len(text) != len("YYYY_MM_DD") {
		return fmt.Errorf("incorrect format for aired date '%s'", text)
	}

	year, err := strconv.Atoi(string(text[:4]))
	if err != nil {
		return err
	}
	month, err := strconv.Atoi(string(text[5:7]))
	if err != nil {
		return err
	}
	day, err := strconv.Atoi(string(text[8:]))
	if err != nil {
		return err
	}

	*date = ShowDate{year, month, day}
	return nil
}
