package download

import (
	"testing"

	"github.com/aws/aws-sdk-go/aws"

	"github.com/unirita/s3dladapter/config"
)

func TestCreateConf_フラグがONの場合は各ログレベルにONの値が設定される(t *testing.T) {
	config.Log.LogDebug = config.Log_Flag_ON
	config.Log.LogSigning = config.Log_Flag_ON
	config.Log.LogHTTPBody = config.Log_Flag_ON
	config.Log.LogRequestRetries = config.Log_Flag_ON
	config.Log.LogRequestErrors = config.Log_Flag_ON

	c := createConf()
	if !c.LogLevel.AtLeast(aws.LogDebug) {
		t.Error("LogDebugがONになっていない。")
	}
	if !c.LogLevel.AtLeast(aws.LogDebugWithSigning) {
		t.Error("LogSigningがONになっていない。")
	}
	if !c.LogLevel.AtLeast(aws.LogDebugWithHTTPBody) {
		t.Error("LogHTTPBodyがONになっていない。")
	}
	if !c.LogLevel.AtLeast(aws.LogDebugWithRequestRetries) {
		t.Error("LogRequestRetriesがONになっていない。")
	}
	if !c.LogLevel.AtLeast(aws.LogDebugWithRequestErrors) {
		t.Error("LogRequestErrorsがONになっていない。")
	}
}

func TestCreateConf_フラグがOFFの場合は各ログレベルにOFFの値が設定される(t *testing.T) {
	config.Log.LogDebug = config.Log_Flag_OFF
	config.Log.LogSigning = config.Log_Flag_OFF
	config.Log.LogHTTPBody = config.Log_Flag_OFF
	config.Log.LogRequestRetries = config.Log_Flag_OFF
	config.Log.LogRequestErrors = config.Log_Flag_OFF

	c := createConf()
	if !c.LogLevel.Matches(aws.LogOff) {
		t.Error("ログを出力しないようになっていない。")
	}
	if c.LogLevel.AtLeast(aws.LogDebug) {
		t.Error("LogDebugがOFFになっていない。")
	}
	if c.LogLevel.AtLeast(aws.LogDebugWithSigning) {
		t.Error("LogSigningがOFFになっていない。")
	}
	if c.LogLevel.AtLeast(aws.LogDebugWithHTTPBody) {
		t.Error("LogHTTPBodyがOFFになっていない。")
	}
	if c.LogLevel.AtLeast(aws.LogDebugWithRequestRetries) {
		t.Error("LogRequestRetriesがOFFになっていない。")
	}
	if c.LogLevel.AtLeast(aws.LogDebugWithRequestErrors) {
		t.Error("LogRequestErrorsがOFFになっていない。")
	}
}

func TestCreateConf_log_onにOFFが設定された場合は他のログレベルはOFFになる(t *testing.T) {
	config.Log.LogDebug = config.Log_Flag_OFF
	config.Log.LogSigning = config.Log_Flag_ON
	config.Log.LogHTTPBody = config.Log_Flag_ON
	config.Log.LogRequestRetries = config.Log_Flag_ON
	config.Log.LogRequestErrors = config.Log_Flag_ON

	c := createConf()
	if !c.LogLevel.Matches(aws.LogOff) {
		t.Error("ログを出力しないようになっていない。")
	}
	if c.LogLevel.AtLeast(aws.LogDebug) {
		t.Error("LogDebugがOFFになっていない。")
	}
	if c.LogLevel.AtLeast(aws.LogDebugWithSigning) {
		t.Error("LogSigningがOFFになっていない。")
	}
	if c.LogLevel.AtLeast(aws.LogDebugWithHTTPBody) {
		t.Error("LogHTTPBodyがOFFになっていない。")
	}
	if c.LogLevel.AtLeast(aws.LogDebugWithRequestRetries) {
		t.Error("LogRequestRetriesがOFFになっていない。")
	}
	if c.LogLevel.AtLeast(aws.LogDebugWithRequestErrors) {
		t.Error("LogRequestErrorsがOFFになっていない。")
	}
}

func TestCreateConf_log_onのみがON設定された場合はdebugのみがON(t *testing.T) {
	config.Log.LogDebug = config.Log_Flag_ON
	config.Log.LogSigning = config.Log_Flag_OFF
	config.Log.LogHTTPBody = config.Log_Flag_OFF
	config.Log.LogRequestRetries = config.Log_Flag_OFF
	config.Log.LogRequestErrors = config.Log_Flag_OFF

	c := createConf()
	if !c.LogLevel.Matches(aws.LogDebug) {
		t.Error("LogDebugがONになっていない。")
	}
	if c.LogLevel.AtLeast(aws.LogDebugWithSigning) {
		t.Error("LogSigningがOFFになっていない。")
	}
	if c.LogLevel.AtLeast(aws.LogDebugWithHTTPBody) {
		t.Error("LogHTTPBodyがOFFになっていない。")
	}
	if c.LogLevel.AtLeast(aws.LogDebugWithRequestRetries) {
		t.Error("LogRequestRetriesがOFFになっていない。")
	}
	if c.LogLevel.AtLeast(aws.LogDebugWithRequestErrors) {
		t.Error("LogRequestErrorsがOFFになっていない。")
	}
}
