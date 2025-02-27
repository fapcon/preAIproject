package types

import (
	"database/sql"
	"github.com/volatiletech/null/v8"
	"time"

	"github.com/google/uuid"
)

func NewNullString(s string) NullString {
	if len(s) == 0 {
		return NullString{}
	}
	return NullString{
		sql.NullString{
			String: s,
			Valid:  true,
		},
	}
}

func NewNullInt64(n int64) NullInt64 {
	if n == 0 {
		return NullInt64{}
	}
	return NullInt64{
		sql.NullInt64{
			Int64: n,
			Valid: true,
		},
	}
}

func NewNullUint64(n uint64) NullUint64 {
	tn := NullUint64{}
	tn.Valid = true
	tn.Uint64.Uint64 = n

	return tn
}

func NewNullFloat64(n float64) NullFloat64 {
	if n == 0 {
		return NullFloat64{}
	}
	return NullFloat64{
		sql.NullFloat64{
			Float64: n,
			Valid:   true,
		},
	}
}

func NewNullBool(b bool) NullBool {
	if !b {
		return NullBool{}
	}
	return NullBool{
		sql.NullBool{
			Bool:  b,
			Valid: true,
		},
	}
}

func NewNullTime(t time.Time) NullTime {
	return NullTime{
		null.Time{
			Time:  t,
			Valid: true,
		},
	}
}

func NewNullUUID(s ...string) NullUUID {
	var uuidRaw uuid.UUID
	var err error
	uuidRaw, err = uuid.NewUUID()
	if len(s) > 0 {
		if len(s[0]) > 0 {
			uuidRaw, err = uuid.Parse(s[0])
			if err != nil {
				return NullUUID{}
			}
		}
	}
	if err != nil {
		return NullUUID{}
	}

	var nullUUID NullUUID

	nullUUID.Binary, err = uuidRaw.MarshalBinary()
	if err != nil {
		return NullUUID{}
	}

	nullUUID.String = uuidRaw.String()

	nullUUID.Valid = true

	return nullUUID
}
