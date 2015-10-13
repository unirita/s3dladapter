// Copyright 2015 unirita Inc.
// Created 2015/10/13 kazami

package util

import (
	"os"
)

//パスの存在チェック
func PathExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}
