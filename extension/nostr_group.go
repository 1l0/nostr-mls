package extension

import (
	"crypto/rand"
	"io"

	"github.com/emersion/go-mls"
	"golang.org/x/crypto/cryptobyte"
)

const (
	ExtensionTypeNostrGroup mls.ExtensionType = 0xF2EE
)

type NostrGroup struct {
	id          []byte
	name        []byte
	description []byte
	admins      [][]byte
	relays      [][]byte
}

// NewNostrGroup creates new NostrGroup.
// admins: admin pubkeys
// relays: relay URLs
func NewNostrGroup(name, description string, admins, relays []string) (*NostrGroup, error) {
	id, err := generateRandomBytes(32)
	if err != nil {
		return nil, err
	}
	al := len(admins)
	a := make([][]byte, al, al)
	for i := 0; i < al; i++ {
		a[i] = []byte(admins[i])
	}
	rl := len(relays)
	r := make([][]byte, rl, rl)
	for i := 0; i < rl; i++ {
		r[i] = []byte(relays[i])
	}
	return &NostrGroup{
		id:          id,
		name:        []byte(name),
		description: []byte(description),
		admins:      a,
		relays:      r,
	}, nil
}

func (n *NostrGroup) Unmarshal(s *cryptobyte.String) error {
	*n = NostrGroup{}

	if !mls.ReadOpaqueVec(s, &n.id) || !mls.ReadOpaqueVec(s, &n.name) || !mls.ReadOpaqueVec(s, &n.description) {
		return io.ErrUnexpectedEOF
	}

	if err := mls.ReadVector(s, func(s *cryptobyte.String) error {
		var pubkey []byte
		if !mls.ReadOpaqueVec(s, &pubkey) {
			return io.ErrUnexpectedEOF
		}
		n.admins = append(n.admins, pubkey)
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

func (n *NostrGroup) Marshal(b *cryptobyte.Builder) {
	b.AddBytes(n.id)
	b.AddBytes(n.name)
	b.AddBytes(n.description)

	mls.WriteVector(b, len(n.admins), func(b *cryptobyte.Builder, i int) {
		mls.WriteOpaqueVec(b, n.admins[i])
	})

	mls.WriteVector(b, len(n.relays), func(b *cryptobyte.Builder, i int) {
		mls.WriteOpaqueVec(b, n.relays[i])
	})
}

func generateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}
	return b, nil
}
