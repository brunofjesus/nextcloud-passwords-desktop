package ncpasswords

import (
	"io/ioutil"
	"net/http"
	"time"
)

func DoRequest(credentials Credentials, method string, path string, body *string) ([]byte, error) {
	url := credentials.Url + "/apps/passwords/api/1.0/" + path

	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(credentials.Username, credentials.Password)

	httpClient := &http.Client{
		Timeout: time.Second * 30,
	}

	res, err := httpClient.Do(req)

	if err != nil {
		return nil, err
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	responseBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return responseBody, err
}
