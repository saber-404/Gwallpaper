### 说明
这是一个win10定时随机更换壁纸程序，配置文件是同目录下的setting.json文件

```
{
  "FolderPath": "D:\\壁纸\\",
  "SleepTime": 900
  "ChangLockWallPaper": false
}

FolderPath 指定壁纸文件夹
SleepTime  是更换间隔,单位是秒
ChangLockWallPaper 是否更改锁屏 布尔值
```

### 编译选项
优化体积，隐藏窗口
```shell
go build -o changewallpaper.exe -ldflags="-s -w -H windowsgui" .\main.go
```

### 版本与gitTag
9构建图片树


### 引用
"github.com/getlantern/systray"

"golang.org/x/sys/windows/registry"

"github.com/lxn/win"

### 测试
有的测试需要管理员权限

`sudo go test -v -run UndoTestSetLockWallpaper .\Gwallpaper_test.go`
