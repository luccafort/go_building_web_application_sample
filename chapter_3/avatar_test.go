package main

import (
	"log"
	"testing"
)

func TestAuthAvatar(t *testing.T) {
	var authAvatar AuthAvatar
	client := new(client)
	url, err := authAvatar.GetAvatarURL(client)
	if err != ErrNoAvatorURL {
		t.Error("値が存在しない場合、AuthAvatar.GetAvatarURL()はErrNoAvatarを返すべき")
	}

	// 値をセット
	testURL := "http://url-to-avatar/"
	client.userData = map[string]interface{}{"avatar_url": testURL}
	url, err = authAvatar.GetAvatarURL(client)
	if err != nil {
		t.Error("値が存在する場合AuthAvatar.GetAvatarURL()はエラーを返すべきでない")
	} else if url != testURL {
		t.Error("AuthAvatar.GetAvatarURL()は正しいURLを返すべき")
	}
}

func TestGravatarAvatar(t *testing.T) {
	var gravatarAvatar GravatarAvatar
	//var testEmail string = "MyEmailAddress@example.com"
	var testURL string = "//www.gravatar.com/avatar/"
	var testUserID string = "0bc83cb571cd1c50ba6f3e8a78ef1346"

	client := new(client)
	client.userData = map[string]interface{}{
		//"email": testEmail,
		"userid": testUserID,
	}
	url, err := gravatarAvatar.GetAvatarURL(client)
	if err != nil {
		t.Error("Gravatar.getAvatarURL()はエラーを返すべきではありません")
	} else if url != testURL+testUserID {
		log.Printf("%s : %s", url, testURL+testUserID)
		t.Error("Gravatar.GetAvatarURL()が'%s'という誤った値を返しました", url)
	}
}
