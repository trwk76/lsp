package lsp

import "github.com/trwk76/jsonrpc"

type PartialResult[PR any] struct {
	port  jsonrpc.Port
	hdrs  *jsonrpc.HeaderSet
	token ProgressToken
}

// Supporting types
type PartialParams interface {
	PartialToken() ProgressToken
}

func NewPartialResult[PR any](port jsonrpc.Port, token ProgressToken, hdrs *jsonrpc.HeaderSet) *PartialResult[PR] {
	return &PartialResult[PR]{
		port:  port,
		hdrs:  hdrs,
		token: token,
	}
}

func (r *PartialResult[PR]) Send(result PR) error {
	var pa *ProgressParams
	var err error

	if pa, err = newProgressParams(r.token, result); err != nil {
		return err
	}

	return jsonrpc.SendNotification(r.port, r.hdrs, ProgressMethod, *pa)
}
