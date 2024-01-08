package client

//import (
//	"context"
//	"fmt"
//	"github.com/basterrus/go_backend_framework/internal/config"
//	"go.mongodb.org/mongo-driver/mongo"
//	"go.mongodb.org/mongo-driver/mongo/options"
//)
//
//func NewMongoClient(ctx context.Context, cfg config.Config) (mongoClient *mongo.Database, err error) {
//	var anonymous bool
//	var mongoConnectionString string
//
//	if cfg.MongoDB.Username == "" && cfg.MongoDB.Password == "" {
//		anonymous = true
//		mongoConnectionString = fmt.Sprintf("mongodb://%s:%s", cfg.MongoDB.Host, cfg.MongoDB.Port)
//	} else {
//		mongoConnectionString = fmt.Sprintf("mongodb://%s:%s@%s:%s",
//			cfg.MongoDB.Username,
//			cfg.MongoDB.Password,
//			cfg.MongoDB.Host,
//			cfg.MongoDB.Port)
//	}
//
//	requestContext, cancel := context.WithTimeout(ctx, cfg.MongoDB.RequestContextTime)
//	defer cancel()
//
//	clientOptions := options.Client().ApplyURI(mongoConnectionString)
//	if !anonymous {
//		clientOptions.SetAuth(options.Credential{
//			AuthSource:  cfg.MongoDB.AuthDB,
//			Username:    cfg.MongoDB.Username,
//			Password:    cfg.MongoDB.Password,
//			PasswordSet: true,
//		})
//	}
//	client, err := mongo.Connect(requestContext, clientOptions)
//	if err != nil {
//		return nil, fmt.Errorf("failed to create client to mongodb due to error %w", err)
//	}
//
//	err = client.Ping(context.Background(), nil)
//	if err != nil {
//		return nil, fmt.Errorf("failed to create client to mongodb due to error %w", err)
//	}
//
//	return client.Database(cfg.MongoDB.Database), nil
//}
