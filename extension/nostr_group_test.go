package extension

import (
	"testing"
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
