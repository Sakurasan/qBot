package appso

import (
	"crypto/tls"
	"encoding/xml"
	"net"
	"net/http"
	"time"
)

var (
	_url = "http://rsshub.ioiox.com/appstore/xianmian"
)

type Rss struct {
	XMLName xml.Name `xml:"rss"`
	Text    string   `xml:",chardata"`
	Atom    string   `xml:"atom,attr"`
	Version string   `xml:"version,attr"`
	Channel struct {
		Text  string `xml:",chardata"`
		Title string `xml:"title"`
		Link  struct {
			Text string `xml:",chardata"`
			Href string `xml:"href,attr"`
			Rel  string `xml:"rel,attr"`
			Type string `xml:"type,attr"`
		} `xml:"link"`
		Description   string `xml:"description"`
		Generator     string `xml:"generator"`
		WebMaster     string `xml:"webMaster"`
		Language      string `xml:"language"`
		LastBuildDate string `xml:"lastBuildDate"`
		Ttl           string `xml:"ttl"`
		Item          []It   `xml:"item"`
	} `xml:"channel"`
}
type It struct {
	Text        string `xml:",chardata"`
	Title       string `xml:"title"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
	Guid        struct {
		Text        string `xml:",chardata"`
		IsPermaLink string `xml:"isPermaLink,attr"`
	} `xml:"guid"`
	Link string `xml:"link"`
}

func XianMian() ([]It, error) {
	req, err := http.NewRequest("GET", _url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/88.0.4324.192 Safari/537.36")
	tr := &http.Transport{
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
			DualStack: true,
		}).DialContext,
		ForceAttemptHTTP2:     true,
		MaxIdleConns:          100,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
		TLSClientConfig:       &tls.Config{InsecureSkipVerify: true},
	}
	// tr.Proxy = http.ProxyURL(func() *url.URL {
	// 	purl, _ := url.Parse("http://127.0.0.1:9000")
	// 	return purl
	// }())
	client := http.Client{
		Transport: tr,
		Timeout:   30 * time.Second,
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	var rss Rss
	xml.NewDecoder(resp.Body).Decode(&rss)
	if err != nil {
		return nil, err
	}
	return rss.Channel.Item, nil
}
