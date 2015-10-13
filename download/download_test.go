package download

import (
	"testing"

	"s3dladapter/config"
)

func generateTestConfig() {
	config.Aws.AccessKeyId = `AKIAI47YWI4JPYD5XKNA`
	config.Aws.SecletAccessKey = `mPZn9lBNdrtnoEbIZWV51pjvd2kySrCOgmYBKqji`
	config.Aws.Region = `ap-northeast-1`
	config.Download.DownloadDir = `c:\TEST`
}

func TestDownload_S3にダウンロードするファイルが存在する場合はエラーじゃない(t *testing.T) {
	generateTestConfig()

	testBucket := "testbucketuniritanewautomation"
	testFile := "test1.txt"

	if err := Download(testBucket, testFile); err != nil {
		t.Errorf("想定外のエラーが発生した： %s", err)
	}
}

func TestDownload_S3にダウンロードするファイルが存在しない場合はエラー(t *testing.T) {
	generateTestConfig()

	testBucket := "noexistBucket"
	testFile := "noexistFile"

	if err := Download(testBucket, testFile); err == nil {
		t.Error("エラーが発生しなかった")
	}
}
