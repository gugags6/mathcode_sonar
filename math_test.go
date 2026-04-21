package main

import "testing"

func TestSoma(t *testing.T) {
	total := Soma(2, 3)
	if total != 5 {
		t.Errorf("Soma(2, 3) = %d; want 5", total)
	}

	}

	func TestSubtrai(t *testing.T) {
	total := Subtrai(3, 3)
	if total != 0 {
		t.Errorf("Subtrai(3, 3) = %d; want 0", total)
	}

	}

	func TestMulti(t *testing.T) {
	total := Multi(3, 3)
	if total != 9 {
		t.Errorf("Subtrai(3, 3) = %d; want 9", total)
	}

	}

	
	func TestDivide(t *testing.T) {
	total := Divide(3, 3)
	if total != 1 {
		t.Errorf("Subtrai(3, 3) = %d; want 1", total)
	}

	}