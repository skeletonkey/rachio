// Package rachio is a Go interface to the Rachio API.
//
// Currently working on support for all v1 endpoints: https://rachio.readme.io/reference/getting-started
//
// # Support
//
// Please refer to the Github Repo for overall coverage and reporting issues: https://github.com/skeletonkey/rachio
//
// If you encounter any issues using this please provide a Pull Request or submit an Issue.
//
// # Rate Limiting
//
// https://rachio.readme.io/reference/rate-limiting
//
// The library is written to the current rate limit of 1700 requests per 24 hours - there's no great effort put into code efficiency.
//
// To see how many request are left and when the counter will reset:
//
//	rachioClient := rachio.New(context.TODO(), "12345678-abcd-1234-abcd-1234567890ab")
//
//	fmt.Printf("Limit: %d\nRemaining: %d\nReset: %s\n", rachioClient.RateInfo.Limit, rachioClient.RateInfo.Remaining, rachioClient.RateInfo.Reset)
//
// # Data Objects
//
// https://rachio.readme.io/reference/data-objects
//
// # Example
//
//	rachioClient := rachio.New(context.TODO(), "12345678-abcd-1234-abcd-1234567890ab")
//
//	id, err := rachioClient.GetPersonInfo()
//
// # Experimental Commands
//
// These are commands that do *not* have documentation and have not been provided with any type of SLA by Rachio. They have been 'discovered' by analysing the Rachio website.
//
// All methods that are 'Experimental' start with 'Exp'.
//
// Please use these commands cautiously!
package rachio
