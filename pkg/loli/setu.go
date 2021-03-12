package loli

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"qBot/tests"

	"github.com/tidwall/gjson"
)

var (
	loliApiKey = tests.LoliApiKey
	_url       = "https://api.lolicon.app/setu/?num=10&apikey="
)

func SetuReq() ([]gjson.Result, error) {
	req, _ := http.NewRequest("GET", _url+loliApiKey, nil)
	req.Header.Set("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/88.0.4324.192 Safari/537.36")

	rsp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	bytersp, _ := ioutil.ReadAll(rsp.Body)
	var datalist = []gjson.Result{}
	if gjson.ValidBytes(bytersp) {
		if code := gjson.GetBytes(bytersp, "code").Int(); code != 0 {
			return nil, errors.New(fmt.Sprintf("error Code:%d", code))
		}
		datalist = gjson.GetBytes(bytersp, "data.#.url").Array()

	}
	return datalist, nil
}

func SetuOne() (string, error) {
	_url = "https://api.lolicon.app/setu/?num=1&apikey="
	req, _ := http.NewRequest("GET", _url+loliApiKey, nil)
	req.Header.Set("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/88.0.4324.192 Safari/537.36")

	rsp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	bytersp, _ := ioutil.ReadAll(rsp.Body)
	var setuUrl string
	if gjson.ValidBytes(bytersp) {
		if code := gjson.GetBytes(bytersp, "code").Int(); code != 0 {
			return "", errors.New(fmt.Sprintf("error Code:%d", code))
		}
		setuUrl = gjson.GetBytes(bytersp, "data.0.url").String()

	}
	return setuUrl, nil
}
