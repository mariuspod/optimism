package dadb

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

const DaDbVersion = 0xdb

// Returns the calldata for a given daTxId from the DA
func DiscoverCallData(daTxId int64) ([]byte, error) {
	return Db.ReadCallData(daTxId)
}

// Publishes the given calldata to the DA and returns the daTxId
func PublishCallData(data []byte) (int64, error) {
	return Db.InsertCallData(data)
}

// encodes the daTxId into bytes
func Encode(daTxId int64) ([]byte, error) {
	buffer := new(bytes.Buffer)
	err := binary.Write(buffer, binary.BigEndian, daTxId)
	if err != nil {
		return nil, fmt.Errorf("encoding of daTxId failed: %w", err)
	}
	return buffer.Bytes(), nil
}

// decode the bytes into the daTxId
func Decode(data []byte) (int64, error) {
	var daTxId int64
	buffer := bytes.NewBuffer(data)
	err := binary.Read(buffer, binary.BigEndian, &daTxId)
	if err != nil {
		return 0, fmt.Errorf("decoding of daTxId failed: %w", err)
	}
	return daTxId, nil
}
