package main

import "time"

// makeMockDB -  Obscure the db implementation with faked latency
func makeMockDB() Database {
	return NewMockDB(300)
}

type MockDB struct {
	latency uint
	data    []DocBundle
}

func NewMockDB(latencyMs uint) *MockDB {
	return &MockDB{latencyMs, nil}
}
func (m *MockDB) netLatency() {
	time.Sleep(time.Duration(m.latency) * time.Millisecond)
}
func (m *MockDB) SetBatch(data []DocBundle) {
	m.netLatency()
	m.data = data
}
func (m *MockDB) FetchBatch(path string) []DocBundle {
	m.netLatency()
	return m.data
}
func (m *MockDB) DeleteCollection(path string) {
	m.data = nil
}
