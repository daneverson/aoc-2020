package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

var (
	numRows = 127
	numCols = 7
)

type seatLocation struct {
	row, col, ID int
}

func main() {
	contents, err := ioutil.ReadFile("./input")
	if err != nil {
		log.Fatalf("failed reading input file: %v", err)
	}
	boardingPasses := strings.Split(string(contents), "\n")
	// boardingPasses := []string{"FBFBBFFRLR"}

	seatLocs := decodeLocations(boardingPasses)

	max := 0
	for _, seat := range seatLocs {
		if seat.ID > max {
			max = seat.ID
		}
	}

	fmt.Printf("The largest seat ID found: %d\n", max)

	// Part 2: find your seat
	seatIDs := map[int]struct{}{}
	for _, loc := range seatLocs {
		seatIDs[loc.ID] = struct{}{}
	}

	for i := 0; i < numRows; i++ {
		for j := 0; j < numCols; j++ {
			thisSeatID := i*8 + j
			if _, thisOccupied := seatIDs[thisSeatID]; !thisOccupied {
				if _, prevExists := seatIDs[thisSeatID-1]; prevExists {
					if _, nextExists := seatIDs[thisSeatID+1]; nextExists {
						fmt.Printf("Found my seat: %d\n", thisSeatID)
						return
					}
				}
			}
		}
	}
}

func decodeLocations(codes []string) []seatLocation {
	locs := []seatLocation{}
	for _, code := range codes {
		row := decodeRow(code[:7], 0, numRows)
		col := decodeCol(code[7:], 0, numCols)
		locs = append(locs, seatLocation{row, col, row*8 + col})
	}
	return locs
}

func decodeRow(rowCode string, min, max int) int {
	if len(rowCode) == 0 {
		return min
	}
	if rowCode[0] == 'F' {
		return decodeRow(rowCode[1:], min, min+(max-min)/2)
	}
	return decodeRow(rowCode[1:], min+(max-min+1)/2, max)
}

func decodeCol(colCode string, min, max int) int {
	if len(colCode) == 0 {
		return min
	}
	if colCode[0] == 'L' {
		return decodeCol(colCode[1:], min, min+(max-min)/2)
	}
	return decodeCol(colCode[1:], min+(max-min+1)/2, max)
}
