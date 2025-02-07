package mls

import (
	"github.com/emersion/go-mls"

	"github.com/1l0/nostr-mls/extension"
)

type CreateGroupResult struct {
	MLSGroup   mls.GroupContext
	Message    []byte
	NostrGroup extension.NostrGroup
}

type SelfUpdateResult struct {
	Message            []byte
	PrevExporterSecret string // hex
	ExporterSecret     string // hex
	Epoch              uint64
}

// creator: pubkey
// admins: list of pubkeys
func (n *NostrMLS) createGroup(name, description, creator string, memberKeyPackages []mls.KeyPackage, admins, relays []string) (*CreateGroupResult, error) {
	// TOOD
	return nil, nil
}

func (n *NostrMLS) createMessageForGroup(groupID []byte, message string) ([]byte, error) {
	// TODO
	return nil, nil
}

func (n *NostrMLS) exporterSecretWithEpoch(groupID []byte) (string, uint64, error) {
	// TODO
	return "", 0, nil
}

func (n *NostrMLS) processMessageForGroup(groupID, message []byte) ([]byte, error) {
	// TODO
	return nil, nil
}

func (n *NostrMLS) memberPubkeys(groupID []byte) ([]string, error) {
	// TODO
	return nil, nil
}

func (n *NostrMLS) selfUpdate(groupID []byte) (*SelfUpdateResult, error) {
	// TODO
	return nil, nil
}
