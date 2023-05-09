package rachio

import (
	"encoding/json"
	"net/url"
)

// GetZone returns the information for a zone
// https://rachio.readme.io/reference/publiczoneid
func (c Client) GetZone(zoneId string) (zone ZoneDetail, err error) {
	fullUrl, err := url.JoinPath(publicURL, "zone", zoneId)
	if err != nil {
		return zone, err
	}
	body, err := c.get(fullUrl)
	if err != nil {
		return zone, err
	}
	err = json.Unmarshal(body, &zone)

	return zone, err
}
