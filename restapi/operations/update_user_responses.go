package operations

import (
	"net/http"
	"swagger/models"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
)

// UpdateUserOKCode is the HTTP code returned for type UpdateUserOK
const UpdateUserOKCode int = 200

// UpdateUserNotFoundCode is the HTTP code returned for type UpdateUserNotFound
const UpdateUserNotFoundCode int = 404

/*
UpdateUserOK User updated successfully

swagger:response updateUserOK
*/
type UpdateUserOK struct {
	/*
	  In: Body
	*/
	Payload *models.User `json:"body,omitempty"`
}

// NewUpdateUserOK creates UpdateUserOK with default headers values
func NewUpdateUserOK() *UpdateUserOK {
	return &UpdateUserOK{}
}

// WithPayload adds the payload to the update user OK response
func (o *UpdateUserOK) WithPayload(payload *models.User) *UpdateUserOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update user OK response
func (o *UpdateUserOK) SetPayload(payload *models.User) {
	o.Payload = payload
}

// WriteResponse writes the response to the client
func (o *UpdateUserOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {
	rw.WriteHeader(UpdateUserOKCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// SuccessResponse constructs a middleware response with the updated user
func (o *UpdateUserOK) SuccessResponse(payload *models.User) middleware.Responder {
	return NewUpdateUserOK().WithPayload(payload)
}

/*
UpdateUserNotFound Error response for user not found

swagger:response updateUserNotFound
*/
type UpdateUserNotFound struct {
	/*
	  In: Body
	*/
	Payload *models.ErrorModel `json:"body,omitempty"`
}

// NewUpdateUserNotFound creates UpdateUserNotFound with default headers values
func NewUpdateUserNotFound() *UpdateUserNotFound {
	return &UpdateUserNotFound{}
}

// WithPayload adds the payload to the update user not found response
func (o *UpdateUserNotFound) WithPayload(payload *models.ErrorModel) *UpdateUserNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update user not found response
func (o *UpdateUserNotFound) SetPayload(payload *models.ErrorModel) {
	o.Payload = payload
}

// WriteResponse writes the response to the client
func (o *UpdateUserNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {
	rw.WriteHeader(UpdateUserNotFoundCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
