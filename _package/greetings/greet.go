package greetings

import "time"

func GoodDay() string {
	return "Good Day"
}

func GoodNight() string {
	return "Good Night"
}

//func IsAM() bool {
//	now := time.Now()
//	noon, _ := time.Parse(time.Kitchen, "12:00PM")
//	noon = noon.AddDate(now.Year(), int(now.Month()), now.Day())
//	return now.Before(noon)
//}
//
//func IsEvening() bool {
//	now := time.Now()
//	evening, _ := time.Parse(time.Kitchen, "7:00PM")
//	evening = evening.AddDate(now.Year(), int(now.Month()), now.Day())
//	return now.Before(evening)
//}
//
//func IsAfternoon() bool {
//	return !IsAM() && IsEvening()
//}

func IsAM() bool {
	localTime := time.Now()
	return localTime.Hour() <= 12
}

func IsAfternoon() bool {
	localTime := time.Now()
	return localTime.Hour() <= 18
}

func IsEvening() bool {
	localTime := time.Now()
	return localTime.Hour() <= 22
}
