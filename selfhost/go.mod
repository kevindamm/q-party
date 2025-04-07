module github.com/kevindamm/q-party/selfhost

go 1.23.4

replace github.com/kevindamm/q-party/schema => ../schema

require (
	github.com/kevindamm/q-party/schema v0.0.0-00010101000000-000000000000
	golang.org/x/net v0.38.0
)
