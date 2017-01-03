# SOON\_ FM Scoreboard

Stores scores for each user based on system events into InfluxDB.

## Scoring

The following events reward the following score values:

* Playing a Track: +1 Point
* Track Skip: -1 Point

## Configuration

Configuration can be loaded from 3 sources in the following order of presedence:

* Config file (`toml` formatted)
* Environment Variables
* Command Line Flags

### 1. From File

File configuration values are loaded first, this file must be `toml` formated,
by default the following directories will be searched for a `config.toml` file:

* `/etc/scoreboard`
* `$HOME/.config/scoreboard`

Please see the example configuration file in this directory.

``` toml
# Logging Configuration
[log]
level = "info" # Logging verbosity (debug, info, warn, error)
logfile = "/path/to/file.log" # Absolute path to log file
format = "json" # Logging format (text, json, logstash)
console_output = true # Enable or disable console log output

# Logstash Configuration
# Only used if log.format is set to "logstash"
[logstash]
type = "foo" # Override logstash type

# Redis Connection Configuration
[redis]
address = "localhost:6379"  # Address of redis server in host:port format
password = "foo" # Optional, remove or leave blank
db = 0 # Optional DB number, remove or leave blank

# Postgres DB Configuration
[db]
host = "localhost" # Datbabase host
port = 5432 # Optional Datbabase port
db = "myDb" # Required DB Name
username = "username" # Optional Username - omit of not required
password = "password" # Optional Password - omit of not required
migration_path = "/path/to/migrations" # Optional Datbase migration path
```

### 2. Environment Variables

The following environment variables can be used to override file configurations.

#### Logging

* `SCOREBOARD_LOG_LEVEL`: Set the logging verbosity
* `SCOREBOARD_LOG_LOGFILE`: Path to log file
* `SCOREBOARD_LOG_FORMAT`: Set the logging format of each log entry

#### Redis Pub/Sub

* `SCOREBOARD_REDIS_ADDRESS`: Redis server address in `host:port` format
* `SCOREBOARD_REDIS_PASSWORD`: Password for Redis server
* `SCOREBOARD_REDIS_DB`: Redis DB Number

#### PostgreSQL DB

* `SCOREBOARD_DB_HOST`: DB Host
* `SCOREBOARD_DB_PORT`: DB Port
* `SCOREBOARD_DB_DB`: DB Name
* `SCOREBOARD_DB_USERNAME`: Username for DB
* `SCOREBOARD_DB_PASSWORD`: Password for DB
* `SCOREBOARD_DB_MIGRATION_PATH`: Database migration path

### 3. CLI Flags

Some command line flags allow you to override file and environment variable
configuration options:

* `-l/--log-level`: Set the logging verbosity
