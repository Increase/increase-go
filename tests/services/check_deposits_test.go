package services

import (
	"context"
	"increase"
	"increase/options"
	"testing"
)

func TestCheckDepositsReject(t *testing.T) {
	t.Skip("Prism incorrectly returns an invalid JSON error")
	c := increase.NewIncrease(options.WithAPIKey("something1234"), options.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.Simulations.CheckDeposits.Reject(context.TODO(), "check_deposit_f06n9gpg7sxn8t19lfc1")
	if err != nil {
		t.Fatal("err should be nil", err)
	}
}

func TestCheckDepositsSubmit(t *testing.T) {
	t.Skip("Prism incorrectly returns an invalid JSON error")
	c := increase.NewIncrease(options.WithAPIKey("something1234"), options.WithBaseURL("http://127.0.0.1:4010"))
	_, err := c.Simulations.CheckDeposits.Submit(context.TODO(), "check_deposit_f06n9gpg7sxn8t19lfc1")
	if err != nil {
		t.Fatal("err should be nil", err)
	}
}
