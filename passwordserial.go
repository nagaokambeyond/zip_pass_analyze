package main

import (
	"strings"
)

// PasswordSerial はパスワードの構造体
type PasswordSerial struct {
	Characters      []byte
	Length          int
	CurrentPosition int
	Arr             []string
	Max             byte
}

// NewPassword はコンストラクタ
func NewPassword(length int) *PasswordSerial {
	p := new(PasswordSerial)
	p.Characters = make([]byte, length, length)
	p.Length = length
	const alpha = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	const num = "1234567890"
	const mark = "!#$%&'()=+-*/{}^~¥|?_<>,."
	p.Arr = strings.Split(alpha+num+mark, "")
	p.Max = byte(len(p.Arr))
	return p
}

// Next 次へ
func (p *PasswordSerial) Next() bool {
	p.Characters[0]++
	isLeap := false
	for wk := 0; wk <= p.CurrentPosition; wk++ {
		if p.Characters[wk] != p.Max {
			continue
		}
		p.Characters[wk] = 0
		if (wk + 1) > p.CurrentPosition {
			isLeap = true
			break
		}
		p.Characters[wk+1]++
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
func (p *PasswordSerial) GetPassword() string {
	var result strings.Builder
	for wk := p.CurrentPosition; wk >= 0; wk-- {
		result.WriteString(p.Arr[p.Characters[wk]])
	}
	return result.String()
}
