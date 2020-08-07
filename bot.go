package main

import (
	"bufio"
	"bytes"
	"fmt"
	"image"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/Mrs4s/MiraiGo/client"
	"github.com/Mrs4s/MiraiGo/message"
	asciiart "github.com/yinghau76/go-ascii-art"
)

var (
	groupMap  = make(map[int64]string)
	myGroupID int64

	uid, _ = strconv.ParseInt(os.Getenv("uid"), 10, 64)

	pwd = os.Getenv("pwd")
)

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
func main() {
	if len(os.Args) == 2 {
		myGroupID, _ = strconv.ParseInt(os.Args[1], 10, 64)
	}
	console := bufio.NewReader(os.Stdin)
	checkDevice()
	qc := client.NewClient(uid, pwd)
	rsp, err := qc.Login()
	if err != nil {
		fmt.Println(err)
	}
	for {
		if !rsp.Success {
			switch rsp.Error {
			case client.NeedCaptcha:
				// f, _ := os.OpenFile("captua.png", os.O_RDWR|os.O_CREATE, os.ModePerm)
				// f.Write(rsp.CaptchaImage)
				img, _, _ := image.Decode(bytes.NewReader(rsp.CaptchaImage))
				fmt.Println(asciiart.New("image", img).Art)
				fmt.Println("请输入验证码： (回车提交)")
				text, _ := console.ReadString('\n')
				rsp, err = qc.SubmitCaptcha(strings.ReplaceAll(text, "\n", ""), rsp.CaptchaSign)
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

	fmt.Printf("登录成功 欢迎使用: %v \n", qc.Nickname)

	qc.ReloadGroupList()
	qc.ReloadFriendList()

	groupList, _ := qc.GetGroupList()
	fmt.Printf("共%2d 个群\n", len(groupList))
	for _, group := range groupList {
		groupMap[group.Code] = group.Name
	}
	fmt.Println(groupMap)

	// friendList, _ := qc.GetFriendList()
	// fmt.Printf("共%2d 个好友\n", friendList.TotalCount)
	// for _, friend := range friendList.List {
	// 	fmt.Println(friend)
	// }

	if _, ok := groupMap[myGroupID]; ok {
		qc.SendGroupMessage(myGroupID, &message.SendingMessage{
			Elements: []message.IMessageElement{message.NewText("测试信息http://www.qq.com")}})
	}

	defer qc.Conn.Close()

}

//判断文件是否存在
func IsFileExist(filename string) bool {
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return true

}
