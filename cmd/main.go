package main

import (
	"context"

	"github.com/dyleme/smog/internal/gen"
	"github.com/dyleme/smog/pkg/drivers/mongo"
	"github.com/dyleme/smog/pkg/utils"
)

const uri = "mongodb://root:example@localhost:27017/"

func main() {
	ctx := context.Background()
	client, err := mongo.New(uri)
	if err != nil {
		panic(err)
	}
	defer client.Disconnect(ctx)
	db := client.Database("local")
	structure := db.Collection("structures")
	generator := gen.NewGenerator(gen.NewInterractor(structure))
	val, err := generator.Gen(ctx, "test")
	utils.NoErr(err)
	utils.Print(val)
}
