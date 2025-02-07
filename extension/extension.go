package extension

import (
	"io"

	"github.com/emersion/go-mls"
	"golang.org/x/crypto/cryptobyte"
)

const (
	ExtensionTypeNostrGroup mls.ExtensionType = 0xF2EE
)

type NostrGroupData struct {
	nostr_group_id []byte
	name           []byte
	description    []byte
	admin_pubkeys  [][]byte
	relays         [][]byte
}

func (n *NostrGroupData) Unmarshal(s *cryptobyte.String) error {
	*n = NostrGroupData{}

	if !mls.ReadOpaqueVec(s, &n.nostr_group_id) || !mls.ReadOpaqueVec(s, &n.name) || !mls.ReadOpaqueVec(s, &n.description) {
		return io.ErrUnexpectedEOF
	}

	if err := mls.ReadVector(s, func(s *cryptobyte.String) error {
		var pubkey []byte
		if !mls.ReadOpaqueVec(s, &pubkey) {
			return io.ErrUnexpectedEOF
		}
		n.admin_pubkeys = append(n.admin_pubkeys, pubkey)
		return nil
	}); err != nil {
		return err
	}

	return mls.ReadVector(s, func(s *cryptobyte.String) error {
		var relay []byte
		if !mls.ReadOpaqueVec(s, &relay) {
			return io.ErrUnexpectedEOF
		}
		n.relays = append(n.relays, relay)
		return nil
	})
}

func (n *NostrGroupData) Marshal(b *cryptobyte.Builder) {
	b.AddBytes(n.nostr_group_id)
	b.AddBytes(n.name)
	b.AddBytes(n.description)

	mls.WriteVector(b, len(n.admin_pubkeys), func(b *cryptobyte.Builder, i int) {
		mls.WriteOpaqueVec(b, n.admin_pubkeys[i])
	})

	mls.WriteVector(b, len(n.relays), func(b *cryptobyte.Builder, i int) {
		mls.WriteOpaqueVec(b, n.relays[i])
	})
}
