// Copyright 2015 unirita Inc.
// Created 2015/10/09 kazami

package config

import (
	"fmt"
	"io"
	"os"

	"github.com/BurntSushi/toml"
)

type config struct {
	Aws      awsTable
	Download downloadTable
	Log      logTable
}

// 設定ファイルのawsテーブル
type awsTable struct {
	AccessKeyId     string `toml:"access_key_id"`
	SecletAccessKey string `toml:"secret_access_key"`
	Region          string `toml:"region"`
}

// 設定ファイルのdownloadテーブル
type downloadTable struct {
	DownloadDir string `toml:"download_dir"`
}

// 設定ファイルのlogテーブル
type logTable struct {
	LogLevel int `toml:"loglevel"`
}

var Aws = new(awsTable)
var Download = new(downloadTable)
var Log = new(logTable)

// 設定ファイルをロードする。
//
// 引数: filePath ロードする設定ファイルのパス
//
// 戻り値： エラー情報
func Load(filePath string) error {
	f, err := os.Open(filePath)
	if err != nil {
		return err
	}

	return loadReader(f)
}

func loadReader(reader io.Reader) error {
	c := new(config)

	if _, err := toml.DecodeReader(reader, c); err != nil {
		return err
	}

	Aws = &c.Aws
	Download = &c.Download
	Log = &c.Log

	return nil
}

// 設定値のエラー検出を行う。
//
// return : エラー情報
func DetectError() error {
	if Aws.AccessKeyId == "" {
		return fmt.Errorf("Aws.access_key_id is blank.")
	}

	if Aws.SecletAccessKey == "" {
		return fmt.Errorf("Aws.seclet_access_key value is not set.")
	}

	if Aws.Region == "" {
		return fmt.Errorf("Aws.region value is not set.")
	}

	if !pathExists(Download.DownloadDir) {
		return fmt.Errorf("Download.download_dir(%s) does not exist.", Download.DownloadDir)
	}

	if Log.LogLevel < 0 || Log.LogLevel > 5 {
		return fmt.Errorf("Log.loglevel (%d) must to be range from 0 to 5", Log.LogLevel)
	}

	return nil
}

//パスの存在チェック
func pathExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}
