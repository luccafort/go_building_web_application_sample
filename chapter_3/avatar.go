package main

import (
	"crypto/md5"
	"errors"
	"fmt"
	"io"
	"strings"
)

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
	if email, ok := c.userData["email"]; ok {
		if emailStr, ok := email.(string); ok {
			m := md5.New()
			io.WriteString(m, strings.ToLower(emailStr))
			return fmt.Sprintf("//www.gravatar.com/avatar/%x", m.Sum(nil)), nil
		}
	}
	return "", ErrNoAvatorURL
}
