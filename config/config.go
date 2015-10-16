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

const (
	Log_Flag_ON  = "on"
	Log_Flag_OFF = "off"
)

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
	LogDebug          string `toml:"log_debug"`
	LogSigning        string `toml:"log_signing"`
	LogHTTPBody       string `toml:"log_loghttp"`
	LogRequestRetries string `toml:"log_request_retries"`
	LogRequestErrors  string `toml:"log_request_errors"`
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

	if !existsDir(Download.DownloadDir) {
		return fmt.Errorf("Download.download_dir(%s) does not exist.", Download.DownloadDir)
	}

	if Log.LogDebug != "on" || Log.LogDebug != "off" {
		return fmt.Errorf("Log.log_debug (%s) does not have the format of on or off.", Log.LogDebug)
	}

	if Log.LogSigning != "on" || Log.LogSigning != "off" {
		return fmt.Errorf("Log.log_signing (%s) does not have the format of on or off.", Log.LogSigning)
	}

	if Log.LogHTTPBody != "on" || Log.LogHTTPBody != "off" {
		return fmt.Errorf("Log.log_httpbody (%s) does not have the format of on or off.", Log.LogHTTPBody)
	}

	if Log.LogRequestRetries != "on" || Log.LogRequestRetries != "off" {
		return fmt.Errorf("Log.log_request_retries (%s) does not have the format of on or off.", Log.LogRequestRetries)
	}

	if Log.LogRequestErrors != "on" || Log.LogRequestErrors != "off" {
		return fmt.Errorf("Log.log_request_errors (%s) does not have the format of on or off.", Log.LogRequestErrors)
	}

	return nil
}

//パスの存在チェック
func existsDir(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}
