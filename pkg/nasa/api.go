package nasa

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const url = "https://api.nasa.gov/planetary/apod?api_key=%s&date=%s"
const DateFormat = "2006-01-02"

type ApodImage struct {
	Copyright      string `json:"copyright,omitempty"`
	Date           string `json:"date,omitempty"`
	Explanation    string `json:"explanation,omitempty"`
	HdUrl          string `json:"hd_url,omitempty"`
	ServiceVersion string `json:"service_version,omitempty"`
	Title          string `json:"title,omitempty"`
	Url            string `json:"url,omitempty"`
}

func PullImage(key string) (*ApodImage, error) {
	now := time.Now().Format(DateFormat)
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf(url, key, now), nil)
	if err != nil {
		return nil, fmt.Errorf("error creatin http request to nasa api:%w", err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending http request to nasa api:%w", err)
	}

	defer func() {
		_ = resp.Body.Close()
	}()

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading image content:%w", err)
	}
	println(string(content))
	image := new(ApodImage)
	if err = json.Unmarshal(content, image); err != nil {
		return nil, err
	}

	return image, nil
}
