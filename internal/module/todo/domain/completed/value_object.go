package completed

type Completed struct {
	value bool
}

func New(value bool) (Completed, error) {
	var vo Completed
	vo.value = value
	return vo, nil
}

func True() Completed {
	n, _ := New(true)
	return n
}

func False() Completed {
	n, _ := New(false)
	return n
}

func (vo Completed) String() string {
	if vo.value {
		return "true"
	}
	return "false"
}

func (vo Completed) Bool() bool {
	return vo.value
}

func (vo Completed) EqualsTo(other Completed) bool {
	return vo.Bool() == other.Bool()
}
