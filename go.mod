module github.com/1l0/nostr-mls

go 1.23.4

// replace github.com/emersion/go-mls => github.com/1l0/go-mls v0.0.0-20250208090403-b38d5b796348
replace github.com/emersion/go-mls => ../../../github.com/emersion/go-mls

require (
	github.com/emersion/go-mls v0.0.0-20250111215207-271336b8374d
	golang.org/x/crypto v0.32.0
)

require (
	github.com/cloudflare/circl v1.3.7 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/stretchr/testify v1.10.0 // indirect
	golang.org/x/sys v0.29.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
