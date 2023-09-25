// File generated from our OpenAPI spec by Stainless.

package increase_test

import (
	"context"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/increase/increase-go"
	"github.com/increase/increase-go/option"
)

func TestContextCancel(t *testing.T) {
	client := increase.NewClient(
		option.WithBaseURL("http://127.0.0.1:4010"),
		option.WithAPIKey("APIKey"),
	)
	cancelCtx, cancel := context.WithCancel(context.Background())
	cancel()
	res, err := client.Accounts.New(cancelCtx, increase.AccountNewParams{
		Name: increase.F("My First Increase Account"),
	})
	if err == nil || res != nil {
		t.Error("Expected there to be a cancel error and for the response to be nil")
	}
}

// neverTransport never completes a request and waits for the Context to be done.
type neverTransport struct{}

func (t *neverTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	<-req.Context().Done()
	return nil, fmt.Errorf("cancelled")
}

func TestContextCancelDelay(t *testing.T) {
	client := increase.NewClient(
		option.WithBaseURL("http://127.0.0.1:4010"),
		option.WithAPIKey("APIKey"),
		option.WithHTTPClient(&http.Client{Transport: &neverTransport{}}),
	)
	cancelCtx, cancel := context.WithCancel(context.Background())
	go func() {
		time.Sleep(time.Millisecond * time.Duration(2))
		cancel()
	}()
	res, err := client.Accounts.New(cancelCtx, increase.AccountNewParams{
		Name: increase.F("My First Increase Account"),
	})
	if err == nil || res != nil {
		t.Error("expected there to be a cancel error and for the response to be nil")
	}
}

func TestContextDeadline(t *testing.T) {
	testTimeout := time.After(3 * time.Second)
	testDone := make(chan bool)

	deadline := time.Now().Add(100 * time.Millisecond)
	deadlineCtx, cancel := context.WithDeadline(context.Background(), deadline)
	defer cancel()

	go func() {
		client := increase.NewClient(
			option.WithBaseURL("http://127.0.0.1:4010"),
			option.WithAPIKey("APIKey"),
			option.WithHTTPClient(&http.Client{Transport: &neverTransport{}}),
		)
		res, err := client.Accounts.New(deadlineCtx, increase.AccountNewParams{
			Name: increase.F("My First Increase Account"),
		})
		if err == nil || res != nil {
			t.Error("expected there to be a deadline error and for the response to be nil")
		}
		testDone <- true
	}()

	select {
	case <-testTimeout:
		t.Fatal("client didn't finish in time")
	case <-testDone:
		diff := time.Now().Sub(deadline)
		if diff < -20*time.Millisecond || 20*time.Millisecond < diff {
			t.Logf("error difference: %v", diff)
			t.Fatal("client did not return within 20ms of context deadline")
		}
	}
}
