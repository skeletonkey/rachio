package rachio

import (
	"encoding/json"
	"fmt"
	"net/url"
)

// GetPersonInfo returns the ID of the owner of the account.
// This is named Info instead of ID to be aligned with the Rachio API
func (c Client) GetPersonInfo() (string, error) {
	fullUrl, err := url.JoinPath(publicURL, "person", "info")
	if err != nil {
		return "", fmt.Errorf("error while attempting to create Info URL: %s", err)
	}
	body, err := c.get(fullUrl)
	if err != nil {
		return "", fmt.Errorf("failed to GET from %s: %s", fullUrl, err)
	}
	var idBody struct {
		ID string `json:"id"`
	}
	err = json.Unmarshal(body, &idBody)
	if err != nil {
		return "", fmt.Errorf("failed to unmarshal body (%v): %s", body, err)
	}

	return idBody.ID, nil
}

// GetPerson retrieves the information for a Person
func (c Client) GetPerson(id string) (Person, error) {
	fullUrl, err := url.JoinPath(publicURL, "person", id)
	if err != nil {
		return Person{}, fmt.Errorf("error while attempting to create Person URL: %s", err)
	}
	body, err := c.get(fullUrl)
	if err != nil {
		return Person{}, fmt.Errorf("failed to GET from %s: %s", fullUrl, err)
	}
	var personData Person
	err = json.Unmarshal(body, &personData)

	return personData, err
}
