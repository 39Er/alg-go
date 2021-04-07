package cache

import (
	"crypto/sha1"
	"math"
	"sort"
	"strconv"
	"sync"
)

const (
	DefaultVirualSpots = 400
	DefaultNodePrefix  = "default-"
)

type node struct {
	nodeKey   string
	spotValue uint32
}

type nodeArray []node

func (p nodeArray) Len() int { return len(p) }

func (p nodeArray) Less(i, j int) bool { return p[i].spotValue < p[j].spotValue }

func (p nodeArray) Swap(i, j int) { p[i], p[j] = p[j], p[i] }

func (p nodeArray) Sort() { sort.Sort(p) }

type HashRing struct {
	vituralSpots int
	nodes        nodeArray
	weights      map[string]int

	lock sync.RWMutex
}

func NewHashRing(spots ...int) *HashRing {
	var sportNum int
	if len(spots) == 0 {
		sportNum = DefaultVirualSpots
	} else {
		sportNum = spots[0]
	}
	return &HashRing{
		vituralSpots: sportNum,
		weights:      make(map[string]int),
	}
}

func (h *HashRing) AddNodes(nodeWeight map[string]int) {
	h.lock.Lock()
	defer h.lock.Unlock()
	for nodeKey, w := range nodeWeight {
		h.weights[nodeKey] = w
	}
	h.generate()
}

func (h *HashRing) AddDefaultNodes(count int) {
	h.lock.Lock()
	defer h.lock.Unlock()
	for i := 0; i < count; i++ {
		h.weights[DefaultNodePrefix+strconv.Itoa(i)] = 1
	}
	h.generate()
}

func (h *HashRing) AddNode(nodeKey string, weight int) {
	h.lock.Lock()
	defer h.lock.Unlock()
	h.weights[nodeKey] = weight
	h.generate()
}

func (h *HashRing) Delete(nodeKey string) {
	h.lock.Lock()
	defer h.lock.Unlock()
	delete(h.weights, nodeKey)
	h.generate()
}

func (h *HashRing) generate() {
	var totalW int
	for _, w := range h.weights {
		totalW += w
	}
	totalVirualSpots := h.vituralSpots * len(h.weights)
	h.nodes = nodeArray{}
	for nodeKey, w := range h.weights {
		spots := int(math.Floor(float64(w) / float64(totalW) * float64(totalVirualSpots)))
		for i := 1; i <= spots; i++ {
			hash := sha1.New()
			hash.Write([]byte(nodeKey + ":" + strconv.Itoa(i)))
			hashBytes := hash.Sum(nil)
			n := node{
				nodeKey:   nodeKey,
				spotValue: getValue(hashBytes[6:10]),
			}
			h.nodes = append(h.nodes, n)
			hash.Reset()
		}
	}
	h.nodes.Sort()
}

func getValue(bs []byte) uint32 {
	if len(bs) < 4 {
		return 0
	}
	return uint32(bs[3])<<24 | uint32(bs[2])<<16 | uint32(bs[1])<<8 | uint32(bs[0])
}

func (h *HashRing) GetNode(s string) string {
	h.lock.RLock()
	defer h.lock.RUnlock()
	if len(h.nodes) == 0 {
		return ""
	}
	hash := sha1.New()
	hash.Write([]byte(s))
	hashBytes := hash.Sum(nil)
	v := getValue(hashBytes[6:10])
	i := sort.Search(len(h.nodes), func(i int) bool { return h.nodes[i].spotValue >= v })
	if i == len(h.nodes) {
		i = 0
	}
	return h.nodes[i].nodeKey
}

func (h *HashRing) ListNodeKeys() []string {
	h.lock.RLock()
	defer h.lock.RUnlock()
	var keys []string
	for key := range h.weights {
		keys = append(keys, key)
	}
	return keys
}
