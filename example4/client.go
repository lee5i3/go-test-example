package example4

import (
	"io/ioutil"
	"net/http"
)

type IClientInterface interface {
	Get(url string) (resp *http.Response, err error)
}

type Client struct {
	HttpClient IClientInterface
}

func (client Client) Get(url string) (*string, error) {
	res, errGet := client.HttpClient.Get(url)
	if errGet != nil {
		return nil, errGet
	}

	defer res.Body.Close()

	body, errRead := ioutil.ReadAll(res.Body)
	if errRead != nil {
		return nil, errRead
	}

	result := string(body)

	return &result, nil
}
