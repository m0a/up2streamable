package main

import (
	"flag"
	"fmt"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"time"

	streamable "github.com/maxkueng/go-streamable"
)

func setuphttpClient(minutes int64) {
	http.DefaultClient = &http.Client{
		Timeout: time.Duration(minutes) * time.Minute,
	}
}

func main() {

	var _ = os.IsExist(nil)
	id := flag.String("u", "", "username")
	pass := flag.String("p", "", "password")
	ext := flag.String("e", ".mpg", "specify file extension")
	searchPath := flag.String("path", "./", "search file path")
	delete := flag.Bool("delete", false, "delete when upload success")
	timeout := flag.Int64("timeout", 10, "specify file upload timeout(minutes)")

	flag.Parse()
	setuphttpClient(*timeout)
	client := streamable.New()
	if *id != "" && *pass != "" {
		client.SetCredentials(*id, *pass)
	}
	var f filepath.WalkFunc = func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}

		if filepath.Ext(path) == *ext {
			fmt.Println("uploading.. ", path)
			video, err := client.UploadVideo(path)
			if err != nil {
				timeout, ok := err.(net.Error)
				if ok && timeout.Timeout() {
					fmt.Println(path, "is timeout. upload skip. ")
					return nil
				}
				return err
			}
			//fmt.Printf("%v\n", video)
			fmt.Println("uploaded  ", path)
			fmt.Println("shortcode = ", video.Shortcode)
			if *delete {
				if err := os.Remove(path); err != nil {
					fmt.Println(err.Error())
					return err
				}
				fmt.Println(path, "deleted")
			}
			return nil
		}
		return nil
	}
	filepath.Walk(*searchPath, f)
}
