package main

import "testing"

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
