package example3

import (
	"io/ioutil"
	"net/http"
)

type Client struct {
	HttpClient *http.Client
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

func Get(url string) (*string, error) {
	client := http.Client{}
	res, errGet := client.Get(url)
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
