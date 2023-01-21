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

var smsTexts = map[string]TestCase{
    testCase{
        "empty sms",
        "",
        0,
        0,
        0,
        true,
    },
//     testCase{
//         "one part in gsm",
//         "",
//         0,
//         1,
//         // ?,
//         true,
//     },
//     testCase{
//         "one part in ucs",
//         "",
//         0,
//         1,
//         // ?,
//         false,
//     },
//     testCase{
//         "several parts in gsm",
//         "",
//         0,
//         2,
//         // ?,
//         true,
//     },
//     testCase{
//         "several parts in ucs",
//         "",
//         0,
//         2,
//         // ?,
//         false,
//     },
}

func TestSmsInfo(t *testing.T) {
	for _, tc := range smsTexts {
		t.Run(tc.description, func(t *testing.T) {
			smsInfo := NewSmsInfo(tc.inputText)
			if smsInfo.Text() != tc.inputText {
				t.Errorf("Expected text: '%s', got: '%s'", tc.inputText, smsInfo.Text())
			}

			if smsInfo.Len() != tc.expectedLen {
				t.Errorf("Expected text len: %d, got: %s", tc.expectedLen, smsInfo.Len())
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
