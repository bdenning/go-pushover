# Pushover Golang Package

[![Documentation](https://godoc.org/github.com/bdenning/pushover?status.svg)](https://godoc.org/github.com/bdenning/pushover)

A golang package for sending notifications via https://api.pushover.net.

This library is **not intended for production use** and was written by its author as an exercise to learn more about golang. Please don't import it and use it in any projects that you care about.

# Package Example
You can use the pushover package within your golang applications as follows:
```Go
// Set your pushover api keys (these are examples)
token := "KzGDORePKggMaC0QOYAMyEEuZJnyUi"
user := "e9e1495ec75826de5983cd1abc8031"
device := "test_device"
title := "Alert"

// Send your message
m := pushover.NewMessage(token, user, device, title)
m.Push("Server exchange01.example.net is in a critical state.")
```
_Note: The pushover.Message struct also implements io.Writer for convenience sake._

# Command Line Tool
A command line tool is provided under cmd/pushover. Build and install the command using:
```Shell
$ make install
```
Before using the command line tool, you must first set the following environment variables.
```Shell
$ export PUSHOVER_TOKEN="KzGDORePKggMaC0QOYAMyEEuZJnyUi"
$ export PUSHOVER_USER="e9e1495ec75826de5983cd1abc8031"
$ export PUSHOVER_DEVICE="test_device"
```
Then messages can be sent by piping output to the pushover command.
```Shell
$ echo "Server exchange01.example.net is in a critcal state" | pushover
```

It is not currently possible to provide a title. At present the title will be set to the hostname of the computer it was sent from.
