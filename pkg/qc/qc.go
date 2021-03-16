package qc

import (
	"bufio"
	"bytes"
	"context"
	"errors"
	"fmt"
	"image"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"qBot/pkg/config"
	"qBot/pkg/errorsType"
	"qBot/pkg/loli"
	"qBot/pkg/qchan"
	"qBot/pkg/tb"
	"qBot/pkg/tts"
	"qBot/tests"
	"strconv"
	"strings"
	"time"

	"github.com/Mrs4s/MiraiGo/client"
	"github.com/Mrs4s/MiraiGo/message"
	"github.com/valyala/fasthttp"
	asciiart "github.com/yinghau76/go-ascii-art"
)

var (
	GroupMap  = make(map[int64]string)
	MyGroupID = make([]int, 2)

	account, _ = strconv.ParseInt(os.Getenv("account"), 10, 64)

	pwd      = os.Getenv("pwd")
	Bot      *client.QQClient
	mjjChan  = make(chan string, 10)
	setuChan = make(chan string, 11)
	setuLock = false
	setuNum  = 0

	heshuiLock    = false
	heshuiUrlList = []string{"https://pic3.zhimg.com/80/v2-679a91c6d249517b6e267cd3b07d7ad7_720w.jpg", "https://pic2.zhimg.com/80/v2-200ebdfc2a49bf29414f504d01f57c22_720w.jpg", "https://pic4.zhimg.com/80/v2-f56975d4faa5332cc24c5b0b37c6cfe9_720w.jpg", "https://pic4.zhimg.com/80/v2-f56975d4faa5332cc24c5b0b37c6cfe9_720w.jpg"}
	moyu          = "https://pic4.zhimg.com/80/v2-981cd99dd2969eaf1ca9b23783eba818_720w.jpg"
	ghsUrl        = "https://pic4.zhimg.com/80/v2-ebe50b205335ad46bb356999146f1106_720w.jpg"
	tigangUrl     = "https://pic.diydoutu.com/bq/2067.jpg"

	ttsmap   = make(map[int64]bool)
	clockmap = make(map[int64]bool)
	watermap = make(map[int64]bool)
	ppmap    = make(map[int64]bool)
)

func Init() {
	if os.Getenv("account") != "" && pwd != "" {
		Bot = client.NewClient(account, pwd)
	} else if config.GlobalConfig.GetInt64("account") != 0 && config.GlobalConfig.GetString("pwd") != "" {
		Bot = client.NewClient(config.GlobalConfig.GetInt64("account"), config.GlobalConfig.GetString("pwd"))
	} else {
		panic("config/qBot.yaml文件 account,pwd 信息未配置")
	}

	MyGroupID = config.GlobalConfig.GetIntSlice("groupID")
	if len(MyGroupID) < 1 {
		panic("config/qBot.yaml文件 groupID 信息未配置")
	}

}

func checkDevice() {
	if !IsFileExist("device.json") {
		log.Println("虚拟设备信息不存在, 将自动生成随机设备.")
		client.GenRandomDevice()
		_ = ioutil.WriteFile("device.json", client.SystemDeviceInfo.ToJson(), 0777)
		log.Println("已生成设备信息并保存到 device.json 文件.")
	} else {
		log.Println("将使用 device.json 内的设备信息运行Bot.")
		devicefile, _ := ioutil.ReadFile("device.json")
		if err := client.SystemDeviceInfo.ReadJson(devicefile); err != nil {
			log.Fatalf("加载设备信息失败: %v", err)
		}
	}
}

