package entity

import "time"

type TODO struct {
	ID          string    // ID
	Title       string    // タイトル
	CreatedTime time.Time // 作成日時
}
