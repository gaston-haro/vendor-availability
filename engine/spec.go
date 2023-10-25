package engine

// Spec defines Engine interface, so it can be mocked
type Spec interface {
	GetInfo() map[string]interface{}
	SayHi(name string) string
}
