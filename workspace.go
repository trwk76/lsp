package lsp

// Supporting types
type WorkspaceFolder struct {
	Uri  Uri    `json:"uri"`
	Name string `json:"name"`
}

type WorkspaceClientCapabilities struct {
	ApplyEdit bool `json:"applyEdit,omitempty"`
	// WorkspaceEdit          *WorkspaceEditClientCapabilities           `json:"workspaceEdit,omitempty"`
	// DidChangeConfiguration *DidChangeConfigurationClientCapabilities  `json:"didChangeConfiguration,omitempty"`
	// DidChangeWatchedFiles  *DidChangeWatchedFilesClientCapabilities   `json:"didChangeWatchedFiles,omitempty"`
	// Symbol                 *WorkspaceSymbolClientCapabilities         `json:"symbol,omitempty"`
	// ExecuteCommand         *ExecuteCommandClientCapabilities          `json:"executeCommand,omitempty"`
	WorkspaceFolders bool `json:"workspaceFolders,omitempty"`
	Configuration    bool `json:"configuration,omitempty"`
	// SemanticTokens         *SemanticTokensWorkspaceClientCapabilities `json:"semanticTokens,omitempty"`
	// CodeLens               *CodeLensWorkspaceClientCapabilities       `json:"codeLens,omitempty"`
	FileOperations *WorkspaceClientFileOperationCapabilities `json:"fileOperations,omitempty"`
	// InlineValue            *InlineValueWorkspaceClientCapabilities    `json:"inlineValue,omitempty"`
	// InlayHint              *InlayHintWorkspaceClientCapabilities      `json:"inlayHint,omitempty"`
	// Diagnostics *DiagnosticWorkspaceClientCapabilities `json:"diagnostics,omitempty"`
}

type WorkspaceClientFileOperationCapabilities struct {
	DynamicRegistration bool `json:"dynamicRegistration,omitempty"`
	DidCreate           bool `json:"didCreate,omitempty"`
	WillCreate          bool `json:"willCreate,omitempty"`
	DidRename           bool `json:"didRename,omitempty"`
	WillRename          bool `json:"willRename,omitempty"`
	DidDelete           bool `json:"didDelete,omitempty"`
	WillDelete          bool `json:"willDelete,omitempty"`
}
type WorkspaceOptions struct {
	WorkspaceFolders *WorkspaceFoldersServerCapabilities        `json:"workspaceFolders,omitempty"`
	FileOperations   *WorkspaceFileOperationsServerCapabilities `json:"fileOperations,omitempty"`
}

type WorkspaceFoldersServerCapabilities struct {
	Supported           bool   `json:"supported,omitempty"`
	ChangeNotifications string `json:"changeNotifications,omitempty"`
}

type WorkspaceFileOperationsServerCapabilities struct {
	DidCreate  *FileOperationRegistrationOptions `json:"didCreate,omitempty"`
	WillCreate *FileOperationRegistrationOptions `json:"willCreate,omitempty"`
	DidRename  *FileOperationRegistrationOptions `json:"didRename,omitempty"`
	WillRename *FileOperationRegistrationOptions `json:"willRename,omitempty"`
	DidDelete  *FileOperationRegistrationOptions `json:"didDelete,omitempty"`
	WillDelete *FileOperationRegistrationOptions `json:"willDelete,omitempty"`
}

type FileOperationRegistrationOptions struct {
	Filters []FileOperationFilter `json:"filters"`
}

type FileOperationFilter struct {
	Scheme  string               `json:"scheme,omitempty"`
	Pattern FileOperationPattern `json:"pattern"`
}

type FileOperationPattern struct {
	Glob    string                       `json:"glob"`
	Matches *FileOperationPatternKind    `json:"matches,omitempty"`
	Options *FileOperationPatternOptions `json:"options,omitempty"`
}

type FileOperationPatternKind string

const (
	FileOperationPatternKind_File   FileOperationPatternKind = "file"
	FileOperationPatternKind_Folder FileOperationPatternKind = "folder"
)

type FileOperationPatternOptions struct {
	IgnoreCase bool `json:"ignoreCase,omitempty"`
}
