// +build awesomebackend

package backend

type AwesomeBackend struct{}

func NewAwesomeBackend() Logger {
	return &AwesomeBackend{}
}

func (b *AwesomeBackend) MyMethod(s string) {
	//DO SOMETHING
}

func (b *AwesomeBackend) MyOtherMethod(n int) {
	//DO SOMETHING
}

func init() {
	New = NewAwesomeBackend
}
