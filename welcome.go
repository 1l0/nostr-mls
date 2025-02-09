package mls

import (
	"github.com/emersion/go-mls"

	"github.com/1l0/nostr-mls/extension"
)

type WelcomePreview struct {
	Welcome        *Welcome
	NostrGroupData *extension.NostrGroupData
}

type Welcome mls.Welcome

func (n *NostrMLS) PreviewWelcome(welcomeMessage []byte) (*WelcomePreview, error) {
	// TODO
	return nil, nil
}

func (p *WelcomePreview) Join() (*Group, error) {
	// TODO
	return nil, nil
}
