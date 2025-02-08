package mls

import (
	"encoding/hex"
	"fmt"

	"github.com/emersion/go-mls"
	"golang.org/x/crypto/cryptobyte"
)

func (n *NostrMLS) CreateKeyPackageHex(pubkey string) (string, error) {
	cred, priv, err := n.updateCredentialWithKey(pubkey)
	if err != nil {
		return "", err
	}
	pkg := &mls.KeyPackage{
		Version:     n.protocolVersion,
		CipherSuite: n.cipherSuite,
		LeafNode: mls.LeafNode{
			SignatureKey: cred.SignaturePublicKey,
			Credential:   *cred.Credential,
			Capabilities: *n.capabilities(),
		},
	}
	var b cryptobyte.Builder
	pkg.MarshalTBS(&b)
	tbs, err := b.Bytes()
	if err != nil {
		return "", err
	}
	result, err := n.cipherSuite.SignWithLabel(priv, []byte{}, tbs)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(result), nil
}

func (n *NostrMLS) ParseKeyPackage(keyPackageHex string) (*KeyPackage, error) {
	pkg := &mls.KeyPackage{}
	b := cryptobyte.String(keyPackageHex)
	if err := pkg.Unmarshal(&b); err != nil {
		return nil, err
	}
	if !pkg.VerifySignature() {
		return nil, fmt.Errorf("invalid signature")
	}
	p := KeyPackage(*pkg)
	return &p, nil
}

func (n *NostrMLS) updateCredentialWithKey(pubkey string) (*CredentialWithKey, []byte, error) {
	priv, pub, err := n.cipherSuite.SignatureScheme().GenerateKeys()
	if err != nil {
		return nil, nil, err
	}
	if err := n.store.Upsert([]byte(pubkey), priv); err != nil {
		return nil, nil, err
	}
	return &CredentialWithKey{
		Credential: &mls.Credential{
			CredentialType: mls.CredentialTypeBasic,
			Identity:       []byte(pubkey),
		},
		SignaturePublicKey: pub,
	}, priv, nil
}

type CredentialWithKey struct {
	Credential         *mls.Credential
	SignaturePublicKey []byte
}
