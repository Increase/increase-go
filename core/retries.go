package core

// func (c *CoreClient) getRetriesRemaining(cr *CoreRequest) *int {
// 	if cr.RetriesRemaining == nil {
// 		retries := c.MaxRetries
// 		cr.RetriesRemaining = &retries
// 	}
// 	return cr.RetriesRemaining
// }

// func (c *CoreClient) canRetry(cr *CoreRequest) bool {
// 	return *c.getRetriesRemaining(cr) > 0
// }

// func (c *CoreClient) calculateRetryTimeoutSeconds(cr *CoreRequest, res *http.Response) int64 {
// 	if rawRetryAfter := res.Header.Get("retry-after"); len(rawRetryAfter) != 0 {
// 		if retryAfter, err := strconv.ParseInt(rawRetryAfter, 0, 64); err == nil && retryAfter <= 60 {
// 			return retryAfter
// 		}
// 	}

// }

// func (c *CoreClient) retry(ctx context.Context, cr *CoreRequest, res *http.Response) (interface{}, error) {
// 	*cr.RetriesRemaining -= 1

// 	// About the Retry-After header: https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Retry-After
// 	//
// 	// TODO: we may want to handle the case where the header is using the http-date syntax: "Retry-After: <http-date>".
// 	// See https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Retry-After#syntax for details.
// 	retryAfter := int64(-1)
// 	if rawRetryAfter := res.Header.Get("retry-after"); len(rawRetryAfter) != 0 {
// 		if retryAfter, err = strconv.ParseInt(rawRetryAfter, 0, 64); err == nil && retryAfter <= 60 {
// 			return retryAfter
// 		}
// 	}
// }
