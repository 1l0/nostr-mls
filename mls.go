package mls

import (
	"github.com/emersion/go-mls"

	"github.com/1l0/nostr-mls/extension"
)

const (
	KindMLSKeyPackage          = 443
	KindMLSWelcome             = 444
	KindMLSGroupMessage        = 445
	KindMLSKeyPackageRelayList = 10051
)

type NostrMLS struct {
	protocolVersion mls.ProtocolVersion
	cipherSuite     mls.CipherSuite
	extensionTypes  []mls.ExtensionType
	store           Store
}

// TODO: wtf is GREASE?
// var GREASE = []uint16{
// 	0x0A0A, 0x1A1A, 0x2A2A, 0x3A3A, 0x4A4A, 0x5A5A, 0x6A6A, 0x7A7A, 0x8A8A, 0x9A9A, 0xAAAA,
// 	0xBABA, 0xCACA, 0xDADA, 0xEAEA,
// }

func NewNostrMLS(store Store) *NostrMLS {
	return &NostrMLS{
		protocolVersion: mls.ProtocolVersionMLS10,
		cipherSuite:     mls.CipherSuiteMLS_128_DHKEMX25519_AES128GCM_SHA256_Ed25519,
		extensionTypes: []mls.ExtensionType{
			mls.ExtensionTypeRequiredCapabilities,
			mls.ExtensionTypeRatchetTree,
			extension.ExtensionTypeLastResort,
			extension.ExtensionTypeNostrGroup,
		},
		store: store,
	}
}

func (n *NostrMLS) ClearStore() error {
	return n.store.Clear()
}

func (n *NostrMLS) capabilities() *mls.Capabilities {
	return &mls.Capabilities{
		Versions:     []mls.ProtocolVersion{n.protocolVersion},
		CipherSuites: []mls.CipherSuite{n.cipherSuite},
		Extensions:   n.extensionTypes,
	}
}

func (n *NostrMLS) GenerateCredentialWithKey(pubkey string) (*mls.Credential, error) {
	// TODO
	return nil, nil
}
