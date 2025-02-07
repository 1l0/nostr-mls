package mls

import (
	"github.com/emersion/go-mls"

	"github.com/1l0/nostr-mls/extension"
)

type WelcomePreview struct {
	Welcome    *mls.Welcome
	NostrGroup *extension.NostrGroup
}

type JoinedGroupResult struct {
	Group      *mls.GroupContext
	NostrGroup *extension.NostrGroup
}

func (n *NostrMLS) parseWelcomeMessage(message []byte) (*mls.Welcome, *extension.NostrGroup, error) {
	// TODO
	return nil, nil, nil
}

func (n *NostrMLS) previewWelcomeEvent(message []byte) (*WelcomePreview, error) {
	// TODO
	return nil, nil
}

func (n *NostrMLS) joinGroupFromWelcome(message []byte) (*JoinedGroupResult, error) {
	// TODO
	return nil, nil
}
