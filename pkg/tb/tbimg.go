package tb

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"github.com/tidwall/gjson"
)

var (
	_tbimgUrl = "https://api.66mz8.com/api/rand.tbimg.php"
	// ?format=images
)

func Tbimg() (map[string]string, error) {
	v := url.Values{}
	v.Add("format", "json")
	req, _ := http.NewRequest("GET", _tbimgUrl+"?"+v.Encode(), nil)
	// bb, _ := httputil.DumpRequest(req, true)
	// fmt.Println(string(bb))
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

	byteRsp, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		return nil, err
	}

	m := make(map[string]string)
	m["pic"] = gjson.GetBytes(byteRsp, "pic_url").String()

	return m, nil

}
