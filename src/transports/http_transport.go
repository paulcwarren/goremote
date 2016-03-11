package remhttp

import (
	"bytes"
	"net/http"

	"fmt"
	"github.com/tedsuo/rata"
	"io/ioutil"

	goremote "../"
)

type HttpTransport struct {
	HttpClient *http.Client
	reqGen     *rata.RequestGenerator
}

func NewHttpTransport(url string, routes rata.Routes) goremote.Transport {
	return &HttpTransport{
		HttpClient: &http.Client{},
		reqGen:     rata.NewRequestGenerator(url, routes),
	}
}

func (t *HttpTransport) Send(msg string, body []byte) ([]byte, error) {

	request, err := t.reqGen.CreateRequest(msg, nil, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}

	response, err := t.HttpClient.Do(request)
	if err != nil {
		return nil, err
	}

	reply, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	if response.StatusCode == 500 { // and possibly 4xx codes as well
		return reply, fmt.Errorf("HTTP Transport error")
	}

	return reply, nil
}