func Login() {
	console := bufio.NewReader(os.Stdin)
	readLine := func() (str string) {
		str, _ = console.ReadString('\n')
		return
	}
	checkDevice()

	rsp, err := Bot.Login()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("使用协议: %v", func() string {
		switch client.SystemDeviceInfo.Protocol {
		case client.IPad:
			return "iPad"
		case client.AndroidPhone:
			return "Android Phone"
		case client.AndroidWatch:
			return "Android Watch"
		case client.MacOS:
			return "MacOS"
		}
		return "未知"
	}())
	Bot.OnServerUpdated(func(c *client.QQClient, e *client.ServerUpdatedEvent) bool {
		fmt.Println(e)
		return false
	})

	for {
		if !rsp.Success {
			switch rsp.Error {
			case client.SliderNeededError:
				if client.SystemDeviceInfo.Protocol == client.AndroidPhone {
					fmt.Println("警告: Android Phone 强制要求暂不支持的滑条验证码, 请开启设备锁或切换到Watch协议验证通过后再使用.")
					fmt.Println("按 Enter 继续....")
					readLine()
					os.Exit(0)
				}
				Bot.AllowSlider = false
				Bot.Disconnect()
				rsp, err = Bot.Login()
				continue
			case client.SMSNeededError:
				fmt.Println("账号已开启设备锁, 按下 Enter 向手机 %v 发送短信验证码.", rsp.SMSPhone)
				readLine()
				if !Bot.RequestSMS() {
					fmt.Println("发送验证码失败，可能是请求过于频繁.")
					time.Sleep(time.Second * 5)
					os.Exit(0)
				}
				fmt.Println("请输入短信验证码： (Enter 提交)")
				text := readLine()
				rsp, err = Bot.SubmitSMS(strings.ReplaceAll(strings.ReplaceAll(text, "\n", ""), "\r", ""))
				continue
			case client.SMSOrVerifyNeededError:
				fmt.Println("账号已开启设备锁，请选择验证方式:")
				fmt.Println("1. 向手机 %v 发送短信验证码", rsp.SMSPhone)
				fmt.Println("2. 使用手机QQ扫码验证.")
				fmt.Println("请输入(1 - 2): ")
				text := readLine()
				if strings.Contains(text, "1") {
					if !Bot.RequestSMS() {
						fmt.Println("发送验证码失败，可能是请求过于频繁.")
						time.Sleep(time.Second * 5)
						os.Exit(0)
					}
					fmt.Println("请输入短信验证码： (Enter 提交)")
					text = readLine()
					rsp, err = Bot.SubmitSMS(strings.ReplaceAll(strings.ReplaceAll(text, "\n", ""), "\r", ""))
					continue
				}
				println("请前往 -> %v <- 验证并重启Bot.", rsp.VerifyUrl)
				println("按 Enter 继续....")
				readLine()
				os.Exit(0)
				return
			case client.NeedCaptcha:
				// f, _ := os.OpenFile("captua.png", os.O_RDWR|os.O_CREATE, os.ModePerm)
				// f.Write(rsp.CaptchaImage)
				img, _, _ := image.Decode(bytes.NewReader(rsp.CaptchaImage))
				fmt.Println(asciiart.New("image", img).Art)
				fmt.Println("请输入验证码： (回车提交)")
				text, _ := console.ReadString('\n')
				rsp, err = Bot.SubmitCaptcha(strings.ReplaceAll(text, "\n", ""), rsp.CaptchaSign)
				continue
			case client.UnsafeDeviceError:
				fmt.Printf("账号已开启设备锁，请前往 -> %v <- 验证并重启Bot.", rsp.VerifyUrl)
				return
			case client.OtherLoginError, client.UnknownLoginError:
				log.Fatalf("登录失败: %v", rsp.ErrorMessage)
			}
		}
		break
	}

	fmt.Printf("登录成功 欢迎使用: %v \n", Bot.Nickname)

	Bot.ReloadGroupList()
	Bot.ReloadFriendList()

	groupList, _ := Bot.GetGroupList()
	fmt.Printf("共%2d 个群\n", len(groupList))
	for _, group := range groupList {
		GroupMap[group.Code] = group.Name
	}
	fmt.Println(GroupMap)

	// friendList, _ := qc.GetFriendList()
	// fmt.Printf("共%2d 个好友\n", friendList.TotalCount)
	// for _, friend := range friendList.List {
	// 	fmt.Println(friend)
	// }
	for _, v := range MyGroupID {
		if _, ok := GroupMap[int64(v)]; ok {
			Bot.SendGroupMessage(int64(v), &message.SendingMessage{
				Elements: []message.IMessageElement{message.NewText("测试信息：消息姬已启动 ")}})
		}
	}
	qchan.SendGroup("测试信息：Qmsg已启动 ", "808468274")
	Bot.OnGroupMessage(msgRoute)
	// defer Bot.Conn.Close()

}

// RadioNews 广播群消息
func RadioNews(msg string) error {
	for _, v := range MyGroupID {
		if _, ok := GroupMap[int64(v)]; ok {
			ret := Bot.SendGroupMessage(int64(v), &message.SendingMessage{
				Elements: []message.IMessageElement{message.NewText(msg)}})
			if ret == nil {
				return errorsType.ErrNilResp
			} else if ret.Id == -1 {
				log.Println("群消息发送失败 message-id", ret.Id,
					"internal-id", ret.InternalId,
					"group", ret.GroupCode,
					"group-name", ret.GroupName,
					"sender", ret.Sender,
					"time", ret.Time)
				return errors.New(fmt.Sprintf("%d", ret.Id))
			}
		}
	}
	return nil
}

