package config

type osxConfigBuilder struct {
	packageManager string
	dependencies   []string
}

func newOSXConfigBuilder() *osxConfigBuilder {
	builder := &osxConfigBuilder{}
	builder.packageManager = "brew"
	builder.dependencies = make([]string, 0)
	return builder
}

func (b *osxConfigBuilder) setDependency(dep string) {
	// TODO: Add logging here
	b.dependencies = append(b.dependencies, dep)
}

func (b *osxConfigBuilder) getConfig() Product {
	return Product{
		packageManager: b.packageManager,
		dependencies:   b.dependencies,
	}
}
