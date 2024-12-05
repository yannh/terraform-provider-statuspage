package statuspage

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/hashicorp/go-retryablehttp"
)

func NewRetryableClient() *retryablehttp.Client {
	client := retryablehttp.NewClient()
	client.CheckRetry = retryPolicy
	client.Backoff = backoffPolicy
	return client
}

const (
	StatusRateLimitExceeded = 420
)

// We're wrapping original retryablehttp.DefaultRetryPolicy and adding retry on 420 HTTP code
func retryPolicy(ctx context.Context, resp *http.Response, err error) (bool, error) {
	defaultShouldRetry, err := retryablehttp.DefaultRetryPolicy(ctx, resp, err)
	if err != nil {
		return false, err
	}

	isRateLimited := resp != nil && resp.StatusCode == StatusRateLimitExceeded
	shouldRetry := defaultShouldRetry || isRateLimited
	return shouldRetry, nil
}

// We're wrapping original retryablehttp.DefaultBackoff and using response header in such format `Retry-After: 60`
// retryablehttp already implements this behaviour, but it's executed only for 429 and 503 HTTP codes, and we need it for 420 HTTP code as well
func backoffPolicy(min, max time.Duration, attemptNum int, resp *http.Response) time.Duration {
	if resp != nil {
		if resp.StatusCode == StatusRateLimitExceeded {
			if sleep, err := parseRetryAfterHeader(resp.Header.Get("Retry-After")); err == nil {
				return sleep
			}
		}
	}
	defaultSleep := retryablehttp.DefaultBackoff(min, max, attemptNum, resp)
	return defaultSleep
}

// Code partially copied from retryablehttp.parseRetryAfterHeader
func parseRetryAfterHeader(retryAfter string) (time.Duration, error) {
	if retryAfter == "" {
		return 0, fmt.Errorf("response Header Retry-After empty or not set")
	}
	// Retry-After: 60
	if sleep, err := strconv.ParseInt(retryAfter, 10, 64); err == nil {
		if sleep < 0 { // a negative sleep doesn't make sense
			return 0, fmt.Errorf("response Header Retry-After set to negative value: %s", retryAfter)
		}
		return time.Second * time.Duration(sleep), nil
	}
	return 0, nil
}
