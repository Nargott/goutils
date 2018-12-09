package goutils

import (
	"database/sql"
	"github.com/go-sql-driver/mysql"
	"github.com/wedancedalot/decimal"
	"time"
)

func ToNullDecimal(value decimal.Decimal) decimal.NullDecimal {
	return decimal.NullDecimal{
		Decimal: value,
		Valid:   true,
	}
}

func ToNullTime(value time.Time) mysql.NullTime {
	return mysql.NullTime{
		Time:  value,
		Valid: !value.IsZero(),
	}
}

func ToNullString(s string) sql.NullString {
	return sql.NullString{
		String: s,
		Valid:  s != "",
	}
}

func ToNullBool(value bool) sql.NullBool {
	return sql.NullBool{
		Bool:  value,
		Valid: true,
	}
}

func ToNullInt64(value int64) sql.NullInt64 {
	return sql.NullInt64{
		Int64: value,
		Valid: true,
	}
}

func ToNullFloat64(value float64) sql.NullFloat64 {
	return sql.NullFloat64{
		Float64: value,
		Valid:   true,
	}
}
