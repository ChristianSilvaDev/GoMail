package repository

import (
	"errors"
	"reflect"
	"time"

	"github.com/google/uuid"

	"github.com/ChristianSilvaDev/GoMail/src/internal/interfaces"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type BaseDynamoRepository[E interfaces.Entity] struct {
	db *dynamodb.DynamoDB
	TableName string
	NewEntity func() *E
}

func (r *BaseDynamoRepository[E]) Create(itemDao *interfaces.DAO) (E, error) {
	var zero E

	daoMap, err := dynamodbattribute.MarshalMap(itemDao)
	if err != nil {
		return zero, err
	}

	entityStruct := r.NewEntity()
	err = dynamodbattribute.UnmarshalMap(daoMap, &entityStruct)
	if err != nil {
		return zero, err
	}

	v := reflect.ValueOf(&entityStruct).Elem()
	v.FieldByName("ID").SetString(uuid.NewString())
	v.FieldByName("CreatedAt").SetString(time.Now().Format(time.RFC3339))

	itemMap, err := dynamodbattribute.MarshalMap(entityStruct)
	if err != nil {
		return zero, err
	}

	input := &dynamodb.PutItemInput{
		TableName: aws.String(r.TableName),
		Item:      itemMap,
	}

	_, err = r.db.PutItem(input)
	if err != nil {
		return zero, err
	}

	return *entityStruct, nil
}

func (r *BaseDynamoRepository[E]) Persist(entity E) error {
	return nil
}

func (r *BaseDynamoRepository[E]) FindById(id string) (E, error) {
	var zero E

	input := &dynamodb.GetItemInput{
		TableName: aws.String(r.TableName),
		Key: map[string]*dynamodb.AttributeValue{
			"id": {S: aws.String(id)},
		},
	}

	result, err := r.db.GetItem(input)
	if err != nil {
		return zero, err
	}

	if result.Item == nil {
		return zero, errors.New("entity not found")
	}
	var entityStruct E
	err = dynamodbattribute.UnmarshalMap(result.Item, &entityStruct)
	if err != nil {
		return zero, err
	}
	return entityStruct, nil

}
