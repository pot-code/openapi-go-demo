// Code generated by goa v3.16.2, DO NOT EDIT.
//
// account HTTP server encoders and decoders
//
// Command:
// $ goa gen flow-editor-server/design

package server

import (
	"context"
	account "flow-editor-server/gen/account"
	"net/http"

	goahttp "goa.design/goa/v3/http"
)

// EncodeGetAccountResponse returns an encoder for responses returned by the
// account getAccount endpoint.
func EncodeGetAccountResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, any) error {
	return func(ctx context.Context, w http.ResponseWriter, v any) error {
		res, _ := v.(*account.AccountOutput)
		enc := encoder(ctx, w)
		body := NewGetAccountResponseBody(res)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}