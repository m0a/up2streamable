package main

import (
	"fmt"
	"os"
	"path/filepath"

	"flag"

	streamable "github.com/maxkueng/go-streamable"
)

func main() {

	id := flag.String("u", "", "username")
	pass := flag.String("p", "", "password")
	ext := flag.String("e", ".mpg", "file extension specify")
	rootpath := flag.String("path", "./", "root file path")
	delete := flag.Bool("delete", false, "delete when upload success")

	flag.Parse()

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
				return err
			}
			//fmt.Printf("%v\n", video)
			fmt.Println("shortcode = ", video.Shortcode)
			if *delete {
				if err := os.Remove(path); err != nil {
					return err
				}
				fmt.Println(path, "deleted")
			}
			return nil
		}
		return nil
	}
	filepath.Walk(*rootpath, f)
}
