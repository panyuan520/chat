package core

import (
	"bytes"
	"encoding/binary"
	"errors"
)

func bti(head []byte) (int, error) {
	var leng int32
	buf := bytes.NewBuffer(head)
	err := binary.Read(buf, binary.BigEndian, &leng)
	if err != nil {
		return 0, errors.New("bti read error")
	}
	return int(leng), nil
}

func itb(body string) ([]byte, error) {
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.BigEndian, uint32(len(body)))
	if err != nil {
		return buf.Bytes(), errors.New("itb write error")
	}
	return buf.Bytes(), nil
}

func encode(body string) ([]byte, error) {
	head, err := itb(body)
	if err != nil {
		return nil, errors.New("encode itb write error")
	}
	head = append(head, []byte(body)...)
	return head, nil
}
