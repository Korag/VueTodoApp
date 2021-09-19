package repository

import (
	"errors"
	"math/rand"
)

type MemoryTodoRepositoryObject struct {
	todoElementsArray []TodoDbModel
}

func CreateInMemoryTodoRepository() ITodoRepository {
	return &MemoryTodoRepositoryObject{todoElementsArray: make([]TodoDbModel, 0)}
}

func (context *MemoryTodoRepositoryObject) GetAllTodoObjects() []TodoDbModel {
	return context.todoElementsArray
}

func (context *MemoryTodoRepositoryObject) AddNewTodoObject(text string, done bool) (*TodoDbModel, error) {
	todo := TodoDbModel{
		ID:   RandomString(10),
		Text: text,
		Done: done,
	}
	context.todoElementsArray = append(context.todoElementsArray, todo)
	return &todo, nil
}

func (context *MemoryTodoRepositoryObject) DeleteTodoObject(id string) (*TodoDbModel, error) {
	for i := 0; i < len(context.todoElementsArray); i++ {
		if context.todoElementsArray[i].ID == id {
			deletedTodo := context.todoElementsArray[i]
			context.todoElementsArray = RemoveIndex(context.todoElementsArray, i)
			return &deletedTodo, nil
		}
	}
	return nil, errors.New("Element of entered id not found in database.")
}

func (context *MemoryTodoRepositoryObject) UpdateTodoObject(id string, done bool) (*TodoDbModel, error) {
	for i := 0; i < len(context.todoElementsArray); i++ {
		if context.todoElementsArray[i].ID == id {
			context.todoElementsArray[i].Done = done
			updatedTodo := context.todoElementsArray[i]
			return &updatedTodo, nil
		}
	}
	return nil, errors.New("Element of entered id not found in database.")
}

func RemoveIndex(s []TodoDbModel, index int) []TodoDbModel {
	return append(s[:index], s[index+1:]...)
}

func RandomString(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	s := make([]rune, n)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}
	return string(s)
}
