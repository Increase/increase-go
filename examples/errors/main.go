package main

import (
	"context"
	"errors"
	"fmt"

	"github.com/increase/increase-go"
	"github.com/increase/increase-go/option"
	"github.com/increase/increase-go/requests"
)

func main() {
	client := increase.NewIncrease(option.WithAPIKey("bad_api_key"))
	_, err := client.Cards.List(context.TODO(), &requests.CardListParams{
		AccountID: increase.F("account_in71c4amph0vgo2qllky"),
	})
	if err == nil {
		panic("expected an error with an invalid API key")
	}

	var apierr *increase.Error
	if errors.As(err, &apierr) {
		fmt.Printf("%+#v\n", apierr)
		println(apierr.Error())
		println(apierr.Detail)
	}
}
