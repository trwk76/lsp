package lsp

import "encoding/json"

// Supporting types
const (
	ProgressMethod               string = "$/progress"
	WorkDoneProgressCreateMethod string = "window/workDoneProgress/create"
	WorkDoneProgressCancelMethod string = "window/workDoneProgress/cancel"
)

type ProgressToken interface{}

type ProgressParams struct {
	Token ProgressToken   `json:"token"`
	Value json.RawMessage `json:"value"`
}

func newProgressParams[VA any](token ProgressToken, value VA) (*ProgressParams, error) {
	var raw json.RawMessage
	var err error

	if raw, err = json.Marshal(&value); err != nil {
		return nil, err
	}

	return &ProgressParams{
		Token: token,
		Value: raw,
	}, nil
}

type WorkDoneProgressCreateParams struct {
	Token ProgressToken `json:"token"`
}

type WorkDoneProgressCreateResult Void

type WorkDoneProgressCancelParams struct {
	Token ProgressToken `json:"token"`
}

type PartialResultTokenParams interface {
	PartialResultToken() ProgressToken
}

type WorkDoneProgressOptions struct {
	WorkDoneProgress *bool `json:"workDoneProgress,omitempty"`
}

type WorkDoneProgressParams struct {
	WorkDoneToken *ProgressToken `json:"workDoneToken,omitempty"`
}

type WorkDoneProgress struct {
	Kind        WorkDoneProgressKind `json:"kind"`
	Cancellable *bool                `json:"cancellable,omitempty"`
	Message     *string              `json:"message,omitempty"`
	Percentage  *uint                `json:"percentage,omitempty"`
	Title       *string              `json:"title,omitempty"`
}

type WorkDoneProgressKind string

const (
	WorkDoneProgressKind_Begin  WorkDoneProgressKind = "begin"
	WorkDoneProgressKind_Report WorkDoneProgressKind = "report"
	WorkDoneProgressKind_End    WorkDoneProgressKind = "end"
)
