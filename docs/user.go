package docs

import (
	"xm/internal/entity/request"
	"xm/internal/entity/response"
)

// swagger:route POST /v1/user/register User idCreateUser
// Creates a User
//
// responses:
// 		200: CreateUserResponseWrapper
//		400: CustomErrorWrapper

// swagger:parameters idCreateUser
type CreateUserRequestWrapper struct {
	// in:body
	Body request.UserRequest
}

// swagger:response CreateUserResponseWrapper
type CreateUserResponseWrapper struct {
	// in:body
	Body response.UserResponse
}

// swagger:route POST /v1/user/login User idLoginUser
// Creates a User
//
// responses:
// 		200: LoginUserResponseWrapper
//		400: CustomErrorWrapper

// swagger:parameters idLoginUser
type LoginUserRequestWrapper struct {
	// in:body
	Body request.UserRequest
}

// swagger:response LoginUserResponseWrapper
type LoginUserResponseWrapper struct {
	// in:body
	Body response.UserResponse
}
