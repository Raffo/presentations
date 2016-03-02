package backend

type Backend interface {
	MyMethod(s string)
	MyOtherMethod(n int)
}

type backendFactory func() Backend

var New backendFactory
