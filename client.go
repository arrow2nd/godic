package godic

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func (g *Godic) client(url string, params url.Values, result interface{}) error {
	req, _ := http.NewRequest("GET", "https://api.codic.jp/"+url, nil)
	req.Header.Add("Authorization", "Bearer "+g.apiKey)

	if params != nil {
		req.URL.RawQuery = params.Encode()
	}

	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		return fmt.Errorf("access failed (%s)", res.Status)
	}

	bytes, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(bytes, result); err != nil {
		return err
	}

	return nil
}
