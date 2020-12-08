package main

import (
	"os"
	"os/signal"
	"qBot/pkg/bot"
	"qBot/pkg/config"
	"qBot/pkg/qc"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	config.Init()
	qc.Init()
	qc.Login()

	go func() {
		for {
			bot.Run()
			time.Sleep(10 * time.Second)
		}
	}()

	r := gin.New()

	r.GET("/debug", func(c *gin.Context) {
		qc.RefreshList()
		c.String(200, "LastOrder:%s\n", bot.LastOrder)
	})
	r.Run()
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt, os.Kill)
	<-ch

}
