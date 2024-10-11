# Webs

![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/madalinpopa/webs) [![Test](https://github.com/madalinpopa/webs/actions/workflows/test.yml/badge.svg)](https://github.com/madalinpopa/webs/actions/workflows/test.yml)


This is a simple HTTP client that can be used to send HTTP requests to a server
and receive the response. The client is implemented in Go and uses the standard
library.

## Usage

To use the library, you need to import the package.

```go
import "github.com/madalinpopa/webs"
```

### Configure and build the client

```go

headers := make(http.Header)
headers.Set("Some-Common-Header", "value-for-all-requests")

client := webs.NewClientBuilder().
	SetHeaders(headers).
	SetConnectionTimeout(2 * time.Second).
	SetResponseTimeout(3 * time.Second).
	Build()
```
### GET

```go

type DeckOfCards struct {
    Id        string `json:"deck_id"`
    Remaining int    `json:"remaining"`
}

func GetDeckOfCards() (*DeckOfCards, error) {
    url := "https://deckofcardsapi.com/api/deck/new/shuffle/?deck_count=1"

    response, err := client.Get(url, nil)
    if err != nil {
        return nil, err
    }   

    var deck DeckOfCards
    if err := response.UnmarshalJson(&deck); err != nil {
        return nil, err
    }

    return &deck, nil
}
```

