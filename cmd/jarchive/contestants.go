package main

import (
	"fmt"
	"log"
	"strconv"
)

type JArchiveContestant struct {
	JCID `json:"id,omitempty"`
	Name string `json:"name"`
	Bio  string `json:"comment"`
}

// Unique numeric value for identifying customers in the archive.
type JCID int

func (id JCID) String() string {
	return fmt.Sprintf("%d", int(id))
}

// Parses the numeric value from a string.
// Fatal error if the value cannot be converted into a number.
func MustParseJCID(numeric string) JCID {
	id, err := strconv.Atoi(numeric)
	if err != nil {
		log.Fatalf("failed to parse JCID from string '%s'\n%s", numeric, err)
	}
	return JCID(id)
}
