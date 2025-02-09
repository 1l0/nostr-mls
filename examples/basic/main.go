package main

import (
	"encoding/hex"
	"encoding/json"
	"log"

	mls "github.com/1l0/nostr-mls"
	"github.com/nbd-wtf/go-nostr"
	"github.com/nbd-wtf/go-nostr/nip44"
)

func main() {
	nostrMLS := mls.NewNostrMLS(mls.NewMemoryStore())
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

	// Bob's key package is published as [mls.KindMLSKeyPackage] event on relays
	bobPkgHex, err := nostrMLS.CreateKeyPackageHex(bobPubkey)

	// ================================
	// We're now acting as Alice
	// ================================

	// fetch Bob's key package event
	bobPkg, err := nostrMLS.ParseKeyPackage(bobPkgHex)
	if err != nil {
		panic(err)
	}

	// create a group Bob can join and publish [mls.KindMLSGroupMessage] with Proposal/Commit on group relays
	group, err := nostrMLS.NewGroup(
		"Bob & Alice",
		"A secret chat between Bob and Alice",
		alicePubkey,
		[]mls.KeyPackage{*bobPkg},
		[]string{aliceKey, bobPubkey},
		[]string{"ws://localhost:8080"},
	)
	if err != nil {
		panic(err)
	}
	groupData := group.NostrGroupData
	welcomeMessage := group.WelcomeMessage

	// wait publishing done on group relays to match the welcome message,
	// then send a gift-wrapped welcome event to Bob's DM relays
	welcomeHex := hex.EncodeToString(welcomeMessage)
	unsignedEvent := &nostr.Event{
		PubKey:    alicePubkey,
		Kind:      mls.KindMLSWelcome,
		CreatedAt: nostr.Now(),
		Content:   welcomeHex,
	}
	unsignedEvent.ID = unsignedEvent.GetID()

	// ================================
	// We're now acting as Bob
	// ================================

	// fetch a gift-wrapped welcome event from DM relays
	info, err := nostrMLS.GroupInfoFromPreview(welcomeMessage, bobPkg, bobPubkey)
	if err != nil {
		panic(err)
	}
	// if we are interested in, let's join the group
	group, err = nostrMLS.Join(info)
	if err != nil {
		panic(err)
	}
	groupData = group.NostrGroupData

	// ================================
	// We're now acting as Alice
	// ================================

	// if Bob joins, [mls.KindMLSGroupMessage] is updated
	// after the comfirmation, send a message on group relays
	unsignedEvent = &nostr.Event{
		PubKey:    alicePubkey,
		Kind:      nostr.KindSimpleGroupChatMessage,
		CreatedAt: nostr.Now(),
		Content:   "Hi Bob!",
	}
	unsignedEvent.ID = unsignedEvent.GetID()
	rawMessage, err := nostrMLS.CreateMessage(groupData.ID, unsignedEvent.String())
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
	chatMessageEvent := &nostr.Event{
		CreatedAt: nostr.Now(),
		Kind:      mls.KindMLSGroupMessage,
		Content:   encryptedMessage,
		Tags: nostr.Tags{
			nostr.Tag{"h", idHex},
		},
	}
	chatMessageEvent.Sign(ephemeralKey)

	// ================================
	// We're now acting as Bob
	// ================================

	// fetch the message from group relays
	decryptedMessage, err := nip44.Decrypt(chatMessageEvent.Content, convKey)
	if err != nil {
		panic(err)
	}
	serializedChatEvent, err := nostrMLS.ParseSerializedEventFromGroupMessage(groupData.ID, decryptedMessage)
	if err != nil {
		panic(err)
	}
	var evt nostr.Event
	if err := json.Unmarshal(serializedChatEvent, &evt); err != nil {
		panic(err)
	}
	log.Printf("Message from Alice:\n%+v\n", evt)
}
