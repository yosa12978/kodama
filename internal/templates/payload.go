package templates

type ErrorPayload struct {
	StatusCode int
	Message    string
}

type Payload struct {
	AppName        string
	AppDescription string
	Content        any
}
