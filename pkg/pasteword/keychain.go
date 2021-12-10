package pasteword

import (
	"errors"

	"github.com/deanishe/awgo/keychain"
)

const (
	keychainService = "cc.crow.pasteword"
)

var (
	ErrSecretNotFound = errors.New("secret not found")

	ErrKeychainGet = errors.New("failed to read from Keychain")
	ErrKeychainSet = errors.New("failed to write to Keychain")
	ErrKeychainDel = errors.New("failed to delete from Keychain")
)

func getSecret(name string) (string, error) {
	chain := keychain.New(keychainService)
	res, err := chain.Get(name)
	if err == keychain.ErrNotFound {
		return "", ErrSecretNotFound
	} else if err != nil {
		return "", ErrKeychainGet
	}
	return res, err
}

func setSecret(name, value string) error {
	chain := keychain.New(keychainService)
	if err := chain.Set(name, value); err != nil {
		return ErrKeychainSet
	}
	return nil
}

func delSecret(name string) error {
	chain := keychain.New(keychainService)
	if err := chain.Delete(name); err != nil {
		return ErrKeychainDel
	}
	return nil
}
