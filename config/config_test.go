package config

import (
	"strings"
	"testing"
)

func generateTestConfig() {
	Aws.AccessKeyId = `testkeyid`
	Aws.SecletAccessKey = `seclettestkey`
	Aws.Region = `ap-northeast-1`
	Download.DownloadDir = `c:\TEST`
	Log.LogLevel = 5
}

func TestLoad_存在しないファイルをロードしようとした場合はエラー(t *testing.T) {
	if err := Load("noexistfilepath"); err == nil {
		t.Error("エラーが発生していない。")
	}
}

func TestLoadByReader_Readerから設定値を取得できる(t *testing.T) {
	conf := `
[aws]
access_key_id='testkeyid'
secret_access_key='seclettestkey'
region='ap-northeast-1'
[download]
download_dir='c:\TEST'
[log]
loglevel=5
`

	r := strings.NewReader(conf)
	err := loadReader(r)
	if err != nil {
		t.Fatalf("想定外のエラーが発生した[%s]", err)
	}

	if Aws.AccessKeyId != `testkeyid` {
		t.Errorf("access_key_idの値[%s]は想定と違っている。", Aws.AccessKeyId)
	}
	if Aws.SecletAccessKey != `seclettestkey` {
		t.Errorf("seclet_access_keyの値[%s]は想定と違っている。", Aws.SecletAccessKey)
	}
	if Aws.Region != `ap-northeast-1` {
		t.Errorf("regionの値[%s]は想定と違っている。", Aws.Region)
	}
	if Download.DownloadDir != `c:\TEST` {
		t.Errorf("download_dirの値[%s]は想定と違っている。", Download.DownloadDir)
	}

}

func TestLoadByReader_tomlの書式に沿っていない場合はエラーが発生する(t *testing.T) {
	conf := `
[aws]
access_key_id=testkeyid
seclet_access_key=seclettestkey
region='ap-northeast-1'
[download]
download_dir='c:\TEST'
[log]
loglevel=5
`

	r := strings.NewReader(conf)
	err := loadReader(r)
	if err == nil {
		t.Error("エラーが発生しなかった")
	}

}

func TestDetectError_設定内容にエラーが無い場合はnilを返す(t *testing.T) {
	generateTestConfig()
	if err := DetectError(); err != nil {
		t.Errorf("想定外のエラーが発生した： %s", err)
	}
}

func TestDetectError_設定ファイルのアクセスキーIDが空の場合はエラー(t *testing.T) {
	generateTestConfig()
	Aws.AccessKeyId = ``
	if err := DetectError(); err == nil {
		t.Error("エラーが発生しなかった。")
	}
}

func TestDetectError_設定ファイルのシークレットアクセスキーが空の場合はエラー(t *testing.T) {
	generateTestConfig()
	Aws.SecletAccessKey = ``
	if err := DetectError(); err == nil {
		t.Error("エラーが発生しなかった。")
	}
}

func TestDetectError_設定ファイルのリージョンが空の場合はエラー(t *testing.T) {
	generateTestConfig()
	Aws.Region = ``
	if err := DetectError(); err == nil {
		t.Error("エラーが発生しなかった。")
	}
}

func TestDetectError_ダウンロード保存先パスが存在しなかった場合はエラー(t *testing.T) {
	generateTestConfig()
	Download.DownloadDir = `C:\EEEE`
	if err := DetectError(); err == nil {
		t.Error("エラーが発生しなかった。")
	}
}

func TestDetectError_loglevelの値が最大(t *testing.T) {
	generateTestConfig()
	Log.LogLevel = 5
	if err := DetectError(); err != nil {
		t.Errorf("期待していないエラーが発生した")
	}

	if Log.LogLevel != 5 {
		t.Errorf("loglevelに正しい値が入っていない[%d]", Log.LogLevel)
	}
}

func TestDetectError_範囲外の数値の場合はエラー_最大越え(t *testing.T) {
	generateTestConfig()
	Log.LogLevel = 6
	if err := DetectError(); err == nil {
		t.Errorf("エラーが発生しなかった。")
	}

}

func TestDetectError_loglevelの値が最小(t *testing.T) {
	generateTestConfig()
	Log.LogLevel = 0
	if err := DetectError(); err != nil {
		t.Errorf("期待していないエラーが発生した")
	}

	if Log.LogLevel != 0 {
		t.Errorf("loglevelに正しい値が入っていない[%d]", Log.LogLevel)
	}
}

func TestDetectError_範囲外の数値の場合はエラー_マイナス(t *testing.T) {
	generateTestConfig()
	Log.LogLevel = -1
	if err := DetectError(); err == nil {
		t.Errorf("エラーが発生しなかった。")
	}

}

func TestpathExists_ローカルパスの存在確認(t *testing.T) {
	nonExistPath := "C:\\HOGEHOGEAAABBB"
	existPath := "C:\\"

	if pathExists(existPath) == false {
		t.Errorf("パスのチェックが間違っています")
	}

	if pathExists(nonExistPath) == true {
		t.Errorf("パスのチェックが間違っています")
	}
}
