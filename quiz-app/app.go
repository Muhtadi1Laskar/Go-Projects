package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
)

type problem struct {
	q string
	a string
}

func main() {
	csvFileName := flag.String("cvs", "problems.csv", "a csv file in the format of 'question, answer'");
	flag.Parse();

	file, err := os.Open(*csvFileName);

	if err != nil {
		exit(fmt.Sprintf("Failed to open the CSV file: %s\n", *csvFileName));
		os.Exit(0);
	}

	r := csv.NewReader(file);
	lines, err := r.ReadAll();

	if err != nil {
		exit("Failed to parse the csv file");
	}

	problems := parseLines(lines);
	correctScore := 0

	for idx, elem := range problems {
		fmt.Printf("Problems # %d: %s = \n", idx+1, elem.q);

		var answer string;
		fmt.Scanf("%s\n", &answer);

		if answer == elem.a {
			correctScore += 1;
		}
	}
	fmt.Printf("You have %v correct answers out of %v", correctScore, len(problems));

}

func parseLines(lines [][]string) []problem {
	ret := make([]problem, len(lines));
	for idx, line := range lines {
		ret[idx] = problem{
			q: line[0],
			a: line[1],
		}
	}
	return ret;
}

func exit(message string) {
	fmt.Println(message);
	os.Exit(1);
}