// Copyright 2015 unirita Inc.
// Created 2015/10/08 kazami

package console

import (
	"testing"

	"s3dladapter/testutil"
)

func TestDisplay_メッセージを出力できる_引数なし(t *testing.T) {
	c := testutil.NewStdoutCapturer()
	c.Start()

	Display("ARG001E")

	output := c.Stop()

	if output != "INVALID ARGUMENT.\n" {
		t.Errorf("stdoutへの出力値[%s]が想定と違います。", output)
	}
}

func TestGetMessage_メッセージを文字列として取得できる_引数なし(t *testing.T) {
	msg := GetMessage("ARG001E")
	if msg != "INVALID ARGUMENT." {
		t.Errorf("取得値[%s]が想定と違います。", msg)
	}
}
