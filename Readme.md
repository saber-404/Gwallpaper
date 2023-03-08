### 说明
这是一个win10定时随机更换壁纸程序，配置文件是同目录下的setting.json文件

```
{
  "RetryTimes": 10, 
  "FolderPath": "D:\\壁纸\\",
  "SleepTime": 900
}

RetryTimes 是重试时间
FolderPath 指定壁纸文件夹 路径中最后一定要是"\\" 或"/"
SleepTime  是更换间隔,单位是秒
```

### 编译选项
优化体积，隐藏窗口
```shell
go build -o changewallpaper.exe -ldflags="-s -w -H windowsgui" .\main.go
```

### 版本与gitTag
3添加一个自己更换的选项

### 引用
"github.com/getlantern/systray"