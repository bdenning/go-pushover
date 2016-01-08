package pushover

import "errors"

var (
	// ErrHTTPStatus is returned when pushover.net response with a non HTTP 200 response code.
	ErrHTTPStatus = errors.New("Recieved a non HTTP 200 (OK) response from the pushover.net API")
	// ErrUnknown is used to indicate that an error has occurred but its underlying cause could not be determined.
	ErrUnknown = errors.New("Recieved a status code indicating an error but did not receive an error message from pushover.net")
)
