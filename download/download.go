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

	"github.com/unirita/s3dladapter/config"
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
//      key        ダウンロード対象のキー名
//
// 戻り値： エラー情報
func Download(bucketName string, key string) error {
	//設定ファイルの情報を与えてS3のインスタンスを作成する
	client := getS3Instance()

	params := &s3.ListObjectsInput{Bucket: &bucketName, Prefix: &key}
	resp, connectErr := client.ListObjects(params)
	if connectErr != nil {
		return connectErr
	}

	manager := s3manager.NewDownloader(nil)
	d := downloader{bucket: bucketName, file: key, dir: config.Download.DownloadDir, Downloader: manager}
	if !exists(key, resp) {
		return fmt.Errorf("Not exists download file.")
	}

	if file, err := d.downlowdFile(key); err != nil {
		if err := os.Remove(file); err != nil {
			fmt.Println(err)
		}
		return err
	}
	return nil
}

// バケット内の複数オブジェクトにダウンロードしたいファイルが存在するか判断。
//
// 引数: resp S3のキー名に部分一致したオブジェクト（複数）
//
// 戻り値： ダウンロードしたいファイルが存在するか　[存在する = true]
func exists(downloadFile string, resp *s3.ListObjectsOutput) bool {
	for _, content := range resp.Contents {
		if *content.Key == downloadFile {
			return true
		}
	}
	return false
}

// ファイルをダウンロードする。
//
// 引数: ダウンロードするキー名
//
// 戻り値： エラー情報
func (d *downloader) downlowdFile(key string) (string, error) {
	buffKeys := strings.Split(key, "/")

	fileName := buffKeys[len(buffKeys)-1]
	file := filepath.Join(d.dir, fileName)

	fs, err := os.Create(file)
	if err != nil {
		return file, err
	}
	defer fs.Close()

	fmt.Printf("Downloading s3://%s/%s to %s...\n", d.bucket, fileName, file)
	params := &s3.GetObjectInput{Bucket: &d.bucket, Key: &fileName}
	if _, err := d.Download(fs, params); err != nil {
		return file, err
	}

	fmt.Printf("Complete download.")
	return file, nil
}

//S3のインスタンスを取得する
func getS3Instance() *s3.S3 {
	defaults.DefaultConfig.Credentials = credentials.NewStaticCredentials(config.Aws.AccessKeyId, config.Aws.SecletAccessKey, "")
	defaults.DefaultConfig.Region = &config.Aws.Region

	return s3.New(createConf())
}

func createConf() *aws.Config {
	conf := aws.NewConfig()

	if config.Log.LogDebug == config.Log_Flag_OFF {
		conf.WithLogLevel(aws.LogOff)
	} else {

		loglevel := aws.LogDebug

		if config.Log.LogSigning == config.Log_Flag_ON {
			loglevel |= aws.LogDebugWithSigning
		}

		if config.Log.LogHTTPBody == config.Log_Flag_ON {
			loglevel |= aws.LogDebugWithHTTPBody
		}

		if config.Log.LogRequestRetries == config.Log_Flag_ON {
			loglevel |= aws.LogDebugWithRequestRetries
		}

		if config.Log.LogRequestErrors == config.Log_Flag_ON {
			loglevel |= aws.LogDebugWithRequestErrors
		}

		conf.WithLogLevel(loglevel)
	}
	return conf
}
