package main

import(
	"testing"
)

func TestGetFebruaryForYear(t *testing.T) {
	if getFebruaryForYear(2012) != 29 {
		t.Error("expected true")
	}
	if getFebruaryForYear(2011) != 28 {
		t.Error("expected true")
	}
	if getFebruaryForYear(2000) != 29 {
		t.Error("expected true")
	}
	if getFebruaryForYear(1900) != 28 {
		t.Error("expected true")
	}
	if getFebruaryForYear(1980) != 29 {
		t.Error("expected true")
	}
	if getFebruaryForYear(1978) != 28 {
		t.Error("expected true")
	}
	if getFebruaryForYear(1984) != 29 {
		t.Error("expected true")
	}
	if getFebruaryForYear(1985) != 28 {
		t.Error("expected true")
	}
}

func TestDayExceeded(t *testing.T) {
	// Should return true - 01/01/1971
	y := 1970
	m := 12
	d := 32
	exceeded := dayExceeded(y, m, d)
	if !exceeded {
		t.Error("expected true")
	}

	// Should return false
	y = 1970
	m = 12
	d = 31
	exceeded = dayExceeded(y, m, d)
	if exceeded {
		t.Error("expected false")
	}

	// Should return true
	y = 1970
	m = 11
	d = 31
	exceeded = dayExceeded(y, m, d)
	if !exceeded {
		t.Error("expected true")
	}

	// Should return true
	y = 1970
	m = 11
	d = 30
	exceeded = dayExceeded(y, m, d)
	if exceeded {
		t.Error("expected false")
	}

	// Should return false
	y = 1970
	m = 2
	d = 30
	exceeded = dayExceeded(y, m, d)
	if !exceeded {
		t.Error("expected true")
	}

	// Should return false
	y = 1970
	m = 2
	d = 29
	exceeded = dayExceeded(y, m, d)
	if !exceeded {
		t.Error("expected true")
	}
}

func TestNextDay(t *testing.T) {
	// Should return 1971/01/01
	y := 1970
	m := 12
	d := 31
	yy, mm, dd := nextDay(y, m, d)
	if yy != 1971 || mm != 1 || dd != 1 {
		t.Error("expected 1971/01/01 but got", yy, mm, dd)
	}

	// Should return 1970/02/29
	y = 1970
	m = 2
	d = 28
	yy, mm, dd = nextDay(y, m, d)
	if yy != 1970 || mm != 3 || dd != 1 {
		t.Error("expected 1970/03/01 but got", yy, mm, dd)
	}

	// Should return 1972/02/29
	y = 1972
	m = 2
	d = 28
	yy, mm, dd = nextDay(y, m, d)
	if yy != 1972 || mm != 2 || dd != 29 {
		t.Error("expected 1970/02/29 but got", yy, mm, dd)
	}
}
