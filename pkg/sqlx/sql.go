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

// ToNullInt64 将int64转为sql.NullInt64
func ToNullInt64(i int64) sql.NullInt64 {
	return sql.NullInt64{
		Int64: int64(i),
		Valid: true,
	}
}

// ToNullTime 将 time.Time 转换为 sql.NullTime
func ToNullTime(t time.Time) sql.NullTime {
	return sql.NullTime{
		Time:  t,
		Valid: true,
	}
}
