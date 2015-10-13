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

func Download(bucketName string, fileName string) error {
	//設定ファイルの情報を与えてS3のインスタンスを作成する
	cred := credentials.NewStaticCredentials(config.Aws.AccessKeyId, config.Aws.SecletAccessKey, "")
	awsConf := aws.Config{Credentials: cred, Region: &config.Aws.Region}
	client := s3.New(&awsConf)

	params := &s3.ListObjectsInput{Bucket: &bucketName, Prefix: &fileName}

	manager := s3manager.NewDownloader(nil)
	d := downloader{bucket: bucketName, file: fileName, dir: config.Download.DownloadDir, Downloader: manager}

	resp, _ := client.ListObjects(params)
	if len(resp.Contents) == 0 {
		return fmt.Errorf("Not Exist download file.")
	}

	if err := d.eachPage(resp); err != nil {
		return err
	}

	return nil
}

func (d *downloader) eachPage(resp *s3.ListObjectsOutput) error {
	for _, obj := range resp.Contents {
		d.downloadToFile(*obj.Key)
	}

	return nil
}

func (d *downloader) downloadToFile(key string) error {
	// Create the directories in the path
	file := filepath.Join(d.dir, key)

	// Setup the local file
	fd, err := os.Create(file)
	if err != nil {
		panic(err)
	}
	defer fd.Close()

	// Download the file using the AWS SDK
	fmt.Printf("Downloading s3://%s/%s to %s...\n", d.bucket, key, file)
	params := &s3.GetObjectInput{Bucket: &d.bucket, Key: &key}
	if _, err := d.Download(fd, params); err != nil {
		return fmt.Errorf("Failed download file %s", file)
	}

	return nil
}
