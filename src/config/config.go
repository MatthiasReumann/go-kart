package config

type Endpoint struct {
	Path       string
	Dockerfile string
}

func NewEndpoint(path string, dockerfile string) Endpoint {
	return Endpoint{
		Path:       path,
		Dockerfile: dockerfile,
	}
}
