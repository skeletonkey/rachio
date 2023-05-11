package rachio

import (
	"encoding/json"
	"net/url"
	"strconv"
)

// GetCurrentSchedule returns the currently running schedule. To determine if there is an actual schedule returned check if Schedule.ScheduleId an empty string or not.
//
// https://rachio.readme.io/reference/publicdeviceidcurrent_schedule
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
//
// https://rachio.readme.io/reference/publicdeviceideventstarttimestarttimeendtimeendtim
func (c Client) GetEvents(deviceId string, start int64, end int64) (events []Event, err error) {
	fullUrl, err := url.JoinPath(publicURL, "device", deviceId, "event")
	if err != nil {
		return events, err
	}
	v := url.Values{}
	v.Add("startTime", strconv.FormatInt(start, 10))
	v.Add("endTime", strconv.FormatInt(end, 10))
	fullUrl += "?" + v.Encode()

	body, err := c.get(fullUrl)
	if err != nil {
		return events, err
	}
	err = json.Unmarshal(body, &events)

	return events, err
}

// GetDeviceInfo returns device information.
//
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

	return device, err
}

// ExpGetDeviceState (EXPERIMENTAL) returns device 'state' information. Among other data, the time/date of next run and last run are returned.
func (c Client) ExpGetDeviceState(deviceID string) (state ExpDeviceState, err error) {
	fullUrl, err := url.JoinPath(experimentalURL, "device", "getDeviceState", deviceID)
	if err != nil {
		return state, err
	}
	body, err := c.get(fullUrl)
	if err != nil {
		return state, err
	}
	err = json.Unmarshal(body, &state)

	return state, err
}
