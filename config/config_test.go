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
