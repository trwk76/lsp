package lsp

import (
	"context"
	"fmt"

	"github.com/trwk76/jsonrpc"
)

// Supporting types
const (
	InitializeMethod  string = "initialize"
	InitializedMethod string = "initialized"
	ShutdownMethod    string = "shutdown"
	ExitMethod        string = "exit"
)

type InitializeParams struct {
	WorkDoneProgressParams

	ProcessId             *int               `json:"processId"`
	ClientInfo            *ProgramInfo       `json:"clientInfo,omitempty"`
	Locale                *string            `json:"locale,omitempty"`
	RootPath              **string           `json:"rootPath,omitempty"`
	RootUri               *DocumentUri       `json:"rootUri"`
	Capabilities          ClientCapabilities `json:"capabilities"`
	InitializationOptions *interface{}       `json:"initializationOptions,omitempty"`
	Trace                 *InitialTraceValue `json:"trace,omitempty"`
	WorkspaceFolders      *[]WorkspaceFolder `json:"workspaceFolders,omitempty"`
}

type InitializeResult struct {
	Capabilities ServerCapabilities `json:"capabilities"`
	ServerInfo   *ProgramInfo       `json:"serverInfo,omitempty"`
}

type InitializedParams Void

type ShutdownParams Void
type ShutdownResult Void

type ExitParams Void

func addLifecycleMethods[Ctxt any](set *MethodSet[Ctxt]) {
	set.Add(NewRequest(InitializeMethod, ClientToServer, processInitialize[Ctxt]))
	set.Add(NewNotification(InitializedMethod, ClientToServer, processInitialized[Ctxt]))
	set.Add(NewRequest(ShutdownMethod, ClientToServer, processShutdown[Ctxt]))
	set.Add(NewNotification(ExitMethod, ClientToServer, processExit[Ctxt]))
}

func processInitialize[Ctxt any](ctx context.Context, srv *Server[Ctxt], port jsonrpc.Port, hdrs *jsonrpc.HeaderSet, id jsonrpc.RequestId, params InitializeParams) (*InitializeResult, error) {
	if srv.State() != ServerState_Uninitialized {
		return nil, jsonrpc.NewError(ErrorCode_RequestFailed, "Server is not uninitialized.", nil)
	}

	if params.ClientInfo != nil {
		srv.clientInfo.Name = params.ClientInfo.Name

		if params.ClientInfo.Version != nil {
			srv.clientInfo.Version = *params.ClientInfo.Version
		}
	}

	if params.ProcessId != nil {
		srv.clientInfo.ProcessId = *params.ProcessId
	}

	srv.clientInfo.Capabilites = params.Capabilities
	srv.setState(ServerState_Initializing)

	return &InitializeResult{}, nil
}

func processInitialized[Ctxt any](ctx context.Context, srv *Server[Ctxt], port jsonrpc.Port, hdrs *jsonrpc.HeaderSet, params InitializedParams) error {
	if srv.State() != ServerState_Initializing {
		return fmt.Errorf("server is not initializing")
	}

	srv.setState(ServerState_Initialized)

	if srv.OnInitialized != nil {
		srv.OnInitialized(srv)
	}

	return nil
}

func processShutdown[Ctxt any](ctx context.Context, srv *Server[Ctxt], port jsonrpc.Port, hdrs *jsonrpc.HeaderSet, id jsonrpc.RequestId, params ShutdownParams) (*ShutdownResult, error) {
	if err := srv.CheckInitialized(); err != nil {
		return nil, err
	}

	if srv.OnShutdown != nil {
		srv.OnShutdown(srv)
	}

	srv.setState(ServerState_Shutdown)
	return &ShutdownResult{}, nil
}

func processExit[Ctxt any](ctx context.Context, srv *Server[Ctxt], port jsonrpc.Port, hdrs *jsonrpc.HeaderSet, params ExitParams) error {
	if srv.State() != ServerState_Shutdown {
		srv.exitCode = 1
	} else {
		srv.exitCode = 0
	}

	port.Close()
	return nil
}

func (srv *Server[Ctxt]) CheckInitialized() error {
	if srv.State() != ServerState_Initialized {
		return jsonrpc.NewError(ErrorCode_ServerNotInitialized, "Server is not initialized.", nil)
	}

	return nil
}

type InitialTraceValue string

const (
	InitialTraceValue_Off      InitialTraceValue = "off"
	InitialTraceValue_Messages InitialTraceValue = "messages"
	InitialTraceValue_Compact  InitialTraceValue = "compact"
	InitialTraceValue_Verbose  InitialTraceValue = "verbose"
)

type ClientCapabilities struct {
	// Workspace        *WorkspaceClientCapabilities        `json:"workspace,omitempty"`
	TextDocument *TextDocumentClientCapabilities `json:"textDocument,omitempty"`
	// NotebookDocument *NotebookDocumentClientCapabilities `json:"notebookDocument,omitempty"`
	Window       *WindowClientCapabilities  `json:"window,omitempty"`
	General      *GeneralClientCapabilities `json:"general,omitempty"`
	Experimental *interface{}               `json:"experimental,omitempty"`
}

