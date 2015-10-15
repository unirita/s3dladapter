package download

import (
	"testing"

	"s3dladapter/config"
)

func TestDownload_認証情報正しくない場合はエラー(t *testing.T) {
	testConfig := "noexistS3.ini"
	if err := config.Load(testConfig); err != nil {
		t.Errorf("テストファイルの読み込みに失敗")
	}

	bucket := "testbucketuniritanewautomation"
	file := "test1.txt"

	if err := Download(bucket, file); err == nil {
		t.Errorf("期待するエラーが起こっていない")
	}
}

func TestDownload_S3に指定したバケット名が存在しない場合はエラー(t *testing.T) {
	testConfig := "existS3.ini"
	if err := config.Load(testConfig); err != nil {
		t.Errorf("テストファイルの読み込みに失敗")
	}

	bucket := "noexistBucket"
	file := "test1.txt"

	if err := Download(bucket, file); err == nil {
		t.Errorf("期待するエラーが起こっていない")
	}
}

func TestDownload_S3に指定したキー名のファイルが存在しない場合はエラー(t *testing.T) {
	testConfig := "existS3.ini"
	if err := config.Load(testConfig); err != nil {
		t.Errorf("テストファイルの読み込みに失敗")
	}

	bucket := "noexistBucket"
	file := "noexist.txt"

	if err := Download(bucket, file); err == nil {
		t.Errorf("期待するエラーが起こっていない")
	}
}
