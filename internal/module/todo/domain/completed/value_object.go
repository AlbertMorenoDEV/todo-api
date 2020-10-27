package completed

type Completed bool

func New(value bool) (Completed, error) {
	return Completed(value), nil
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
	if vo {
		return "true"
	}
	return "false"
}

func (vo Completed) Bool() bool {
	return bool(vo)
}

func (vo Completed) EqualsTo(other Completed) bool {
	return vo.Bool() == other.Bool()
}
