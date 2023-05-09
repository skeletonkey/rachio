package rachio

import (
	"net/http"
)

// Event holds specific 'event' information for a device
type Event struct {
	Category  string `json:"category"`
	DeviceId  string `json:"deviceId"`
	EventDate int64  `json:"eventDate"`
	Hidden    bool   `json:"hidden"`
	IconUrl   string `json:"iconUrl"`
	Id        string `json:"id"`
	SubType   string `json:"subType"`
	Summary   string `json:"summary"`
	Topic     string `json:"topic"`
	Type      string `json:"type"`
}

// Schedule holds information about a specific scheduled run
type Schedule struct {
	CycleCount      int    `json:"cycleCount"`
	Cycling         bool   `json:"cycling"`
	DeviceId        string `json:"deviceId"`
	Duration        int    `json:"duration"`
	DurationNoCycle int    `json:"durationNoCycle"`
	ScheduleId      string `json:"scheduleId"`
	StartDate       int64  `json:"startDate"`
	Status          string `json:"status"`
	TotalCycleCount int    `json:"totalCycleCount"`
	Type            string `json:"type"`
	ZoneDuration    int    `json:"zoneDuration"`
	ZoneId          string `json:"zoneId"`
	ZoneStartDate   int64  `json:"zoneStartDate"`
}

type nameType struct {
	Name string `json:"name"`
}
type customCropType struct {
	nameType
	Coefficient float64 `json:"coefficient"`
}
type customNozzleType struct {
	nameType
	InchesPerHour float64 `json:"inchesPerHour"`
}
type customSlopeType struct {
	nameType
	SortOrder int `json:"sortOrder"`
}

// ZoneDetail holds all relevant details for a given zone
type ZoneDetail struct {
	AvailableWater             float64          `json:"availableWater"`
	CustomCrop                 customCropType   `json:"customCrop"`
	CustomNozzle               customNozzleType `json:"customNozzle"`
	CustomShade                nameType         `json:"customShade"`
	CustomSlope                customSlopeType  `json:"customSlope"`
	CustomSoil                 nameType         `json:"customSoil"`
	DepthOfWater               float64          `json:"depthOfWater"`
	Efficiency                 float64          `json:"efficiency"`
	Enabled                    bool             `json:"enabled"`
	FixedRuntime               int              `json:"fixedRuntime"`
	Id                         string           `json:"id"`
	ImageUrl                   string           `json:"imageUrl"`
	LastWateredDate            int64            `json:"lastWateredDate"`
	ManagementAllowedDepletion float64          `json:"managementAllowedDepletion"`
	MaxRuntime                 int              `json:"maxRuntime"`
	Name                       string           `json:"name"`
	RootZoneDepth              float32          `json:"rootZoneDepth"`
	Runtime                    int              `json:"runtime"`
	RuntimeNoMultiplier        int              `json:"runtimeNoMultiplier"`
	SaturatedDepthOfWater      float64          `json:"saturatedDepthOfWater"`
	ScheduleDataModified       bool             `json:"scheduleDataModified"`
	WateringAdjustmentRuntimes map[int]int      `json:"wateringAdjustmentRuntimes"`
	YardAreaSquareFeet         int              `json:"yardAreaSquareFeet"`
	ZoneNumber                 int              `json:"zoneNumber"`
}

// Zone holds information for a device's Zones. Use the ID to find more details in ZoneDetail
type Zone struct {
	Duration  int    `json:"duration"`
	SortOrder int    `json:"sortOrder"`
	ZoneId    string `json:"zoneId"`
}

// ScheduleRule contains information about the schedule set up on the device
type ScheduleRule struct {
	CycleSoak        bool     `json:"cycleSoak"`
	CycleSoakStatus  string   `json:"cycleSoakStatus"`
	Enabled          bool     `json:"enabled"`
	EndDate          int64    `json:"endDate"`
	EtSkip           bool     `json:"etSkip"`
	ExternalName     string   `json:"externalName"`
	Id               string   `json:"id"`
	Name             string   `json:"name"`
	Operator         string   `json:"operator"`
	ScheduleJobTypes []string `json:"scheduleJobTypes"`
	StartDate        int64    `json:"startDate"`
	StartDay         int      `json:"startDay"`
	StartMonth       int      `json:"startMonth"`
	StartYear        int      `json:"startYear"`
	Summary          string   `json:"summary"`
	TotalDuration    int      `json:"totalDuration"`
	Type             string   `json:"type"`
	Zones            []Zone   `json:"zones"`
}

// Device contains information about the device attached to the account
type Device struct {
	CreateDate              int64          `json:"createDate"`
	Deleted                 bool           `json:"deleted"`
	Elevation               float64        `json:"elevation"`
	FlexScheduleRules       []ScheduleRule `json:"flexScheduleRules"`
	HomeKitCompatible       bool           `json:"homeKitCompatible"`
	Id                      string         `json:"id"`
	Latitude                float64        `json:"latitude"`
	Longitude               float64        `json:"longitude"`
	MacAddress              string         `json:"macAddress"`
	Model                   string         `json:"model"`
	Name                    string         `json:"name"`
	On                      bool           `json:"on"`
	Paused                  bool           `json:"paused"`
	RainDelayExpirationDate int64          `json:"rainDelayExpirationDate"`
	RainDelayStartDate      int64          `json:"rainDelayStartDate"`
	RainSensorTripped       bool           `json:"rainSensorTripped"`
	ScheduleModeType        string         `json:"scheduleModeType"`
	ScheduleRules           []ScheduleRule `json:"scheduleRules"`
	SerialNumber            string         `json:"serialNumber"`
	Status                  string         `json:"status"`
	TimeZone                string         `json:"timeZone"`
	UtcOffset               int            `json:"utcOffset"`
	Webhooks                []Webhook      `json:"webhooks"`
	Zip                     string         `json:"zip"`
	Zones                   []ZoneDetail   `json:"zones"`
}

// Person contains information about the owner of the account
type Person struct {
	Email      string `json:"email"`
	Id         string `json:"id"`
	Username   string `json:"username"`
	Devices    []Device
	CreateDate int64  `json:"createDate"`
	Deleted    bool   `json:"deleted"`
	FullName   string `json:"fullName"`
}

// Webhook appears to be deprecated - keeping it here for completeness
type Webhook struct {
	Category       string `json:"category"`
	CreateDate     int64  `json:"createDate"`
	DeviceId       string `json:"deviceId"`
	EventDate      int64  `json:"eventDate"`
	ExternalId     string `json:"externalId"`
	Hidden         bool   `json:"hidden"`
	Id             string `json:"id"`
	LastUpdateDate int64  `json:"lastUpdateDate"`
	SubType        string `json:"subType"`
	Summary        string `json:"summary"`
	Type           string `json:"type"`
}

// RateInfo holds the 'rate' information for the current client
type RateInfo struct {
	Limit     int   // 'Daily' limit - should by 1700
	Remaining int   // remaining request available until reset
	Reset     int64 // reset time - epoch format
}

// Client holds everything required to an instance of Rachio
type Client struct {
	bearerToken string
	RateInfo    *RateInfo
	httpClient  *http.Client
}
