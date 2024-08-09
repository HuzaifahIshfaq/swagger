package operations

import (
	"swagger/models"
	"sync"
)

// In-memory user storage (shared between handlers)
var userStore = struct {
	sync.RWMutex
	users map[int64]*models.User
}{users: make(map[int64]*models.User)}
