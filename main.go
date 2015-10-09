package main

import (
	"flag"
	"fmt"
	"os"

	"s3dladapter/config"
	"s3dladapter/console"
	"s3dladapter/download"
)

// 実行時引数のオプション
type arguments struct {
	versionFlag bool   // バージョン情報表示フラグ
	bucketName  string //バケット名
	fileName    string //S3からダウンロードするファイル名
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

	if args.versionFlag == flag_OFF && args.bucketName == "" && args.fileName == "" {
		showUsage()
		return rc_OK
	}

	if args.bucketName == "" || args.fileName == "" {
		console.Display("ARG001E")
		return rc_ERROR
	}

	//TODO:　設定ファイル読み込み
	conf := config.Config{
		AccessKeyId:      "<AWS_ACCESS_KEY_ID>",
		SecletAccessKey:  "<AWS_SECRET_ACCESS_KEY>",
		Region:           "ap-northeast-1",
		DownloadLocation: "C:\\TEST",
	}

	//設定ファイルを読み込んだ情報でS3に接続してダウンロード
	download.Download(conf, args.bucketName, args.fileName)

	rc := rc_OK
	return rc
}

// コマンドライン引数を解析し、arguments構造体を返す。
func fetchArgs() *arguments {
	args := new(arguments)
	flag.Usage = showUsage
	flag.BoolVar(&args.versionFlag, "v", false, "version option")
	flag.StringVar(&args.bucketName, "b", "", "Designate bucket option")
	flag.StringVar(&args.fileName, "f", "", "Designate download file option")
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
