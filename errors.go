package lsp

import "github.com/trwk76/jsonrpc"

const (
	ErrorCode_ServerNotInitialized       jsonrpc.ErrorCode = -32002
	ErrorCode_UnknownErrorCode           jsonrpc.ErrorCode = -32001
	ErrorCode_LspReservedErrorRangeStart jsonrpc.ErrorCode = -32899
	ErrorCode_RequestFailed              jsonrpc.ErrorCode = -32803
	ErrorCode_ServerCancelled            jsonrpc.ErrorCode = -32802
	ErrorCode_ContentModified            jsonrpc.ErrorCode = -32801
	ErrorCode_RequestCancelled           jsonrpc.ErrorCode = -32800
)
