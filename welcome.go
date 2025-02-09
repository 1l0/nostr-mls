package mls

import (
	"github.com/emersion/go-mls"
	"golang.org/x/crypto/cryptobyte"

	"github.com/1l0/nostr-mls/extension"
)

func (n *NostrMLS) GroupInfoFromWelcome(welcomeMessage []byte, keyPackage *mls.KeyPackage, pubkey string) (*mls.GroupInfo, error) {
	w := mls.Welcome{}
	b := cryptobyte.String(welcomeMessage)
	if err := w.Unmarshal(&b); err != nil {
		return nil, err
	}
	ref, err := keyPackage.GenerateRef()
	if err != nil {
		return nil, err
	}
	priv, err := n.store.Get([]byte(pubkey))
	if err != nil {
		return nil, err
	}
	groupSecrets, err := w.DecryptGroupSecrets(ref, priv)
	if err != nil {
		return nil, err
	}
	groupInfo, err := w.DecryptGroupInfo(groupSecrets.JoinerSecret, nil)
	if err != nil {
		return nil, err
	}
	return groupInfo, nil
}

func (n *NostrMLS) Join(info *mls.GroupInfo) (*Group, error) {
	groupData, err := extension.NostrGroupDataFromContext(&info.GroupContext)
	if err != nil {
		return nil, err
	}

	return &Group{
		Context:        &info.GroupContext,
		NostrGroupData: groupData,
	}, nil
}
