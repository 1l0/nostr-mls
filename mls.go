package mls

import (
	"fmt"
	"io"

	"github.com/emersion/go-mls"
	"golang.org/x/crypto/cryptobyte"

	"github.com/1l0/nostr-mls/extension"
)

type NostrMLS struct {
	CipherSuite    mls.CipherSuite
	ExtensionTypes []mls.ExtensionType
}

func NewNostrMLS() *NostrMLS {
	return &NostrMLS{
		CipherSuite: mls.CipherSuiteMLS_128_DHKEMX25519_AES128GCM_SHA256_Ed25519,
		ExtensionTypes: []mls.ExtensionType{
			mls.ExtensionTypeRatchetTree,
			mls.ExtensionTypeRequiredCapabilities,
			extension.ExtensionTypeNostrGroup,
		},
	}
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
