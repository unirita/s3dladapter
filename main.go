package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/unirita/s3dladapter/config"
	"github.com/unirita/s3dladapter/console"
	"github.com/unirita/s3dladapter/download"
)

// 実行時引数のオプション
type arguments struct {
	versionFlag bool   // バージョン情報表示フラグ
	bucketName  string //バケット名
	keyName     string //S3からダウンロードするキー名
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

type downloadFunc func(string, string) error

var doDownload downloadFunc = download.Do

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

	if args.configPath == "" || args.bucketName == "" || args.keyName == "" {
		showUsage()
		return rc_ERROR
	}

	if strings.HasSuffix(args.keyName, "/") {
		console.Display("DLA001E")
		return rc_ERROR
	}

	if err := config.Load(args.configPath); err != nil {
		console.Display("DLA002E", err)
		return rc_ERROR
	}

	if err := config.DetectError(); err != nil {
		console.Display("DLA003E", err)
		return rc_ERROR
	}

	//設定ファイルを読み込んだ情報でS3に接続してダウンロード
	if err := doDownload(args.bucketName, args.keyName); err != nil {
		console.Display("DLA004E", err)
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
	flag.StringVar(&args.keyName, "k", "", "Designate download key option")
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
