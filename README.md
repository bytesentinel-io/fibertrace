<div align="center">
    <img src="./.assets/bytesentinel.png" width="250px" style="margin-left: 10px" />
</div>

<h1 align="center">
  FiberTrace
</h1>

<div align="center">
    <img src="https://img.shields.io/github/downloads/bytesentinel-io/REPO/total?style=for-the-badge" />
    <img src="https://img.shields.io/github/last-commit/bytesentinel-io/REPO?color=%231BCBF2&style=for-the-badge" />
    <img src="https://img.shields.io/github/issues/bytesentinel-io/REPO?style=for-the-badge" />
</div>

<br />

FiberTrace is a logging package for Go applications. It provides a flexible and modular logging solution with support for writing logs to console and file in JSON or text format.

[![ko-fi](https://ko-fi.com/img/githubbutton_sm.svg)](https://ko-fi.com/Z8Z8JPE9P)

# Installation

To use FiberTrace in your Go project, you need to import the package:

```shell
import "github.com/bbenouarets/bytesentinel/fibertrace"
```

Then, run the following command to fetch the package:

```shell
go get -u github.com/bbenouarets/bytesentinel/fibertrace
```

# Usage

## Creating a Logger
To create a logger instance, use the NewLogger function:

```go
logger, err := fibertrace.NewLogger(logFilePath, application, jsonFormat)
if err != nil {
    // Handle error
}
```

- `logFilePath` _(string)_: The path to the log file. Set it to an empty string to disable file logging.
- `application` _(string)_: The name of your application.
- `jsonFormat` _(bool)_: Set it to true to log in JSON format, or false to log in text format.

## Logging Messages
The logger provides three logging methods: `Info`, `Error`, and `Debug`. Use them to log messages of different levels:

```go
logger.Info("This is an info message.")
logger.Error("This is an error message.")
logger.Debug("This is a debug message.")
```

The log messages will be written to the console and/or the log file based on the logger configuration.