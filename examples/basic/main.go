package main

import (
	"context"
	"fmt"

	"github.com/increase/increase-go"
	"github.com/increase/increase-go/option"
	"github.com/increase/increase-go/requests"
)

func main() {
	client := increase.NewIncrease(
		option.WithAPIKey("my api key"), // defaults to os.LookupEnv("INCREASE_API_KEY")
	)
	page, err := client.Cards.List(context.TODO(), &requests.CardListParams{
		AccountID: increase.F("account_in71c4amph0vgo2qllky"),
	})
	for page != nil {
		fmt.Printf("%+#v %+#v\n", page, err)
		for _, item := range page.Data {
			println(item.ID)
		}
		page, err = page.GetNextPage()
	}
	if err != nil {
		panic(err.Error())
	}
}
