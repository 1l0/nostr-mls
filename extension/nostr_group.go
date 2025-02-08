package extension

import (
	"crypto/rand"
	"fmt"
	"io"

	"github.com/emersion/go-mls"
	"golang.org/x/crypto/cryptobyte"
)

const (
	ExtensionTypeNostrGroup mls.ExtensionType = 0xF2EE
)

type NostrGroupData struct {
	ID          []byte
	Name        []byte
	Description []byte
	Admins      [][]byte
	Relays      [][]byte
}

// NewNostrGroupData creates new NostrGroup.
// admins: admin pubkeys
// relays: relay URLs
func NewNostrGroupData(name, description string, admins, relays []string) (*NostrGroupData, error) {
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
	return &NostrGroupData{
		ID:          id,
		Name:        []byte(name),
		Description: []byte(description),
		Admins:      a,
		Relays:      r,
	}, nil
}

func NostrGroupDataFromContext(ctx *mls.GroupContext) (*NostrGroupData, error) {
	var ext *mls.Extension = nil
	for _, ex := range ctx.Extensions {
		if ex.ExtensionType == ExtensionTypeNostrGroup {
			ext = &ex
		}
	}
	if ext == nil {
		return nil, fmt.Errorf("NostrGroup extension not found")
	}
	ng := NostrGroupData{}
	cs := cryptobyte.String(ext.ExtensionData)
	if err := ng.Unmarshal(&cs); err != nil {
		return nil, err
	}
	return &ng, nil
}

func (n *NostrGroupData) Unmarshal(s *cryptobyte.String) error {
	*n = NostrGroupData{}

	if !mls.ReadOpaqueVec(s, &n.ID) || !mls.ReadOpaqueVec(s, &n.Name) || !mls.ReadOpaqueVec(s, &n.Description) {
		return io.ErrUnexpectedEOF
	}

	if err := mls.ReadVector(s, func(s *cryptobyte.String) error {
		var pubkey []byte
		if !mls.ReadOpaqueVec(s, &pubkey) {
			return io.ErrUnexpectedEOF
		}
		n.Admins = append(n.Admins, pubkey)
		return nil
	}); err != nil {
		return err
	}

	return mls.ReadVector(s, func(s *cryptobyte.String) error {
		var relay []byte
		if !mls.ReadOpaqueVec(s, &relay) {
			return io.ErrUnexpectedEOF
		}
		n.Relays = append(n.Relays, relay)
		return nil
	})
}

func (n *NostrGroupData) Marshal(b *cryptobyte.Builder) {
	mls.WriteOpaqueVec(b, n.ID)
	mls.WriteOpaqueVec(b, n.Name)
	mls.WriteOpaqueVec(b, n.Description)

	mls.WriteVector(b, len(n.Admins), func(b *cryptobyte.Builder, i int) {
		mls.WriteOpaqueVec(b, n.Admins[i])
	})

	mls.WriteVector(b, len(n.Relays), func(b *cryptobyte.Builder, i int) {
		mls.WriteOpaqueVec(b, n.Relays[i])
	})
}

func generateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n, n)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}
	return b, nil
}
