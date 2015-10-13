// Copyright 2015 unirita Inc.
// Created 2015/10/09 kazami

package download

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
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

//os.Getenv()

func Download(bucket string, file string) {
	//TODO: 設定ファイルの情報を与えてS3のインスタンスを作成する
	cred := credentials.NewStaticCredentials(config.Aws.AccessKeyId, config.Aws.SecletAccessKey, "")
	awsConf := aws.Config{Credentials: cred, Region: &config.Aws.Region}
	client := s3.New(&awsConf)

	params := &s3.ListObjectsInput{Bucket: &bucket, Prefix: &file}

	manager := s3manager.NewDownloader(nil)
	d := downloader{bucket: bucket, file: file, dir: config.Download.DownloadDir, Downloader: manager}

	client.ListObjectsPages(params, d.eachPage)
}

func (d *downloader) eachPage(page *s3.ListObjectsOutput, more bool) bool {
	for _, obj := range page.Contents {
		d.downloadToFile(*obj.Key)
	}

	return true
}

func (d *downloader) downloadToFile(key string) {
	// Create the directories in the path
	file := filepath.Join(d.dir, key)
	if err := os.MkdirAll(filepath.Dir(file), 0775); err != nil {
		panic(err)
	}

	// Setup the local file
	fd, err := os.Create(file)
	if err != nil {
		panic(err)
	}
	defer fd.Close()

	// Download the file using the AWS SDK
	fmt.Printf("Downloading s3://%s/%s to %s...\n", d.bucket, key, file)
	params := &s3.GetObjectInput{Bucket: &d.bucket, Key: &key}
	d.Download(fd, params)

	return
}
