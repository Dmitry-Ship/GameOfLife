package gameOfLife

import "testing"

func TestNewWorldHighDensityNoPanic(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Fatalf("NewWorld panicked with high density: %v", r)
		}
	}()
	w := NewWorld(5, 5, 20)
	if w == nil {
		t.Fatal("NewWorld returned nil")
	}
}
