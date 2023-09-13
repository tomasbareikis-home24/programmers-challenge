package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	bytes, _ := io.ReadAll(os.Stdin)

	teamsOrder := []string{}

	teamsResults := map[string]*studentsAndDistance{}

	for _, l := range parseInput(bytes) {
		stepsRecord, err := parseInputLine(l)
		if err != nil {
			panic(err)
		}

		if stepsRecord == nil {
			continue
		}

		dist := stepsRecord.distance()

		// the distance == 0 if at least one of the steps is 0 or the steps array is empty
		// we do not count such records
		if dist == 0 {
			continue
		}

		if _, ok := teamsResults[stepsRecord.team]; !ok {
			teamsResults[stepsRecord.team] = &studentsAndDistance{}
			teamsOrder = append(teamsOrder, stepsRecord.team)
		}

		teamsResults[stepsRecord.team].students++
		teamsResults[stepsRecord.team].distance += dist
	}

	for _, team := range teamsOrder {
		fmt.Printf("%s %d %.2f\n", team, teamsResults[team].students, teamsResults[team].distance)
	}
}

func parseInput(input []byte) []string {
	return strings.Split(string(input), "\n")
}

type studentsAndDistance struct {
	students int
	distance float64
}

type stepsRecord struct {
	team       string
	stepLength int
	steps      []int
}

func (sr *stepsRecord) distance() float64 {
	distance := 0
	for _, step := range sr.steps {
		if step == 0 {
			return 0
		}
		distance += step * sr.stepLength
	}

	return float64(distance) / 100 / 1000 // convert to km
}

func parseInputLine(line string) (*stepsRecord, error) {
	lineFields := strings.Fields(line)

	if len(lineFields) < 3 { // at least 3 fields are required: team, stepLength, steps
		return nil, nil
	}

	sr := stepsRecord{
		team: lineFields[0],
	}

	stepLength, err := strconv.Atoi(lineFields[1])
	if err != nil {
		return nil, err
	}

	sr.stepLength = stepLength

	steps := make([]int, 0, len(lineFields)-2)

	for _, strVal := range lineFields[2:] {
		intVal, err := strconv.Atoi(strVal)
		if err != nil {
			return nil, err
		}
		steps = append(steps, intVal)
	}

	sr.steps = steps

	return &sr, nil
}
