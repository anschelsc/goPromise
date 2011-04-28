//Copyright 2010 Anschel Schaffer-Cohen
//Under go's BSD-style license
//at $GOROOT/LICENSE

//The promise package implements Scheme-style
//delayed evaluation
package promise

import (
	"reflect"
	"sync"
)

//A Promise contains a function, which will be evaluated
//at most once, whenever Force() is called.
type Promise struct {
	f      interface{}
	args   []interface{}
	values []interface{}
	once   sync.Once
}

//Delay(f, args) takes a function and returns a Promise, which will
//evaluate f(args) if and when necessary. f WILL ONLY BE RUN ONCE,
//so take this into account if it has side-effects.
func Delay(f interface{}, args ...interface{}) *Promise {
	return &Promise{f: f, args: args}
}

//p.Force() evaluates p only if it has never been evaluated
//before--otherwise it returns the memoized result.
func (p *Promise) Force() []interface{} {
	p.once.Do(func() {
		fV := reflect.ValueOf(p.f)
		in := make([]reflect.Value, len(p.args))
		for i, v := range p.args {
			in[i] = reflect.ValueOf(v)
		}
		out := fV.Call(in)
		p.values = make([]interface{}, len(out))
		for i, v := range out {
			p.values[i] = v.Interface()
		}
	})
	return p.values
}
