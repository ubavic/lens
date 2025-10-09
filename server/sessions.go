package server

import (
	"fmt"
	"sync"
	"time"

	"github.com/ubavic/lens/view/component"
)

type Session struct {
	uid        string
	lastAccess time.Time
	handlers   component.HandlerMap
}

var sessionsMu sync.Mutex
var sessions map[string]Session

func newSession(uid string) {
	sessionsMu.Lock()
	defer sessionsMu.Unlock()

	fmt.Println("session created", uid)

	sessions[uid] = Session{
		uid:        uid,
		lastAccess: time.Now(),
		handlers:   make(map[component.Uid]component.Handler),
	}
}

func removeSession(uid string) {
	sessionsMu.Lock()
	defer sessionsMu.Unlock()

	delete(sessions, uid)
}
