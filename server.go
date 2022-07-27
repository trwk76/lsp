package lsp

import (
	"sync"

	"github.com/trwk76/jsonrpc"
)

type EventHandler[Ctxt any] func(srv *Server[Ctxt])

type Server[Ctxt any] struct {
	lock          sync.Mutex
	client        *jsonrpc.Client
	server        *jsonrpc.Server
	ctxt          Ctxt
	state         ServerState
	clientInfo    ClientInfo
	methods       *MethodSet[Ctxt]
	exitCode      int
	OnInitialized EventHandler[Ctxt]
	OnShutdown    EventHandler[Ctxt]
}

func NewServer[Ctxt any]() *Server[Ctxt] {
	return &Server[Ctxt]{
		state:   ServerState_Uninitialized,
		methods: NewStandardMethodSet[Ctxt](),
	}
}

func (s *Server[Ctxt]) Server() *jsonrpc.Server {
	return s.server
}

func (s *Server[Ctxt]) Client() *jsonrpc.Client {
	return s.client
}

func (s *Server[Ctxt]) State() ServerState {
	s.lock.Lock()
	defer s.lock.Unlock()

	return s.state
}

func (s *Server[Ctxt]) ClientInfo() ClientInfo {
	s.lock.Lock()
	defer s.lock.Unlock()

	return s.clientInfo
}

func (s *Server[Ctxt]) Context() Ctxt {
	return s.ctxt
}

func (s *Server[Ctxt]) ExitCode() int {
	return s.exitCode
}

func (s *Server[Ctxt]) setState(state ServerState) {
	s.lock.Lock()
	s.state = state
	s.lock.Unlock()
}

type ServerState uint8

const (
	ServerState_Uninitialized ServerState = 0
	ServerState_Initializing  ServerState = 1
	ServerState_Initialized   ServerState = 2
	ServerState_Shutdown      ServerState = 3
)

type ClientInfo struct {
	Name        string
	Version     string
	ProcessId   int
	Capabilites ClientCapabilities
}
