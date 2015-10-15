// Copyright 2015 unirita Inc.
// Created 2015/10/09 kazami

package download

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	//"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/defaults"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"

	"s3dladapter/config"
)

//ダウンロードするオブジェクトの構造体
type downloader struct {
	*s3manager.Downloader
	bucket string
	file   string
	dir    string
}

// S3からファイルをダウンロードする
//
// 引数: bucketName ダウンロード対象のファイルが入ったバケット名
//      fileName ダウンロード対象のファイル
//
// 戻り値： エラー情報
func Download(bucketName string, fileName string) error {
	//設定ファイルの情報を与えてS3のインスタンスを作成する
	client := getS3Instance()

	params := &s3.ListObjectsInput{Bucket: &bucketName, Prefix: &fileName}
	resp, connectErr := client.ListObjects(params)
	if connectErr != nil {
		return connectErr
	}

	manager := s3manager.NewDownloader(nil)
	downloadManager := downloader{bucket: bucketName, file: fileName, dir: config.Download.DownloadDir, Downloader: manager}
	key, err := downloadManager.searchFile(resp)
	if err != nil {
		return err
	}

	if err := downloadManager.downloadToFile(key); err != nil {
		return err
	}

	return nil
}

func (downloadManager *downloader) searchFile(resp *s3.ListObjectsOutput) (string, error) {
	var foundKey string

	for _, content := range resp.Contents {
		if (*content.Key == downloadManager.file) && (!strings.Contains(*content.Key, "/")) {
			fmt.Println(*content.Key)
			foundKey = *content.Key
		}
	}

	if foundKey == "" {
		return "", fmt.Errorf("Specified file not found.")
	}

	return foundKey, nil
}

func (downloadManager *downloader) downloadToFile(key string) error {
	file := filepath.Join(downloadManager.dir, key)

	fs, err := os.Create(file)
	if err != nil {
		return err
	}
	defer fs.Close()

	fmt.Printf("Downloading s3://%s/%s to %s...\n", downloadManager.bucket, key, file)
	params := &s3.GetObjectInput{Bucket: &downloadManager.bucket, Key: &key}
	if totalByte, err := downloadManager.Download(fs, params); err != nil {
		fmt.Println(totalByte)
		return err
	}

	return nil
}

//S3のインスタンスを取得する
func getS3Instance() *s3.S3 {
	defaults.DefaultConfig.Credentials = credentials.NewStaticCredentials(config.Aws.AccessKeyId, config.Aws.SecletAccessKey, "")
	defaults.DefaultConfig.Region = &config.Aws.Region

	return s3.New(nil)
}
