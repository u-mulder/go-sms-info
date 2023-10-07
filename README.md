# sms_info 

`sms-info` provides essential data about sms message using text of this message as input data.

## Installation

-- TODO

> go get github.com/u-mulder/go-sms-info

## Usage

```golang
import (
    "fmt"
    smsInfo "github.com/u-mulder/go-sms-info"
)

smsText := "Some sms text goes here"

// Create new SmsInfo struct using smsText as input
smsInfo := smsInfo.NewSmsInfo(tc.inputText)
			
// Output struct
fmt.Printf("SmsInfo %v", smsInfo)

// Output struct's fields
fmt.Printf(
    "Sms with text '%s' has len %d, is split in %d part(s) with %d max runes per part. All chars of sms are encoded using gsm charset: %t",
    smsInfo.Text(),
    smsInfo.Len(),
    smsInfo.PartsCount(),
    smsInfo.RunesPerPart(),
    smsInfo.IsGsm(),
)
```

## Misc

Detecting charset in php can be found [here](https://github.com/u-mulder/sms-charset-detector).