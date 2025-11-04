// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package spotted_test

import (
	"net/http"
	"strconv"
	"testing"
	"time"

	"github.com/stainless-sdks/spotted-go"
	"github.com/stainless-sdks/spotted-go/option"
	standardwebhooks "github.com/standard-webhooks/standard-webhooks/libraries/go"
)

func TestWebhookUnwrap(t *testing.T) {
	client := spotted.NewClient(
		option.WithWebhookKey("whsec_c2VjcmV0Cg=="),
		option.WithClientID("My Client ID"),
		option.WithClientSecret("My Client Secret"),
	)
	payload := []byte("{}")
	wh, err := standardwebhooks.NewWebhook("whsec_c2VjcmV0Cg==")
	if err != nil {
		t.Error("Failed to sign test webhook message")
	}
	msgID := "1"
	now := time.Now()
	sig, err := wh.Sign(msgID, now, payload)
	if err != nil {
		t.Error("Failed to sign test webhook message:", err)
	}
	headers := make(http.Header)
	headers.Set("webhook-signature", sig)
	headers.Set("webhook-id", msgID)
	headers.Set("webhook-timestamp", strconv.FormatInt(now.Unix(), 10))
	_, err = client.Webhooks.Unwrap(payload, headers)
	if err != nil {
		t.Error("Failed to unwrap webhook:", err)
	}
}
