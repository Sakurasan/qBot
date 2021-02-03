package qchan

import (
	"log"
	"net/http"
	"net/url"
	"strings"
)

var (
	key       = "49965840c480ba7ed1dfaa5723dd2bda"
	qmsgurl   = "https://qmsg.zendee.cn/send/"
	qgroupurl = "https://qmsg.zendee.cn/group/"
)

func Send(msg string) {
	v := url.Values{}
	v.Add("msg", msg)
	req, _ := http.NewRequest(http.MethodPost, qmsgurl+key, strings.NewReader(v.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	_, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err)
	}

}

func SendGroup(msg string, qq string) {
	v := url.Values{}
	v.Add("msg", msg)
	v.Add("qq", qq)
	req, _ := http.NewRequest(http.MethodPost, qgroupurl+key, strings.NewReader(v.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	_, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err)
	}

}
