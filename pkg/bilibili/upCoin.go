package bilibili

import (
	"encoding/json"
	"net/http"
	"net/url"
)

type UpCoinRspType struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	TTL     int    `json:"ttl"`
	Data    []struct {
		Aid       int    `json:"aid"` //
		Videos    int    `json:"videos"`
		Tid       int    `json:"tid"`
		Tname     string `json:"tname"`
		Copyright int    `json:"copyright"`
		Pic       string `json:"pic"`
		Title     string `json:"title"` //
		Pubdate   int    `json:"pubdate"`
		Ctime     int    `json:"ctime"`
		Desc      string `json:"desc"`
		State     int    `json:"state"`
		Duration  int    `json:"duration"`
		Rights    struct {
			Bp            int `json:"bp"`
			Elec          int `json:"elec"`
			Download      int `json:"download"`
			Movie         int `json:"movie"`
			Pay           int `json:"pay"`
			Hd5           int `json:"hd5"`
			NoReprint     int `json:"no_reprint"`
			Autoplay      int `json:"autoplay"`
			UgcPay        int `json:"ugc_pay"`
			IsCooperation int `json:"is_cooperation"`
			UgcPayPreview int `json:"ugc_pay_preview"`
			NoBackground  int `json:"no_background"`
		} `json:"rights"`
		Owner struct {
			Mid  int    `json:"mid"`
			Name string `json:"name"`
			Face string `json:"face"`
		} `json:"owner"`
		Stat struct {
			Aid      int `json:"aid"` //
			View     int `json:"view"`
			Danmaku  int `json:"danmaku"`
			Reply    int `json:"reply"`
			Favorite int `json:"favorite"`
			Coin     int `json:"coin"`
			Share    int `json:"share"`
			NowRank  int `json:"now_rank"`
			HisRank  int `json:"his_rank"`
			Like     int `json:"like"`
			Dislike  int `json:"dislike"`
		} `json:"stat"`
		Dynamic   string `json:"dynamic"`
		Cid       int    `json:"cid"`
		Dimension struct {
			Width  int `json:"width"`
			Height int `json:"height"`
			Rotate int `json:"rotate"`
		} `json:"dimension"`
		Bvid       string `json:"bvid"` //
		Coins      int    `json:"coins"`
		Time       int    `json:"time"`
		IP         string `json:"ip"`
		InterVideo bool   `json:"inter_video"`
		MissionID  int    `json:"mission_id,omitempty"`
		SeasonID   int    `json:"season_id,omitempty"`
	} `json:"data"`
}

// 获取up投币视频
func UpCoinVideo(vmid string) (*UpCoinRspType, error) {
	_url := "http://api.bilibili.com/x/space/coin/video"
	v := url.Values{}
	v.Add("vmid", vmid)
	req, err := http.NewRequest("GET", _url+"?"+v.Encode(), nil)
	if err != nil {
		return nil, err
	}

	rsp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	rspType := new(UpCoinRspType)
	if err := json.NewDecoder(rsp.Body).Decode(rspType); err != nil {
		return nil, err
	}
	return rspType, nil
}
