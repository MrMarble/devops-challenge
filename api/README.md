# Air quality API

API REST in Go that returns a JSON containing data from air quality measurements with the following structure:

```json
{
    data: [
        {
            "timestamp":"2016-10-01 00:00:00.004",
            "id_entity":"aq_salvia",
            "so2":6.80117094260474,
            "no2":48.398337879833704,
            "co":0.657363926741451,
            "o3":48.49706558445371,
            "pm10":20.1015302324903,
            "pm2_5":9.137353903174679
        }
        ...
    ]
}
```

The API also serves as a entry point for data, you can pass a path to a csv file that contains the measurements and it will insert them to the database provided by a environment variable.

## Required Environment variables

```bash
POSTGRES_HOST
POSTGRES_PORT
POSTGRES_USER
POSTGRES_PASSWORD
POSTGRES_DB
GIN_MODE # Switch to "release" mode in production. Default "debug"
```
