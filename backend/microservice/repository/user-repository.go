package repository

import (
	"context"
	"direst/domain"

	"cloud.google.com/go/firestore"
)

//IUserRepository ..
type IUserRepository interface {
	FindByID(id string) (*firestore.DocumentSnapshot, error)
	Save(id string, data domain.User) (*firestore.WriteResult, error)
}

//UserRepository ..
type UserRepository struct {
	context           context.Context
	userCollectionRef *firestore.CollectionRef
}

//NewUserRepository ..
func NewUserRepository(ctx context.Context, fc *firestore.CollectionRef) *UserRepository {
	return &UserRepository{context: ctx, userCollectionRef: fc}
}

//FindByID ..
func (r *UserRepository) FindByID(id string) (*firestore.DocumentSnapshot, error) {
	return r.userCollectionRef.Doc(id).Get(r.context)
}

//Save ..
func (r *UserRepository) Save(id string, data domain.User) (*firestore.WriteResult, error) {
	return r.userCollectionRef.Doc(id).Set(r.context, data)
}
