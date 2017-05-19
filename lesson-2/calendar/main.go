package main

import (
"fmt"
"time"
)

func main() {
	cases := []struct {
		month   string
		quarter int
	}{
		{month: "01", quarter: 1},
		{month: "02", quarter: 1},
		{month: "03", quarter: 1},
		{month: "04", quarter: 2},
		{month: "05", quarter: 2},
		{month: "06", quarter: 2},
		{month: "07", quarter: 3},
		{month: "08", quarter: 3},
		{month: "09", quarter: 3},
		{month: "10", quarter: 4},
		{month: "11", quarter: 4},
		{month: "12", quarter: 4},
	}

	//TODO Реализовать Календарь

	for _, test := range cases {
		parsed, _ := time.Parse("2006-01-02", fmt.Sprintf("2015-%s-15", test.month))
		fmt.Println(parsed)
		//calendar := NewCalendar(parsed)
		//actual := calendar.CurrentQuarter()
		//if actual != test.quarter {
		//	t.Error("Month:", test.month,
		//		"Expected Quarter:", test.quarter,
		//		"Actual Quarter:", actual)
		//}
	}
}

func