type GeneralClientCapabilities struct {
	StaleRequestSupport *struct {
		Cancel                 bool     `json:"cancel"`
		RetryOnContentModified []string `json:"retryOnContentModified"`
	} `json:"staleRequestSupport,omitempty"`
	RegularExpressions *RegularExpressionsClientCapabilities `json:"regularExpressions,omitempty"`
	Markdown           *MarkdownClientCapabilities           `json:"markdown,omitempty"`
	PositionEncodings  *[]PositionEncodingKind               `json:"positionEncodings,omitempty"`
}

type RegularExpressionsClientCapabilities struct {
	Engine  string  `json:"engine"`
	Version *string `json:"version,omitempty"`
}

type MarkdownClientCapabilities struct {
	Parser      string    `json:"parser"`
	Version     *string   `json:"version,omitempty"`
	AllowedTags *[]string `json:"allowedTags,omitempty"`
}

type PositionEncodingKind string

const (
	PositionEncodingKind_UTF8  PositionEncodingKind = "utf-8"
	PositionEncodingKind_UTF16 PositionEncodingKind = "utf-16"
	PositionEncodingKind_UTF32 PositionEncodingKind = "utf-32"
)

type ServerCapabilities struct {
	PositionEncoding *PositionEncodingKind    `json:"positionEncoding,omitempty"`
	TextDocumentSync *TextDocumentSyncOptions `json:"textDocumentSync,omitempty"`
	// NotebookDocumentSync             *NotebookDocumentSyncRegistrationOptions   `json:"notebookDocumentSync,omitempty"`
	// CompletionProvider               *CompletionOptions                                                               `json:"completionProvider,omitempty"`
	// HoverProvider                    *HoverOptions                                                     `json:"hoverProvider,omitempty"`
	// SignatureHelpProvider            *SignatureHelpOptions                                                            `json:"signatureHelpProvider,omitempty"`
	// DeclarationProvider *DeclarationRegistrationOptions `json:"declarationProvider,omitempty"`
	// DefinitionProvider               *DefinitionOptions                                                `json:"definitionProvider,omitempty"`
	// TypeDefinitionProvider           *TypeDefinitionRegistrationOptions         `json:"typeDefinitionProvider,omitempty"`
	// ImplementationProvider           *ImplementationRegistrationOptions         `json:"implementationProvider,omitempty"`
	// ReferencesProvider               *ReferenceOptions                                                 `json:"referencesProvider,omitempty"`
	// DocumentHighlightProvider        *DocumentHighlightOptions                                         `json:"documentHighlightProvider,omitempty"`
	// DocumentSymbolProvider           *DocumentSymbolOptions                                            `json:"documentSymbolProvider,omitempty"`
	// CodeActionProvider               *CodeActionOptions                                                `json:"codeActionProvider,omitempty"`
	// CodeLensProvider                 *CodeLensOptions                                                                 `json:"codeLensProvider,omitempty"`
	// DocumentLinkProvider             *DocumentLinkOptions                                                             `json:"documentLinkProvider,omitempty"`
	// ColorProvider                    *DocumentColorRegistrationOptions           `json:"colorProvider,omitempty"`
	// WorkspaceSymbolProvider          *WorkspaceSymbolOptions                                           `json:"workspaceSymbolProvider,omitempty"`
	// DocumentFormattingProvider       *DocumentFormattingOptions                                        `json:"documentFormattingProvider,omitempty"`
	// DocumentRangeFormattingProvider  *DocumentRangeFormattingOptions                                   `json:"documentRangeFormattingProvider,omitempty"`
	// DocumentOnTypeFormattingProvider *DocumentOnTypeFormattingOptions                                                 `json:"documentOnTypeFormattingProvider,omitempty"`
	// RenameProvider                   *RenameOptions                                                    `json:"renameProvider,omitempty"`
	// FoldingRangeProvider             *FoldingRangeRegistrationOptions             `json:"foldingRangeProvider,omitempty"`
	// SelectionRangeProvider           *SelectionRangeRegistrationOptions         `json:"selectionRangeProvider,omitempty"`
	// ExecuteCommandProvider           *ExecuteCommandOptions                                                           `json:"executeCommandProvider,omitempty"`
	// CallHierarchyProvider            *CallHierarchyRegistrationOptions           `json:"callHierarchyProvider,omitempty"`
	// LinkedEditingRangeProvider       *LinkedEditingRangeRegistrationOptions `json:"linkedEditingRangeProvider,omitempty"`
	// SemanticTokensProvider           *SemanticTokensRegistrationOptions               `json:"semanticTokensProvider,omitempty"`
	// MonikerProvider                  *MonikerRegistrationOptions                       `json:"monikerProvider,omitempty"`
	// TypeHierarchyProvider            *TypeHierarchyRegistrationOptions           `json:"typeHierarchyProvider,omitempty"`
	// InlineValueProvider              *InlineValueRegistrationOptions               `json:"inlineValueProvider,omitempty"`
	// InlayHintProvider                *InlayHintRegistrationOptions                   `json:"inlayHintProvider,omitempty"`
	// DiagnosticProvider               *DiagnosticRegistrationOptions                       `json:"diagnosticProvider,omitempty"`
	Workspace    *WorkspaceOptions `json:"workspace,omitempty"`
	Experimental *interface{}      `json:"experimental,omitempty"`
}

type ProgramInfo struct {
	Name    string  `json:"name"`
	Version *string `json:"version,omitempty"`
}
