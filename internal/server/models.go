package server

type Metrics struct {
	Name  string
	Value string
}
type Context struct {
	Gauge   []Metrics
	Counter []Metrics
}
