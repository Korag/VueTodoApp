package repository

type ITodoRepository interface {
	GetAllTodoObjects() []TodoDbModel
	AddNewTodoObject(text string, done bool) (*TodoDbModel, error)
	UpdateTodoObject(id string, done bool) (*TodoDbModel, error)
	DeleteTodoObject(id string) (*TodoDbModel, error)
}
