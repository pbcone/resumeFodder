package main

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

// FpModel a struct to represend an item in the FilmPac Dynamo Model table
type FpModel struct {
	ModelID     string `json:"model_id"`
	Frn         string `json:"frn"`
	Email       string `json:"email"`
	Gender      string `json:"gender"`
	Age         int    `json:"age"`
	Nationality string `json:"nationality"`
}

func addFpModelToDynamo(FpModel FpModel) {
	fmt.Println("Adding new FP Model to Database", FpModel)
	svc := dynamodb.New(sess)
	av, err := dynamodbattribute.MarshalMap(FpModel)
	dbInput := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(config.fpModelTable),
	}
	_, err = svc.PutItem(dbInput)

	if err != nil {
		fmt.Println("Got error adding FpModel to Dynamo")
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

func scanDynamoFpmodelTable() dynamodb.ScanOutput {
	fmt.Println("Scanning FP Model table")
	svc := dynamodb.New(sess)
	scanInput := &dynamodb.ScanInput{
		TableName: aws.String(config.fpModelTable),
	}
	result, err := svc.Scan(scanInput)
	if err != nil {
		fmt.Println("Error Performing FP Model Scan")
		fmt.Println(err.Error())
		os.Exit(1)
	}
	return *result
}
