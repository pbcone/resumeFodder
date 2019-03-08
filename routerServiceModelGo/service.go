package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"

	"github.com/google/uuid"
	"github.com/julienschmidt/httprouter"
)

//CreateFpModel POST /fpmodel/create/
func CreateFpModel(resp http.ResponseWriter, r *http.Request, data httprouter.Params) {
	start := time.Now()

	id := uuid.New()
	modelRecord := FpModel{}

	json.NewDecoder(r.Body).Decode(&modelRecord)
	modelRecord.Frn = id.String()
	addFpModelToDynamo(modelRecord)

	elapsed := time.Since(start)
	fmt.Println(`finished adding into Dynamo`)
	fmt.Fprint(resp, elapsed)
}

//GetAllFPModel GET /fpmodel/all
func GetAllFPModel(resp http.ResponseWriter, r *http.Request, data httprouter.Params) {
	dynamoResult := scanDynamoFpmodelTable()
	result := []FpModel{}
	resp.Header().Set(`Content-Type`, `application/json`)
	err := dynamodbattribute.UnmarshalListOfMaps(dynamoResult.Items, &result)
	if err != nil {
		fmt.Println("Error unmarshaling FpModel scan results", err.Error())
	}
	json.NewEncoder(resp).Encode(result)

	fmt.Println(`finished adding into Dynamo`)
}
