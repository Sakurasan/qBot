package bot

import (
	"fmt"
	"qBot/pkg/mjj"
	"qBot/pkg/qc"
	"sync"

	"github.com/Sakurasan/to"
)

var (
	LastOrder   string
	UrlTemplate = `https://hostloc.com/thread-%s-1-1.html`
)

func Run() {
	var lock sync.Mutex
	lock.Lock()
	newsList, err := mjj.MjjCdpMobile()
	if err != nil {
		fmt.Println(err)
		return
	}
	if LastOrder != "" && LastOrder != newsList[0][0] {
		// for _, v := range newsList {
		// 	if LastOrder != v[0] {
		// 		qc.RadioNews(v[1] + fmt.Sprintf(UrlTemplate, v[0]))
		// 	}
		// }
		qc.RadioNews(newsList[0][1] + fmt.Sprintf(UrlTemplate, newsList[0][0]))
	} else {
		for i := len(newsList) - 1; i > 0; i-- {
			if to.Int64(LastOrder) < to.Int64(newsList[i][0]) {
				qc.RadioNews(newsList[i][1] + fmt.Sprintf(UrlTemplate, newsList[i][0]))
			}
		}
	}

	LastOrder = newsList[0][0]
	lock.Unlock()
	return
}