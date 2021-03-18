package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssm"
)

func main() {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
		Config: aws.Config{
			Region: aws.String("us-east-1"),
		},
	}))

	svc := ssm.New(sess)
	fmt.Println(svc)

	paramsMap := marshalMapParams()

	for key, value := range paramsMap {
		putParams(svc, key, value)
	}
}

func marshalMapParams() map[string]string {
	data, err := ioutil.ReadFile("../Project-8a8c500b8c6d.json")
	if err != nil {
		log.Fatal(err)
	}
	var paramsMap map[string]string
	json.Unmarshal(data, &paramsMap)

	return paramsMap
}

func putParams(svc *ssm.SSM, key string, value string) {
	ssmInput := &ssm.PutParameterInput{
		Name:  aws.String("/bqconfig/" + key),
		Value: aws.String(value),
		Type:  aws.String("String"),
	}

	_, err := svc.PutParameter(ssmInput)
	if err != nil {
		log.Fatal(err)
	}
}
