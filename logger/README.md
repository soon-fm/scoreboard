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

The `config` package has more information on how to configure the `logger` and
other packages.

## Global Usage

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

### Logging Context

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
