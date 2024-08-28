package gen

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type SchemaInterractor interface {
	Get(ctx context.Context, name string) (Schema, error)
}

type Interractor struct {
	schemas *mongo.Collection
}

func NewInterractor(schemas *mongo.Collection) *Interractor {
	return &Interractor{
		schemas: schemas,
	}
}

type dbSchema struct {
	Name   string `json:"name"`
	Schema Schema `json:"schema"`
}

func (i *Interractor) Get(ctx context.Context, name string) (Schema, error) {
	res := i.schemas.FindOne(ctx, bson.D{{"name", name}})
	if err := res.Err(); err != nil {
		return Schema{}, err
	}

	var dbSch dbSchema
	err := res.Decode(&dbSch)
	if err != nil {
		return Schema{}, err
	}

	return dbSch.Schema, nil
}

type Generator struct {
	schema SchemaInterractor
}

func NewGenerator(interractor SchemaInterractor) *Generator {
	return &Generator{
		schema: interractor,
	}
}

func (g *Generator) Gen(ctx context.Context, name string) (any, error) {
	schema, err := g.schema.Get(ctx, name)
	if err != nil {
		return nil, fmt.Errorf("get schema: %w", err)
	}

	res := genBySchema(schema)
	return res, nil
}
