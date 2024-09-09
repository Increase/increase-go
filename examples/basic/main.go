package main

import (
	"context"
	"fmt"

	"github.com/Increase/increase-go"
)

func main() {
	client := increase.NewClient()
	page, err := client.Cards.List(context.TODO(), increase.CardListParams{
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
