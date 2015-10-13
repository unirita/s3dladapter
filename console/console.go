// Copyright 2015 unirita Inc.
// Created 2015/10/09 kazami

package console

import (
	"fmt"
)

// USAGE表示用の定義メッセージ
const USAGE = `Usage :
    s3dladapter.exe [-v] [-b Bucket] [-f File] [-c configPath]
Option :
    -v                   :   Print s3dladapter version.
    -b bucket name       :   Designate a buket name.
    -f downloadFile name :   Designate a file name.(Without extensions.)
    -c configFile Path   :   Designate configFile Path.
	
    -c default is same path as the executable file.
    -b, -f, is a required input.
Copyright 2015 unirita Inc.
`

// コンソールメッセージ一覧
var msgs = map[string]string{
	"ARG001E": "INVALID ARGUMENT.",
	"CON001E": "FAILED TO READ CONFIG FILE.",
	"CON002E": "CONFIG PARM IS NOT EXACT FORMAT.",
}

// 標準出力へメッセージコードcodeに対応したメッセージを表示する。
//
// param : code メッセージコードID。
//
// return : 出力文字数。
//
// return : エラー情報。
//func Display(code string) {
//	msg := GetMessage(code)
//	for _, s := range stack_msg {
//		if code == s {
//			fmt.Println(msg)
//		}
//	}
//}

func Display(code string, a ...interface{}) (int, error) {
	msg := GetMessage(code, a...)

	return fmt.Println(msg)
}

// 出力メッセージを文字列型で取得する。
//
// param : code メッセージコードID。
//
//
// return : 取得したメッセージ
//func GetMessage(code string) string {
//	return fmt.Sprintf("%s", msgs[code])
//}

func GetMessage(code string, a ...interface{}) string {
	return fmt.Sprintf("%s %s", code, fmt.Sprintf(msgs[code], a...))
}
