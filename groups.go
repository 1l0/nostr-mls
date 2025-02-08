package mls

import (
	"github.com/emersion/go-mls"

	"github.com/1l0/nostr-mls/extension"
)

type CreateGroup struct {
	groupContext   *mls.GroupContext
	Message        []byte
	NostrGroupData *extension.NostrGroupData
}

type SelfUpdate struct {
	Message            []byte
	PrevExporterSecret string // hex
	ExporterSecret     string // hex
	Epoch              uint64
}

type KeyPackage mls.KeyPackage

// creator: pubkey
// admins: list of pubkeys
func (n *NostrMLS) NewCreateGroup(name, description, creator string, memberKeyPackages []KeyPackage, admins, relays []string) (*CreateGroup, error) {
	// cap := n.Capabilities()
	// cred, err := n.GenerateCredentialWithKey(creator)
	// if err != nil {
	// 	return nil, err
	// }
	// ctx := &mls.GroupContext{}

	// data, err := extension.NewNostrGroupData(name, description, admins, relays)
	// if err != nil {
	// 	return nil, err
	// }
	// var builder cryptobyte.Builder
	// data.Marshal(&builder)
	// rawData, err := builder.Bytes()
	// if err != nil {
	// 	return nil, err
	// }
	// nostrGroupExt := mls.Extension{
	// 	ExtensionType: extension.ExtensionTypeNostrGroup,
	// 	ExtensionData: rawData,
	// }
	return nil, nil
}

func (n *NostrMLS) CreateMessageForGroup(groupID []byte, message string) (string, error) {
	// TODO
	return "", nil
}

func (n *NostrMLS) ExporterSecretWithEpoch(groupID []byte) (string, uint64, error) {
	// TODO
	return "", 0, nil
}

func (n *NostrMLS) ProcessMessageForGroup(groupID []byte, message string) ([]byte, error) {
	// TODO
	return nil, nil
}

func (n *NostrMLS) MemberPubkeys(groupID []byte) ([]string, error) {
	// TODO
	return nil, nil
}

func (n *NostrMLS) NewSelfUpdate(groupID []byte) (*SelfUpdate, error) {
	// TODO
	return nil, nil
}
