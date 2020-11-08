package config

// Director orchestrates configuration
type Director struct {
	builder Builder
}

// NewDirector returns a director for configuration
func NewDirector(b Builder) *Director {
	return &Director{
		builder: b,
	}
}

// MakeConfig exposes configProduct
func (d *Director) MakeConfig() Product {
	d.builder.setDependency("stow")
	d.builder.setDependency("git")
	d.builder.setDependency("shellcheck")
	d.builder.setDependency("socat")
	d.builder.setDependency("git-extras")
	d.builder.setDependency("libnotify-bin")
	return d.builder.getConfig()
}
