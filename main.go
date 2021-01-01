package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/yeka/zip"
)

func main() {
	r, err := zip.OpenReader(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer r.Close()

	start := time.Now()
	a := NewPassword(4)

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
