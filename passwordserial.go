package main

import (
	"fmt"
	"strings"
)

// PasswordSerial はパスワードの構造体
type PasswordSerial struct {
	Characters      []int
	Length          int
	CurrentPosition int
	Arr             []string
	Max             int
}

// NewPassword はコンストラクタ
func NewPassword(length int) *PasswordSerial {
	p := new(PasswordSerial)
	p.Characters = make([]int, length, length)
	p.Length = length
	const alpha = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	const num = "1234567890"
	const mark = "!#$%&'()=+-*/{}^~¥|?_<>,."
	p.Arr = strings.Split(alpha+num+mark, "")
	p.Max = len(p.Arr)
	return p
}

// ToString 文字列へ
func (p *PasswordSerial) ToString() string {
	return fmt.Sprintf("Length=%d", p.Length)
}

// Next 次へ
func (p *PasswordSerial) Next() bool {
	p.Characters[0]++
	isLeap := false
	for wk := 0; wk <= p.CurrentPosition; wk++ {
		if p.Characters[wk] == p.Max {
			p.Characters[wk] = 0
			if (wk + 1) > p.CurrentPosition {
				isLeap = true
				break
			}
			p.Characters[wk+1]++
		}

	}
	if isLeap {
		p.CurrentPosition++
		if p.Length <= p.CurrentPosition {
			return false
		}
	}

	return true
}

// GetPassword もじ
func (p PasswordSerial) GetPassword() string {
	result := ""
	for wk := p.CurrentPosition; wk >= 0; wk-- {
		result += p.Arr[p.Characters[wk]]
	}
	return result
}
