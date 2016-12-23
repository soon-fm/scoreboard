# SOON\_ FM Scoreboard Configuration

Configuration can be loaded from 3 sources in the following order of presedence:

* Config file (`toml` formatted)
* Environment Variables
* Command Line Flags

## 1. From File

File configuration values are loaded first, this file must be `toml` formated,
by default the following directories will be searched for a `config.toml` file:

* `/etc/scoreboard`
* `$HOME/.config/scoreboard`

Please see the example configuration file in this directory.

## 2. Environment Variables

The following environment variables can be used to override file configurations.

* `SCOREBOARD_LOG_LEVEL`: Set the logging verbosity
* `SCOREBOARD_LOG_LOGFILE`: Path to log file
* `SCOREBOARD_LOG_FORMAT`: Set the logging format of each log entry

## 3. CLI Flags

Some command line flags allow you to override file and environment variable
configuration options:

* `-l/--log-level`: Set the logging verbosity
