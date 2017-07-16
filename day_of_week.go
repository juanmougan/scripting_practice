// Program to calculate which day of the week is a date 
// passed as a parameter.
// This date has to be after 01/01/1970.

package main

import "fmt"

// Returns 28 or 29 for February
func getFebruaryForYear(y int) int {
    if ((y % 4 == 0) && (y % 100 != 0)) || (y % 400 == 0) {
       return 29
    }
    return 28
}

// Returns true if the day passed should be converted to the next month's 1st.
// E.g. passing "1970/12/32" should return true - it should be converted to "1971/01/01"
func dayExceeded(y int, m int, d int) bool {
    days := make(map[int]int)
    days[1] = 31
    days[2] = getFebruaryForYear(y)
    days[3] = 31
    days[4] = 30
    days[5] = 31
    days[6] = 30
    days[7] = 31
    days[8] = 31
    days[9] = 30
    days[10] = 31
    days[11] = 30
    days[12] = 31
    return d > days[m]
}

// Returns the next day for a given day in format (year, month, day)
func nextDay(y int, m int, d int) (int, int, int) {
    var yy, mm, dd int = y, m, d
    dd = dd + 1
    if dayExceeded(y, m, dd) {
        dd = 1
        mm = mm + 1
    }
    if mm > 12 {
        dd = 1
        mm = 1
        yy = yy + 1
    }
    return yy, mm, dd
}

// For a given day, returns the number of days from 01/01/1970.
func dateToDays(y int, m int, d int) (int) {
    dd := 1
    mm := 1
    yy := 1970
    count := 0
    for y != yy && m != mm && d != dd {
        y, m, d = nextDay(y, m, d)
        count = count + 1
    }
    return count
}

// Gets the dayshift (e.g. sunday = 0, monday = 1, etc.) 
// for a given day from 01/01/1970.
func getDayshift(dayAsNum int) int {
    const baseShift = 4
    return (dayAsNum + baseShift) % 7
}

// Converts a day shifting after 01/01/1970 (as a Number)
// to a day of the week (e.g. "Monday").
func convertNumToWeekday(dayAsNum int) string {
    m := make(map[int]string)
    m[0] = "Sunday"
    m[1] = "Monday"
    m[2] = "Tuesday"
    m[3] = "Wednesday"
    m[4] = "Thursday"
    m[5] = "Friday"
    m[6] = "Saturday"
    return m[dayAsNum]
}

// For a given day, returns which day of the week it is.
// The date passed as a parameter has to be after 01/01/1970.
func DayOfWeek(year int, month int, day int) string {
    dayAsNum := dateToDays(year, month, day)
    weekdayAsNum := getDayshift(dayAsNum)
    return convertNumToWeekday(weekdayAsNum)
}

func main() {

    // Test
    var y, m, d int = 1970, 12, 31
    fmt.Println(DayOfWeek(y, m, d))
}
