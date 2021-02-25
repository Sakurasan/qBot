package main

import (
	"fmt"
	"qBot/pkg/qchan"
	"qBot/pkg/qqmusic"
	"time"
)

var (
	list567   = []int{5567, 5670, 6567, 7567, 13567, 50567, 55670, 56700, 56713, 56767, 60567, 65670, 70567, 75670}
	sendcount = 0
)

func main() {
	for range time.Tick(1 * time.Second) {
		// fmt.Println(time.Now().Second())
		album_count := qqmusic.Get_album_count()
		// fmt.Println(album_count)
		for _, v := range list567 {
			if album_count+10 == v && sendcount <= 13 {
				qchan.Send(fmt.Sprintf("%s 伍六七之玄武国篇 原声大碟 已售 %d 张", time.Now().Format("2006-01-02 15:04:05"), album_count))
				sendcount++
			} else {
				sendcount = 0
			}
		}

	}
}
