package countingsundays

import "fmt"

func ExecProjectEulerProblem() {
	dateBegin := Date{1, 1, 1901}
	dateEnd := Date{31, 12, 2000}
	println(sundayAmountsOnFirstOfMonthIn(dateBegin, dateEnd))
}

func sundayAmountsOnFirstOfMonthIn(dateBegin, dateEnd Date) int {
	date := Date{7, 1, 1900} //
	amount := 0
	for dateEnd.After(date) {
		date = addDays(date, 7)
		if date.After(dateBegin) {
			if date.date == 1 {
				amount++
			}
		}
	}
	return amount
}

// Date -- Ein Datum
type Date struct {
	date  int
	month int
	year  int
}

// Date.After - ermittelt ob das Datum später ist
func (a *Date) After(b Date) bool {
	if a.year > b.year {
		return true
	}
	if a.year < b.year {
		return false
	}
	if a.month > b.month {
		return true
	}
	if a.month < b.month {
		return false
	}
	return a.date > b.date
}

func addDays(date Date, addDays int) Date {
	newDate := date.date + addDays
	newMonth := date.month
	newYear := date.year
	var dateLimit int
	switch date.month {
	case 1:
		dateLimit = 31
	case 2:
		if isLeapYear(date.year) {
			dateLimit = 29
		} else {
			dateLimit = 28
		}
	case 3:
		dateLimit = 31
	case 4:
		dateLimit = 30
	case 5:
		dateLimit = 31
	case 6:
		dateLimit = 30
	case 7:
		dateLimit = 31
	case 8:
		dateLimit = 31
	case 9:
		dateLimit = 30
	case 10:
		dateLimit = 31
	case 11:
		dateLimit = 30
	case 12:
		dateLimit = 31
	default:
		fmt.Println("ERROR - Month ", date.month, " nicht bekannt")
	}
	if newDate > dateLimit {
		newDate -= dateLimit
		newMonth++
	}
	if newMonth > 12 {
		newYear++
		newMonth -= 12
	}

	return Date{newDate, newMonth, newYear}
}

func isLeapYear(year int) bool {
	if year%4 == 0 {
		if year%100 == 0 {
			if year%400 == 0 {
				return true
			}
			return false
		}
		return true
	}
	return false
}
