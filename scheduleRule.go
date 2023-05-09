package rachio

import (
	"encoding/json"
	"net/url"
)

// GetScheduleRule returns the information for a schedule
// https://rachio.readme.io/reference/publicscheduleruleid
func (c Client) GetScheduleRule(id string) (rule ScheduleRule, err error) {
	fullUrl, err := url.JoinPath(publicURL, "schedulerule", id)
	if err != nil {
		return rule, err
	}
	body, err := c.get(fullUrl)
	if err != nil {
		return rule, err
	}
	err = json.Unmarshal(body, &rule)

	return rule, err
}

// GetFlexScheduleRule returns information for a flexscheduleRule entity
// https://rachio.readme.io/reference/publicflexscheduleruleid
func (c Client) GetFlexScheduleRule(id string) (rule ScheduleRule, err error) {
	fullUrl, err := url.JoinPath(publicURL, "flexschedulerule", id)
	if err != nil {
		return rule, err
	}
	body, err := c.get(fullUrl)
	if err != nil {
		return rule, err
	}
	err = json.Unmarshal(body, &rule)

	return rule, err
}
