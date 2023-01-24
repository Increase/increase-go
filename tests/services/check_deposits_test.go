package services

import (
	"context"
	"testing"

	client "increase"
)

func TestCheckDepositsReject(t *testing.T) {
	t.Skip("Prism incorrectly returns an invalid JSON error")
	c := client.NewIncreaseWithOptions(client.ClientOptions{
		APIKey:  "something1234",
		BaseURL: "http://127.0.0.1:4010",
	})
	_, err := c.CheckDeposits.Simulations.Reject(context.TODO(), "check_deposit_f06n9gpg7sxn8t19lfc1")
	if err != nil {
		t.Fatal("err should not be nil", err)
	}
}

func TestCheckDepositsSubmit(t *testing.T) {
	t.Skip("Prism incorrectly returns an invalid JSON error")
	c := client.NewIncreaseWithOptions(client.ClientOptions{
		APIKey:  "something1234",
		BaseURL: "http://127.0.0.1:4010",
	})
	_, err := c.CheckDeposits.Simulations.Submit(context.TODO(), "check_deposit_f06n9gpg7sxn8t19lfc1")
	if err != nil {
		t.Fatal("err should not be nil", err)
	}
}
