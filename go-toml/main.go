package main

import (
	"fmt"
	"log"
	"path"
	"path/filepath"

	"github.com/fsnotify/fsnotify"
	toml "github.com/pelletier/go-toml"
)

type Level int32

const (
	DEBUG Level = iota + 1
	INFO
	WARN
	ERROR
	PANIC
)

func main() {
	filePath := "/home/tian/workspace/golang/src/github.com/tiandh987/gopractise-demo/go-toml/test.toml"
	var fileLevel Level

	// config, _ := toml.LoadFile(filePath) //加载toml文件
	// key := config.Get("os").(string)     //读取key对应的值.括号为指定数据类型，也可以忽略
	// fmt.Println(key)

	dir := filepath.Dir(filePath)
	ext := path.Ext(filePath)

	//fmt.Println("dir", dir)
	//fmt.Println("ext", ext)
	//fmt.Println("filePath", filePath)

	// config, _ := toml.LoadFile(filePath)
	// level := config.Get("log.file_level")
	// fmt.Println(level)

	resChan := make(chan Level, 100)
	done := make(chan bool)
	go func() {
		watcher, err := fsnotify.NewWatcher()
		if err != nil {
			log.Fatal("NewWatcher failed: ", err)
		}
		defer watcher.Close()

		err = watcher.Add(dir)
		if err != nil {
			log.Fatal("Add failed:", err)
		}
		defer close(done)

		for {
			select {
			case event, _ := <-watcher.Events:
				if event.Name == filePath && event.Op&fsnotify.Write == fsnotify.Write {
					switch ext {
					case ".toml":
						//go-toml.LoadFile(filePath)
						config, _ := toml.LoadFile(filePath)
						level := config.Get("log.file_level")
						if level == nil {
							continue
						}
						//fmt.Println("toml level : ", level)

						switch t := level.(type) {
						case int:
							fmt.Printf("int : %d\n", t)
							fileLevel = Level(int(t))
						case int32:
							fmt.Printf("int : %d\n", t)
							fileLevel = Level(int32(t))
						case int64:
							fmt.Printf("int64 : %d\n", t)
							fileLevel = Level(int64(t))

						case string:
							fmt.Printf("string : %s\n", t)
							fileLevel = StrToLevel(string(t))

						default:
							fmt.Println("default ", t)
							continue
						}

					case ".ini":

					}
				}
			}

			if fileLevel == DEBUG || fileLevel == INFO || fileLevel == WARN || fileLevel == ERROR || fileLevel == PANIC {
				resChan <- fileLevel
			}
		}
	}()

	for res := range resChan {
		fmt.Println("resChan ", res)
	}

	<-done
}

func StrToLevel(level string) Level {
	var result Level

	switch level {
	case "debug":
		result = DEBUG
	case "info":
		result = INFO
	case "warn":
		result = WARN
	case "error":
		result = ERROR
	case "panic":
		result = PANIC
	}
	return result
}
