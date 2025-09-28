package main

import (
    "context"
    "encoding/json"
    "log"
    "time"

    "github.com/aws/aws-lambda-go/events"
    "github.com/aws/aws-lambda-go/lambda"
    "serverless/common"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "go.mongodb.org/mongo-driver/mongo"
)

type Post struct {
    Title     string `json:"title" bson:"title"`
    Thumbnail string `json:"thumbnail" bson:"thumbnail"`
    Views     int    `json:"views" bson:"views"`
}

func handler(ctx context.Context, req events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
    ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
    defer cancel()

    var p Post
    if err := json.Unmarshal([]byte(req.Body), &p); err != nil {
        log.Println("invalid body:", err)
        return apiError(400, "invalid JSON payload"), nil
    }

    if p.Title == "" {
        return apiError(400, "title is required"), nil
    }

    // default views
    p.Views = 0

    client, err := common.GetMongoClient(ctx)
    if err != nil {
        log.Println("mongo connect:", err)
        return apiError(500, "database connection error"), nil
    }

    collection := client.Database("gw2style").Collection("gw2style")
    res, err := collection.InsertOne(ctx, p)
    if err != nil {
        // if the client is in a bad state, log and return 500
        if err == mongo.ErrClientDisconnected {
            log.Println("mongo client disconnected:", err)
        }
        log.Println("insert error:", err)
        return apiError(500, "failed to create post"), nil
    }

    id := ""
    if oid, ok := res.InsertedID.(primitive.ObjectID); ok {
        id = oid.Hex()
    } else {
        // fallback to formatting whatever was returned
        id = stringifyInsertedID(res.InsertedID)
    }

    respBody, _ := json.Marshal(map[string]string{"message": "Post created successfully", "id": id})
    return &events.APIGatewayProxyResponse{
        StatusCode: 201,
        Headers: map[string]string{
            "Content-Type":                 "application/json",
            "Access-Control-Allow-Origin":  "*",
            "Access-Control-Allow-Methods": "POST, OPTIONS",
        },
        Body: string(respBody),
    }, nil
}

func stringifyInsertedID(id interface{}) string {
    b, _ := json.Marshal(id)
    return string(b)
}

func apiError(status int, msg string) *events.APIGatewayProxyResponse {
    body, _ := json.Marshal(map[string]string{"error": msg})
    return &events.APIGatewayProxyResponse{
        StatusCode: status,
        Headers: map[string]string{
            "Content-Type":                "application/json",
            "Access-Control-Allow-Origin": "*",
        },
        Body: string(body),
    }
}

func withRecover(h func(context.Context, events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error)) func(context.Context, events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
    return func(ctx context.Context, req events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
        defer func() {
            if r := recover(); r != nil {
                log.Printf("panic recovered: %v", r)
            }
        }()
        return h(ctx, req)
    }
}

func main() {
    lambda.Start(withRecover(handler))
}