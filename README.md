# Pushover Golang Package

[![Documentation](https://godoc.org/github.com/invisiblethreat/go-pushover?status.svg)](https://godoc.org/github.com/invisiblethreat/go-pushover) [![Build Status](https://drone.io/github.com/invisiblethreat/go-pushover/status.png)](https://drone.io/github.com/invisiblethreat/go-pushover/latest) [![Coverage Status](https://coveralls.io/repos/invisiblethreat/go-pushover/badge.svg?branch=master&service=github)](https://coveralls.io/github/invisiblethreat/go-pushover?branch=master)

A Golang package for sending notifications via https://api.pushover.net.

This library is a mashup of https://github.com/bdenning/go-pushover and
https://github.com/gregdel/pushover. Neither did exactly what I wanted, in the
way that I wanted.

# Package Example
You can use the pushover package within your golang applications as follows:
```Go
// Set your pushover API keys
token := "KzGDORePKggMaC0QOYAMyEEuZJnyUi"
user := "e9e1495ec75826de5983cd1abc8031"

// Send your message
m := pushover.NewMessage(token, user)
m.Push("Server exchange01.example.net is in a critical state.")
```

# Command Line Tool
A command line tool is provided under cmd/pushover. Build and install the command using:
```Shell
$ make install
```
Before using the command line tool, you must first set the following environment variables.
```Shell
$ export PUSHOVER_TOKEN="KzGDORePKggMaC0QOYAMyEEuZJnyUi"
$ export PUSHOVER_USER="e9e1495ec75826de5983cd1abc8031"
```
Then messages can be sent by piping output to the pushover command.
```Shell
$ echo "Server exchange01.example.net is in a critical state" | pushover
```
