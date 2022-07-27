package lsp

type TextDocumentProvider interface {
	DidOpen(document TextDocumentItem)
	DidClose(document TextDocumentIdentifier)
	DidChange(document VersionedTextDocumentIdentifier, changes []TextDocumentContentChangeEvent)
	WillSave(document TextDocumentIdentifier, reason TextDocumentSaveReason)
	DidSave(document TextDocumentIdentifier, text *string)
}

// Supporting types
const (
	Method_DidOpenTextDocument   string = "textDocument/didOpen"
	Method_DidChangeTextDocument string = "textDocument/didChange"
	Method_WillSaveTextDocument  string = "textDocument/willSave"
	Method_DidSaveTextDocument   string = "textDocument/didSave"
	Method_DidCloseTextDocument  string = "textDocument/didClose"
	Method_WillSaveWaitUntil     string = "textDocument/willSaveWaitUntil"
)

type DidOpenTextDocumentParams struct {
	TextDocument TextDocumentItem `json:"textDocument"`
}

type DidChangeTextDocumentParams struct {
	TextDocument   VersionedTextDocumentIdentifier  `json:"textDocument"`
	ContentChanges []TextDocumentContentChangeEvent `json:"contentChanges"`
}

type WillSaveTextDocumentParams struct {
	TextDocument TextDocumentIdentifier `json:"textDocument"`
	Reason       TextDocumentSaveReason `json:"reason"`
}

type DidSaveTextDocumentParams struct {
	TextDocument TextDocumentIdentifier `json:"textDocument"`
	Text         *string                `json:"text,omitempty"`
}

type DidCloseTextDocumentParams struct {
	TextDocument TextDocumentIdentifier `json:"textDocument"`
}

type TextDocumentClientCapabilities struct {
	Synchronization *TextDocumentSyncClientCapabilities `json:"synchronization,omitempty"`
	// Completion         *CompletionClientCapabilities               `json:"completion,omitempty"`
	// Hover              *HoverClientCapabilities                    `json:"hover,omitempty"`
	// SignatureHelp      *SignatureHelpClientCapabilities            `json:"signatureHelp,omitempty"`
	// Declaration        *DeclarationClientCapabilities              `json:"declaration,omitempty"`
	// Definition         *DefinitionClientCapabilities               `json:"definition,omitempty"`
	// TypeDefinition     *TypeDefinitionClientCapabilities           `json:"typeDefinition,omitempty"`
	// Implementation     *ImplementationClientCapabilities           `json:"implementation,omitempty"`
	// References         *ReferenceClientCapabilities                `json:"references,omitempty"`
	// DocumentHighlight  *DocumentHighlightClientCapabilities        `json:"documentHighlight,omitempty"`
	// DocumentSymbol     *DocumentSymbolClientCapabilities           `json:"documentSymbol,omitempty"`
	// CodeAction         *CodeActionClientCapabilities               `json:"codeAction,omitempty"`
	// CodeLens           *CodeLensClientCapabilities                 `json:"codeLens,omitempty"`
	// DocumentLink       *DocumentLinkClientCapabilities             `json:"documentLink,omitempty"`
	// ColorProvider      *DocumentColorClientCapabilities            `json:"colorProvider,omitempty"`
	// Formatting         *DocumentFormattingClientCapabilities       `json:"formatting,omitempty"`
	// RangeFormatting    *DocumentRangeFormattingClientCapabilities  `json:"rangeFormatting,omitempty"`
	// OnTypeFormatting   *DocumentOnTypeFormattingClientCapabilities `json:"onTypeFormatting,omitempty"`
	// Rename             *RenameClientCapabilities                   `json:"rename,omitempty"`
	// PublishDiagnostics *PublishDiagnosticsClientCapabilities       `json:"publishDiagnostics,omitempty"`
	// FoldingRange       *FoldingRangeClientCapabilities             `json:"foldingRange,omitempty"`
	// SelectionRange     *SelectionRangeClientCapabilities           `json:"selectionRange,omitempty"`
	// LinkedEditingRange *LinkedEditingRangeClientCapabilities       `json:"linkedEditingRange,omitempty"`
	// CallHierarchy      *CallHierarchyClientCapabilities            `json:"callHierarchy,omitempty"`
	// SemanticTokens     *SemanticTokensClientCapabilities           `json:"semanticTokens,omitempty"`
	// Moniker            *MonikerClientCapabilities                  `json:"moniker,omitempty"`
	// TypeHierarchy      *TypeHierarchyClientCapabilities            `json:"typeHierarchy,omitempty"`
	// InlineValue        *InlineValueClientCapabilities              `json:"inlineValue,omitempty"`
	// InlayHint          *InlayHintClientCapabilities                `json:"inlayHint,omitempty"`
	// Diagnostic *DiagnosticClientCapabilities `json:"diagnostic,omitempty"`
}

