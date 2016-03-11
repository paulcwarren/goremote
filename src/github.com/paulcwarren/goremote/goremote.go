package goremote

type Transport interface {
	Send(tgt string, message interface{}, reply interface{}) (error)
}
