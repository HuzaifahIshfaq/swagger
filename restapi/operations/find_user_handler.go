package operations

import (
	"swagger/models"

	"github.com/go-openapi/runtime/middleware"
)

// In-memory user storage for demonstration purposes

// FindUsersHandlerImpl implementation for finding users
type FindUsersHandlerImpl struct{}

// NewFindUsersHandler creates a new handler for finding users
func NewFindUsersHandler() FindUsersHandler {
	return &FindUsersHandlerImpl{}
}

// Handle the find users request
func (h *FindUsersHandlerImpl) Handle(params FindUsersParams) middleware.Responder {
	userStore.RLock()
	defer userStore.RUnlock()

	users := make([]*models.User, 0, len(userStore.users))
	for _, user := range userStore.users {
		users = append(users, user)
	}

	return NewFindUsersOK().WithPayload(users)
}
