package nosql

import (
	"context"
	"github.com/GeovaneCavalcante/api-anime-golang/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

const collectionName = "animes"

func (n *NoSqlRepo) Get(id string) (*entity.Anime, error){

	var a entity.Anime

	collectionAnime := n.db.Collection(collectionName)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	ID, _ := primitive.ObjectIDFromHex(id)
	err := collectionAnime.FindOne(ctx, bson.D{{"_id", ID}}).Decode(&a)

	if err != nil {
		return nil, err
	}

	return &a, err
}

