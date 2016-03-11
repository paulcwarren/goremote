// Another im
package remhttp

import (
	goremote "../"
)

type GrpcTransport struct {
	//todo
}

func NewGrpcTransport() goremote.Transport {
	return &GrpcTransport{}
}

func (t *HttpTransport) Send(msg string, body []byte) ([]byte, error) {

	// todo

	return nil, nil
}
