package mls

import (
	"github.com/emersion/go-mls"

	"github.com/1l0/nostr-mls/extension"
)

type NostrMLS struct {
	CipherSuite   mls.CipherSuite
	ExtensionType mls.ExtensionType
}

func NewNostrMLS() *NostrMLS {
	return &NostrMLS{
		CipherSuite:   mls.CipherSuiteMLS_128_DHKEMX25519_AES128GCM_SHA256_Ed25519,
		ExtensionType: extension.ExtensionTypeNostrGroup,
	}
}
