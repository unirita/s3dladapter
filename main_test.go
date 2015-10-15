package main

import (
	"os"
	"strings"
	"testing"

	"s3dladapter/console"
	"s3dladapter/testutil"
)

var testDataDir string

func TestFetchArgs_コマンドラインオプションを取得できる(t *testing.T) {
	os.Args = append(os.Args, "-v", "-b", "bucket", "-f", "file", "-c", "test.ini")
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

	if args.configPath != "test.ini" {
		t.Error("-cオプションの指定を検出できなかった。")
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
	args.fileName = "file"

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

func TestRealMain_存在しない設定ファイルが指定された場合(t *testing.T) {
	c := testutil.NewStdoutCapturer()

	args := new(arguments)
	args.bucketName = "testbucket"
	args.fileName = "testfile"
	args.configPath = "noexists.ini"

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
	args.fileName = "testfile"
	args.configPath = "configerror.ini"

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

func TestRealMain_s3にダウンロードファイルが無い場合(t *testing.T) {
	c := testutil.NewStdoutCapturer()

	args := new(arguments)
	args.bucketName = "testbucket"
	args.fileName = "testfile"
	args.configPath = "s3dladapter.ini"

	c.Start()
	rc := realMain(args)
	out := c.Stop()

	if rc != rc_ERROR {
		t.Errorf("想定外のrc[%d]が返された。", rc)
	}
	if !strings.Contains(out, "DOWNLOAD FAILED.") {
		t.Error("出力内容が想定と違っている。")
		t.Logf("出力: %s", out)
	}
}
