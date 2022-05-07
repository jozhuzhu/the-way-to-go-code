package _map

import "fmt"

func PrintWeek() {
	week := map[int]string{
		1: "Sunday",
		2: "Monday",
		3: "Tuesday",
		4: "Wednesday",
		5: "Thursday",
		6: "Friday",
		7: "Saturday",
	}

	findHollyDay := false

	for key, value := range week {
		fmt.Printf("day %d: %s\n", key, value)
	}

	for key, day := range week {
		if day == "Tuesday" || day == "Hollyday" {
			fmt.Println(day, " is the ", key, "th day in the week")
			if day == "Hollyday" {
				findHollyDay = true
			}
		}
	}

	if !findHollyDay {
		fmt.Println("Hollyday is not a day in week")
	}
}
