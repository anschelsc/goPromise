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
	f      *reflect.FuncValue
	args   []reflect.Value
	values []interface{}
	once   sync.Once
}

//Delay(f) takes a function and returns a Promise, which will
//evaluate f if and when necessary. f WILL ONLY BE RUN ONCE,
//so take this into account if it has side-effects.
func Delay(f interface{}, args ...interface{}) *Promise {
	p := new(Promise)
	p.f = reflect.NewValue(f).(*reflect.FuncValue)
	p.args = make([]reflect.Value, len(args))
	for i, arg := range args {
		p.args[i] = reflect.NewValue(arg)
	}
	p.values = make([]interface{}, p.f.Type().(*reflect.FuncType).NumOut())
	return p
}

//p.Force() evaluates p only if it has never been evaluated
//before--otherwise it returns the cached result.
func (p *Promise) Force() []interface{} {
	p.once.Do(func() {
		for i, v := range p.f.Call(args) {
			values[i] = v.Interface()
		}
	})
	return p.values
}
