package atomic

import (
	"sync"
)

type String struct {
	m sync.Mutex
	v string
}

func NewString() *String {
	return &String{}
}

func (v *String) Set(value string) bool {
	v.m.Lock()
	old := v.v
	v.v = value
	v.m.Unlock()

	return old != value
}