type TextDocumentSyncClientCapabilities struct {
	DynamicRegistration bool `json:"dynamicRegistration,omitempty"`
	WillSave            bool `json:"willSave,omitempty"`
	WillSaveWaitUntil   bool `json:"willSaveWaitUntil,omitempty"`
	DidSave             bool `json:"didSave,omitempty"`
}

type TextDocumentSyncOptions struct {
	OpenClose         bool                 `json:"openClose,omitempty"`
	Change            TextDocumentSyncKind `json:"change,omitempty"`
	WillSave          bool                 `json:"willSave,omitempty"`
	WillSaveWaitUntil bool                 `json:"willSaveWaitUntil,omitempty"`
	Save              *SaveOptions         `json:"save,omitempty"`
}

type SaveOptions struct {
	IncludeText bool `json:"includeText,omitempty"`
}

type TextDocumentSyncKind int

const (
	TextDocumentSyncKind_None        TextDocumentSyncKind = 0
	TextDocumentSyncKind_Full        TextDocumentSyncKind = 1
	TextDocumentSyncKind_Incremental TextDocumentSyncKind = 2
)

type TextDocumentItem struct {
	Uri        DocumentUri `json:"uri"`
	LanguageId string      `json:"languageId"`
	Version    int         `json:"version"`
	Text       string      `json:"text"`
}

type TextDocumentSaveReason uint

const (
	TextDocumentSaveReason_Manual     TextDocumentSaveReason = 1
	TextDocumentSaveReason_AfterDelay TextDocumentSaveReason = 2
	TextDocumentSaveReason_FocusOut   TextDocumentSaveReason = 3
)

type TextDocumentRegistrationOptions struct {
	DocumentSelector *DocumentSelector `json:"documentSelector"`
}

type TextDocumentSaveRegistrationOptions struct {
	TextDocumentRegistrationOptions
	SaveOptions
}

type TextDocumentChangeRegistrationOptions struct {
	TextDocumentRegistrationOptions

	SyncKind TextDocumentSyncKind `json:"syncKind"`
}

type TextDocumentContentChangeEvent struct {
	Range       *Range `json:"range,omitempty"`
	RangeLength uint   `json:"rangeLength,omitempty"`
	Text        string `json:"text"`
}

type TextDocumentEdit struct {
	TextDocument OptionalVersionedTextDocumentIdentifier `json:"textDocument"`
	Edits        []AnnotatedTextEdit                     `json:"edits"`
}

type TextDocumentPositionParams struct {
	TextDocument TextDocumentIdentifier `json:"textDocument"`
	Position     Position               `json:"position"`
}

type AnnotatedTextEdit struct {
	Range        Range                       `json:"range"`
	NewText      string                      `json:"newText"`
	AnnotationId *ChangeAnnotationIdentifier `json:"annotationId,omitempty"`
}

type TextEdit struct {
	Range   Range  `json:"range"`
	NewText string `json:"newText"`
}

type OptionalVersionedTextDocumentIdentifier struct {
	TextDocumentIdentifier
	Version *int `json:"version"`
}

type VersionedTextDocumentIdentifier struct {
	TextDocumentIdentifier

	Version int `json:"version"`
}

type TextDocumentIdentifier struct {
	Uri DocumentUri `json:"uri"`
}

type DocumentSelector []Choice2[string, DocumentFilter]

type DocumentFilter Choice2[TextDocumentFilter, NotebookCellTextDocumentFilter]

type NotebookCellTextDocumentFilter struct {
	Notebook Choice2[string, NotebookDocumentFilter] `json:"notebook"`
	Language *string                                 `json:"language,omitempty"`
}

type TextDocumentFilter struct {
	Language string `json:"language"`
	Scheme   string `json:"scheme,omitempty"`
	Pattern  string `json:"pattern,omitempty"`
}

type NotebookDocumentFilter struct {
	NotebookType string `json:"notebookType"`
	Scheme       string `json:"scheme,omitempty"`
	Pattern      string `json:"pattern,omitempty"`
}
