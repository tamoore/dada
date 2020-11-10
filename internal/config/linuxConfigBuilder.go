package config

type linuxConfigBuilder struct {
	packageManager string
	dependencies   []string
}

func newLinuxConfigBuilder() *linuxConfigBuilder {
	builder := &linuxConfigBuilder{}
	builder.packageManager = "apt"
	builder.dependencies = make([]string, 0)
	return builder
}

func (b *linuxConfigBuilder) setDependency(dep string) {
	// TODO: Add logging here
	b.dependencies = append(b.dependencies, dep)
}

func (b *linuxConfigBuilder) getConfig() Product {
	return Product{
		PackageManager: b.packageManager,
		Dependencies:   b.dependencies,
	}
}
