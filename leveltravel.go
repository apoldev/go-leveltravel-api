package leveltravel

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type LevelTravelApi struct {
	apikey  string
	log     LoggerInterface
	Fetcher DataFetcher
}

type LoggerInterface interface {
	Debug(...interface{})
}

type DataFetcher interface {
	Fetch(url string, headers map[string]string) (resp io.ReadCloser, err error)
}

type HttpDataFetcher struct{}

func (h *HttpDataFetcher) Fetch(url string, headers map[string]string) (resp io.ReadCloser, err error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return
	}

	for k, v := range headers {
		req.Header.Add(k, v)
	}

	client := &http.Client{}
	httpResp, err := client.Do(req)
	if err != nil {
		return
	}

	resp = httpResp.Body
	return
}

func (l *LevelTravelApi) SetLogger(logger LoggerInterface) {
	l.log = logger
}

func (l *LevelTravelApi) getJson(path string, args map[string]string, v interface{}) error {
	apiUrl, err := url.Parse("https://api.level.travel" + path)
	if err != nil {
		return err
	}
	params := url.Values{}
	for k, v := range args {
		if v == "" {
			continue
		}
		params.Add(k, v)
	}

	apiUrl.RawQuery = params.Encode()

	if l.log != nil {
		l.log.Debug("API Send: " + apiUrl.String())
	}

	resp, err := l.Fetcher.Fetch(apiUrl.String(), map[string]string{
		"Accept":        "application/vnd.leveltravel.v3",
		"Authorization": fmt.Sprintf("Token token=\"%s\"", l.apikey),
	})
	if err != nil {
		return err
	}
	defer resp.Close()
	return json.NewDecoder(resp).Decode(&v)
}

func NewLevelTravelApi(apikey string) *LevelTravelApi {
	return &LevelTravelApi{
		apikey:  apikey,
		Fetcher: &HttpDataFetcher{},
	}
}
