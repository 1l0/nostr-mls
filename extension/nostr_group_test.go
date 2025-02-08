package extension

import (
	"testing"

	"github.com/emersion/go-mls"
	"golang.org/x/crypto/cryptobyte"
)

func TestNewNostrGroup(t *testing.T) {
	group, err := NewNostrGroupData(
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
	if len(group.ID) != 32 {
		t.Fatalf("group id != 32 length: %d", len(group.ID))
	}
	if string(group.Name) != "bla" {
		t.Fatal("failed to set group name")
	}
	if string(group.Description) != "foo" {
		t.Fatal("failed to set group description")
	}
	if string(group.Admins[1]) != "0000000000000000000000000000002" {
		t.Fatal("failed to set group admins")
	}
	if string(group.Relays[1]) != "wss://relay.efgh.net" {
		t.Fatal("failed to set group relays")
	}
}

func TestNostrGroupFromGroupContext(t *testing.T) {
	group, err := NewNostrGroupData(
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

	result, err := NostrGroupDataFromContext(ctx)
	if err != nil {
		t.Fatal(err)
	}
	if len(result.ID) != 32 {
		t.Fatalf("group id != 32 length: %d", len(result.ID))
	}
	if string(result.Name) != "bla" {
		t.Fatal("failed to set group name")
	}
	if string(result.Description) != "foo" {
		t.Fatal("failed to set group description")
	}
	if string(result.Admins[1]) != "0000000000000000000000000000002" {
		t.Fatal("failed to set group admins")
	}
	if string(result.Relays[1]) != "wss://relay.efgh.net" {
		t.Fatal("failed to set group relays")
	}
}
