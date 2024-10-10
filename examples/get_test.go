package examples

import (
	"errors"
	"github.com/madalinpopa/webs"
	"net/http"
	"testing"
)

func TestGetDeckOfCards(t *testing.T) {
	t.Run("GetDeckOfCards", func(t *testing.T) {
		deck, err := GetDeckOfCards()

		mock := webs.Mock{
			Method: http.MethodGet,
			Url:    "https://deckofcardsapi.com/api/deck/new/shuffle/?deck_count=1",
			Error:  errors.New("error"),
		}
		webs.AddMock(mock)

		if err == nil {
			t.Errorf("expected error, got nil")
		}

		if deck != nil {
			t.Errorf("expected nil, got %+v", deck)
		}

	})
}
