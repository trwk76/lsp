package lsp

type Void struct{}
type Uri string
type DocumentUri string

type ChangeAnnotationIdentifier string

type Position struct {
	Line      uint `json:"line"`
	Character uint `json:"character"`
}

type Range struct {
	Start Position `json:"start"`
	End   Position `json:"end"`
}

type Choice2[OPT1 any, OPT2 any] struct {
	Opt1 *OPT1
	Opt2 *OPT2
}
