package mls

// Store is used for [NostrMLS]
type Store interface {
	// Upsert updates or inserts an entry
	Upsert(key, value []byte) error
	// Get gets an entry
	Get(key []byte) ([]byte, error)
	// Remove removes an entry
	Remove(key []byte) error
	// Clear clears entire store
	Clear() error
}
