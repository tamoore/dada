package config

// GetBuilder return a config builder
func GetBuilder(t configType) Builder {
	if t == OSX {
		return newOSXConfigBuilder()
	}

	if t == Linux {
		return newLinuxConfigBuilder()
	}

	return nil
}
