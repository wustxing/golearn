package server

import (
	"fmt"
	"github.com/ankur-anand/simple-go-rpc/src/dataserial"
	"log"
	"reflect"
)

type RPCServer struct {
	addr  string
	funcs map[string]reflect.Value
}

func NewServer(addr string) *RPCServer {
	return &RPCServer{addr: addr, funcs: map[string]reflect.Value{}}
}

func (p *RPCServer) Register(fnName string, fFunc interface{}) {
	if _, ok := p.funcs[fnName]; ok {
		return
	}

	p.funcs[fnName] = reflect.ValueOf(fFunc)
}

func (p *RPCServer) Execute(req dataserial.RPCdata) dataserial.RPCdata {
	f, ok := p.funcs[req.Name]
	if !ok {
		e := fmt.Sprintf("func %s not Registered", req.Name)
		log.Println(e)
		return dataserial.RPCdata{Name: req.Name, Args: nil, Err: e}
	}
	log.Printf("func %s is called\n", req.Name)

	inArgs := make([]reflect.Value, len(req.Name))
	for i := range req.Args {
		inArgs[i] = reflect.ValueOf(req.Args[i])
	}

	out := f.Call(inArgs)
	resArgs := make([]interface{}, len(out)-1)
	for i := 0; i < len(out)-1; i++ {
		resArgs[i] = out[i].Interface()
	}

	var er string
	if _, ok := out[len(out)-1].Interface().(error); ok {
		er = out[len(out)-1].Interface().(error).Error()
	}
	return dataserial.RPCdata{Name: req.Name, Args: resArgs, Err: er}
}
