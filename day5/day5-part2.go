package main

import (
	"strconv"
	"strings"
	"sort"
)

// Range struct to hold start and end of a numerical range
type Range struct {
	start int
	end   int
}


func allFresh(rangesStr []string) {
	var ranges []Range
	for _, rStr := range rangesStr {
		parts := strings.Split(rStr, "-")
		if len(parts) != 2 {
			continue
		}

		start, err := strconv.Atoi(parts[0])
		if err != nil {
			continue
		}
		end, err := strconv.Atoi(parts[1])
		if err != nil {
			continue
		}
		ranges = append(ranges, Range{start: start, end: end})
	}

	if len(ranges) == 0 {
		println(0)
		return
	}

	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i].start < ranges[j].start
	})

	mergedRanges := []Range{ranges[0]}

	for i := 1; i < len(ranges); i++ {
		lastMerged := &mergedRanges[len(mergedRanges)-1]
		current := ranges[i]

		if current.start <= lastMerged.end+1 { // +1 to account for contiguous ranges
			if current.end > lastMerged.end {
				lastMerged.end = current.end
			}
		} else {
			mergedRanges = append(mergedRanges, current)
		}
	}

	totalUnique := 0
	for _, r := range mergedRanges {
		totalUnique += (r.end - r.start + 1)
	}

	println(totalUnique)
}
