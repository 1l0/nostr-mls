package mls

import (
	"encoding/hex"

	"github.com/emersion/go-mls"
	"golang.org/x/crypto/cryptobyte"
)

func (n *NostrMLS) CreateKeyPackageHex(pubkey string) (string, error) {
	cred, err := n.GenerateCredentialWithKey(pubkey)
	if err != nil {
		return "", err
	}
	pkg := &mls.KeyPackage{
		LeafNode:  mls.LeafNode{},
		Signature: cred.SignaturePublicKey,
	}
	var b cryptobyte.Builder
	pkg.Marshal(&b)
	result, err := b.Bytes()
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(result), nil
}

func (n *NostrMLS) ParseKeyPackage(keyPackageHex string) (*KeyPackage, error) {
	// TODO
	return nil, nil
}

func (n *NostrMLS) DeleteKeyPackageFromStore(keyPackage *KeyPackage) error {
	// TODO
	return nil
}

func (n *NostrMLS) GenerateCredentialWithKey(pubkey string) (*CredentialWithKey, error) {
	priv, pub, err := n.cipherSuite.SignatureScheme().GenerateKeys()
	if err != nil {
		return nil, err
	}
	if err := n.store.Upsert([]byte(pubkey), priv); err != nil {
		return nil, err
	}
	return &CredentialWithKey{
		Credential: &mls.Credential{
			CredentialType: mls.CredentialTypeBasic,
			Identity:       []byte(pubkey),
		},
		SignaturePublicKey: pub,
	}, nil
}

type CredentialWithKey struct {
	Credential         *mls.Credential
	SignaturePublicKey []byte
}
