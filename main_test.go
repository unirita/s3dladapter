package main

import (
	"errors"
	"flag"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/unirita/s3dladapter/console"
	"github.com/unirita/s3dladapter/download"
	"github.com/unirita/s3dladapter/testutil"
)

func makeDownloadSuccess() {
	doDownload = func(bucket string, key string) error {
		return nil
	}
}

func makeDownloadFail() {
	doDownload = func(bucket string, key string) error {
		return errors.New("error")
	}
}

func restoreDownloadFunc() {
	doDownload = download.Do
}

func TestFetchArgs_コマンドラインオプションを取得できる(t *testing.T) {
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.PanicOnError)
	os.Args = os.Args[:1]
	os.Args = append(os.Args, "-v", "-b", "bucket", "-k", "file", "-c", "test.ini")
	args := fetchArgs()

	if args.versionFlag != flag_ON {
		t.Error("-vオプションの指定を検出できなかった。")
	}

	if args.bucketName != "bucket" {
		t.Error("-bオプションの指定を検出できなかった。")
	}

	if args.keyName != "file" {
		t.Error("-kオプションの指定を検出できなかった。")
	}

	if args.configPath != "test.ini" {
		t.Error("-cオプションの指定を検出できなかった。")
	}

}

func TestFetchArgs_コマンドラインオプションに値が指定されなかった場合(t *testing.T) {
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.PanicOnError)
	os.Args = os.Args[:1]
	args := fetchArgs()

	if args.versionFlag != flag_OFF {
		t.Error("-vオプションの値が想定と異なっている。")
	}

	if args.bucketName != "" {
		t.Error("-bオプションの値が想定と異なっている。")
	}

	if args.keyName != "" {
		t.Error("-kオプションの値が想定と異なっている。")
	}

	if args.configPath != "" {
		t.Error("-cオプションの値が想定と異なっている。")
	}
}

func TestRealMain_バージョン出力オプションが指定された場合(t *testing.T) {
	c := testutil.NewStdoutCapturer()

	args := new(arguments)
	args.versionFlag = flag_ON

	c.Start()
	rc := realMain(args)
	out := c.Stop()

	if rc != rc_OK {
		t.Errorf("想定外のrc[%d]が返された。", rc)
	}
	if !strings.Contains(out, Version) {
		t.Error("出力内容が想定と違っている。")
		t.Logf("出力: %s", out)
	}
}

func TestRealMain_引数に何も指定されなかった場合(t *testing.T) {
	c := testutil.NewStdoutCapturer()

	args := new(arguments)

	c.Start()
	rc := realMain(args)
	out := c.Stop()

	if rc != rc_ERROR {
		t.Errorf("想定外のrc[%d]が返された。", rc)
	}
	if !strings.Contains(out, console.USAGE) {
		t.Error("出力内容が想定と違っている。")
		t.Logf("出力: %s", out)
	}
}

func TestRealMain_引数にバケット名が指定されていない場合(t *testing.T) {
	c := testutil.NewStdoutCapturer()

	args := new(arguments)
	args.keyName = "file.txt"
	args.configPath = "config.ini"

	c.Start()
	rc := realMain(args)
	out := c.Stop()

	if rc != rc_ERROR {
		t.Errorf("想定外のrc[%d]が返された。", rc)
	}
	if !strings.Contains(out, console.USAGE) {
		t.Error("出力内容が想定と違っている。")
		t.Logf("出力: %s", out)
	}
}

func TestRealMain_引数にファイル名が指定されていない場合(t *testing.T) {
	c := testutil.NewStdoutCapturer()

	args := new(arguments)
	args.bucketName = "bucket"
	args.configPath = "config.ini"

	c.Start()
	rc := realMain(args)
	out := c.Stop()

	if rc != rc_ERROR {
		t.Errorf("想定外のrc[%d]が返された。", rc)
	}
	if !strings.Contains(out, console.USAGE) {
		t.Error("出力内容が想定と違っている。")
		t.Logf("出力: %s", out)
	}
}

