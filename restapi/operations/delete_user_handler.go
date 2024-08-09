package operations

import (
	"net/http"

	"swagger/models"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"
)

// DeleteUserHandlerImpl implementation for deleting a user
type DeleteUserHandlerImpl struct{}

// NewDeleteUserHandler creates a new handler for deleting a user
func NewDeleteUserHandler() DeleteUserHandler {
	return &DeleteUserHandlerImpl{}
}

// Handle the delete user request
func (h *DeleteUserHandlerImpl) Handle(params DeleteUserParams) middleware.Responder {
	// Adjust this to the correct field in your params
	userID := params.ID

	// Lock the store for writing
	userStore.Lock()
	defer userStore.Unlock()

	// Check if the user exists
	if _, exists := userStore.users[userID]; !exists {
		return NewDeleteUserDefault(http.StatusNotFound).WithPayload(&models.ErrorModel{
			Code:    swag.Int32(http.StatusNotFound),
			Message: swag.String("User not found"),
		})
	}

	// Delete the user
	delete(userStore.users, userID)

	// Return success response (adjust if custom success model isn't defined)
	return middleware.ResponderFunc(func(rw http.ResponseWriter, pr runtime.Producer) {
		rw.WriteHeader(http.StatusOK)
		rw.Write([]byte("User successfully deleted"))
	})
}
