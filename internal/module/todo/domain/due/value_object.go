package due

import (
	"time"
)

type Due struct {
	value time.Time
}

func New(v time.Time) (*Due, error) {
	var vo Due
	vo.value = v
	return &vo, nil
}

func FromMilliseconds(v int64) (*Due, error) {
	return New(time.Unix(v, 0))
}

func (vo Due) Time() time.Time {
	return vo.value
}

func (vo Due) IsPast() bool {
	return vo.value.Before(time.Now())
}

func (vo Due) EqualsTo(other Due) bool {
	return vo.Time().Unix() == other.Time().Unix()
}
