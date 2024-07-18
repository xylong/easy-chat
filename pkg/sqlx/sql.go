package sqlx

import (
	"database/sql"
	"time"
)

// ToNullString 将Go的string转换为sql.NullString
func ToNullString(s string) sql.NullString {
	if s == "" {
		return sql.NullString{Valid: false}
	}

	return sql.NullString{Valid: true, String: s}
}

// ToNullTime 将 time.Time 转换为 sql.NullTime
func ToNullTime(t time.Time) sql.NullTime {
	return sql.NullTime{
		Time:  t,
		Valid: true,
	}
}
