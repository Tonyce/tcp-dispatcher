package channel

import (
	"log"
	"sync"
)

type Manager struct {
	mu       sync.RWMutex
	channels map[string]*Channel
}

var (
	ChannelManager *Manager
	mutex          sync.Mutex
)

func newManager() *Manager {
	return &Manager{
		channels: make(map[string]*Channel),
	}
}

func GetChannelManager() *Manager {
	mutex.Lock()
	if ChannelManager == nil {
		ChannelManager = newManager()
	}
	mutex.Unlock()
	return ChannelManager
}

func (m *Manager) AddChannel(name string, channel *Channel) {
	var (
		exist bool
	)
	m.mu.Lock()
	if _, exist = m.channels[name]; !exist {
		m.channels[name] = channel
		log.Println("add new room", name)
	}
	m.mu.Unlock()
}

func (m *Manager) GetChannel(name string) *Channel {
	m.mu.RLock()
	defer m.mu.RUnlock()
	channel, exist := m.channels[name]
	if !exist {
		return nil
	}
	return channel
}
