// Package rachio is a Go interface to the Rachio API.
//
// Currently working on support for all v1 endpoints: https://rachio.readme.io/reference/getting-started
//
// Please refer to the Github Repo for overall coverage and reporting issues: https://github.com/skeletonkey/rachio
//
// # Rate Limiting
//
// https://rachio.readme.io/reference/rate-limiting
//
// The library is written to the current rate limit of 1700 requests per 24 hours - there's no great effort put into code efficiency.
//
// To see how many request are left and when the counter will reset:
//
//	rachioClient := rachioNew.New(context.TODO(), "12345678-abcd-1234-abcd-1234567890ab")
//	fmt.Printf("Limit: %d\nRemaining: %d\nReset: %s\n", rachioClient.RateInfo.Limit, rachioClient.RateInfo.Remaining, rachioClient.RateInfo.Reset)
//
// # Data Objects
//
// https://rachio.readme.io/reference/data-objects
//
// # Example
//
//	rachioClient := rachioNew.New(context.TODO(), "12345678-abcd-1234-abcd-1234567890ab")
//
//	id, err := rachioClient.GetPersonInfo()
package rachio
