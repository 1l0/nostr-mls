package mls

import (
	"github.com/emersion/go-mls"

	"github.com/1l0/nostr-mls/extension"
)

type WelcomePreview struct {
	Welcome    *Welcome
	NostrGroup *extension.NostrGroupData
}

type Welcome mls.Welcome

func (n *NostrMLS) PreviewWelcome(welcomeMessage []byte) (*WelcomePreview, error) {
	// TODO
	return nil, nil
}

func (n *NostrMLS) Join(welcomeMessage []byte) (*Group, error) {
	// TODO
	return nil, nil
}
