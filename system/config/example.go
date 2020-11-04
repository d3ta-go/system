package config

// Config represent Config (Example)
type Config struct {
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
