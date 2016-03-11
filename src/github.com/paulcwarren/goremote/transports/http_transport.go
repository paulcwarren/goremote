package transports

import (
	"encoding/json"
	"bytes"
	"net/http"

	"fmt"
	"github.com/tedsuo/rata"
	"io/ioutil"

	"github.com/paulcwarren/goremote"
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

func (t *HttpTransport) Send(tgt string, message interface{}, reply interface{}) (error) {

	body, err := json.Marshal(message)
	if err != nil {
		return err
	}

	request, err := t.reqGen.CreateRequest(tgt, nil, bytes.NewReader(body))
	if err != nil {
		return err
	}

	response, err := t.HttpClient.Do(request)
	if err != nil {
		return err
	}

	replyBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}

	if response.StatusCode == 500 { // and possibly 4xx codes as well
		return fmt.Errorf("HTTP Transport error")
	} else {
		err = json.Unmarshal(replyBytes, response)
		if err != nil {
			return err
		}
	}

	return nil
}
