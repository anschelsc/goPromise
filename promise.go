//Copyright 2010 Anschel Schaffer-Cohen
//Under go's BSD-style license
//at $GOROOT/LICENSE

//The promise package implements Scheme-style
//delayed evaluation
package promise

type Promise struct {
	f     func() interface{}
	value interface{}
	done  bool
}

func Delay(f func() interface{}) *Promise {
	return &Promise{f: f}
}

func (p *Promise) Force() interface{} {
	if !p.done {
		p.value = p.f()
		p.done = true
	}
	return p.value
}
