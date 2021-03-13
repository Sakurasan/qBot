package qc

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"

	"github.com/Mrs4s/MiraiGo/client"
	"github.com/Mrs4s/MiraiGo/message"
	"github.com/robfig/cron/v3"
	"github.com/valyala/fasthttp"
)

func Qclock() {
	c := cron.New()
	c.AddFunc("0 0,9-23 * * *", func() {
		for groupCode, clock := range watermap {
			var once sync.Once
			once.Do(func() {
				if clock {
					sm, err := clockImgByUrl(Bot, groupCode, heshuiUrlList[rand.Intn(len(heshuiUrlList)-1)])
					if err != nil {
						return
					}
					Bot.SendGroupMessage(groupCode, sm.Append(message.NewText("\n"+time.Now().Format("2006-01-02 15:04:05"))))
				}
			})
		}
		for groupCode, clock := range ppmap {
			var once sync.Once
			once.Do(func() {
				if clock {
					sm, err := clockImgByUrl(Bot, groupCode, tigangUrl)
					if err != nil {
						return
					}
					Bot.SendGroupMessage(groupCode, sm.Append(message.NewText("\n"+time.Now().Format("2006-01-02 15:04:05"))))
				}
			})
		}
		defer func() {
			if err := recover(); err != nil {
				log.Println(err)
			}
			return
		}()

	})
	c.Start()
}

func clockImgByUrl(c *client.QQClient, groupCode int64, url string) (*message.SendingMessage, error) {
	_, cc := context.WithTimeout(context.Background(), 10*time.Second)
	defer cc()
	sm := message.NewSendingMessage()
	req := &fasthttp.Request{}
	req.Header.Set("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/88.0.4324.192 Safari/537.36")
	req.SetRequestURI(url)
	rsp := &fasthttp.Response{}
	if err := fasthttp.Do(req, rsp); err != nil {
		return nil, err
	}
	img, err := c.UploadGroupImage(groupCode, bytes.NewReader(rsp.Body()))
	if err != nil {
		sm.Append(message.NewText(fmt.Sprintf("上传失败 %s ", err.Error())))
		return sm, nil
	}
	sm.Append(img)
	return sm, nil
}
