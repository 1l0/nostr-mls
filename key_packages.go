package mls

import (
	"github.com/emersion/go-mls"
)

func (n *NostrMLS) CreateKeyPackageForEvent(pubkey string) (string, error) {
	// TODO
	return "", nil
}

func (n *NostrMLS) ParseKeyPackage(keyPackageHex string) (*mls.KeyPackage, error) {
	// TODO
	return nil, nil
}

func (n *NostrMLS) DeleteKeyPackageFromStorage(keyPackage *mls.KeyPackage) error {
	// TODO
	return nil
}
