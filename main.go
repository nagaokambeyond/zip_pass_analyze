package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/yeka/zip"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("zip_pass_analyze ファイル名 パスワードの長さ")
		return
	}
	filename := os.Args[1]
	passwordLength, _ := strconv.Atoi(os.Args[2])
	if f, err := os.Stat(filename); os.IsNotExist(err) || f.IsDir() {
		fmt.Println("zipファイルは存在しません！")
		return
	}

	r, err := zip.OpenReader(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer r.Close()

	start := time.Now()
	a := NewPassword(passwordLength)

	isSuccess := false
	for _, f := range r.File {
		if !f.IsEncrypted() {
			log.Fatal("no pass zip")
		}
		for {
			pass := a.GetPassword()
			f.SetPassword(pass)

			r, err := f.Open()
			if err != nil {
				if !a.Next() {
					break
				}
				continue
			}
			defer r.Close()
			buf, err := ioutil.ReadAll(r)
			if err != nil {
				if !a.Next() {
					break
				}
				continue
			}

			fmt.Printf("pass=%v\n", pass)
			fmt.Printf("Size of %v: %v byte(s)\n", f.Name, len(buf))
			isSuccess = true
			break
		}
		break
	}
	if isSuccess {
		fmt.Println("成功")
	}
	end := time.Now()
	fmt.Println(start)
	fmt.Println(end)

	return
}
