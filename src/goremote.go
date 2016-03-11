package goremote

type Transport interface {
	Send(msg string, body []byte) ([]byte, error)
}
