package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"time"

	"github.com/fatih/color"
	"github.com/srishanbhattarai/nepcal/nepcal"
)

// Copied from "nepcal/constants.go" as these are not public but are needed for this specific use case.
const (
	adLBoundY = 1943
	adLBoundM = int(time.April)
	adLBoundD = 14
)

// DateMapEntry is each entry in the array of results returned by the reference URL.
type DateMapEntry struct {
	NpYear  int `json:"npYear"`
	NpMonth int `json:"npMonth"`
	NpDay   int `json:"npDay"`
	EnYear  int `json:"enYear"`
	EnMonth int `json:"enMonth"`
	EnDay   int `json:"enDay"`
}

// Latest reference data.
const refEntriesFile = "ref_67989f277816566067a53ac446e558ed94cc489c.min.json"

// Get the reference entries, properly serialized, and sliced for only the range for which we support.
func getRefEntries() ([]DateMapEntry, error) {
	// Read file.
	b, err := ioutil.ReadFile(refEntriesFile)
	if err != nil {
		return nil, err
	}

	// Marshal
	var entries []DateMapEntry
	err = json.Unmarshal(b, &entries)
	if err != nil {
		return nil, err
	}

	// Figure out the index at which the first entry for adLBoundY exists.
	index := 0
	for k, v := range entries {
		if v.EnYear == adLBoundY && v.EnMonth == adLBoundM && v.EnDay == adLBoundD {
			index = k
			break
		}
	}

	cov := (float64(len(entries)-index+1) / float64(len(entries))) * 100
	fmt.Printf("Starting at index: %d, percentage coverage of reference: %.2f\n", index, cov)

	return entries[index:], nil
}

// Get entries for the same date range as above but from the nepcal library, as specified using 'count'.
// Panics if date happens to be out of bounds, should never happen.
func getNepcalEntries(count int) []DateMapEntry {
	entries := make([]DateMapEntry, count)

	t := time.Date(adLBoundY, time.Month(adLBoundM), adLBoundD, 0, 0, 0, 0, time.UTC)
	i := 0

	for i < count {
		bs, err := nepcal.FromGregorian(t)
		if err != nil {
			panic(fmt.Sprintf("Invariant violation: %v, %s\n", err, t))
		}

		entry := DateMapEntry{
			NpYear:  bs.Year(),
			NpMonth: int(bs.Month()),
			NpDay:   bs.Day(),
			EnYear:  t.Year(),
			EnMonth: int(t.Month()),
			EnDay:   t.Day(),
		}

		entries[i] = entry

		t = t.AddDate(0, 0, 1)

		i++
	}

	return entries
}

// Format entry into string.
func fmtEntry(e DateMapEntry) string {
	us := fmt.Sprintf("(en) %d-%d-%d", e.EnYear, e.EnMonth, e.EnDay)
	np := fmt.Sprintf("(np) %d-%d-%d", e.NpYear, e.NpMonth, e.NpDay)

	return fmt.Sprintf("%s ==> %s", us, np)
}

// diff the two data sources and print to stdout.
// us = entries created by this library
// them = external entries for validation.
func diffEntries(us, them []DateMapEntry) {
	if len(us) != len(them) {
		panic(fmt.Sprintf("Invariant violation mismatching lengths; us = %v, them = %v\n", us, them))
	}

	type failure struct {
		index int
		us    DateMapEntry
		them  DateMapEntry
	}
	var failures []failure

	// Find cases where dates don't match
	for i := 0; i < len(us); i++ {
		if us[i] != them[i] {
			failures = append(failures, failure{i, us[i], them[i]})
		}
	}

	fmt.Printf(fmt.Sprintf("Number of inconsistencies: %s\n", color.YellowString(strconv.Itoa(len(failures)))))

	// Pretty print diffs
	for _, v := range failures {
		fmt.Printf("Inconsistency: %s\n", color.BlueString(strconv.Itoa(v.index)))

		us := fmt.Sprintf("- (actual)   %s", fmtEntry(v.us))
		them := fmt.Sprintf("+ (expected) %s", fmtEntry(v.them))

		fmt.Printf(fmt.Sprintf("%s\n", color.RedString(us)))
		fmt.Printf(fmt.Sprintf("%s\n\n", color.GreenString(them)))
	}
}

func main() {
	refEntries, err := getRefEntries()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	libEntries := getNepcalEntries(len(refEntries))
	diffEntries(libEntries, refEntries)
}
