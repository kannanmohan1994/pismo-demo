package utils

import (
	"github.com/leonelquinteros/gorand"
)

func GenerateUUID() (string, error) {
	uuid, err := gorand.UUIDv4()
	if err != nil {
		return "", err
	}

	uuidStr, err := gorand.MarshalUUID(uuid)
	if err != nil {
		return "", err
	}
	return uuidStr, nil
}

func VerifyUUID(uuid string) bool {
	_, err := gorand.UnmarshalUUID(uuid)
	return err == nil
}
