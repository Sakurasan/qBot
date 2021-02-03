package bot

import (
	"fmt"
	"qBot/pkg/mjj"
	"qBot/pkg/qc"
	"qBot/pkg/qchan"
	"sync"

	"github.com/Sakurasan/to"
)

var (
	LastOrder   string
	UrlTemplate = `https://hostloc.com/thread-%s-1-1.html`
)

func Run() {
	var lock sync.RWMutex
	lock.RLock()
	newsList, err := mjj.MjjCdpMobile()
	if err != nil {
		fmt.Println(err)
		return
	}

	for i := len(newsList) - 1; i >= 0; i-- {
		if to.Int64(LastOrder) < to.Int64(newsList[i][0]) {
			err := qc.RadioNews(newsList[i][1] + fmt.Sprintf(UrlTemplate, newsList[i][0]))
			if err != nil {
				qchan.SendGroup(newsList[i][1]+fmt.Sprintf(UrlTemplate, newsList[i][0]), "808468274")
			}
		}
	}

	LastOrder = newsList[0][0]
	lock.RUnlock()
	return
}
