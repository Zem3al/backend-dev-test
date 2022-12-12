package service

type Command interface {
	Valid() error
	Run() error
}
