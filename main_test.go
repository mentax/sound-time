package main

import (
	"testing"
)

func TestDuration(t *testing.T) {
	// Test with the first MP3 file
	duration1, err := duration("data/1.mp3")
	if err != nil {
		t.Errorf("Error getting duration for 1.mp3: %v", err)
	}
	// This is a sample duration, the actual duration might be different
	// and may need adjustment after running the test.
	expectedDuration1 := 7
	if duration1 != expectedDuration1 {
		t.Errorf("Duration for 1.mp3 is incorrect, got: %d, want: %d", duration1, expectedDuration1)
	}

	// Test with the second MP3 file
	duration2, err := duration("data/2.mp3")
	if err != nil {
		t.Errorf("Error getting duration for 2.mp3: %v", err)
	}
	// This is a sample duration, the actual duration might be different
	// and may need adjustment after running the test.
	expectedDuration2 := 131
	if duration2 != expectedDuration2 {
		t.Errorf("Duration for 2.mp3 is incorrect, got: %d, want: %d", duration2, expectedDuration2)
	}
}