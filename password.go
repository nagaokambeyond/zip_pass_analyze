package main

import (
	"strings"
)

// Password はパスワードの構造体
type Password struct {
	DigitCharacter      []byte
	PasswordLength      int
	DigitPosition       int
	CharacterType       []string
	CharacterTypeLength byte
}

// NewPassword はコンストラクタ
func NewPassword(length int) *Password {
	p := new(Password)
	p.DigitCharacter = make([]byte, length, length)
	p.PasswordLength = length
	const alpha = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	const num = "1234567890"
	const mark = "!#$%&'()=+-*/{}^~¥|?_<>,."
	p.CharacterType = strings.Split(alpha+num+mark, "")
	p.CharacterTypeLength = byte(len(p.CharacterType))
	return p
}

// Next 次へ
func (p *Password) Next() bool {
	p.DigitCharacter[0]++
	isLeap := false
	for wk := 0; wk <= p.DigitPosition; wk++ {
		if p.DigitCharacter[wk] != p.CharacterTypeLength {
			continue
		}
		p.DigitCharacter[wk] = 0
		if (wk + 1) > p.DigitPosition {
			isLeap = true
			break
		}
		p.DigitCharacter[wk+1]++
	}
	if isLeap {
		p.DigitPosition++
		if p.PasswordLength <= p.DigitPosition {
			return false
		}
	}

	return true
}

// GetPassword もじ
func (p *Password) GetPassword() string {
	var result strings.Builder
	for wk := p.DigitPosition; wk >= 0; wk-- {
		result.WriteString(p.CharacterType[p.DigitCharacter[wk]])
	}
	return result.String()
}
