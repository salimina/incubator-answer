package main

import (
	"testing"
)

func TestValidateQuestionSubmission(t *testing.T) {
	tests := []struct {
		name      string
		title     string
		tags      *string
		body      *string
		expectErr bool
	}{
		{"Valid Title Only", "Valid Title", nil, nil, false},
		{"Optional Fields Null", "Another Title", nil, nil, false},
		{"Optional Fields Empty", "Another Title", strPtr(""), strPtr(""), false},
		{"Missing Title", "", strPtr("tag1"), strPtr("Sample body"), true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validateQuestionSubmission(tt.title, tt.tags, tt.body)
			if (err != nil) != tt.expectErr {
				t.Errorf("Expected error: %v, got: %v", tt.expectErr, err)
			}
		})
	}
}

func strPtr(s string) *string {
	return &s
}

// Mock function to simulate validation logic
func validateQuestionSubmission(title string, tags, body *string) error {
	if title == "" {
		return fmt.Errorf("title is required")
	}
	return nil
}
