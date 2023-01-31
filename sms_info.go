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

var gsmChars = map[string]int{
	"@": 1, "Δ": 1, " ": 1, "0": 1, "¡": 1, "P": 1, "¿": 1, "p": 1,
	"£": 1, "_": 1, "!": 1, "1": 1, "A": 1, "Q": 1, "a": 1, "q": 1,
	"$": 1, "Φ": 1, `"`: 1, "2": 1, "B": 1, "R": 1, "b": 1, "r": 1,
	"¥": 1, "Γ": 1, "#": 1, "3": 1, "C": 1, "S": 1, "c": 1, "s": 1,
	"è": 1, "Λ": 1, "¤": 1, "4": 1, "D": 1, "T": 1, "d": 1, "t": 1,
	"é": 1, "Ω": 1, "%": 1, "5": 1, "E": 1, "U": 1, "e": 1, "u": 1,
	"ù": 1, "Π": 1, "&": 1, "6": 1, "F": 1, "V": 1, "f": 1, "v": 1,
	"ì": 1, "Ψ": 1, "'": 1, "7": 1, "G": 1, "W": 1, "g": 1, "w": 1,
	"ò": 1, "Σ": 1, "(": 1, "8": 1, "H": 1, "X": 1, "h": 1, "x": 1,
	"Ç": 1, "Θ": 1, ")": 1, "9": 1, "I": 1, "Y": 1, "i": 1, "y": 1,
	"\n": 1, "Ξ": 1, "*": 1, ":": 1, "J": 1, "Z": 1, "j": 1, "z": 1,
	"Ø": 1, "\x1B": 1, "+": 1, ";": 1, "K": 1, "Ä": 1, "k": 1, "ä": 1,
	"ø": 1, "Æ": 1, ",": 1, "<": 1, "L": 1, "Ö": 1, "l": 1, "ö": 1,
	"\r": 1, "æ": 1, "-": 1, "=": 1, "M": 1, "Ñ": 1, "m": 1, "ñ": 1,
	"Å": 1, "ß": 1, ".": 1, ">": 1, "N": 1, "Ü": 1, "n": 1, "ü": 1,
	"å": 1, "É": 1, "/": 1, "?": 1, "O": 1, "§": 1, "o": 1, "à": 1,
	//  Extension set characters, each occupies two positions in sms
	"|": 1, "^": 1, "€": 1, "{": 1, "}": 1, "[": 1, "~": 1, "]": 1, "\\": 1,
}

var extensionCharsRe *regexp.Regexp

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

func (s *SmsInfo) setProps() {
	if "" == s.text {
		return
	}

	s.len = countTextLen(s.text)
	var maxLen int

	if hasUcsChar(s.text) {
		s.runesPerPart = runes_per_part_ucs
		s.charset = charset_ucs
		maxLen = runes_max_ucs
	} else {
		s.runesPerPart = runes_per_part_gsm
		s.charset = charset_gsm
		maxLen = runes_max_gsm
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

func hasUcsChar(text string) bool {
	for _, symbol := range text {
		if _, ok := gsmChars[string(symbol)]; !ok {
			return true
		}
	}

	return false
}

func countTextLen(text string) int {
	textLen := utf8.RuneCountInString(text)

	if nil == extensionCharsRe {
		extensionCharsRe = regexp.MustCompile(`[~€|^\[\]{}\\]+`)
	}

	for _, match := range extensionCharsRe.FindAllString(text, -1) {
		textLen += utf8.RuneCountInString(match)
	}

	return textLen
}
