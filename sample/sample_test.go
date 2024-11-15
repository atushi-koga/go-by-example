package sample

import "testing"

func TestSample1(t *testing.T) {
	got := Sample1()
	if got != 1 {
		t.Errorf("Sample1() = %v, want 1", got)
	}
}

func TestSample2(t *testing.T) {
	got := sample2()
	if got != 2 {
		t.Errorf("sample2() = %v, want 2", got)
	}
}
