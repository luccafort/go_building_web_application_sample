package main

import "errors"

// ErrNoAvator はAvatarインスタンスがアバターのURLを返すことができない場合に発生するエラー
var ErrNoAvatorURL = errors.New("chat: アバターのURLを取得できません")

// Avatar はユーザのプロフィール画像を表す型
type Avatar interface {
	/*
	 * GetAvatarURLは指定されたクライアントのアバターのURLを返します
	 * 問題が発生した場合にはエラーを返す。
	 * 特にURLを取得できなかった場合にErrNoAvatarを返す
	 */
	GetAvatarURL(c *client) (string, error)
}

type AuthAvatar struct{}

var UseAuthAvatar AuthAvatar

// オブジェクトの参照がないため"_"で参照しないことを明示("_"は省略も可)
func (_ AuthAvatar) GetAvatarURL(c *client) (string, error) {
	if url, ok := c.userData["avatar_url"]; ok {
		if retURL, ok := url.(string); ok {
			return retURL, nil
		}
	}
	return "", ErrNoAvatorURL

}

// Gravatar構造体を設定
type GravatarAvatar struct{}

var UseGravatar GravatarAvatar

func (_ GravatarAvatar) GetAvatarURL(c *client) (string, error) {
	var key string = "userid"
	if _userid, ok := c.userData[key]; ok {
		if userid, ok := _userid.(string); ok {
			return "//www.gravatar.com/avatar/" + userid, nil
		}
	}
	return "", ErrNoAvatorURL
}

type FileSystemAvatar struct{}

var UseFileSystemAvatar FileSystemAvatar

func (_ FileSystemAvatar) GetAvatarURL(c *client) (string, error) {
	if _userid, ok := c.userData["userid"]; ok {
		if userid, ok := _userid.(string); ok {
			return "/avatars/" + userid + ".jpg", nil
		}
	}
	return "", ErrNoAvatorURL
}
