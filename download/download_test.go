package download

import (
	"testing"

	"github.com/aws/aws-sdk-go/aws"

	"github.com/unirita/s3dladapter/config"
)

func TestCreateConf(t *testing.T) {
	config.Log.LogDebug = config.Log_Flag_ON
	config.Log.LogSigning = config.Log_Flag_ON
	config.Log.LogHTTPBody = config.Log_Flag_OFF
	config.Log.LogRequestRetries = config.Log_Flag_OFF
	config.Log.LogRequestErrors = config.Log_Flag_ON

	c := createConf()
	if !c.LogLevel.AtLeast(aws.LogDebug) {
		t.Error("LogDebugがONになっていない。")
	}
	if !c.LogLevel.Matches(aws.LogDebugWithSigning) {
		t.Error("LogSigningがONになっていない。")
	}
	if c.LogLevel.Matches(aws.LogDebugWithHTTPBody) {
		t.Error("LogHTTPBodyがOFFになっていない。")
	}
	if c.LogLevel.Matches(aws.LogDebugWithRequestRetries) {
		t.Error("LogRequestRetriesがOFFになっていない。")
	}
	if !c.LogLevel.Matches(aws.LogDebugWithRequestErrors) {
		t.Error("LogRequestErrorsがONになっていない。")
	}
}
