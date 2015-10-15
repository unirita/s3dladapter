package download

import (
	"testing"

	"s3dladapter/config"
)

func TestGetS3Instance_認証情報に該当するアカウントが存在する場合はインスタンスを返す(t *testing.T) {
	testConfig := "existS3.ini"
	if err := config.Load(testConfig); err != nil {
		t.Errorf("テストファイルの読み込みに失敗")
	}

	if _, err := GetS3Instance(); err != nil {
		t.Errorf("認証が通っていない")
	}

}

func TestGetS3Instance_認証情報に該当するアカウントが存在しない場合はエラー(t *testing.T) {
	testConfig := "noexistS3.ini"
	if err := config.Load(testConfig); err != nil {
		t.Errorf("テストファイルの読み込みに失敗")
	}

	if _, err := GetS3Instance(); err == nil {
		t.Errorf("期待していない認証が通っている")
	}
}
