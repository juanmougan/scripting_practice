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