// RefreshList 刷新联系人
func RefreshList() error {
	err := Bot.ReloadGroupList()
	if err != nil {
		return err
	}
	return Bot.ReloadFriendList()
}
func Printmap() {
	fmt.Println("ttsmap", ttsmap)
	fmt.Println("clockmap", clockmap)
	fmt.Println("watermap", watermap)
	fmt.Println("ppmap", ppmap)
}

//判断文件是否存在
func IsFileExist(filename string) bool {
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return true

}

func Renew() {
	Bot.Conn.Close()
	Bot.Login()
	RefreshList()
}

func msgRoute(c *client.QQClient, msg *message.GroupMessage) {
	s := msg.ToString()
	// log.Printf("%s[%d]:%s[%d] %s\n", msg.GroupName, msg.GroupCode, msg.Sender.Nickname, msg.Sender.Uin, s)
	for _, m := range msg.Elements {
		switch m.(type) {
		case *message.TextElement:
			s := m.(*message.TextElement).Content
			if (strings.Contains(s, "。") || strings.Contains(s, ",") || strings.Contains(s, "，")) && ttsmap[msg.GroupCode] != false {
				// if msg.Sender.Uin == tests.QType.Uin {
				// 	return
				// }
				var tc tts.Config
				if strings.Contains(s, ",") || strings.Contains(s, "，") {
					tc.Language = "ja"
					transt := tts.NewTransT()
					transt.Target = "ja"
					transt.Text = s
					tc.Speak = notNillValve(s, tts.Trans(transt))
				} else if strings.Contains(s, "。") {
					tc.Language = "zh-CN"
					tc.Speak = s
				}

				voice, err := tts.Speak(tc)
				if err != nil {
					log.Println(err)
					return
				}
				gv, err := c.UploadGroupPtt(msg.GroupCode, bytes.NewReader(voice.Bytes()))
				if err != nil {
					fmt.Println(err)
					return
				}
				c.SendGroupMessage(msg.GroupCode, message.NewSendingMessage().Append(gv))
				return
			} else if strings.Contains(s, "clock") {
				args := strings.Split(s, " ")
				if len(args) < 2 {
					return
					c.SendGroupMessage(msg.GroupCode, message.NewSendingMessage().Append(message.NewText("无效指令")))
					return
				}
				switch args[1] {
				case "喝水":
					watermap[msg.GroupCode] = true
				case "喝水关":
					watermap[msg.GroupCode] = false
				case "提肛":
					ppmap[msg.GroupCode] = true
				case "提肛关":
					ppmap[msg.GroupCode] = false
				}
				return
			}
			switch s {
			case "menu":
				menu := `喝水
提肛
ghs
涩图, 瑟图
女菩萨, 绅士
老算盘
, 。 //? 日语：中文
clock 
`
				c.SendGroupMessage(
					msg.GroupCode, message.NewSendingMessage().Append(message.NewText(menu)))
			case "喝水":
				sm, err := upLoadImgByUrl(c, msg, heshuiUrlList[rand.Intn(len(heshuiUrlList)-1)])
				if err != nil {
					return
				}
				c.SendGroupMessage(msg.GroupCode, sm)
			case "提肛":
				sm, err := upLoadImgByUrl(c, msg, tigangUrl)
				if err != nil {
					return
				}
				img, _ := upImgByUrl(c, msg.GroupCode, "https://i.loli.net/2021/03/13/REphylvQiAsYdPf.gif")
				c.SendGroupMessage(msg.GroupCode, sm.Append(img))
			case "ghs":
				sm, err := upLoadImgByUrl(c, msg, ghsUrl)
				if err != nil {
					return
				}
				c.SendGroupMessage(msg.GroupCode, sm)
			case "涩图", "瑟图":
				if !setuLock {
					var (
						m   map[string]string
						err error
					)
					if rand.Intn(10)%2 == 0 {
						m, err = tb.Tao()
					} else {
						m, err = tb.Tbimg()
					}

					if err != nil {
						c.SendGroupMessage(msg.GroupCode, message.NewSendingMessage().Append(message.NewText("没有，滚")))
						return
					}
					imsg, err := upLoadFlashImgByUrl(c, msg, m["pic"])
					if err != nil {
						log.Println(err)
						c.SendGroupMessage(msg.GroupCode, message.NewSendingMessage().Append(message.NewText("ghs?哪呢哪呢？")))
						return
					}
					// c.SendGroupMessage(msg.GroupCode, &message.SendingMessage{Elements: imsg})
					c.SendGroupMessage(msg.GroupCode, imsg)
				}
			case "女菩萨", "绅士":
				if setuLock {
					c.SendGroupMessage(msg.GroupCode, message.NewSendingMessage().Append(message.NewText("ghs?哪呢哪呢？")))
					return
				}
				var (
					m   map[string]string
					err error
				)

				if rand.Intn(10)%2 == 0 {
					m, err = tb.Tao()
				} else {
					m, err = tb.Tbimg()
				}

				if err != nil {
					c.SendGroupMessage(msg.GroupCode, message.NewSendingMessage().Append(message.NewText("ghs?哪呢哪呢？")))
					return
				}
				c.SendGroupMessage(msg.GroupCode, message.NewSendingMessage().Append(message.NewText(m["pic"])))
			case "老算盘":
				_url, err := loli.SetuOne()
				if err != nil {
					c.SendGroupMessage(msg.GroupCode, message.NewSendingMessage().Append(message.NewText("ghs?哪呢哪呢？")))
					return
				}
				c.SendGroupMessage(msg.GroupCode, message.NewSendingMessage().Append(message.NewText(_url)))
				return
			case "ping":
				c.SendGroupMessage(msg.GroupCode, message.NewSendingMessage().Append(message.NewText("pong")))
				return
			}
		}
		if msg.Sender.Uin == tests.QType.Uin {
			if s == "涩图开" {
				setuLock = false
				return
			}
			if s == "涩图关" {
				setuLock = true
				return
			}
			if s == "tts关" {
				ttsmap[msg.GroupCode] = false
				return
			}
			if s == "tts开" {
				ttsmap[msg.GroupCode] = true
				return
			}
			// transt := tts.NewTransT()
			// transt.Target = "ja"
			// transt.Text = s
			// var tc tts.Config
			// tc.Speak = notNillValve(s, tts.Trans(transt))
			// if strings.Contains(s, ",") || strings.Contains(s, "，") {
			// 	tc.Language = "ja"
			// } else if strings.Contains(s, "。") {
			// 	tc.Language = "zh-CN"
			// }

			// voice, err := tts.Speak(tc)
			// if err != nil {
			// 	log.Println(err)
			// 	return
			// }
			// gv, err := c.UploadGroupPtt(msg.GroupCode, bytes.NewReader(voice.Bytes()))
			// if err != nil {
			// 	fmt.Println(err)
			// 	return
			// }
			// c.SendGroupMessage(msg.GroupCode, message.NewSendingMessage().Append(gv))
			return

		}
	}
	// if msg.GroupCode == 808468274 {}
}

