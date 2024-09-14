package poller

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/benosborntech/feedme/cdc/consts"
	commonConsts "github.com/benosborntech/feedme/common/consts"
	"github.com/benosborntech/feedme/common/types"
	"github.com/benosborntech/feedme/common/utils"
	"github.com/bsm/redislock"
	"github.com/redis/go-redis/v9"
)

type Poller struct {
	db     *sql.DB
	client *redis.Client
}

func NewPoller(db *sql.DB, client *redis.Client) *Poller {
	return &Poller{
		db:     db,
		client: client,
	}
}

func (p *Poller) fetchAndPublish(ctx context.Context) error {
	locker := redislock.New(p.client)

	lockKey := utils.GenerateKey(commonConsts.POLLER_PREFIX, commonConsts.LOCK_PREFIX, "fetchandpublish")

	log.Printf("obtaining lock, lock=%v", lockKey)

	lock, err := locker.Obtain(ctx, lockKey, commonConsts.LOCK_DURATION, nil)
	if err != nil {
		return err
	}

	log.Printf("obtained lock, lock=%v", lockKey)

	defer lock.Release(ctx)

	// First we need to query the most recent counter
	counterKey := utils.GenerateKey(commonConsts.POLLER_PREFIX, "counter")

	ctr := 0

	res, err := p.client.Get(ctx, counterKey).Result()
	if err == nil {
		resCtr, err := strconv.Atoi(res)
		if err == nil {
			ctr = resCtr

			log.Printf("fetched counter, ctr=%v", ctr)
		}
	} else {
		log.Printf("failed to fetch counter")
	}

	prevTimestamp := time.Now().Add(-consts.DB_HISTORY_WINDOW)

	rows, err := p.db.QueryContext(ctx, "SELECT * FROM items WHERE id >= ? AND created_at > ? ORDER BY created_at DESC", ctr, prevTimestamp)
	if err != nil {
		return err
	}
	defer rows.Close()

	items := []types.Item{}

	for rows.Next() {
		var item types.Item
		if err := rows.Scan(&item.Id, &item.Location, &item.ItemType, &item.Quantity, &item.CreatedAt); err != nil {
			log.Printf("failed to parse item, err=%v", err)

			continue
		}
		items = append(items, item)
	}

	if len(items) == 0 {
		log.Printf("no items to retrieve")

		return nil
	}

	// Publish the items to the stream for all precision levels
	for _, item := range items {
		for i := 1; i < len(item.Location); i++ {
			key := item.Location[:i]
			body, err := json.Marshal(types.Event{
				Item: item,
			})
			if err != nil {
				log.Printf("failed to marshal body, err=%v", err)

				continue
			}

			if err := p.client.Publish(ctx, key, string(body)).Err(); err != nil {
				log.Printf("failed to publish key, key=%v", string(body))

				continue
			}
		}
	}

	// Set the new counter - we use >= 0 since we need to start with zero, so we need the + 1
	newCtr := items[0].Id + 1
	if err := p.client.Set(ctx, counterKey, fmt.Sprint(newCtr), 0).Err(); err != nil {
		return err
	}

	return nil
}

func (p *Poller) Poll(ctx context.Context) error {
	ticker := time.NewTicker(consts.POLL_FREQ)

	for {
		select {
		case <-ticker.C:
			err := p.fetchAndPublish(ctx)
			if err != nil {
				log.Printf("fetch and publish failed, err=%v", err)

				continue
			}

			log.Printf("published new data, err=%v", err)
		case <-ctx.Done():
			log.Printf("poller received shutdown signal")
		}
	}
}
