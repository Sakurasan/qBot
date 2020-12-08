使用官方Docker安装
```
下载docker image : docker pull chromedp/headless-shell
运行docker : docker run -d -p 9222:9222 --rm --name headless-shell chromedp/headless-shell
浏览器访问 http://10.202.255.220:9222/json 
```

yum安装chronium-headless
```
搜索chrome 的yum源
yum search chromium

sudo yum install chromium-headless.x86_64

```
[Linux CentOS 7 安装字体库 & 中文字体](https://blog.csdn.net/wlwlwlwl015/article/details/51482065)

启动
```
nohup /usr/lib64/chromium-browser/headless_shell --no-first-run --no-default-browser-check --headless --disable-gpu --remote-debugging-port=9222 --no-sandbox --disable-plugins --remote-debugging-address=0.0.0.0 --window-size=1920,1080 &
```
检查是否启动
```
netstat -lntp
```
headless_shell(chrome) Flag 参数说明
```
--no-first-run 第一次不运行
---default-browser-check 不检查默认浏览器
--headless 不开启图像界面
--disable-gpu 关闭gpu,服务器一般没有显卡
remote-debugging-port chrome-debug工具的端口(golang chromepd 默认端口是9222,建议不要修改)
--no-sandbox 不开启沙盒模式可以减少对服务器的资源消耗,但是服务器安全性降低,配和参数 --remote-debugging-address=127.0.0.1 一起使用
--disable-plugins 关闭chrome插件
--remote-debugging-address 远程调试地址 0.0.0.0 可以外网调用但是安全性低,建议使用默认值 127.0.0.1
--window-size 窗口尺寸
更多参数说明详解headless-chrome官方文档
```


cmd /c "C:\Program Files (x86)\Google\Chrome\Application\chrome.exe"  --remote-debugging-port=9222 --no-sandbox --disable-plugins --remote-debugging-address=0.0.0.0

https://www.dazhuanlan.com/2019/10/25/5db209ebbcc11/

[example] https://www.cnblogs.com/pu369/p/12330074.html