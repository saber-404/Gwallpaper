### 说明
Win10 定时随机更换壁纸程序，支持锁屏壁纸同步。

### 配置文件 `setting.json`

```json
{
    "SleepTime": 900,
    "ChangeLockWallPaper": false,
    "Cache": {
        "Name": "D:\\壁纸",
        "Children": [
            { "Name": "test.png", "Children": null }
        ]
    }
}
```

- `SleepTime` — 更换间隔（秒）
- `ChangeLockWallPaper` — 是否同步更改锁屏壁纸
- `Cache` — 缓存的图片树（自动生成）

### 编译

```shell
# 隐藏窗口（默认，适合开机启动）
make

# 带控制台窗口（调试用）
make gui

# 清理编译产物
make clean

# 普通启动
make run

# 管理员权限启动（支持更改锁屏壁纸）
make powerRun
```

### 手动编译

```shell
go build -o deploy/changewallpaper.exe -ldflags="-s -w -H windowsgui" main.go
```

### 引用

- `github.com/getlantern/systray`
- `golang.org/x/sys/windows/registry`

### 版本

- 10 — 加入图片树缓存机制
