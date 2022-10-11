package sms_charset

import (
	"errors"
)

const (
	charset_gsm = iota
	charset_ucs
)

type DetectResult {
	text string
	partsCount int
	symbolsPerPart int
	charset int
}

func IsGsm(dr DetectResult) bool {
	return dr.charset == charset_gsm
}

func PartsCount(dr DetectResult) int {
    return dr.partsCount
}

func (dr *DetectResult) detect() {
    // TODO
}

func Detect(smsText string) (DetectResult, error) {
	// TODO implement
	var dr := DetectResult{
    	smsText,
        1,
        0,
	    charset_gsm,
	}
	dr.detect()

	return dr, nil
}
