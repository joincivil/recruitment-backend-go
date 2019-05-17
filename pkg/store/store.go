package store

// Persister is an interface to the diff data store
type Persister interface {
	Load(id string) (interface{}, error)
	Save(id string, data interface{}) error
}

// TODO: Implement a data persister
