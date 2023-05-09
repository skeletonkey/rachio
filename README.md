# rachio

Rachio API library

A smart watering device.

This is a Go implementation of the API for better control of the Rachio device through the API that Rachio offers.

## Purpose

The purpose of this library is to get to 100% coverage. Having said that my priorities are:

* GET commands
* properly represent the data returned

I'm happy to add DELETE/POST/PUT commands; I ask that if it is requested you agree to test the command before it is added to the source code.

There is a concerted effort to properly expose all data returned via the API; however, the documentation is not 100%.  If you find a discrepancy please create an Issue.

## Commands

### v1 - Summary

|  Type  | Complete |
|:------:|:---------|
| DELETE | 0%       |
|  GET   | 67%      | 
|  POST  | 0%       | 
|  PUT   | 0%       | 
| TOTAL  | 30%      |

### v1 - Endpoints

|     Category     | Method | End Point                                                      | Supported |
|:----------------:|:------:|:---------------------------------------------------------------|:---------:|
|      Person      |  GET   | /public/person/info                                            |     ✅     |
|                  |  GET   | /public/person/:id                                             |     ✅     |
|      Device      |  GET   | /public/device/:id/current_schedule                            |     ✅     |
|                  |  GET   | /public/device/:id/event?startTime=:startTime&endTime=:endTime |     ✅     |
|                  |  GET   | /public/device/:id/forecast?units=:units                       |     ❌     |
|                  |  PUT   | /public/device/stop_water                                      |     ❌     |
|                  |  PUT   | /public/device/rain_delay                                      |     ❌     |
|                  |  PUT   | /public/device/on                                              |     ❌     |
|                  |  PUT   | /public/device/off                                             |     ❌     |
|                  |  PUT   | /public/device/pause_zone_run                                  |     ❌     |
|                  |  PUT   | /public/device/resume_zone_run                                 |     ❌     |
|                  |  GET   | /public/device/:id                                             |     ✅     |
|       Zone       |  GET   | /public/zone/:id                                               |     ✅     |
|                  |  PUT   | /public/zone/start                                             |     ❌     |
|                  |  PUT   | /public/zone/start_multiple                                    |     ❌     |
|                  |  PUT   | /public/zone/setMoisturePercent                                |     ❌     |
|                  |  PUT   | /public/zone/setMoistureLevel                                  |     ❌     |
|   ScheduleRule   |  GET   | /public/schedulerule/:id                                       |     ✅     |
|                  |  PUT   | /public/schedulerule/skip                                      |     ❌     |
|                  |  PUT   | /public/schedulerule/start                                     |     ❌     |
|                  |  PUT   | /public/schedulerule/seasonal_adjustment                       |     ❌     |
| FlexScheduleRule |  GET   | /public/flexschedulerule/:id                                   |     ✅     |
|   Notification   |  GET   | /public/notification/webhook_event_type                        |     ❌     |
|                  |  GET   | /public/notification/:deviceId/webhook                         |     ❌     |
|                  |  GET   | /public/notification/webhook/:id                               |     ❌     |
|                  |  PUT   | /public/notification/webhook                                   |     ❌     |
|                  | DELETE | /public/notification/webhook/:id                               |     ❌     |
|                  |  POST  | /public/notification/webhook                                   |     ❌     |

*NOTES*

* `forecast` will probably not be implemented.  It's a different data structure and if you need weather data you should probably look to a specific weather app.