func TestRealMain_引数に設定ファイルのパスが指定されていない場合(t *testing.T) {
	c := testutil.NewStdoutCapturer()

	args := new(arguments)
	args.bucketName = "bucket"
	args.keyName = "file.txt"

	c.Start()
	rc := realMain(args)
	out := c.Stop()

	if rc != rc_ERROR {
		t.Errorf("想定外のrc[%d]が返された。", rc)
	}
	if !strings.Contains(out, console.USAGE) {
		t.Error("出力内容が想定と違っている。")
		t.Logf("出力: %s", out)
	}
}

func TestRealMain_引数がバケットのみの場合(t *testing.T) {
	c := testutil.NewStdoutCapturer()

	args := new(arguments)
	args.bucketName = "bucket"

	c.Start()
	rc := realMain(args)
	out := c.Stop()

	if rc != rc_ERROR {
		t.Errorf("想定外のrc[%d]が返された。", rc)
	}
	if !strings.Contains(out, console.USAGE) {
		t.Error("出力内容が想定と違っている。")
		t.Logf("出力: %s", out)
	}
}

func TestRealMain_引数がダウンロードファイルのみの場合(t *testing.T) {
	c := testutil.NewStdoutCapturer()

	args := new(arguments)
	args.keyName = "file"

	c.Start()
	rc := realMain(args)
	out := c.Stop()

	if rc != rc_ERROR {
		t.Errorf("想定外のrc[%d]が返された。", rc)
	}
	if !strings.Contains(out, console.USAGE) {
		t.Error("出力内容が想定と違っている。")
		t.Logf("出力: %s", out)
	}
}

func TestRealMain_引数が設定ファイルのみの場合(t *testing.T) {
	c := testutil.NewStdoutCapturer()

	args := new(arguments)
	args.configPath = "testconfig.ini"

	c.Start()
	rc := realMain(args)
	out := c.Stop()

	if rc != rc_ERROR {
		t.Errorf("想定外のrc[%d]が返された。", rc)
	}
	if !strings.Contains(out, console.USAGE) {
		t.Error("出力内容が想定と違っている。")
		t.Logf("出力: %s", out)
	}
}

func TestRealMain_引数にディレクトリが指定された場合(t *testing.T) {
	c := testutil.NewStdoutCapturer()

	args := new(arguments)
	args.bucketName = "bucket"
	args.keyName = "test/"
	args.configPath = "config.ini"

	c.Start()
	rc := realMain(args)
	out := c.Stop()

	if rc != rc_ERROR {
		t.Errorf("想定外のrc[%d]が返された。", rc)
	}
	if !strings.Contains(out, "DIRECTORY CAN NOT BE SPECIFIED.") {
		t.Error("出力内容が想定と違っている。")
		t.Logf("出力: %s", out)
	}
}

func TestRealMain_存在しない設定ファイルが指定された場合(t *testing.T) {
	c := testutil.NewStdoutCapturer()

	args := new(arguments)
	args.bucketName = "testbucket"
	args.keyName = "testfile"
	args.configPath = "noexistsconf.ini"

	c.Start()
	rc := realMain(args)
	out := c.Stop()

	if rc != rc_ERROR {
		t.Errorf("想定外のrc[%d]が返された。", rc)
	}
	if !strings.Contains(out, "FAILED TO READ CONFIG FILE.") {
		t.Error("出力内容が想定と違っている。")
		t.Logf("出力: %s", out)
	}
}

func TestRealMain_不正な内容の設定ファイルが指定された場合(t *testing.T) {
	c := testutil.NewStdoutCapturer()

	args := new(arguments)
	args.bucketName = "testbucket"
	args.keyName = "testfile"
	args.configPath = filepath.Join("testdata", "error.ini")

	c.Start()
	rc := realMain(args)
	out := c.Stop()

	if rc != rc_ERROR {
		t.Errorf("想定外のrc[%d]が返された。", rc)
	}
	if !strings.Contains(out, "CONFIG PARM IS NOT EXACT FORMAT.") {
		t.Error("出力内容が想定と違っている。")
		t.Logf("出力: %s", out)
	}
}
