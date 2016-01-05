# Pushover Golang Package
A golang package for sending notifications via https://pushover.net.

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

# Commandline Tool
A command line tool is provided under cmd/pushover. Build and install the command using:
```
$ make install
```
Before using the commandline too, you must first set the following environment variables.
```
$ export PUSHOVER_TOKEN="[your device token]"
$ export PUSHOVER_USER="[your user token]"
$ export PUSHOVER_DEVICE="[your device id]"
```
Then messages can be sent by piping output to the pushover command.
```
$ echo "Server exchange01.example.net is in a critcal state" | pushover
```

It is not currently possible to provide a title. At present the title will be set to the hostname of the computer it was sent from.

# Useful Links
* Pushover API Documentation: https://pushover.net/api
* Documentation for this package:  https://gowalker.org/github.com/bdenning/pushover
