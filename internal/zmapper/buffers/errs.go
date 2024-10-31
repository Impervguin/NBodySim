package buffers

type ZmapperError struct {
	str string
}

var ErrOutOfBounds ZmapperError = ZmapperError{str: "point is out of bounds"}

func (zme ZmapperError) Error() string {
	return zme.str
}
