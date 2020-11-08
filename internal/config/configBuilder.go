package config

type configType int

const (
	// OSX platform targets mac osx
	OSX configType = iota
	// Linux platform targets linux os
	Linux
)

// Builder used to create a configProduct platform
type Builder interface {
	setDependency(dep string)
	getConfig() Product
}
