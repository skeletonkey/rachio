package rachio

import (
	"errors"
	"io"
	"net/http"
	"strconv"
	"time"
)

// get does an HTTP get to the provided URL.  It will return the body (as a string), headers (as a map[string]string), and an error if something went wrong.
func (c Client) get(url string) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return []byte{}, err
	}
	c.setHeaders(req)
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return []byte{}, err
	}
	if resp.StatusCode != 200 {
		return []byte{}, errors.New(resp.Status)
	}
	defer func() {
		err := resp.Body.Close()
		if err != nil {
			panic(err)
		}
	}()
	err = c.processHeaders(resp)
	if err != nil {
		return []byte{}, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, err
	}

	return body, nil
}

func (c Client) setHeaders(req *http.Request) {
	req.Header.Add("Authorization", "Bearer "+c.bearerToken)
	req.Header.Add("Content-Type", "application/json")
}

func (c Client) processHeaders(resp *http.Response) error {
	var err error
	if c.RateInfo.Limit, err = strconv.Atoi(resp.Header.Get("X-RateLimit-Limit")); err == nil {
		if c.RateInfo.Remaining, err = strconv.Atoi(resp.Header.Get("X-RateLimit-Remaining")); err == nil {
			t, err := time.Parse(time.RFC3339, resp.Header.Get("X-RateLimit-Reset"))
			if err == nil {
				c.RateInfo.Reset = t.Unix()
			}

		}
	}

	return err
}
