package sms_info

import (
	"testing"
)

type testCase struct {
	description          string
	inputText            string
	expectedLen          int
	expectedPartsCount   int
	expectedRunesPerPart int
	expectedCharsetIsGsm bool
}

var smsTexts = []testCase{
	{
		"empty sms",
		"",
		0,
		0,
		0,
		true,
	},
	{
		"one part in gsm",
		"Hello! This short message tells you nothing",
		43,
		1,
		153,
		true,
	},
	{
		"one part in ucs",
		"Привет! Это короткое сообщение не имеет смысла",
		46,
		1,
		67,
		false,
	},
	{
		"several parts in gsm",
		"Hello! This is a very-very long message for sms service. Though it tells you nothing, it still must be split into 2 parts to be sent over sms as it's length exceeds the limit for single sms",
		189,
		2,
		153,
		true,
	},
	{
		"several parts in ucs",
		"Привет! Это довольно длинное смс-сообщение. Хотя оно не имеет смысла, оно все равно состоит из двух частей",
		106,
		2,
		67,
		false,
	},
	{
		"two positions chars in gsm, one sms",
		"He[]o! This short message tells you ^thing, but costs 30€",
		61,
		1,
		153,
		true,
	},
	{
		"two positions chars in ucs, one sms",
		"Привет| Это короткое смс~сообщение. Но оно стоит 30€",
		55,
		1,
		67,
		false,
	},
	{
		"two positions chars in gsm, several parts",
		"He[]o! This is a very long message for sms service. Though it tells you nothing, it still must be split into 2 parts to be sent over sms as it~s length exceeds the limit for single sms",
		187,
		2,
		153,
		true,
	},
	{
		"two positions chars in ucs, several parts",
		"Привет| Это довольно длинное смс~сообщение. Оно не имеет смысла, но все равно состоит из двух частей и стоит 30€",
		115,
		2,
		67,
		false,
	},
}

func TestSmsInfo(t *testing.T) {
	for _, tc := range smsTexts {
		t.Run(tc.description, func(t *testing.T) {
			smsInfo := NewSmsInfo(tc.inputText)
			if smsInfo.Text() != tc.inputText {
				t.Errorf("Expected text: '%s', got: '%s'", tc.inputText, smsInfo.Text())
			}

			if smsInfo.Len() != tc.expectedLen {
				t.Errorf("Expected text len: %d, got: %d", tc.expectedLen, smsInfo.Len())
			}

			if smsInfo.PartsCount() != tc.expectedPartsCount {
				t.Errorf("Expected parts count: %d, got: %d", tc.expectedPartsCount, smsInfo.PartsCount())
			}

			if smsInfo.RunesPerPart() != tc.expectedRunesPerPart {
				t.Errorf("Expected runes per part: %d, got: %d", tc.expectedRunesPerPart, smsInfo.RunesPerPart())
			}

			if smsInfo.IsGsm() != tc.expectedCharsetIsGsm {
				t.Errorf("Expected isGsm to be %v, got: %v", tc.expectedCharsetIsGsm, smsInfo.IsGsm())
			}
		})
	}
}
