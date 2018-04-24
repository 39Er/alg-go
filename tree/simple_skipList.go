package tree

import (
	"math/rand"
)

const DefaultLevel = 32

type nodeStructure struct {
	key, value interface{}
	forward    []*nodeStructure
}

// next returns the next node in the skip list containing n.
func (n *nodeStructure) next() *nodeStructure {
	if len(n.forward) == 0 {
		return nil
	}
	return n.forward[0]
}

type simple_skipList struct {
	header   *nodeStructure
	MaxLevel int
	length   int
	lessThan func(l, r interface{}) bool
}

func InitSkipList(lessThan func(l, r interface{}) bool, maxLevel ...int) *simple_skipList {
	var level int
	if len(maxLevel) == 0 {
		level = DefaultLevel
	} else {
		level = maxLevel[0]
	}
	header := new(nodeStructure)
	//for i := 0; i < MaxLevel; i++ {
	//	header.forward[i] = nil
	//}
	return &simple_skipList{
		header:   header,
		lessThan: lessThan,
		MaxLevel: level,
	}
}

func (s *simple_skipList) Len() int {
	return s.length
}

func (s *simple_skipList) level() int {
	return len(s.header.forward) - 1
}

func (s *simple_skipList) getPath(current *nodeStructure, update []*nodeStructure, key interface{}) *nodeStructure {
	depth := len(current.forward) - 1
	for i := depth; i >= 0; i-- {
		for current.forward[i] != nil && s.lessThan(current.forward[i].key, key) {
			current = current.forward[i]
		}
		if update != nil {
			update[i] = current
		}
	}
	return current.next()
}

func (s *simple_skipList) Seek(key interface{}) bool {
	candidate := s.getPath(s.header, nil, key)
	if candidate == nil || candidate.key != key {
		return false
	}
	return true
}

func (s *simple_skipList) Get(key interface{}) (interface{}, bool) {
	candidate := s.getPath(s.header, nil, key)
	if candidate == nil || candidate.key != key {
		return nil, false
	}
	return candidate.value, true
}

func (s *simple_skipList) Set(key interface{}, value interface{}) {
	if key == nil {
		panic("goskiplist: nil keys are not supported")
	}
	// s.level starts from 0, so we need to allocate one.
	update := make([]*nodeStructure, s.level()+1, s.effectiveLevel()+1)
	candidate := s.getPath(s.header, update, key)

	if candidate != nil && candidate.key == key {
		candidate.value = value
		return
	}

	newLevel := s.randomLevel()

	if currentLevel := s.level(); newLevel > currentLevel {
		// there are no pointers for the higher levels in
		// update. Header should be there. Also add higher
		// level links to the header.
		for i := currentLevel + 1; i <= newLevel; i++ {
			update = append(update, s.header)
			s.header.forward = append(s.header.forward, nil)
		}
	}

	newNode := &nodeStructure{
		forward: make([]*nodeStructure, newLevel+1, s.effectiveLevel()+1),
		key:     key,
		value:   value,
	}

	for i := 0; i <= newLevel; i++ {
		newNode.forward[i] = update[i].forward[i]
		update[i].forward[i] = newNode
	}

	s.length++
}

func (s *simple_skipList) Delete(key interface{}) (interface{}, bool) {
	if key == nil {
		panic("goskiplist: nil keys are not supported")
	}
	// s.level starts from 0, so we need to allocate one.
	update := make([]*nodeStructure, s.level()+1, s.effectiveLevel()+1)
	candidate := s.getPath(s.header, update, key)
	if candidate == nil || candidate.key != key {
		return nil, false
	}
	for i := 0; i < s.level(); i++ {
		update[i].forward[i] = candidate.forward[i]
	}
	for s.level() > 0 && s.header.forward[s.level()] == nil {
		s.header.forward = s.header.forward[:s.level()]
	}
	s.length--
	return candidate.value, true
}

func maxValue(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func (s *simple_skipList) effectiveLevel() int {
	return maxValue(s.level(), s.MaxLevel)
}

// TODO:Returns a new random level.
func (s *simple_skipList) randomLevel() (n int) {
	for n = 0; n < s.effectiveLevel() && rand.Float64() < p; n++ {
	}
	return
}

// Ordered is an interface which can be linearly ordered by the
// LessThan method, whereby this instance is deemed to be less than
// other. Additionally, Ordered instances should behave properly when
// compared using == and !=.
type OrderValue interface {
	LessThan(other OrderValue) bool
}

// New returns a new SkipList.
//
// Its keys must implement the Ordered interface.
func NewSkipList(level ...int) *simple_skipList {
	comparator := func(left, right interface{}) bool {
		return left.(OrderValue).LessThan(right.(OrderValue))
	}
	return InitSkipList(comparator, level...)

}

// NewIntKey returns a SkipList that accepts int keys.
func NewIntSkipList(level ...int) *simple_skipList {
	return InitSkipList(func(l, r interface{}) bool {
		return l.(int) < r.(int)
	}, level...)
}

// NewStringMap returns a SkipList that accepts string keys.
func NewStringSkip(level ...int) *simple_skipList {
	return InitSkipList(func(l, r interface{}) bool {
		return l.(string) < r.(string)
	}, level...)
}
