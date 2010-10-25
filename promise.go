//Copyright 2010 Anschel Schaffer-Cohen
//Under go's BSD-style license
//at $GOROOT/LICENSE

//The promise package implements Scheme-style
//delayed evaluation
package promise

import "sync"

//A Promise contains a function, which will be evaluated
//at most once, whenever Force() is called.
type Promise struct {
	f     func() interface{}
	value interface{}
	once  sync.Once
}

//Delay(f) takes a function and returns a Promise, which will
//evaluate f if and when necessary. f WILL ONLY BE RUN ONCE,
//so take this into account if it has side-effects.
func Delay(f func() interface{}) *Promise {
	return &Promise{f: f}
}

//p.Force() evaluates p only if it has never been evaluated
//before--otherwise it returns the cached result.
func (p *Promise) Force() interface{} {
	p.once.Do(func() { p.value = p.f() })
	return p.value
}
