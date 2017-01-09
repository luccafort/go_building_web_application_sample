package main

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"testing"

	gomniauthtest "github.com/stretchr/gomniauth/test"
)

func TestAuthAvatar(t *testing.T) {
	var authAvatar AuthAvatar
	testUser := &gomniauthtest.TestUser{}
	testUser.On("AvatarURL").Return("", ErrNoAvatorURL)
	testChatUser := &chatUser{User: testUser}

	url, err := authAvatar.GetAvatarURL(testChatUser)
	if err != ErrNoAvatorURL {
		t.Error("値が存在しない場合、AuthAvatar.GetAvatarURL()はErrNoAvatarを返すべき")
	}

	// 値をセット
	testURL := "http://url-to-avatar/"
	testUser = &gomniauthtest.TestUser{}
	testChatUser.User = testUser
	testUser.On("AvatarURL").Return(testURL, nil)

	url, err = authAvatar.GetAvatarURL(testChatUser)
	if err != nil {
		t.Error("値が存在する場合AuthAvatar.GetAvatarURL()はエラーを返すべきでない", "-", err)
	} else if url != testURL {
		t.Error("AuthAvatar.GetAvatarURL()は正しいURLを返すべき", "-", err)
	}
}

func TestGravatarAvatar(t *testing.T) {
	var gravatarAvatar GravatarAvatar
	var testURL string = "//www.gravatar.com/avatar/"
	var testUserID string = "abc"

	user := &chatUser{uniqueID: "abc"}
	url, err := gravatarAvatar.GetAvatarURL(user)
	if err != nil {
		t.Error("Gravatar.getAvatarURL()はエラーを返すべきではありません")
	} else if url != testURL+testUserID {
		log.Printf("%s : %s", url, testURL+testUserID)
		t.Errorf("Gravatar.GetAvatarURL()が'%s'という誤った値を返しました", url)
	}
}

func TestFileSystemAvatar(t *testing.T) {
	filename := filepath.Join("avatars", "abc.jpg") // テスト用のアバター画像パスを生成
	ioutil.WriteFile(filename, []byte{}, 0777)      // テスト用のファイルを出力
	defer func() { os.Remove(filename) }()          // テスト終了後に削除する

	var fileSystemAvatar FileSystemAvatar
	var testFile = "/avatars/abc.jpg"
	user := &chatUser{uniqueID: "abc"}

	url, err := fileSystemAvatar.GetAvatarURL(user)
	if err != nil {
		t.Error("FileSystemAvatar.GetAvatarURL()はエラーを返すべきでない")
	} else if url != testFile {
		t.Errorf("FileSystemAvatar.GetAvatarURL()が'%s'という誤った値を返しました", url)
	}
}
