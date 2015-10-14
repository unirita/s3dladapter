package download

import (
	"testing"

	"s3dladapter/config"
)

<<<<<<< HEAD
=======
func generateTestConfig() {
	config.Aws.AccessKeyId = ``
	config.Aws.SecletAccessKey = ``
	config.Aws.Region = ``
	config.Download.DownloadDir = `c:\TEST`
}

>>>>>>> c7749dee8ed85f50a6503549c9a941164558a612
func TestDownload_S3にダウンロードするファイルが存在する場合はエラーじゃない(t *testing.T) {
	testConfig := `downloadtest.ini`
	if err := config.Load(testConfig); err != nil {
		t.Error("テスト用のconfigが存在しない")
	}

	testBucket := "testbucketuniritanewautomation"
	testFile := "test1.txt"

	if err := Download(testBucket, testFile); err != nil {
		t.Errorf("想定外のエラーが発生した： %s", err)
	}
}

func TestDownload_S3にダウンロードするファイルが存在しない場合はエラー(t *testing.T) {
	testConfig := `downloadtest.ini`
	if err := config.Load(testConfig); err != nil {
		t.Error("テスト用のconfigが存在しない")
	}

	testBucket := "noexistBucket"
	testFile := "noexistFile"

	if err := Download(testBucket, testFile); err == nil {
		t.Error("エラーが発生しなかった")
	}
}
