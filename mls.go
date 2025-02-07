package mls

import (
	"fmt"
	"io"

	"github.com/emersion/go-mls"
	"golang.org/x/crypto/cryptobyte"

	"github.com/1l0/nostr-mls/extension"
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

func (n *NostrMLS) Capabilities() *mls.Capabilities {
	return &mls.Capabilities{
		Versions:     []mls.ProtocolVersion{n.protocolVersion},
		CipherSuites: []mls.CipherSuite{n.cipherSuite},
		Extensions:   n.extensionTypes,
	}
}

func (n *NostrMLS) ClearStore() error {
	return n.store.Clear()
}

// TODO: delete this
type ExampleExtensions struct {
	extensions []mls.Extension
}

func (exts *ExampleExtensions) UnmarshalExample(s *cryptobyte.String) error {

	*exts = ExampleExtensions{}

	l, err := mls.UnmarshalExtensionVec(s)
	if err != nil {
		return err
	}
	exts.extensions = l

	return nil
}

func (exts *ExampleExtensions) UnmarshalExample2(s *cryptobyte.String) error {
	*exts = ExampleExtensions{}

	return mls.ReadVector(s, func(s *cryptobyte.String) error {
		var ext mls.Extension
		if !s.ReadUint16((*uint16)(&ext.ExtensionType)) || !mls.ReadOpaqueVec(s, &ext.ExtensionData) {
			return io.ErrUnexpectedEOF
		}
		if ext.ExtensionType != extension.ExtensionTypeNostrGroup {
			return fmt.Errorf("mismatched extension type to NostrGroup: %d", ext.ExtensionType)
		}
		exts.extensions = append(exts.extensions, ext)
		return nil
	})
}

func (exts *ExampleExtensions) MarshalExample(b *cryptobyte.Builder) {
	mls.MarshalExtensionVec(b, exts.extensions)
}
