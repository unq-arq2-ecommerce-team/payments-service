package application

type UseCase interface {
	Do(input interface{}) (interface{}, error)
}
