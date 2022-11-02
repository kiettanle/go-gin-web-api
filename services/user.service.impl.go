package services

import (
	"context"
	"errors"
	"fmt"
	"go-gin-web-api/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserServiceImpl struct {
	userCollection *mongo.Collection
	ctx            context.Context
}

func NewUserService(userCollection *mongo.Collection, ctx context.Context) UserService {
	return &UserServiceImpl{
		userCollection: userCollection,
		ctx:            ctx,
	}
}

func (u *UserServiceImpl) Create(user *models.User) (*models.User, error) {

	result, err := u.userCollection.InsertOne(u.ctx, user)

	fmt.Println(result.InsertedID)
	if err != nil {
		return nil, err
	}

	var createdUser *models.User
	query := bson.D{bson.E{Key: "_id", Value: result.InsertedID}}

	e := u.userCollection.FindOne(u.ctx, query).Decode(&createdUser)
	return createdUser, e
}

func (u *UserServiceImpl) Get(id *string) (*models.User, error) {
	var user *models.User
	objectId, err := primitive.ObjectIDFromHex(*id)
	if err != nil {
		return nil, err
	}

	query := bson.M{"_id": objectId}

	err = u.userCollection.FindOne(u.ctx, query).Decode(&user)
	return user, err
}

func (u *UserServiceImpl) GetAll() ([]*models.User, error) {
	var users []*models.User
	cursor, err := u.userCollection.Find(u.ctx, bson.D{{}})

	if err != nil {
		return nil, err
	}

	for cursor.Next(u.ctx) {
		var user *models.User
		err := cursor.Decode((&user))
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	cursor.Close((u.ctx))

	return users, nil
}

func (u *UserServiceImpl) Update(id *string, user *models.User) (*models.User, error) {
	objectId, err := primitive.ObjectIDFromHex(*id)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": objectId}
	query := bson.D{primitive.E{Key: "$set", Value: bson.D{primitive.E{Key: "name", Value: user.Name}, primitive.E{Key: "age", Value: user.Age}, primitive.E{Key: "address", Value: user.Address}}}}
	result, _ := u.userCollection.UpdateOne(u.ctx, filter, query)
	if result.MatchedCount != 1 {
		return nil, errors.New("User not found.")
	}

	return u.Get(id)
}

func (u *UserServiceImpl) Delete(id *string) error {
	objectId, err := primitive.ObjectIDFromHex(*id)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": objectId}
	result, _ := u.userCollection.DeleteOne(u.ctx, filter)
	if result.DeletedCount != 1 {
		return errors.New("User not found.")
	}
	return nil
}
