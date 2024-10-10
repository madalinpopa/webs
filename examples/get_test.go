package examples

import (
	"testing"
)

// TestGetDeckOfCards verifies the GetDeckOfCards function by checking for errors, non-nil deck, and non-empty deck ID.
func TestGetDeckOfCards(t *testing.T) {
	deck, err := GetDeckOfCards()

	if err != nil {
		t.Error("expected no error, got ", err)
	}

	if deck == nil {
		t.Error("expected deck to be not nil")
	}

	if deck != nil && deck.Id == "" {
		t.Error("expected deck id to be not empty")
	}

}
