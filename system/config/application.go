package config

// Applications represent Applications
type Applications struct {
	Servers Servers `json:"servers" yaml:"servers"`
}

// Servers represent all available application server
type Servers struct {
	RestAPI    RestAPI    `json:"restapi" yaml:"restapi"`
	GraphQLAPI GraphQLAPI `json:"graphqlAPI" yaml:"graphqlAPI"`
}

// RestAPI server
type RestAPI struct {
	Name        string      `json:"name" yaml:"name"`
	Version     string      `json:"version" yaml:"version"`
	Description string      `json:"description" yaml:"description"`
	Options     RESTOptions `json:"options" yaml:"options"`
}

// GraphQLAPI server
type GraphQLAPI struct {
	Name        string         `json:"name" yaml:"name"`
	Version     string         `json:"version" yaml:"version"`
	Description string         `json:"description" yaml:"description"`
	Options     GraphQLOptions `json:"options" yaml:"options"`
}

// RESTOptions represent RESTOptions
type RESTOptions struct {
	ShowEngineHeader  bool                  `json:"showEngineHeader" yaml:"showEngineHeader"`
	DisplayOpenAPI    bool                  `json:"displayOpenAPI" yaml:"displayOpenAPI"`
	Listener          Listener              `json:"listener" yaml:"listener"`
	Middlewares       Middlewares           `json:"middlewares" yaml:"middlewares"`
	OpenAPIDefinition RootOpenAPIDefinition `json:"openAPIDefinition" yaml:"openAPIDefinition"`
}

// GraphQLOptions represent GraphQLOptions
type GraphQLOptions struct {
	ShowEngineHeader bool           `json:"showEngineHeader" yaml:"showEngineHeader"`
	DisplaySWAPI     bool           `json:"displaySWAPI" yaml:"displaySWAPI"`
	Listener         Listener       `json:"listener" yaml:"listener"`
	Middlewares      Middlewares    `json:"middlewares" yaml:"middlewares"`
	SWAPIConfig      RootGraphQLAPI `json:"SWAPIConfig" yaml:"SWAPIConfig"`
}

// Listener represent Listener
type Listener struct {
	Port string `json:"port" yaml:"port"`
}

// Middlewares type
type Middlewares struct {
	Logger MiddlewareLogger `json:"logger" yaml:"logger"`
}

// MiddlewareLogger type
type MiddlewareLogger struct {
	Enable bool `json:"enable" yaml:"enable"`
}
