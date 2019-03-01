package goutils

import (
	"encoding/base64"
	"fmt"
	"github.com/btcsuite/btcutil/base58"
	"github.com/satori/go.uuid"
)

const GoutilsBase58Version = 42

// StringUUIDToBase64 returns packed and Base64-encoded UUID as string
func StringUUIDToBase64(su string) (string, error) {
	u, err := uuid.FromString(su)
	if err != nil {
		return "", err
	}

	bu, _ := u.MarshalBinary() // MarshalBinary does not returns error at all, so ignoring it
	return base64.StdEncoding.EncodeToString(bu), nil
}

// Base64ToUuid returns unpacked base64-decoded and converted from binary UUID
func Base64ToUuidString(bs string) (string, error) {
	b, err := base64.StdEncoding.DecodeString(bs)
	if err != nil {
		return "", err
	}

	return uuid.FromBytesOrNil(b).String(), nil
}

// UUIDToBase64 returns packed and Base64-encoded UUID
func UUIDToBase64(u uuid.UUID) string {
	bu, _ := u.MarshalBinary() // MarshalBinary does not returns error at all, so ignoring it
	return base64.StdEncoding.EncodeToString(bu)
}

// Base64ToUuid returns unpacked base64-decoded and converted from binary UUID
func Base64ToUuid(bs string) (uuid.UUID, error) {
	b, err := base64.StdEncoding.DecodeString(bs)
	if err != nil {
		return uuid.Nil, err
	}

	return uuid.FromBytesOrNil(b), nil
}

// UUIDToBase58 returns packed and Base58-encoded UUID
func UUIDToBase58(u uuid.UUID) string {
	bu, _ := u.MarshalBinary() // MarshalBinary does not returns error at all, so ignoring it
	return base58.Encode(bu)
}

// Base58ToUuid returns unpacked base58-decoded and converted from binary UUID
func Base58ToUuid(bs string) uuid.UUID {
	return uuid.FromBytesOrNil(base58.Decode(bs))
}

// UUIDToBase58Check returns packed and Base58-encoded with version UUID
func UUIDToBase58Check(u uuid.UUID) string {
	bu, _ := u.MarshalBinary() // MarshalBinary does not returns error at all, so ignoring it
	return base58.CheckEncode(bu, GoutilsBase58Version)
}

// Base58ToUuid returns unpacked base58-decoded with version and converted from binary UUID
func Base58CheckToUuid(bs string) (uuid.UUID, error) {
	b, v, err := base58.CheckDecode(bs)
	if err != nil {
		return uuid.Nil, err
	}
	if v != GoutilsBase58Version {
		return uuid.Nil, fmt.Errorf("bad base58 version %d, when expects %d", v, GoutilsBase58Version)
	}

	return uuid.FromBytesOrNil(b), nil
}
