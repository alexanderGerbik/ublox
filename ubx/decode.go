package ubx

import (
	"bytes"
	"encoding/binary"
	"errors"
	"io"
)

var (
	errInvalidFrame  = errors.New("invalid UBX frame")
	errInvalidChkSum = errors.New("invalid UBX checksum")
)

func Decode(frame []byte) (msg Message, err error) {

	buf := bytes.NewReader(frame)

	var header struct {
		Preamble uint16
		ClassID  uint16
		Length   uint16
	}

	if err := binary.Read(buf, binary.LittleEndian, &header); err != nil {
		return nil, err
	}

	if header.Preamble != 0x62B5 {
		return nil, errInvalidFrame
	}

	if buf.Len()+2 < int(header.Length) {
		return nil, io.ErrShortBuffer
	}

	var a, b byte
	for _, v := range frame[2 : header.Length+6] {
		a += byte(v)
		b += a
	}

	if frame[header.Length+6] != a || frame[header.Length+7] != b {
		return nil, errInvalidChkSum
	}

	switch header.ClassID {
	case (AckAck{}).classID():
		msg = &AckAck{}
	case (AckNak{}).classID():
		msg = &AckNak{}
	case (CfgCfg1{}).classID():
		msg = &CfgCfg1{}
	case (CfgCfg2{}).classID():
		msg = &CfgCfg2{}
	case (CfgHnr{}).classID():
		msg = &CfgHnr{}
	case (CfgMsg1{}).classID():
		msg = &CfgMsg1{}
	case (CfgMsg2{}).classID():
		msg = &CfgMsg2{}
	case (CfgMsg3{}).classID():
		msg = &CfgMsg3{}
	case (CfgRate{}).classID():
		msg = &CfgRate{}
	case (NavPvt{}).classID():
		msg = &NavPvt{}
	default:
	}
	if msg != nil {
		err = binary.Read(buf, binary.LittleEndian, msg)
	} else {
		msg = &RawMessage{ClassID: header.ClassID, Data: append([]byte(nil), frame[6:len(frame)-2]...)}
	}

	return msg, err

}
