package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"TodoApp/graph/generated"
	"TodoApp/graph/model"
	"TodoApp/repository"
	"context"
)

var repositoryContext = repository.CreateInMemoryTodoRepository()

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.TodoAddNew) (*model.Todo, error) {
	done := false
	addedTodo, err := repositoryContext.AddNewTodoObject(input.Text, done)
	if addedTodo == nil {
		return nil, err
	} else {
		todo := model.Todo{
			ID:   addedTodo.ID,
			Text: addedTodo.Text,
			Done: addedTodo.Done}
		return &todo, nil
	}
}

func (r *mutationResolver) UpdateTodo(ctx context.Context, input model.TodoUpdateExisting) (*model.Todo, error) {
	updatedTodo, err := repositoryContext.UpdateTodoObject(input.ID, input.Done)
	if updatedTodo == nil {
		return nil, err
	} else {
		return &model.Todo{
			ID:   updatedTodo.ID,
			Text: updatedTodo.Text,
			Done: updatedTodo.Done}, nil
	}
}

func (r *mutationResolver) DeleteTodo(ctx context.Context, input model.TodoDeleteExisting) (*model.Todo, error) {
	deletedTodo, err := repositoryContext.DeleteTodoObject(input.ID)
	if deletedTodo == nil {
		return nil, err
	} else {
		return &model.Todo{
			ID:   deletedTodo.ID,
			Text: deletedTodo.Text,
			Done: deletedTodo.Done}, nil
	}
}

func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	data := repositoryContext.GetAllTodoObjects()
	todoModelList := make([]*model.Todo, 0)
	for i := 0; i < len(data); i++ {
		todoModelList = append(todoModelList, &model.Todo{
			ID:   data[i].ID,
			Text: data[i].Text,
			Done: data[i].Done})
	}
	return todoModelList, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
