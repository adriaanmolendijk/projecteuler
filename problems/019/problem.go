package main

import "fmt"

func main() {
	date := 1
	day := 1
	month := 1
	year := 1900
	monthDayMap := map[int]int{
		1:  31, // Jan
		2:  28, // Feb
		3:  31, // Mar
		4:  30, // Apr
		5:  31, // May
		6:  30, // Jun
		7:  31, // Jul
		8:  31, // Aug
		9:  30, // Sep
		10: 31, // Oct
		11: 30, // Nov
		12: 31, // Dec
	}

	count := 0
	for year < 2001 {
		if firstMonthSunday(day, date, month) && year >= 1901 {
			count++
		}

		day++
		if day > 7 {
			day = 1
		}

		date++
		if date > monthDayMap[month] {
			if month == 2 && isLeapYear(year) {
				if date > 29 {
					date = 1
					month++
				}
			} else {
				date = 1
				month++
			}
			if month > 12 {
				month = 1
				year++
			}
		}
	}

	fmt.Println(count)
}

func firstMonthSunday(day, date, month int) bool {
	return date == 1 && day == 7
}

func isLeapYear(year int) bool {
	if year%4 == 0 && (year%100 != 0 || year%400 == 0) {
		return true
	}
	return false
}
