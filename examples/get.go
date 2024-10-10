package examples

type DeckOfCards struct {
	Id        string `json:"deck_id"`
	Remaining int    `json:"remaining"`
}

func GetDeckOfCards() (*DeckOfCards, error) {
	url := "https://deckofcardsapi.com/api/deck/new/shuffle/?deck_count=1"

	response, err := client.Get(url, nil)
	if err != nil {
		return &DeckOfCards{}, err
	}

	var deck DeckOfCards
	if err := response.UnmarshalJson(&deck); err != nil {
		return &DeckOfCards{}, err
	}

	return &deck, nil
}
