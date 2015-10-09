package main

import (
	"os"
	"strings"
	"testing"

	"s3dladapter/testutil"
)

var testDataDir string

func TestFetchArgs_コマンドラインオプションを取得できる(t *testing.T) {
	os.Args = append(os.Args, "-v", "-b", "bucket", "-f", "file")
	args := fetchArgs()

	if args.versionFlag != flag_ON {
		t.Error("-vオプションの指定を検出できなかった。")
	}

	if args.bucketName != "bucket" {
		t.Error("-bオプションの指定を検出できなかった。")
	}

	if args.fileName != "file" {
		t.Error("-fオプションの指定を検出できなかった。")
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

}

func TestRealMain_引数が指定された場合(t *testing.T) {

}
