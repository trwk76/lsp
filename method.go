package lsp

import (
	"context"
	"encoding/json"
	"reflect"

	"github.com/trwk76/jsonrpc"
)

func NewStandardMethodSet[Ctxt any]() *MethodSet[Ctxt] {
	set := NewMethodSet[Ctxt]()

	addLifecycleMethods(set)

	return set
}

/**
 *	MethodDirection is an enumeration specifying if a method is meant to be sent from Client to Server or Server to Client.
 */
type MethodDirection uint8

const (
	ClientToServer MethodDirection = 1
	ServerToClient MethodDirection = 2
)

/**
 *	MethodDefinition interface describes a generic JSONRPC method.
 */
type MethodDefinition[Ctxt any] interface {
	Method() string
	Direction() MethodDirection
	ParamsType() reflect.Type
}

type RequestDefinition[Ctxt any] interface {
	MethodDefinition[Ctxt]

	ResultType() reflect.Type
	Process(context.Context, *Server[Ctxt], jsonrpc.Port, *jsonrpc.HeaderSet, jsonrpc.RequestId, json.RawMessage) (json.RawMessage, error)
}

type NotificationDefinition[Ctxt any] interface {
	MethodDefinition[Ctxt]

	Process(context.Context, *Server[Ctxt], jsonrpc.Port, *jsonrpc.HeaderSet, json.RawMessage) error
}

/**
 *	MethodSet struct holds a set of JSONRPC MethodDefinition
 */
type MethodSet[Ctxt any] struct {
	names map[string]MethodDefinition[Ctxt]
}

func NewMethodSet[Ctxt any]() *MethodSet[Ctxt] {
	return &MethodSet[Ctxt]{
		names: make(map[string]MethodDefinition[Ctxt]),
	}
}

func (s *MethodSet[Ctxt]) Get(method string) MethodDefinition[Ctxt] {
	return s.names[method]
}

func (s *MethodSet[Ctxt]) Add(definition MethodDefinition[Ctxt]) {
	s.names[definition.Method()] = definition
}

/**
 *	request struct holds the information of a JSONRPC request method
 */
type RequestHandler[Ctxt any] func(context.Context, *Server[Ctxt], jsonrpc.Port, *jsonrpc.HeaderSet, jsonrpc.RequestId, json.RawMessage) (json.RawMessage, error)

type request[Ctxt any] struct {
	method  string
	dir     MethodDirection
	params  reflect.Type
	result  reflect.Type
	process RequestHandler[Ctxt]
}

func NewRawRequest[Ctxt any](method string, dir MethodDirection, params reflect.Type, result reflect.Type, process RequestHandler[Ctxt]) RequestDefinition[Ctxt] {
	return request[Ctxt]{
		method:  method,
		dir:     dir,
		params:  params,
		result:  result,
		process: process,
	}
}

type TypedRequestHandler[Ctxt any, PA any, RE any] func(ctx context.Context, srv *Server[Ctxt], port jsonrpc.Port, hdrs *jsonrpc.HeaderSet, id jsonrpc.RequestId, params PA) (*RE, error)

func NewRequest[Ctxt any, PA any, RE any](method string, dir MethodDirection, process TypedRequestHandler[Ctxt, PA, RE]) RequestDefinition[Ctxt] {
	return NewRawRequest(
		method,
		dir,
		typeOf[PA](),
		typeOf[RE](),
		func(ctx context.Context, srv *Server[Ctxt], port jsonrpc.Port, hdrs *jsonrpc.HeaderSet, id jsonrpc.RequestId, params json.RawMessage) (json.RawMessage, error) {
			var par PA
			var res *RE
			var resdata json.RawMessage
			var err error

			if params != nil {
				if err = json.Unmarshal(params, &par); err != nil {
					return nil, jsonrpc.NewInvalidParamsError(nil)
				}
			}

			if res, err = process(ctx, srv, port, hdrs, id, par); err != nil {
				return nil, err
			}

			if res != nil {
				if resdata, err = json.Marshal(res); err != nil {
					return nil, err
				}
			}

			return resdata, nil
		},
	)
}

type TypedRequestWithPartialHandler[Ctxt any, PA PartialParams, RE any, PR any] func(ctx context.Context, srv *Server[Ctxt], port jsonrpc.Port, hdrs *jsonrpc.HeaderSet, id jsonrpc.RequestId, params PA, partial *PartialResult[PR]) (*RE, error)

