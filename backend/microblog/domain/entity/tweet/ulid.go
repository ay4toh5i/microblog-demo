package tweet

import (
	"database/sql/driver"
	"errors"
	"math/rand"
	"time"

	"github.com/oklog/ulid/v2"
)

type ULID ulid.ULID

func NewULID() ULID {
	t := time.Now()
	entropy := ulid.Monotonic(rand.New(rand.NewSource(t.UnixNano())), 0)
	return ULID(ulid.MustNew(ulid.Timestamp(t), entropy))
}

func (id ULID) Value() (driver.Value, error) {
	return ulid.ULID(id).String(), nil
}

func (id *ULID) Scan(value interface{}) error {
	v, ok := value.([]byte)
	if !ok {
		return errors.New("invalid data")
	}

	return (*ulid.ULID)(id).Scan(string(v))
}
