//go:generate go-extpoints . Provider
package providers

type Provider interface {
	// Setup
	Setup()
	// Hosts
	Query(namespace, query string) ([]Host, error)
}

type Host struct {
	Name  string
	Addr  string
	Index int
	Group int
}
