package mls

import (
	"github.com/emersion/go-mls"

	"github.com/1l0/nostr-mls/extension"
)

type WelcomePreview struct {
	Welcome    *Welcome
	NostrGroup *extension.NostrGroupData
}

type JoinedGroup struct {
	Group      *mls.GroupContext
	NostrGroup *extension.NostrGroupData
}

type Welcome mls.Welcome

func (n *NostrMLS) ParseWelcomeMessage(message []byte) (*Welcome, *extension.NostrGroupData, error) {
	// TODO
	return nil, nil, nil
}

func (n *NostrMLS) PreviewWelcomeEvent(message []byte) (*WelcomePreview, error) {
	// TODO
	return nil, nil
}

func (n *NostrMLS) JoinGroupFromWelcome(message []byte) (*JoinedGroup, error) {
	// TODO
	return nil, nil
}
