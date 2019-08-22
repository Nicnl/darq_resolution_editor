package ll_nrbf

import (
	"io"
)

type Decoder struct {
	r   io.Reader
	buf []byte
}

type Header struct {
	RecordTypeEnumeration uint8
	ObjectId              int32
	HeaderId              int32
	MajorVersion          int32
	MinorVersion          int32
}

func NewDecoder(r io.Reader) *Decoder {
	return &Decoder{r: r}
}

func (d *Decoder) nextByte() (b byte, err error) {
	var p [1]byte
	_, err = d.r.Read(p[:])
	if err == nil {
		b = p[0]
	}
	return
}
