package extension

import (
	"testing"

	"github.com/emersion/go-mls"
	"golang.org/x/crypto/cryptobyte"
)

func TestNewNostrGroup(t *testing.T) {
	group, err := NewNostrGroup(
		"bla",
		"foo",
		[]string{
			"0000000000000000000000000000001",
			"0000000000000000000000000000002"},
		[]string{
			"wss://relay.abcd.com",
			"wss://relay.efgh.net",
		},
	)
	if err != nil {
		t.Fatal(err)
	}
	if len(group.id) != 32 {
		t.Fatalf("group id != 32 length: %d", len(group.id))
	}
	if string(group.name) != "bla" {
		t.Fatal("failed to set group name")
	}
	if string(group.description) != "foo" {
		t.Fatal("failed to set group description")
	}
	if string(group.admins[1]) != "0000000000000000000000000000002" {
		t.Fatal("failed to set group admins")
	}
	if string(group.relays[1]) != "wss://relay.efgh.net" {
		t.Fatal("failed to set group relays")
	}
}

func TestNostrGroupFromGroupContext(t *testing.T) {
	group, err := NewNostrGroup(
		"bla",
		"foo",
		[]string{
			"0000000000000000000000000000001",
			"0000000000000000000000000000002"},
		[]string{
			"wss://relay.abcd.com",
			"wss://relay.efgh.net",
		},
	)
	if err != nil {
		t.Fatal(err)
	}

	var b cryptobyte.Builder
	group.Marshal(&b)
	data, err := b.Bytes()
	if err != nil {
		t.Fatal(err)
	}

	id, err := generateRandomBytes(32)
	if err != nil {
		t.Fatal(err)
	}

	ctx := &mls.GroupContext{
		Version:                 1,
		CipherSuite:             mls.CipherSuiteMLS_128_DHKEMX25519_AES128GCM_SHA256_Ed25519,
		GroupID:                 id,
		Epoch:                   1615215381,
		TreeHash:                nil,
		ConfirmedTranscriptHash: nil,
		Extensions: []mls.Extension{
			{
				ExtensionType: ExtensionTypeNostrGroup,
				ExtensionData: data,
			},
		},
	}

	result, err := NostrGroupFromGroupContext(ctx)
	if err != nil {
		t.Fatal(err)
	}
	if len(result.id) != 32 {
		t.Fatalf("group id != 32 length: %d", len(result.id))
	}
	if string(result.name) != "bla" {
		t.Fatal("failed to set group name")
	}
	if string(result.description) != "foo" {
		t.Fatal("failed to set group description")
	}
	if string(result.admins[1]) != "0000000000000000000000000000002" {
		t.Fatal("failed to set group admins")
	}
	if string(result.relays[1]) != "wss://relay.efgh.net" {
		t.Fatal("failed to set group relays")
	}
}
