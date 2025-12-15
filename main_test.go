package main

import "testing"

// TestHello is the simplest possible test
func TestSystemMonitor(t *testing.T) {
	// This test always passes - it's just to verify the test setup
	t.Log("Test framework is working")

	// You can add assertions as you write more code
	expected := "test"
	if len(expected) != 4 {
		t.Errorf("Expected 'test' to have 4 letters")
	}
}
