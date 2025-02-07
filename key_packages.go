package mls

import (
	"github.com/emersion/go-mls"
)

func (n *NostrMLS) createKeyPackageForEvent(pubkey string) (string, error) {
	// TODO
	return "", nil
}

func (n *NostrMLS) parseKeyPackage(keyPackageHex string) (*mls.KeyPackage, error) {
	// TODO
	return nil, nil
}

func (n *NostrMLS) deleteKeyPackageFromStorage(keyPackage *mls.KeyPackage) error {
	// TODO
	return nil
}

func (n *NostrMLS) generateCredentialWithKey(pubkey string) (*mls.Credential, error) {
	// TODO
	return nil, nil
}
