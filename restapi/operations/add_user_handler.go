package operations

import (
	"net/http"

	"swagger/models"

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"
)

// In-memory user storage

// AddUserHandlerImpl implementation for adding a new user
type AddUserHandlerImpl struct{}

// NewAddUserHandler creates a new handler for adding a user
func NewAddUserHandler() AddUserHandler {
	return &AddUserHandlerImpl{}
}

// Handle the add user request
func (h *AddUserHandlerImpl) Handle(params AddUserParams) middleware.Responder {
	user := params.User

	// Validate if user ID already exists
	userStore.RLock()
	if _, exists := userStore.users[*user.ID]; exists {
		userStore.RUnlock()
		return NewAddUserDefault(http.StatusConflict).WithPayload(&models.ErrorModel{
			Code:    swag.Int32(http.StatusConflict),
			Message: swag.String("User already exists"),
		})
	}
	userStore.RUnlock()

	// Add user to the store
	userStore.Lock()
	userStore.users[*user.ID] = user
	userStore.Unlock()

	// Return success response
	return NewAddUserOK().WithPayload(user)
}
