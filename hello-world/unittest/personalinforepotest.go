package repository

import (
	"testing"

	"github.com/rohanbr/testresume/models"
	"github.com/rohanbr/testresume/repository"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type mockDynamoDB struct {
	dynamodbiface.DynamoDBAPI
}

func (m *mockDynamoDB) PutItem(input *dynamodb.PutItemInput) (*dynamodb.PutItemOutput, error) {
	return &dynamodb.PutItemOutput{}, nil
}

func TestCreate(t *testing.T) {
	
	repo := &repository.PersonalInfoRepo{
		Dynamodb: &mockDynamoDB{},
		TableName: "PersonalInfo",
	}

	personalInfo := models.PersonalInfo{
		Name: "John",
		Mobile:  "8197937705",
		Email:     "johndoe@example.com",
		Objective:     "1234567890",
	}

	result := repo.Create(personalInfo)
	if !result {
		t.Fatalf("Expected create to return true, but got false")
	}
}
