package util

import (
	"testing"
)

func TestPathExists_ローカルパスの存在確認(t *testing.T) {
	nonExistPath := "C:\\HOGEHOGEAAABBB"
	existPath := "C:\\"

	if PathExists(existPath) == false {
		t.Errorf("パスのチェックが間違っています")
	}

	if PathExists(nonExistPath) == true {
		t.Errorf("パスのチェックが間違っています")
	}
}
