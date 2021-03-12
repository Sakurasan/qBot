package tts

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"regexp"

	"github.com/axgle/mahonia"
)

var (
	transReg = regexp.MustCompile(`"(.*?)"`)
	_url     = "https://translate.googleapis.com/translate_a/single?client=gtx&sl=%s&tl=%s&dt=t&q=%s"
)

type transType struct {
	Source string
	Target string
	Text   string
}

func NewTransT() *transType {
	return &transType{
		Source: "auto",
		Target: "zh-CN",
	}
}
func Trans(t *transType) string {
	if t.Text == "" {
		return ""
	}
	req, err := http.NewRequest("GET", fmt.Sprintf(_url, url.QueryEscape(t.Source), url.QueryEscape(t.Target), url.QueryEscape(t.Text)), nil)
	if err != nil {
		log.Println(err)
		return ""
	}
	req.Header.Set("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/88.0.4324.192 Safari/537.36")
	rsp, err := http.DefaultClient.Do(req)
	if err != nil {
		return ""
	}
	byteRsp, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		return ""
	}
	result := regexp.MustCompile(`"(.*?)"`).FindStringSubmatch(string(byteRsp))
	if len(result) > 1 {
		return result[1]
	}
	return ""
}

func toutf8(s string) string {
	enc := mahonia.NewEncoder("utf8")
	return enc.ConvertString(s)
}
