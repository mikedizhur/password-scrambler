package main

import (
	"cmp"
	"errors"
	"slices"
)

var defaultTranslations = [][]string{
	{"_", " ", "-"},
	{"!", "1", "I", "l", "|", "i"},
	{"#", "4", "H", "h"},
	{"$", "5", "S", "s"},
	{"&", "8", "B", "b"},
	{"@", "a", "A"},
	{"0", "D", "O"},
	{"0", "O", "o"},
	{"2", "Z", "z"},
	{"3", "E", "e"},
	{"4", "A"},
	{"6", "G", "b"},
	{"7", "T", "t"},
	{"9", "G", "g"},
	{"<", "C", "c"},
}

func fcharCmp(a, b []string) int {
	return cmp.Compare(a[0], b[0])
}

func formatTranslations(translations [][]string) ([][]string, error) {
	// just to know that everything is sorted if list is custom
	if len(translations) == 0 {
		return [][]string{}, errors.New("Empty translations list")
	}


	var outList [][]string

	for _, list := range translations {
		slices.Sort(list)
		outList = append(outList, list)
	}

	slices.SortFunc(outList, fcharCmp)

	return outList, nil
}

// should optimize one day
type aliasMap map[string][]string

func mergeSets[K cmp.Ordered](a, b []K) []K {
	// merge 2 ordered slices, while removing duplicates

	if len(a) == 0 {
		return b
	} else if len(b) == 0 {
		return a
	}

	var merged []K

	sliderA := 0
	sliderB := 0

	for {
		if sliderA >= len(a) {
			merged = append(merged, b[sliderB:]...)
			return merged
		} else if sliderB >= len(b) {
			merged = append(merged, a[sliderA:]...)
			return merged
		}
		if a[sliderA] < b[sliderB] {
			merged = append(merged, a[sliderA])
			sliderA++
		} else if b[sliderB] < a[sliderA] {
			merged = append(merged, b[sliderB])
			sliderB++
		} else {
			sliderA++
		}
	}
}

/*
 * Take symbol from list
 * Add it as key to aliasMap
 * add the list as value
 * if key already in map:
 * 	check if lists match (they should not)
 * 	if they do, pass
 * 	if they do not, merge slices
*/


func genTransMaps(translations [][]string) aliasMap {
	outMap := make(aliasMap) // change if type is changed !!
	for _, symList := range translations {
		for _, sym := range symList {
			if _, exists := outMap[sym]; exists {
				outMap[sym] = mergeSets(outMap[sym], symList)
			} else {
				outMap[sym] = symList
			}
		}
	}
	return outMap
}
