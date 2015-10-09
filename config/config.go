// Copyright 2015 unirita Inc.
// Created 2015/10/09 kazami

package config

// 設定ファイルのadapterテーブル
type Config struct {
	AccessKeyId      string
	SecletAccessKey  string
	Region           string
	DownloadLocation string
}

var Adapter = new(Config)
