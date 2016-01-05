# Pushover Golang Package
A golang package for sending notifications via https://pushover.net.

This library is not intended for production use and was written by it's author as an excerise to learn more about golang. Please don't import it and us it in any projects that you care about.

# An Example
```golang
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

# Useful Links
* Pushover API Documentation: https://pushover.net/api
* Documentation for this package:  https://gowalker.org/github.com/bdenning/pushover
