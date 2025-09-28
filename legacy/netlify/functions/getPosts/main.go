package main

import (
    "context"
    "encoding/json"
    "log"
    "runtime/debug"
    "time"

    "github.com/aws/aws-lambda-go/events"
    "github.com/aws/aws-lambda-go/lambda"
    "serverless/common"
    "go.mongodb.org/mongo-driver/bson"
)

func handler(ctx context.Context, req events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
    ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
    defer cancel()

    client, err := common.GetMongoClient(ctx)
    if err != nil {
        log.Println("mongo connect:", err)
        return &events.APIGatewayProxyResponse{StatusCode: 500, Body: "db connect error"}, nil
    }

    collection := client.Database("gw2style").Collection("gw2style")
    cursor, err := collection.Find(ctx, bson.M{})
    if err != nil {
        log.Println("find:", err)
        return &events.APIGatewayProxyResponse{StatusCode: 500, Body: "db find error"}, nil
    }
    defer cursor.Close(ctx)

    var posts []map[string]interface{}
    if err := cursor.All(ctx, &posts); err != nil {
        log.Println("cursor all:", err)
        return &events.APIGatewayProxyResponse{StatusCode: 500, Body: "decode error"}, nil
    }

    b, _ := json.Marshal(posts)
    return &events.APIGatewayProxyResponse{
        StatusCode: 200,
        Headers: map[string]string{"Content-Type": "application/json"},
        Body:       string(b),
    }, nil
}

func withRecover(h func(context.Context, events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error)) func(context.Context, events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
    return func(ctx context.Context, req events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
        defer func() {
            if r := recover(); r != nil {
                log.Printf("panic recovered: %v\n%s", r, debug.Stack())
            }
        }()
        return h(ctx, req)
    }
}

func main() {
    lambda.Start(withRecover(handler))
}