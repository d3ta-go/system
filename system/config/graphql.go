package config

// RootGraphQLAPI represent RootGraphQLAPI
type RootGraphQLAPI struct {
	Config GraphQLAPIConfig `json:"config"  yaml:"config"`
}

// GraphQLAPIConfig represent GraphQLAPIConfig
type GraphQLAPIConfig struct {
	AppTitle       string `json:"appTitle" yaml:"appTitle"`
	SchemaLocation string `json:"schemaLocation" yaml:"schemaLocation"`
}
