package loli

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/tidwall/gjson"
)

var (
	_mzurl = "https://api.func.ws/api/img/mz?format=json"
)

func Mzt() (string, error) {
	req, _ := http.NewRequest("GET", _mzurl, nil)
	req.Header.Set("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/88.0.4324.192 Safari/537.36")

	rsp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	bytersp, _ := ioutil.ReadAll(rsp.Body)
	var setuUrl string
	if gjson.ValidBytes(bytersp) {
		if code := gjson.GetBytes(bytersp, "code").Int(); code != 200 {
			return "", errors.New(fmt.Sprintf("error Code:%d", code))
		}
		setuUrl = gjson.GetBytes(bytersp, "data.img").String()

	}
	return setuUrl, nil
}
