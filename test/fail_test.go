package test

import "testing"

func TestFail(t *testing.T) {
	t.Errorf("Error here")
}

func TestFail2(t *testing.T) {
	t.Errorf("Error 2 here")
}
