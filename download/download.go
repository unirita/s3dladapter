// Copyright 2015 unirita Inc.
// Created 2015/10/09 kazami

package download

import (
	"fmt"
	"os"
	"path"
	"path/filepath"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"

	"github.com/unirita/s3dladapter/config"
)

// S3からファイルをダウンロードする
//
// 引数: bucketName ダウンロード対象のファイルが入ったバケット名
//      key        ダウンロード対象のキー名
//
// 戻り値： エラー情報
func Do(bucket string, key string) error {
	//設定ファイルの情報を与えてS3のインスタンスを作成する
	client := s3.New(createConf())

	if err := tryToFindFile(client, bucket, key); err != nil {
		return err
	}

	localPath, err := downlowdFile(bucket, key, config.Download.DownloadDir)
	if err != nil {
		if err := os.Remove(localPath); err != nil {
			fmt.Println(err)
		}
		return err
	}

	fmt.Println(localPath)
	return nil
}

func tryToFindFile(client *s3.S3, bucket, key string) error {
	params := &s3.ListObjectsInput{Bucket: &bucket, Prefix: &key}
	objectList, err := client.ListObjects(params)
	if err != nil {
		return err
	}
	if !exists(objectList, key) {
		return fmt.Errorf("Not exists download file.")
	}

	return nil
}

// バケット内の複数オブジェクトにダウンロードしたいファイルが存在するか判断。
//
// 引数: resp S3のキー名に部分一致したオブジェクト（複数）
//
// 戻り値： ダウンロードしたいファイルが存在するか　[存在する = true]
func exists(objectList *s3.ListObjectsOutput, downloadFile string) bool {
	for _, content := range objectList.Contents {
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
func downlowdFile(bucket, key, localDir string) (string, error) {
	fileName := path.Base(key)
	localPath := filepath.Join(localDir, fileName)

	file, err := os.Create(localPath)
	if err != nil {
		return localPath, err
	}
	defer file.Close()

	fmt.Printf("Downloading s3://%s/%s to %s...\n", bucket, key, localPath)
	d := s3manager.NewDownloader(nil)
	params := &s3.GetObjectInput{Bucket: &bucket, Key: &key}
	if _, err := d.Download(file, params); err != nil {
		return localPath, err
	}

	fmt.Println("Download complete .")
	return localPath, nil
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
