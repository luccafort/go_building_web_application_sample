package backup

import (
	"path/filepath"
	"time"
)

// Monitor モニタリング対象の変更検知とバックアップを実行管理
type Monitor struct {
	Paths       map[string]string // ファイルパス(絶対パス)
	Archiver    Archiver          // アーカイバ
	Destination string            // アーカイバの保存先
}

// Now 対象をチェック、変更がある個数を返す
func (m *Monitor) Now() (int, error) {
	var counter int
	for path, lastHash := range m.Paths {
		newHash, err := DirHash(path)
		if err != nil {
			return 0, err
		}
		if newHash != lastHash {
			err := m.act(path)
			if err != nil {
				return counter, err
			}
			m.Paths[path] = newHash
			counter++
		}
	}
	return counter, nil
}

// act バックアップを開始
func (m *Monitor) act(path string) error {
	dirname := filepath.Base(path)
	filename := m.Archiver.DestFmt()(time.Now().UnixNano())
	return m.Archiver.Archive(path, filepath.Join(m.Destination, dirname, filename))
}
