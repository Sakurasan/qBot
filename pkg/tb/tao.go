package tb

import (
	"encoding/json"
	"net/http"
	"time"
)

var (
	_taoUrl = "https://api.vvhan.com/api/tao"
)

func Tao() (map[string]string, error) {
	req, _ := http.NewRequest("GET", _taoUrl+"?type=json", nil)
	req.Header.Set("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/88.0.4324.192 Safari/537.36")
	c := http.Client{
		// CheckRedirect: func(req *http.Request, via []*http.Request) error {
		// 	return http.ErrAbortHandler
		// },
		Timeout: 30 * time.Second,
	}
	rsp, err := c.Do(req)
	if err != nil {
		return nil, err
	}

	m := make(map[string]string)
	if err := json.NewDecoder(rsp.Body).Decode(&m); err != nil {
		return nil, err
	}

	return m, nil

}
