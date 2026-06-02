package internal

type Npmrc struct {
	Path    string
	Content string
}

type NpmrcProfile struct {
	Name    string
	Path    string
	Active  bool
	Content string
}
