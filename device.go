package rachio

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
)

// GetCurrentSchedule returns the currently running schedule. To determine if there is an actual schedule returned check if Schedule.ScheduleId an empty string or not.
func (c Client) GetCurrentSchedule(deviceID string) (schedule Schedule, err error) {
	fullUrl, err := url.JoinPath(publicURL, "device", deviceID, "current_schedule")
	if err != nil {
		return schedule, err
	}
	body, err := c.get(fullUrl)
	if err != nil {
		return schedule, err
	}
	err = json.Unmarshal(body, &schedule)

	return schedule, err
}

// GetEvents returns events for a device between end and start date (both millisecond epoch time)
func (c Client) GetEvents(deviceId string, start int64, end int64) (events []Event, err error) {
	fullUrl, err := url.JoinPath(publicURL, "device", deviceId, "event")
	if err != nil {
		return events, err
	}
	v := url.Values{}
	v.Add("startTime", strconv.FormatInt(start, 10))
	v.Add("endTime", strconv.FormatInt(end, 10))
	fullUrl += "?" + v.Encode()
	fmt.Printf("URL: %s\n", fullUrl)

	body, err := c.get(fullUrl)
	if err != nil {
		return events, err
	}
	fmt.Println(string(body))
	err = json.Unmarshal(body, &events)

	return events, err
}

// GetDeviceInfo returns device information.
// https://rachio.readme.io/reference/publicdeviceid
func (c Client) GetDeviceInfo(deviceID string) (device Device, err error) {
	fullUrl, err := url.JoinPath(publicURL, "device", deviceID)
	if err != nil {
		return device, err
	}
	body, err := c.get(fullUrl)
	if err != nil {
		return device, err
	}
	err = json.Unmarshal(body, &device)
	fmt.Printf("Body: %s\n\n\n", body)

	return device, err
}
