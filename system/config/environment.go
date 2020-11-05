package config

// Environment represent Environment config
type Environment struct {
	Stage              string   `json:"stage" yaml:"stage"`
	RunTestEnvironment []string `json:"runTestEnvironment" yaml:"runTestEnvironment"`
}
