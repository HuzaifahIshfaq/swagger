package operations

import (
	"net/http"
	"swagger/models"

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"
)

// UpdateUserHandlerImpl implements the handler for updating a user
type UpdateUserHandlerImpl struct{}

// NewUpdateUserHandler creates a new handler for updating a user
func NewUpdateUserHandler() UpdateUserHandler {
	return &UpdateUserHandlerImpl{}
}

// Handle handles the update user request
func (h *UpdateUserHandlerImpl) Handle(params UpdateUserParams) middleware.Responder {
	userID := params.ID
	updatedUser := params.User

	// Access the global user store (assuming it's already declared in store.go)
	userStore.Lock()
	defer userStore.Unlock()

	// Check if the user exists
	existingUser, exists := userStore.users[userID]
	if !exists {
		// Return 404 Not Found with error payload
		return NewUpdateUserNotFound().WithPayload(&models.ErrorModel{
			Code:    swag.Int32(http.StatusNotFound),
			Message: swag.String("User not found"),
		})
	}

	// Update the user's details
	existingUser.Name = updatedUser.Name
	existingUser.Place = updatedUser.Place

	// Return success response with the updated user
	return NewUpdateUserOK().WithPayload(existingUser) // Use WithPayload to set the payload
}
