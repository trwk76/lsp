package lsp

// Supporting types
type WindowClientCapabilities struct {
	WorkDoneProgress bool                                  `json:"workDoneProgress,omitempty"`
	ShowMessage      *ShowMessageRequestClientCapabilities `json:"showMessage,omitempty"`
	ShowDocument     *ShowDocumentClientCapabilities       `json:"showDocument,omitempty"`
}

type ShowMessageRequestClientCapabilities struct {
	MessageActionItem *ShowMessageRequestClientCapabilitiesMessageActionItem `json:"messageActionItem,omitempty"`
}

type ShowMessageRequestClientCapabilitiesMessageActionItem struct {
	AdditionalPropertiesSupport bool `json:"additionalPropertiesSupport,omitempty"`
}

type ShowDocumentClientCapabilities struct {
	Support bool `json:"support"`
}
