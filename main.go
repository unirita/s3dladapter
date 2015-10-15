package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"s3dladapter/config"
	"s3dladapter/console"
	"s3dladapter/download"
)

// 実行時引数のオプション
type arguments struct {
	versionFlag bool   // バージョン情報表示フラグ
	bucketName  string //バケット名
	fileName    string //S3からダウンロードするファイル名
	configPath  string //設定ファイルのパス
}

//バッチプログラムの戻り値
const (
	rc_OK    = 0
	rc_ERROR = 1
)

// フラグ系実行時引数のON/OFF
const (
	flag_ON  = true
	flag_OFF = false
)

func main() {
	args := fetchArgs()
	rc := realMain(args)
	os.Exit(rc)
}

func realMain(args *arguments) int {
	if args.versionFlag == flag_ON {
		showVersion()
		return rc_OK
	}

	if args.configPath == "" || args.bucketName == "" || args.fileName == "" {
		showUsage()
		return rc_ERROR
	}

	if strings.HasSuffix(args.fileName, "/") {
		console.Display("ADP001E")
	}

	if err := config.Load(args.configPath); err != nil {
		console.Display("ADP002E", err)
		return rc_ERROR
	}

	if err := config.DetectError(); err != nil {
		console.Display("ADP003E", err)
		return rc_ERROR
	}

	//設定ファイルを読み込んだ情報でS3に接続してダウンロード
	if err := download.Download(args.bucketName, args.fileName); err != nil {
		console.Display("ADP004E", err)
		return rc_ERROR
	}

	return rc_OK
}

// コマンドライン引数を解析し、arguments構造体を返す。
func fetchArgs() *arguments {
	args := new(arguments)
	flag.Usage = showUsage
	flag.BoolVar(&args.versionFlag, "v", false, "version option")
	flag.StringVar(&args.bucketName, "b", "", "Designate bucket option")
	flag.StringVar(&args.fileName, "f", "", "Designate download file option")
	flag.StringVar(&args.configPath, "c", "", "Designate config file option")
	flag.Parse()
	return args
}

// バージョンを表示する。
func showVersion() {
	fmt.Printf("%s\n", Version)
}

// オンラインヘルプを表示する。
func showUsage() {
	fmt.Print(console.USAGE)
}
