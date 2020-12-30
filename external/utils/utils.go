package utils

import (
	"bytes"
	"encoding/binary"

	"github.com/syndtr/goleveldb/leveldb"
)

func StringOut(bye []byte) string {
	return string(bye)
}

func StringIn(strings string) []byte {
	return []byte(strings)
}

func IntIn(n int) []byte {
	data := int64(n)
	bytebuf := bytes.NewBuffer([]byte{})
	binary.Write(bytebuf, binary.BigEndian, data)
	return bytebuf.Bytes()
}

func IntOut(bye []byte) int {
	bytebuff := bytes.NewBuffer(bye)
	var data int64
	binary.Read(bytebuff, binary.BigEndian, &data)
	return int(data)
}

func UintIn(n uint) []byte {
	data := uint64(n)
	bytebuf := bytes.NewBuffer([]byte{})
	binary.Write(bytebuf, binary.BigEndian, data)
	return bytebuf.Bytes()
}

func UintOut(bye []byte) uint {
	bytebuff := bytes.NewBuffer(bye)
	var data uint64
	binary.Read(bytebuff, binary.BigEndian, &data)
	return uint(data)
}

func ErrorToErrCode(err error) int32 {
	var errorCode int32
	errorCode = -1
	switch err {
	case leveldb.ErrNotFound:
		errorCode = 0
	case leveldb.ErrReadOnly:
		errorCode = 1
	case leveldb.ErrSnapshotReleased:
		errorCode = 2
	case leveldb.ErrIterReleased:
		errorCode = 3
	case leveldb.ErrClosed:
		errorCode = 4
	}
	return errorCode
}

func ErrCodeToError(errorCode int32) error {
	switch errorCode {
	case 0:
		return leveldb.ErrNotFound
	case 1:
		return leveldb.ErrReadOnly
	case 2:
		return leveldb.ErrSnapshotReleased
	case 3:
		return leveldb.ErrIterReleased
	case 4:
		return leveldb.ErrClosed
	}
	return nil
}
