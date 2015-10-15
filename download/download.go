// Copyright 2015 unirita Inc.
// Created 2015/10/09 kazami

package download

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
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
	d := downloader{bucket: bucketName, file: fileName, dir: config.Download.DownloadDir, Downloader: manager}
	if !existsDownloadFile(fileName, resp) {
		return fmt.Errorf("Not exists download file.")
	}

	if err := d.downlowdFile(fileName); err != nil {
		return err
	}

	return nil
}

// バケット内の複数オブジェクトにダウンロードしたいファイルが存在するか判断。
//
// 引数: resp S3のキー名に部分一致したオブジェクト（複数）
//
// 戻り値： ダウンロードしたいファイルが存在するか　[存在する = true]
func existsDownloadFile(downloadFile string, resp *s3.ListObjectsOutput) bool {
	for _, content := range resp.Contents {
		if *content.Key == downloadFile {
			return true
		}
	}
	return false
}

// ファイルをダウンロードする。
//
// 引数: ダウンロードするファイル名
//
// 戻り値： エラー情報
func (d *downloader) downlowdFile(key string) error {
	buffKeys := strings.Split(key, "/")

	fileName := buffKeys[len(buffKeys)-1]
	file := filepath.Join(d.dir, fileName)

	fs, err := os.Create(file)
	if err != nil {
		return err
	}
	defer fs.Close()

	fmt.Printf("Downloading s3://%s/%s to %s...\n", d.bucket, fileName, file)
	params := &s3.GetObjectInput{Bucket: &d.bucket, Key: &fileName}
	if _, err := d.Download(fs, params); err != nil {
		if err := os.Remove(file); err != nil {
			return err
		}
		return err
	}

	fmt.Printf("Complete download.")
	return nil
}

//S3のインスタンスを取得する
func getS3Instance() *s3.S3 {
	defaults.DefaultConfig.Credentials = credentials.NewStaticCredentials(config.Aws.AccessKeyId, config.Aws.SecletAccessKey, "")
	defaults.DefaultConfig.Region = &config.Aws.Region

	conf := aws.NewConfig()

	if config.Log.LogLevel == 0 {
		conf.WithLogLevel(aws.LogOff)
	} else if config.Log.LogLevel == 1 {
		conf.WithLogLevel(aws.LogDebugWithSigning)
	} else if config.Log.LogLevel == 2 {
		conf.WithLogLevel(aws.LogDebugWithHTTPBody)
	} else if config.Log.LogLevel == 3 {
		conf.WithLogLevel(aws.LogDebugWithRequestRetries)
	} else if config.Log.LogLevel == 4 {
		conf.WithLogLevel(aws.LogDebugWithRequestErrors)
	} else if config.Log.LogLevel == 5 {
		conf.WithLogLevel(aws.LogDebug)
	}

	return s3.New(conf)
}
