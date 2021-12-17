package codic

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func (c *Codic) client(url string, params url.Values) ([]byte, error) {
	req, _ := http.NewRequest("GET", "https://api.codic.jp/"+url, nil)
	req.Header.Add("Authorization", "Bearer "+c.apiKey)

	if params != nil {
		req.URL.RawQuery = params.Encode()
	}

	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		return nil, fmt.Errorf("access failed (%s)", res.Status)
	}

	return io.ReadAll(res.Body)
}
