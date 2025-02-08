package mls

import (
	"github.com/emersion/go-mls"

	"github.com/1l0/nostr-mls/extension"
)

type Group struct {
	groupContext   *mls.GroupContext
	WelcomeMessage []byte
	NostrGroupData *extension.NostrGroupData
}

type GroupUpdateInfo struct {
	WelcomeMessage     []byte
	PrevExporterSecret string // hex
	ExporterSecret     string // hex
	Epoch              uint64
}

type KeyPackage mls.KeyPackage

// creator: pubkey
// admins: list of pubkeys
func (n *NostrMLS) NewGroup(name, description, creator string, memberKeyPackages []KeyPackage, admins, relays []string) (*Group, error) {
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

func (n *NostrMLS) CreateMessage(groupID []byte, message string) (string, error) {
	// TODO
	return "", nil
}

func (n *NostrMLS) ExporterSecretWithEpoch(groupID []byte) (string, uint64, error) {
	// TODO
	return "", 0, nil
}

func (n *NostrMLS) ParseSerializedEventFromGroupMessage(groupID []byte, message string) ([]byte, error) {
	// TODO
	return nil, nil
}

func (n *NostrMLS) memberPubkeys(groupID []byte) ([]string, error) {
	// TODO
	return nil, nil
}

func (n *NostrMLS) updateInfo(groupID []byte) (*GroupUpdateInfo, error) {
	// TODO
	return nil, nil
}