func NewRequestWithPartial[Ctxt any, PA PartialParams, RE any, PR any](method string, dir MethodDirection, port jsonrpc.Port, hdrs *jsonrpc.HeaderSet, process TypedRequestWithPartialHandler[Ctxt, PA, RE, PR]) RequestDefinition[Ctxt] {
	return NewRawRequest(
		method,
		dir,
		typeOf[PA](),
		typeOf[RE](),
		func(ctx context.Context, srv *Server[Ctxt], port jsonrpc.Port, hdrs *jsonrpc.HeaderSet, id jsonrpc.RequestId, params json.RawMessage) (json.RawMessage, error) {
			var par PA
			var partial *PartialResult[PR] = nil
			var res *RE
			var resdata json.RawMessage
			var err error

			if params != nil {
				if err = json.Unmarshal(params, &par); err != nil {
					return nil, jsonrpc.NewInvalidParamsError(nil)
				}

				if par.PartialToken() != nil {
					partial = NewPartialResult[PR](port, par.PartialToken(), hdrs)
				}
			}

			if res, err = process(ctx, srv, port, hdrs, id, par, partial); err != nil {
				return nil, err
			}

			if res != nil {
				if resdata, err = json.Marshal(res); err != nil {
					return nil, err
				}
			}

			return resdata, nil
		},
	)
}

func (r request[Ctxt]) Method() string {
	return r.method
}

func (r request[Ctxt]) Direction() MethodDirection {
	return r.dir
}

func (r request[Ctxt]) ParamsType() reflect.Type {
	return r.params
}

func (r request[Ctxt]) ResultType() reflect.Type {
	return r.result
}

func (r request[Ctxt]) Process(ctx context.Context, srv *Server[Ctxt], port jsonrpc.Port, hdrs *jsonrpc.HeaderSet, id jsonrpc.RequestId, params json.RawMessage) (json.RawMessage, error) {
	return r.process(ctx, srv, port, hdrs, id, params)
}

/**
 *	notification struct holds the information of a JSONRPC request notification
 */
type NotificationHandler[Ctxt any] func(ctx context.Context, srv *Server[Ctxt], port jsonrpc.Port, hdrs *jsonrpc.HeaderSet, params json.RawMessage) error

type notification[Ctxt any] struct {
	method  string
	dir     MethodDirection
	params  reflect.Type
	process NotificationHandler[Ctxt]
}

func NewRawNotification[Ctxt any](method string, dir MethodDirection, params reflect.Type, process NotificationHandler[Ctxt]) NotificationDefinition[Ctxt] {
	return notification[Ctxt]{
		method:  method,
		dir:     dir,
		params:  params,
		process: process,
	}
}

type TypedNotificationHandler[Ctxt any, PA any] func(ctx context.Context, srv *Server[Ctxt], port jsonrpc.Port, hdrs *jsonrpc.HeaderSet, params PA) error

func NewNotification[Ctxt any, PA any](method string, dir MethodDirection, process TypedNotificationHandler[Ctxt, PA]) NotificationDefinition[Ctxt] {
	return NewRawNotification(
		method,
		dir,
		typeOf[PA](),
		func(ctx context.Context, srv *Server[Ctxt], port jsonrpc.Port, hdrs *jsonrpc.HeaderSet, params json.RawMessage) error {
			var par PA
			var err error

			if params != nil {
				if err = json.Unmarshal(params, &par); err != nil {
					return err
				}
			}

			if err = process(ctx, srv, port, hdrs, par); err != nil {
				return err
			}

			return nil
		},
	)
}

func (n notification[Ctxt]) Method() string {
	return n.method
}

func (n notification[Ctxt]) Direction() MethodDirection {
	return n.dir
}

func (n notification[Ctxt]) ParamsType() reflect.Type {
	return n.params
}

func (n notification[Ctxt]) Process(ctx context.Context, srv *Server[Ctxt], port jsonrpc.Port, hdrs *jsonrpc.HeaderSet, params json.RawMessage) error {
	return n.process(ctx, srv, port, hdrs, params)
}

func typeOf[T any]() reflect.Type {
	var tmp *T
	return reflect.TypeOf(tmp).Elem()
}
