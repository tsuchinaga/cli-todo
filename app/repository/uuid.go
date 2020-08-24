package repository

type UUIDRepository interface {
	New() (string, error)
}
