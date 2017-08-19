package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

const (
	FirstWeekend = 6 //2000 year 1 month 1 day is saturday
	MonthNum     = 12
	WeekendNum   = 7
	MaxYear      = 10000
	MinYear      = -10000
)

var num = []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

func findK(year int) (k int) {
	sub := year - 2000
	if sub == 0 {
		k = FirstWeekend
	} else if sub > 0 {
		k = (sub + FirstWeekend + (sub+3)/4 - (sub-1)/100 + (sub-1)/400) % 7
	} else {
		k = (sub + FirstWeekend + sub/4 - sub/100 + sub/400) % 7
	}
	if k < 0 {
		k += 7
	}
	return k
}

func isLeapYear(year int) (ok bool) {
	if year%400 == 0 || year%4 == 0 && year%100 != 0 {
		return true
	}
	return false
}

func ByYear(year int) {
	if year > MinYear && year < MaxYear {
		var i, j, k int
		k = findK(year)
		if isLeapYear(year) {
			num[1] = 29
		} else {
			num[1] = 28
		}
		fmt.Printf("************%d", year, "year calendar***********\n")
		for i = 0; i < MonthNum; i++ {
			fmt.Printf("               %30d      year            %-20s     ",
				year, time.Month(i))
		}
		for j = 0; j < WeekendNum; j++ {
			fmt.Printf("    %s", time.Weekday(j))
		}
		fmt.Println()
		for j = 0; j < k; j++ {
			fmt.Printf("           ")
		}
		for j = 0; j < num[i]; j++ {
			if k == 0 {
				fmt.Println()
			}
			if j > 9 {
				fmt.Printf("%8d", j+1)
			} else {
				fmt.Printf("%9d", j+1)
			}
			k = (k + 1) % 7
		}
		fmt.Println()
	} else {
		fmt.Println("error")
	}
}

func ByMonth(year int, mon int) {
	if year > MinYear && year < MaxYear && mon > 0 && mon < MonthNum {
		var i, j, k int
		k = findK(year)
		if isLeapYear(year) {
			num[1] = 29
		} else {
			num[1] = 28
		}
		for i = 0; i < mon-1; i++ {
			k = k + num[i]
		}
		k %= 7
		fmt.Printf("              %d   %s        \n", year, time.Month(mon))
		for j = 0; j < WeekendNum; j++ {
			fmt.Printf("%10s", time.Weekday(j))
		}
		for j = 0; j < k; j++ {
			if j == 0 {
				fmt.Printf("\n%10s", " ")
			} else {
				fmt.Printf("%10s", " ")
			}
		}
		for j = 0; j < num[mon-1]; j++ {
			if k == 0 {
				fmt.Println()
			}
			if j > 9 {
				fmt.Printf("%10d", j+1)
			} else {
				fmt.Printf("%10d", j+1)
			}
			k = (k + 1) % 7
		}
		fmt.Println()
	} else {
		fmt.Println("error")
	}
}

func GetWeekDay(year, mon, d int) {
	if year > MinYear && year < MaxYear && mon > 0 && mon < MonthNum+1 {
		var i, k int
		k = findK(year)
		if isLeapYear(year) {
			num[1] = 29
		} else {
			num[1] = 28
		}
		if d > 0 && d <= num[mon-1] {
			for i = 0; i < mon-1; i++ {
				k = k + num[i]
			}
			k = (k + d - 1) % 7
			fmt.Printf("                         %s %d, %d is : %s \n", time.Month(mon), d, year, time.Weekday(k))

		} else {
			fmt.Println("error")
		}

	} else {
		fmt.Println("error")
	}
}

func main() {
	arg_num := len(os.Args)
	var digit []int
	for i := 1; i < arg_num; i++ {
		curr, err := strconv.Atoi(os.Args[i])
		if err != nil {
			fmt.Println("error happened ,exit")
			return
		}
		digit = append(digit, curr)
	}
	if arg_num == 4 {
		GetWeekDay(digit[0], digit[1], digit[2])
	} else if arg_num == 3 {
		ByMonth(digit[0], digit[1])
	} else if arg_num == 2 {
		ByYear(digit[0])
	} else {
		fmt.Println("arg num err")
	}
	return
}
