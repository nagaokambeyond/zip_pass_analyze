package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	a := NewPassword(4)
	//for idx := 1; idx <= 100; idx++ {
	for {
		//fmt.Println(a.GetPassword())
		if !a.Next() {
			break
		}
	}
	end := time.Now()
	fmt.Println(start)
	fmt.Println(end)

	return

	//a: = Password.New(10)

	// r, err := zip.OpenReader(os.Args[1])
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer r.Close()
	// const alpha = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	// const num = "1234567890"
	// const mark = "!#$%&'()=+-*/{}^~¥|?_<>,."
	// var arr = strings.Split(alpha+num+mark, "")

	// for _, f := range r.File {
	// 	if !f.IsEncrypted() {
	// 		log.Fatal("no pass zip")
	// 	}
	// 	var characters [4]int
	// 	const last = len(characters) - 1
	// 	var max = len(arr) - 1

	// 	var pass string
	// 	var col int
	// 	var isSuccess bool
	// 	col = 0
	// 	isSuccess = false
	// 	fmt.Println(time.Now())
	// 	for {

	// 		for idx := 0; idx <= max; idx++ {
	// 			characters[0] = idx
	// 			pass = ""
	// 			for wk := col; wk >= 0; wk-- {
	// 				pass = pass + arr[characters[wk]]
	// 			}
	// 			f.SetPassword(pass)

	// 			r, err := f.Open()
	// 			if err != nil {
	// 				continue
	// 			}
	// 			buf, err := ioutil.ReadAll(r)
	// 			if err != nil {
	// 				continue
	// 			}

	// 			fmt.Printf("pass=%v\n", pass)
	// 			fmt.Printf("Size of %v: %v byte(s)\n", f.Name, len(buf))
	// 			isSuccess = true
	// 			break
	// 		}
	// 		if isSuccess {
	// 			break
	// 		}
	// 		if col == 0 {
	// 			col++
	// 		} else if col == 1 {
	// 			characters[col]++
	// 			if characters[col] > max {
	// 				for wk := col; wk >= 0; wk-- {
	// 					characters[wk] = 0
	// 				}
	// 				col++
	// 			}
	// 		} else if col == 2 {
	// 			characters[col-1]++
	// 			if characters[col-1] > max {
	// 				for wk := col - 1; wk >= 0; wk-- {
	// 					characters[wk] = 0
	// 				}
	// 				characters[col]++
	// 			}
	// 			if characters[col] > max {
	// 				for wk := col; wk >= 0; wk-- {
	// 					characters[wk] = 0
	// 				}
	// 				col++
	// 			}
	// 		} else if col >= 3 {
	// 			characters[col-2]++
	// 			if characters[col-2] > max {
	// 				for wk := col - 2; wk >= 0; wk-- {
	// 					characters[wk] = 0
	// 				}
	// 				characters[col-1]++
	// 			}
	// 			if characters[col-1] > max {
	// 				for wk := col - 1; wk >= 0; wk-- {
	// 					characters[wk] = 0
	// 				}
	// 				characters[col]++
	// 			}
	// 			if characters[col] > max {
	// 				for wk := col; wk >= 0; wk-- {
	// 					characters[wk] = 0
	// 				}
	// 				col++
	// 			}
	// 		}
	// 		if col > last {
	// 			break
	// 		}
	// 	}
	// 	fmt.Println(time.Now())
	// }
}
