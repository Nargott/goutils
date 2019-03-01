package goutils

import (
	"encoding/base64"
	"github.com/satori/go.uuid"
)

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