func upLoadImgByUrl(c *client.QQClient, msg *message.GroupMessage, url string) (*message.SendingMessage, error) {
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
	img, err := c.UploadGroupImage(msg.GroupCode, bytes.NewReader(rsp.Body()))
	if err != nil {
		sm.Append(message.NewText(fmt.Sprintf("上传失败 %s ", err.Error())))
		return sm, nil
	}
	sm.Append(img)
	return sm, nil
}
func upLoadFlashImgByUrl(c *client.QQClient, msg *message.GroupMessage, url string) (*message.SendingMessage, error) {
	_, cc := context.WithTimeout(context.Background(), 10*time.Second)
	defer cc()
	sm := message.NewSendingMessage()
	req := new(fasthttp.Request)
	req.Header.Set("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/88.0.4324.192 Safari/537.36")
	req.SetRequestURI(url)
	rsp := new(fasthttp.Response)
	if err := fasthttp.Do(req, rsp); err != nil {
		log.Println("fasthttp.Do", err)
		return nil, err
	}

	img, err := c.UploadGroupImage(msg.GroupCode, bytes.NewReader(rsp.Body()))
	if err != nil {
		sm.Append(message.NewText(fmt.Sprintf("上传失败 %s ", err.Error())))
		return sm, nil
	}
	// if i, ok := img.(*message.GroupImageElement); ok {}
	sm.Append(&message.GroupFlashPicElement{GroupImageElement: *img})

	return sm, nil
}

func SetuMsg(c *client.QQClient, msg *message.GroupMessage) {

}

func notNillValve(a, b string) string {
	if b != "" {
		return b
	}
	return a
}
