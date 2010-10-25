//Copyright 2010 Anschel Schaffer-Cohen
//Under go's BSD-style license
//at $GOROOT/LICENSE

//The promise package implements Scheme-style
//delayed evaluation
package promise

//A Promise contains a function, which will be evaluated
//at most once, whenever Force() is called.
type Promise struct {
	f     func() interface{}
	value interface{}
	done  bool
}

//Delay(f) takes a function and returns a Promise, which will
//evaluate f if and when necessary.
func Delay(f func() interface{}) *Promise {
	return &Promise{f: f}
}

//p.Force() evaluates p only if it has never been evaluated
//before--otherwise it returns the cached result.
func (p *Promise) Force() interface{} {
	if !p.done {
		p.value = p.f()
		p.done = true
	}
	return p.value
}
