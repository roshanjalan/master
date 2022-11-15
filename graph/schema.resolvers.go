package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"fmt"

	"github.com/Roshan/New/graph/generated"
	"github.com/Roshan/New/graph/model"
	"github.com/Roshan/New/test"
)

// CreateTodo is the resolver for the createTodo field.
func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	panic(fmt.Errorf("not implemented"))
}

// UpsertStudent is the resolver for the upsertStudent field.
func (r *mutationResolver) UpsertStudent(ctx context.Context, input *model.StudentInput) (*model.StudentResponse, error) {
	result, err := test.TableInsert(ctx, input)
	if err != nil {
		return nil, err
	}
	return result, err
}

// Todos is the resolver for the todos field.
func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	panic(fmt.Errorf("not implemented"))
}

// Addition is the resolver for the addition field.
func (r *queryResolver) Addition(ctx context.Context, input *model.AdditionInput) (*model.AdditionResponse, error) {
	var num1 int
	var num2 int
	var num3 int
	num1 = *input.Number1
	num2 = *input.Number2
	var mod model.AdditionResponse
	num3 = num1 + num2
	mod = model.AdditionResponse{Number3: (&num3)}
	return &mod, nil
}

// Subtraction is the resolver for the subtraction field.
func (r *queryResolver) Subtraction(ctx context.Context, input *model.SubInput) (*model.SubResponse, error) {
	var num1 int
	var num2 int
	var num3 int
	num1 = *input.Number1
	num2 = *input.Number2
	var mod1 model.SubResponse
	num3 = num1 - num2
	mod1 = model.SubResponse{Number3: (&num3)}
	return &mod1, nil
}

// Multiplication is the resolver for the multiplication field.
func (r *queryResolver) Multiplication(ctx context.Context, input *model.MulInput) (*model.MulResponse, error) {
	var num1 int
	var num2 int
	var num3 int
	num1 = *input.Number1
	num2 = *input.Number2
	var mod2 model.MulResponse
	num3 = num1 * num2
	mod2 = model.MulResponse{Number3: (&num3)}
	return &mod2, nil
}

// Division is the resolver for the division field.
func (r *queryResolver) Division(ctx context.Context, input *model.DivInput) (*model.DivResponse, error) {
	var num1 int
	var num2 int
	var num3 int
	num1 = *input.Number1
	num2 = *input.Number2
	var mod3 model.DivResponse
	num3 = num1 / num2
	mod3 = model.DivResponse{Number3: (&num3)}
	return &mod3, nil
}

// Calculation is the resolver for the calculation field.
func (r *queryResolver) Calculation(ctx context.Context, input *model.CalInput) (*model.CalResponse, error) {
	var num1 int
	var num2 int
	var result int
	//var operator string
	operator := *input.Operator
	num1 = *input.Number1
	num2 = *input.Number2
	var mod model.CalResponse
	if operator == "+" {
		result = num1 + num2
	}
	if operator == "-" {
		result = num1 - num2
	}
	if operator == "*" {
		result = num1 * num2
	}
	if operator == "/" {
		if num2 == 0 {
			return nil, errors.New("can't divide by 0")
		}
		result = num1 / num2
	}

	mod = model.CalResponse{Result: (&result)}
	return &mod, nil
}

// Datafetch is the resolver for the datafetch field.
func (r *queryResolver) Datafetch(ctx context.Context, input *model.StudentData) ([]*model.StudentOutput, error) {
	result, err := test.FetchData(ctx, input)
	if err != nil {
		return nil, err
	}
	return result, err
}

// ExcelData is the resolver for the excelData field.
func (r *queryResolver) ExcelData(ctx context.Context) (*model.ExcelOutput, error) {
	response, err := test.UploadExcel(ctx)
	if err != nil {
		return nil, err
	}
	return response, err
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
