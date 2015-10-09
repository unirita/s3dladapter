package console

import (
	"fmt"
)

// USAGE表示用の定義メッセージ
const USAGE = `Usage :
    s3dladapter.exe [-v] [-b Bucket] [-f File]
Option :
    -v                   :   Print s3dladapter version.
	-b bucket name       :   Designate a buket name.
    -f downloadFile name :   Designate a file name.(Without extensions.)
	
	-b, -f, is a required input.
Copyright 2015 unirita Inc.
`

var stack_msg = []string{"ARG001E"}

// コンソールメッセージ一覧
var msgs = map[string]string{
	"ARG001E": "INVALID ARGUMENT.",
}

// 標準出力へメッセージコードcodeに対応したメッセージを表示する。
//
// param : code メッセージコードID。
//
// return : 出力文字数。
//
// return : エラー情報。
func Display(code string) {
	msg := GetMessage(code)
	for _, s := range stack_msg {
		if code == s {
			fmt.Println(msg)
		}
	}
}

// 出力メッセージを文字列型で取得する。
//
// param : code メッセージコードID。
//
//
// return : 取得したメッセージ
func GetMessage(code string) string {
	return fmt.Sprintf("%s", msgs[code])
}
