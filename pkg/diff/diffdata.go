package diff

// Differ is an interface for the diff data producer
type Differ interface {
	Run(newsroomName string) error
}

// TODO: Implement a Differ
