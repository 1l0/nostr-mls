package main

import (
	"encoding/hex"
	"encoding/json"
	"log"
	"time"

	nmls "github.com/1l0/nostr-mls"
	"github.com/nbd-wtf/go-nostr"
	"github.com/nbd-wtf/go-nostr/nip44"
)

func main() {
	nostrMLS := nmls.NewNostrMLS(nmls.NewMemoryStore())
	aliceKey := nostr.GeneratePrivateKey()
	alicePubkey, err := nostr.GetPublicKey(aliceKey)
	if err != nil {
		panic(err)
	}
	bobKey := nostr.GeneratePrivateKey()
	bobPubkey, err := nostr.GetPublicKey(bobKey)
	if err != nil {
		panic(err)
	}

	// bob's key package is published as [nmls.KindMLSKeyPackage] event on relays
	bobPkgHex, err := nostrMLS.CreateKeyPackageForEvent(bobPubkey)

	// ================================
	// We're now acting as Alice
	// ================================

	// fetch bob's key package event
	bobPkg, err := nostrMLS.ParseKeyPackage(bobPkgHex)
	if err != nil {
		panic(err)
	}

	// create a group bob can join
	group, err := nostrMLS.NewCreateGroup(
		"Bob & Alice",
		"A secret chat between Bob and Alice",
		alicePubkey,
		[]nmls.KeyPackage{*bobPkg},
		[]string{aliceKey, bobPubkey},
		[]string{"ws://localhost:8080"},
	)
	if err != nil {
		panic(err)
	}
	groupData := group.NostrGroupData
	groupMessage := group.Message

	// TODO: send a gift-wrapped welcome event to bob's DM inboxes

	// send a message to group relays
	unsignedEvent := &nostr.Event{
		PubKey:    alicePubkey,
		Kind:      nostr.KindSimpleGroupChatMessage,
		CreatedAt: nostr.Timestamp(time.Now().Unix()),
		Content:   "Hi Bob!",
	}
	unsignedEvent.ID = unsignedEvent.GetID()
	rawMessage, err := nostrMLS.CreateMessageForGroup(groupData.ID, unsignedEvent.String())
	if err != nil {
		panic(err)
	}
	exporterKey, _, err := nostrMLS.ExporterSecretWithEpoch(groupData.ID)
	if err != nil {
		panic(err)
	}
	exporterPubkey, err := nostr.GetPublicKey(exporterKey)
	if err != nil {
		panic(err)
	}
	convKey, err := nip44.GenerateConversationKey(exporterPubkey, exporterKey)
	if err != nil {
		panic(err)
	}
	encryptedMessage, err := nip44.Encrypt(rawMessage, convKey)
	if err != nil {
		panic(err)
	}
	ephemeralKey := nostr.GeneratePrivateKey()
	idHex := hex.EncodeToString(groupData.ID)
	messageEvent := &nostr.Event{
		Kind:    nmls.KindMLSGroupMessage,
		Content: encryptedMessage,
		Tags: nostr.Tags{
			nostr.Tag{"h", idHex},
		},
	}
	messageEvent.Sign(ephemeralKey)

	// ================================
	// We're now acting as Bob
	// ================================

	// fetch a welcome event from DM inboxes
	_, err = nostrMLS.PreviewWelcomeEvent(groupMessage)
	if err != nil {
		panic(err)
	}
	// if you like it, join the group
	_, err = nostrMLS.JoinGroupFromWelcome(groupMessage)
	if err != nil {
		panic(err)
	}

	// fetch a group message event from group relays
	decryptedMessage, err := nip44.Decrypt(messageEvent.Content, convKey)
	if err != nil {
		panic(err)
	}
	rawChatEvent, err := nostrMLS.ProcessMessageForGroup(groupData.ID, decryptedMessage)
	if err != nil {
		panic(err)
	}
	var evt nostr.Event
	if err := json.Unmarshal(rawChatEvent, &evt); err != nil {
		panic(err)
	}
	log.Printf("Bob received from Alice:\n%+v\n", evt)
}
