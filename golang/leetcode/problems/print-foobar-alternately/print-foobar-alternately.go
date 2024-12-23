package print_foobar_alternately

// https://leetcode.com/problems/print-foobar-alternately/

type FooBar struct {
	n            int
	fooCh, barCh chan struct{}
}

func NewFooBar(n int) *FooBar {
	fooCh := make(chan struct{}, 1)
	barCh := make(chan struct{}, 1)
	barCh <- struct{}{}
	return &FooBar{n: n, fooCh: fooCh, barCh: barCh}
}

func (fb *FooBar) Foo(printFoo func()) {
	for i := 0; i < fb.n; i++ {
		<-fb.barCh
		printFoo()
		fb.fooCh <- struct{}{}
	}
	close(fb.fooCh)
}

func (fb *FooBar) Bar(printBar func()) {
	for i := 0; i < fb.n; i++ {
		<-fb.fooCh
		printBar()
		fb.barCh <- struct{}{}
	}
	close(fb.barCh)
}
