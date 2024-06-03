// Code generated by goa v3.16.2, DO NOT EDIT.
//
// account HTTP server types
//
// Command:
// $ goa gen flow-editor-server/design

package server

import (
	account "flow-editor-server/gen/account"
)

// GetAccountResponseBody is the type of the "account" service "getAccount"
// endpoint HTTP response body.
type GetAccountResponseBody struct {
	Activated  *bool `form:"activated,omitempty" json:"activated,omitempty" xml:"activated,omitempty"`
	Membership *int  `form:"membership,omitempty" json:"membership,omitempty" xml:"membership,omitempty"`
}

// NewGetAccountResponseBody builds the HTTP response body from the result of
// the "getAccount" endpoint of the "account" service.
func NewGetAccountResponseBody(res *account.AccountOutput) *GetAccountResponseBody {
	body := &GetAccountResponseBody{
		Activated:  res.Activated,
		Membership: res.Membership,
	}
	return body
}
