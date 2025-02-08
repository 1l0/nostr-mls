module github.com/1l0/nostr-mls

go 1.23.4

replace github.com/emersion/go-mls => github.com/1l0/go-mls v0.0.0-20250208005016-53cd9c163136

// replace github.com/emersion/go-mls => ../../../github.com/emersion/go-mls

require (
	github.com/emersion/go-mls v0.0.0-00010101000000-000000000000
	golang.org/x/crypto v0.32.0
)

require (
	github.com/cloudflare/circl v1.3.7 // indirect
	golang.org/x/sys v0.29.0 // indirect
)
