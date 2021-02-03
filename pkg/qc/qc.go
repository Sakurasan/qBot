package qc

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"image"
	"io/ioutil"
	"log"
	"os"
	"qBot/pkg/config"
	"strconv"
	"strings"
	"time"

	"github.com/Mrs4s/MiraiGo/client"
	"github.com/Mrs4s/MiraiGo/message"
	asciiart "github.com/yinghau76/go-ascii-art"
)

var (
	GroupMap  = make(map[int64]string)
	MyGroupID = make([]int, 2)

	account, _ = strconv.ParseInt(os.Getenv("account"), 10, 64)

	pwd = os.Getenv("pwd")
	Bot *client.QQClient
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

	// defer Bot.Conn.Close()

}

// RadioNews 广播群消息
func RadioNews(msg string) error {
	for _, v := range MyGroupID {
		if _, ok := GroupMap[int64(v)]; ok {
			ret := Bot.SendGroupMessage(int64(v), &message.SendingMessage{
				Elements: []message.IMessageElement{message.NewText(msg)}})
			if ret == nil || ret.Id == -1 {
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
