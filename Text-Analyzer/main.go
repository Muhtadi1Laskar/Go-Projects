package main

import (
	"fmt"
	"strings"
)

func count(text, operation string) int {
	operation = strings.ToLower(operation)
	var result []string

	switch operation {
		case "word-count":
			result = tokenize(text, " ")
		case "character-count":
			text = strings.Join(strings.Fields(text), "")
			result = tokenize(text, "")
		case "sentence-count":
			result = tokenize(text, ".")
		default:
			return -1
	}

	return len(result)
}

func tokenize(text string, separator string) []string {
	text = strings.Trim(text, " ")
	return strings.Split(strings.ToLower(text), separator)
}

func main() {
var data string = `But as I walked down the steps I saw that the evening was not quite
over.Fifty feet from the door a dozen headlights illuminated a
bizarre and tumultuous scene. In the ditch beside the road, right side
up, but violently shorn of one wheel, rested a new coupé which had
left Gatsby’s drive not two minutes before. The sharp jut of a wall
accounted for the detachment of the wheel, which was now getting
considerable attention from half a dozen curious chauffeurs. However,
as they had left their cars blocking the road, a harsh, discordant din
from those in the rear had been audible for some time, and added to
the already violent confusion of the scene.`

	fmt.Println(count(data, "word-count"))
	fmt.Println(count(data, "character-count"))
	fmt.Println(count(data, "sentence-count"))
}