package repository

import (
    "github.com/aws/aws-sdk-go/aws"
    "github.com/rohanbr/testresume/models"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/dynamodb"
    "github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
    "github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"


    "log"
)

type PersonalInfoRepo struct {
	Dynamodb dynamodbiface.DynamoDBAPI
    TableName string
}
 
func NewPersonalInfoRepo() *PersonalInfoRepo{
    sess := session.Must(session.NewSessionWithOptions(session.Options{
        SharedConfigState: session.SharedConfigEnable,
    }))
	return &PersonalInfoRepo{
		Dynamodb: dynamodb.New(sess),
        TableName: "PersonalInfo",
	}
}

func (personalInfoRepo *PersonalInfoRepo) Create(personalInfo models.PersonalInfo) bool{
    result, err := dynamodbattribute.MarshalMap(personalInfo)
    input := &dynamodb.PutItemInput{
        Item:      result,
        TableName: aws.String(personalInfoRepo.TableName),
    }
    _, err = personalInfoRepo.Dynamodb.PutItem(input)
    if err != nil {
        log.Fatalf("Got error calling PutItem: %s", err)
        return false
    } else {
        return true
    }
}
