package repository

import (
	"context"
	"errors"

	"github.com/starrysilk/immortal/pkg/logger"
	"github.com/starrysilk/immortal/types/event"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (h *Handler) HandleEvent(e *event.Event) error {
	collName, _ := getCollectionName(e.Kind)
	coll := h.db.Client.Database(h.db.DBName).Collection(collName)

	ctx, cancel := context.WithTimeout(context.Background(), h.db.QueryTimeout)
	defer cancel()

	if e.Kind.IsRegular() {
		_, err := coll.InsertOne(ctx, e)
		if err != nil {
			return err
		}

		return nil
	}

	var filter bson.D

	if e.Kind.IsReplaceable() {
		filter = bson.D{
			bson.E{
				Key:   "pubkey",
				Value: e.PublicKey,
			},
			{
				Key:   "kind",
				Value: e.Kind,
			},
			{
				Key: "created_at",
				Value: bson.M{
					"$lte": e.CreatedAt,
				},
			},
		}
	}

	if e.Kind.IsAddressable() {
		var dTag string
		for _, t := range e.Tags {
			if len(t) < 2 {
				continue
			}

			if t[0] == "d" {
				dTag = t[1]

				break
			}
		}

		if dTag == "" {
			return errors.New("no d tag found")
		}

		filter = bson.D{
			bson.E{
				Key:   "pubkey",
				Value: e.PublicKey,
			},
			{
				Key:   "kind",
				Value: e.Kind,
			},
			{
				Key: "created_at",
				Value: bson.M{
					"$lte": e.CreatedAt,
				},
			},
			bson.E{
				Key: "$and",
				Value: bson.A{
					bson.M{
						"tags": bson.M{
							"$elemMatch": bson.M{
								"0": "d",
								"1": dTag,
							},
						},
					},
				},
			},
		}
	}

	opts := options.Replace().SetUpsert(true)
	_, err := coll.ReplaceOne(ctx, filter, e, opts)
	if err != nil {
		_, err := h.grpc.AddLog(context.Background(),
			"database error while adding new event", err.Error())
		if err != nil {
			logger.Error("can't send log to manager", "err", err)
		}

		return err
	}

	return nil
}
