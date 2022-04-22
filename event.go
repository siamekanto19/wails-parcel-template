package main

import (
	"math/rand"
	"sync"
)

type Event struct {
	Type string
	Data string
	Seq  int
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

var seq int
var mu sync.Mutex
var eventTypes = []string{"hello", "world", "phew", "oops"}

func NewDummyEvent() Event {
	mu.Lock()
	seq++
	mu.Unlock()
	return Event{
		Type: eventTypes[rand.Intn(len(eventTypes))],
		Data: RandStringRunes(32),
		Seq:  seq,
	}
}
