## Blynker

Some defenitions:

`Blynker` is a backend for hardware device (called ***Sensor***) and telegram bot (called ***Blynkerbot***). 

`Blynkerbot` - telegram bot, not implemented yet.

`Sensor` is a device (ESP8266 based) ***with one or more detectors*** 
(temperature, light, passive infrared motion sensor etc).

_**Sensor**_ sends data to _**Blynker**_. _**Blynker**_ saves incoming data.
_**Blynkerbot**_ gets data from _**Blynker**_ and shows to user.
At this time  _**Blynker**_ is using InfluxDB as a sensor values storage.

#### Now there are 3 endpoints:
`POST \get_data` - returns last values of sensor 

`POST \set_data` - saves values to storage

`GET \` - checks the status of sensor

#### TODO:

- Dockercompose - server + InfluxDB
- Add storage for sensors credentials
- Add logging middleware
- Add swagger
- Add linters
- Refactor `\get_data` to `\get_data\{sensorID}`
- Add endpoint `\get_sensors`
- Add github actions for auto deployment

- Refactor server to work with more than one sensor

#### DONE:

- Remove CSV repo

