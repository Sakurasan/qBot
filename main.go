package main

import (
	"log"
	"os"
	"os/signal"
	"qBot/pkg/bot"
	"qBot/pkg/config"
	"qBot/pkg/qc"
	"time"

	"github.com/gin-gonic/gin"
)

// var msgchan = make(chan string, 100)

func main() {
	config.Init()
	qc.Init()
	qc.Login()

	go func() {
		for {
			bot.Run()
			time.Sleep(15 * time.Second)
		}
	}()

	r := gin.New()

	r.GET("/debug", func(c *gin.Context) {
		qc.RefreshList()
		c.String(200, "LastOrder:%s\n", bot.LastOrder)
	})
	r.GET("/send", func(c *gin.Context) {
		msg, _ := c.Params.Get("msg")
		qc.RadioNews("MiraiGo " + msg)
	})
	port := ":9999"
	if config.GlobalConfig.GetString("port") != "" {
		port = config.GlobalConfig.GetString("port")
	}
	r.Run(port)
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt, os.Kill)
	<-ch
	qc.RadioNews(time.Now().Format("2006-01-02 15:04:05.9999" + " AWSL"))
	log.Println("AWSL")
}
