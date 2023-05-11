package rachio

import (
	"context"
	"net/http"
	"time"
)

const defaultRateLimit = 1700

// These are hardcoded instead of in a config file as they shouldn't change
// often and to reduce the need for more dependencies/complications.
const experimentalURL = "https://cloud-rest.rach.io"
const publicURL = "https://api.rach.io/1/public"

// New returns a Rachio client allowing for interaction with the Rachio API.
// Context is currently ignored.  It is here for future implementation to be a good citizen within larger programs.
// Token is retrieved from Rachio: https://rachio.readme.io/reference/authentication
func New(ctx context.Context, token string) Client {
	return Client{
		token,
		&RateInfo{
			Limit:     defaultRateLimit,
			Remaining: defaultRateLimit,
			Reset:     time.Now().Unix(),
		},
		&http.Client{},
	}
}
