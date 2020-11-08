package config

import "encoding/json"

// Config represent Config (Example)
type Config struct {
	Environment  Environment  `json:"environment" yaml:"environment"`
	Applications Applications `json:"applications" yaml:"applications"`
	IAM          IAM          `json:"IAM" yaml:"IAM"`
	Securities   Securities   `json:"securities" yaml:"securities"`
	DirLocations DirLocations `json:"dirLocations" yaml:"dirLocations"`
	Databases    Databases    `json:"databases" yaml:"databases"`
	Connectors   Connectors   `json:"connectors" yaml:"connectors"`
	SMTPServers  SMTPServers  `json:"SMTPServers" yaml:"SMTPServers"`
	Caches       Caches       `json:"caches" yaml:"caches"`
	Indexers     Indexers     `json:"indexers" yaml:"indexers"`
}

// ToJSON convert Config to JSON
func (c *Config) ToJSON() []byte {
	JSON, err := json.Marshal(c)
	if err != nil {
		return nil
	}
	return JSON
}

// CanRunTest can Run Test
func (c *Config) CanRunTest() bool {
	can := false
	for _, v := range c.Environment.RunTestEnvironment {
		if string(v) == c.Environment.Stage {
			can = true
			break
		}
	}
	return can
}
