package main

import "errors"
import "io/ioutil"
import "path/filepath"

// ErrNoAvator はAvatarインスタンスがアバターのURLを返すことができない場合に発生するエラー
var ErrNoAvatorURL = errors.New("chat: アバターのURLを取得できません")

// Avatar はユーザのプロフィール画像を表す型
type Avatar interface {
	/*
	 * GetAvatarURLは指定されたクライアントのアバターのURLを返します
	 * 問題が発生した場合にはエラーを返す。
	 * 特にURLを取得できなかった場合にErrNoAvatarを返す
	 */
	GetAvatarURL(ChatUser) (string, error)
}

type AuthAvatar struct{}

var UseAuthAvatar AuthAvatar

// オブジェクトの参照がないため"_"で参照しないことを明示("_"は省略も可)
func (_ AuthAvatar) GetAvatarURL(u ChatUser) (string, error) {
	url := u.AvatarURL()
	if url != "" {
		return url, nil
	}
	return "", ErrNoAvatorURL

}

// Gravatar構造体を設定
type GravatarAvatar struct{}

var UseGravatar GravatarAvatar

func (_ GravatarAvatar) GetAvatarURL(u ChatUser) (string, error) {
	return "//www.gravatar.com/avatar/" + u.UniqueID(), nil
}

type FileSystemAvatar struct{}

var UseFileSystemAvatar FileSystemAvatar

func (_ FileSystemAvatar) GetAvatarURL(u ChatUser) (string, error) {
	if files, err := ioutil.ReadDir("avatars"); err == nil {
		for _, file := range files {
			if file.IsDir() {
				continue
			}
			// ユーザIDにマッチするか精査
			if match, _ := filepath.Match(u.UniqueID()+"*", file.Name()); match {
				return "/avatars/" + file.Name(), nil
			}
		}
	}
	return "", ErrNoAvatorURL
}
