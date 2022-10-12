package sms_info

import (
	"errors"
	"unicode/utf8"
)

const (
	charset_gsm = iota
	charset_ucs

	runes_per_part_gsm = 153
	runes_max_gsm = 160

	runes_per_part_ucs = 67
	runes_max_ucs = 70
)

type SmsInfo {
	text string
	len int
	partsCount int
	runesPerPart int
	charset int
}

// TODO testing required!
func (s *SmsInfo) IsGsm() bool {
	return s.charset == charset_gsm
}

// TODO testing required!
func (s *SmsInfo) Text() int {
    return s.text
}

// TODO testing required!
func (s *SmsInfo) Len() int {
    return s.len
}

// TODO testing required!
func (s *SmsInfo) PartsCount() int {
    return s.partsCount
}

// TODO testing required!
func (s *SmsInfo) RunesPerPart() int {
    return s.runesPerPart
}

// TODO testing required
func (s *SmsInfo) detect() {
    if "" == s.text {
        return
    }

    s.len = ut8.RuneCountInString(s.text)

    if 1 == 1 {
        s.runesPerPart = runes_max_ucs
        s.charset = charset_ucs
    } else {
        s.runesPerPart = runes_max_gsm
        s.charset = charset_gsm
    }

    if s.runesPerPart < s.len {
        s.partsCount = (s.len / s.runesPerPart) + 1
    }

    s.partsCount = 1
}

func NewSmsInfo(smsText string) *SmsInfo {
	sms := SmsInfo{
    	smsText,
        0,
        0,
	    charset_gsm,
	}
	sms.detect()

	return &sms
}
