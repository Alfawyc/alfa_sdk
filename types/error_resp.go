package types

import "alfa_sdk/sdkerr"

type ErrResp struct {
	ErrorDescription string `json:"error_description"`
	Error            string `json:"error"`
}

func (r *ErrResp) ToSdkError() error {
	return sdkerr.NewCustomError(r.Error, r.ErrorDescription)
}
