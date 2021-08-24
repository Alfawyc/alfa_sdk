package alfa_sdk

import (
	"alfa_sdk/sdkerr"
)

type Error struct {
	Code    string `json:"code"`
	Details string `json:"details"`
	Message string `json:"message"`
}

type ErrorList []Error

type ErrorResponse struct {
	Errors ErrorList `json:"errors,omitempty"`
}

func (e *Error) ToSdkError() error {
	return sdkerr.NewCustomErrorf("Api Error:", "Code = %s , Details = %s", e.Code, e.Details)
}
