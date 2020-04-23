package repo

// $ mockery -name=Repo
type Repo interface {

	GetNames() ([]string, error)
	PutName(name string) (bool, error)
}