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
