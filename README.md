### This is backend for hardware device (called sensor) and telegram bot.
#### Now there is 2 endpoints:
`/get_data` - returns last values of sensor 

`/set_data` - save values to storage

Swagger is gonna be soon.

#### TODO:

Dockercompose - server + InfluxDB

Storage for sensors credentials

Remove CSV repo

Add github actions for auto deployment

Refactor server to work with more than one sensor
