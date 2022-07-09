package core

import (
	"sort"
)

func GetSerialNumber(key int, episodes map[int]bool) int {
	var sortedIDs []int
	for k := range episodes {
		sortedIDs = append(sortedIDs, k)
	}
	sort.SliceStable(sortedIDs, func(i, j int) bool { return sortedIDs[i] < sortedIDs[j] })
	serial := 1
	for _, k := range sortedIDs {
		if k == key {
			return serial
		}
		serial++
	}
	return -1
}

func getIDFromSeason(season int, epNum int, episodes []int) int {
	// Fallback method to if getEpisodeInfoByNumber is down
	return 0
}
