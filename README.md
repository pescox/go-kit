# go-kit
go开发常用的工具库。

## 导入库
```bash
go get github.com/pescox/go-kit@latest
```

## log
终端日志工具，可以打印带时间戳的，格式化的，Json结构化的日志。
```golang
func main() {
    log.Debug("debug message")
	log.DebugF("debug message %s", "arg")
	log.DebugJ("debug message")

	log.Info("info message")
	log.InfoF("info message %s", "arg")
	log.InfoJ("info message")

	log.Warn("warn message")
	log.WarnF("warn message %s", "arg")
	log.WarnJ("warn message")

	log.Error("error message")
	log.ErrorF("error message %s", "arg")
	log.ErrorJ("error message")

	// log.Fatal("fatal message")
	// log.FatalF("fatal message %s", "arg")
	// log.FatalJ("fatal message")

	log.Print("some message")
	log.Printf("some message %s", "arg")
}
```

## cast
```golang
func main() {
    var i int = 10
    pi := cast.Ptr(i)

    vi := cast.Val(pi)
}
```