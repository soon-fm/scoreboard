# Logger Package

The `logger` package provides a global logger setup at application
initialisation.

You can also instanciate your own speerate logger instances if needed.

## Configuration

The logger supports configuration via a `Config` interface. By default this is loadded
from:

* A `toml` formatted file
* Environment Variables
* CLI Flags

### Config File

``` toml
[log]
level = "info" # Logging verbosity (debug, info, warn, error)
logfile = "/path/to/file.log" # Absolute path to log file
format = "json" # Logging format (text, json, logstash)
console_output = true # Enable or disable console log output

# Logstash Configuration
# Only used if log.format == "logstash"
[logstash]
type = "foo" # Override logstash type
```

###  Environment Variables

The following environment variables can be used to override file configurations.

* `SCOREBOARD_LOG_LEVEL`: Set the logging verbosity
* `SCOREBOARD_LOG_LOGFILE`: Path to log file
* `SCOREBOARD_LOG_FORMAT`: Set the logging format of each log entry

### CLI Flags

Some command line flags allow you to override file and environment variable
configuration options:

* `-l/--log-level`: Set the logging verbosity

## Usage

### Simple

Simply call the `logger` with the designed log level method:

``` go
import "scoreboard/logger"

func main() {
    logger.Debug("a debug message")
    logger.Info("a info message")
    logger.Watn("a warning message")
    logger.Error("an error message")
    logger.Fatal("a fatal message, exits the application")
    logger.Panic("a panic message, panics the application")
}
```

### With Context

Context can also be logged by using the `WithField`, `WithFields`, and `WithError`
methods.

``` go
import "scoreboard/logger"

func main() {
    logger.WithField("foo", "bar").Debug("a debug message")
    logger.WithFields(logger.F{
        "foo": "bar",
        "fizz": "buzz",
    }).Debug("a debug message")
    logger.WithError(errors.New("an error")).Error("an error message")
}
```
