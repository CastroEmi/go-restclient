package rest

import (
	"container/list"
	"sync"
)

// ResourceCache, is an LRU-TTL Cache, that caches Responses base on headers
// It uses 3 goroutines -> one for LRU, and the other two for TTL.

// Cache
var resourceCache *resourceTTLLRUMap

// ByteSize is a helper for configuring MaxCacheSize
type ByteSize int64

const (
	_ = iota

	// KB = KiloBytes
	KB ByteSize = 1 << (10 * iota)

	// MB = MegaBytes
	MB

	// GB = GygaBytes
	GB
)

// MaxCacheSize is the maximum byte size to be hold the ResourceCache
// Default is 1 Gigabyte
// Type: rest.ByteSize
var MaxCacheSize = 1 * GB

// Current cache Size
var cacheSize int64

type lruOperation int

const (
	move lruOperation = iota
	push
	delete
	last
)

type lruMsg struct {
	operation lruOperation
	resp      *Response
}

type resourceTTLLRUMap struct {
	cache    map[string]*Response
	skipList *skipList    // skipList for TTL
	lruList  *list.List   // List for LRU
	lruChan  chan *lruMsg // Channel for LRU messages
	ttlChan  chan bool    // Channel for TTL messages
	popChan  chan string
	rwMutex  sync.RWMutex // Read Write Locking Mutex
}
