package sms_info

import (
	"regexp"
	"unicode/utf8"
)

const (
	charset_gsm = iota
	charset_ucs

	runes_per_part_gsm = 153
	runes_max_gsm      = 160

	runes_per_part_ucs = 67
	runes_max_ucs      = 70
)

type SmsInfo struct {
	text         string
	len          int
	partsCount   int
	runesPerPart int
	charset      int
}

var gsmRegexp *regexp.Regexp

func (s *SmsInfo) IsGsm() bool {
	return s.charset == charset_gsm
}

func (s *SmsInfo) Text() string {
	return s.text
}

func (s *SmsInfo) Len() int {
	return s.len
}

func (s *SmsInfo) PartsCount() int {
	return s.partsCount
}

func (s *SmsInfo) RunesPerPart() int {
	return s.runesPerPart
}

// TODO implement and testing required
func (s *SmsInfo) setProps() {
	if "" == s.text {
		return
	}

	s.len = utf8.RuneCountInString(s.text)
	var maxLen int

	if getGsmRegexp().MatchString(s.text) {
		s.runesPerPart = runes_per_part_gsm
		s.charset = charset_gsm
		maxLen = runes_max_gsm
	} else {
		s.runesPerPart = runes_per_part_ucs
		s.charset = charset_ucs
		maxLen = runes_max_ucs
	}

	if maxLen < s.len {
		s.partsCount = (s.len / s.runesPerPart) + 1
	} else {
		s.partsCount = 1
	}
}

func NewSmsInfo(smsText string) *SmsInfo {
	sms := SmsInfo{
		smsText,
		0,
		0,
		0,
		charset_gsm,
	}
	sms.setProps()

	return &sms
}

func getGsmRegexp() *regexp.Regexp {
	if gsmRegexp == nil {
		gsmRegexp = regexp.MustCompile(`[a-z]+`)
	}

	return gsmRegexp
}
