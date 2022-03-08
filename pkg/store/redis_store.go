package store

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/mhoc/urlshortener/pkg/config"
	"github.com/mhoc/urlshortener/pkg/util"
)

type Redis struct {
	client *redis.Client
}

func NewRedis(cfg *config.Config) Redis {
	client := redis.NewClient(&redis.Options{
		Addr: cfg.RedisURL,
	})
	return Redis{
		client: client,
	}
}

func (r Redis) Create(ctx context.Context, redirectToUrl string, expiresIn time.Duration) (string, error) {
	id := util.NewID()
	// Store the inverse relationship first; that a url points to a given id.
	// URLs can theoretically be pretty long, and redis recommends against excessively large key
	// sizes, but their technical max is 512mb so... could always hash this if it became an
	// issue.
	setUrlResult, err := r.client.SetNX(ctx, fmt.Sprintf("shortlink:url:%v", redirectToUrl), id, expiresIn).Result()
	if err != nil {
		return "", err
	}
	if !setUrlResult {
		// If that SETNX returns false, nothing was set, which means that URL is already in redis.
		// So, we retrieve the ID it corresponds to, and return that rather than the id we generated
		// above.
		storedId, err := r.client.Get(ctx, fmt.Sprintf("shortlink:url:%v", redirectToUrl)).Result()
		if err != nil {
			return "", err
		}
		return storedId, nil
	}
	// Otherwise, that was set, so set the primary relationship.
	err = r.client.SetNX(ctx, fmt.Sprintf("shortlink:id:%v", id), redirectToUrl, expiresIn).Err()
	if err != nil {
		return "", err
	}
	return id, nil
}

func (r Redis) Get(ctx context.Context, id string) (string, error) {
	result, err := r.client.Get(ctx, fmt.Sprintf("shortlink:id:%v", id)).Result()
	switch {
	case err == redis.Nil:
		return "", nil
	case err != nil:
		return "", err
	default:
		return result, nil
	}
}

func (r Redis) Remove(ctx context.Context, id string) (bool, error) {
	// In this bi-directional relationship setup, we unfortunately have to query for the URL in
	// order to know the key the url->id entry is stored under.
	url, err := r.Get(ctx, id)
	if err != nil {
		return false, err
	}
	err = r.client.Del(ctx,
		fmt.Sprintf("shortlink:id:%v", id),
		fmt.Sprintf("shortlink:url:%v", url),
	).Err()
	switch {
	case err == redis.Nil:
		return false, nil
	case err != nil:
		return false, err
	default:
		return true, nil
	}
}

func (r Redis) Stop() {
	r.client.Close()
}
