package due

import "time"

type Due struct {
	value time.Time
}

func New(value time.Time) (Due, error) {
	var vo Due
	vo.value = value
	return vo, nil
}

func (vo Due) String() string {
	return vo.value.Format("yyyy-MM-dd HH:mm:ss")
}

func (vo Due) Time() time.Time {
	return vo.value
}

func (vo Due) IsPast() bool {
	return vo.value.Before(time.Now())
}

func (vo Due) EqualsTo(other Due) bool {
	return vo.String() == other.String()
}
