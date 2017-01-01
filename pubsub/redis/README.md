# Redis Pub/Sub Implementation

This package wraps around Redis Pub/Sub functionality that implements the
`pubsub` package interfaces.

## Configuration

This package is configured from a `toml` formatted configuration file and
optional environment variables.

### Config File

The configuration file is loaded at run time, please see the main `README.md`
on config file locations.

``` toml
[redis]
address = "localhost:6379"  # Address of redis server in host:port format
password = "foo" # Optional, remove or leave blank
db = 0 # Optional DB number, remove or leave blank
```

### Environment Variables

Environment variables are read at run time.

```
SCOREBOARD_REDIS_ADDRESS = "localhost:6379"
SCOREBOARD_REDIS_PASSWORD = "foo"
SCOREBOARD_REDIS_DB = "0"
```
