package listener

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"sync"

	"github.com/benosborntech/feedme/common/consts"
	"github.com/benosborntech/feedme/common/types"
	"github.com/benosborntech/feedme/common/utils"
	"github.com/redis/go-redis/v9"
)

type Listener struct {
	keys   sync.Map
	client *redis.Client
}

type keyChannel struct {
	counter  int
	channels map[int]chan *types.Event
	cancel   context.CancelFunc
	mutex    sync.RWMutex
}

func NewListener(client *redis.Client) *Listener {
	return &Listener{
		client: client,
		keys:   sync.Map{},
	}
}

// Will be used to listen for our events that get started from our various handlers and serve the requests
func (l *Listener) serveEvents(ctx context.Context, key string, keyChannel *keyChannel) {
	log.Printf("starting serve events, key=%v", key)

	redisKey := utils.GenerateKey(consts.CHANNEL_PREFIX, key)
	subscription := l.client.Subscribe(ctx, redisKey)

	ch := subscription.Channel()

	for {
		select {
		case msg := <-ch:
			func() {
				keyChannel.mutex.RLock()
				defer keyChannel.mutex.RUnlock()

				var event types.Event
				if err := json.Unmarshal([]byte(msg.String()), &event); err != nil {
					log.Printf("failed to parse event, err=%v", err)

					return
				}

				for _, channel := range keyChannel.channels {
					channel <- &event
				}
			}()
		case <-ctx.Done():
			log.Printf("serve events received exit signal, key=%v", key)

			return
		}
	}
}

func (l *Listener) Subscribe(key string) (int, chan *types.Event, error) {
	kChannel := &keyChannel{
		counter:  0,
		channels: map[int]chan *types.Event{},
	}
	var ok bool

	value, _ := l.keys.LoadOrStore(key, kChannel)
	kChannel, ok = value.(*keyChannel)
	if !ok {
		return 0, nil, fmt.Errorf("failed to convert value, value=%v", value)
	}

	kChannel.mutex.Lock()
	defer kChannel.mutex.Unlock()

	log.Printf("retrieved key channel, key=%v, kChannel=%v", key, kChannel)

	if len(kChannel.channels) == 0 {
		// Start the listener for the new events
		ctx, cancel := context.WithCancel(context.Background())
		kChannel.cancel = cancel

		go l.serveEvents(ctx, key, kChannel) // Execute this asynchronously as a new thread - it is OK to have recursive locking here
	}

	id := kChannel.counter
	ch := make(chan *types.Event)

	kChannel.channels[id] = ch
	kChannel.counter += 1

	return id, ch, nil
}

func (l *Listener) Unsubscribe(key string, id int) error {
	kChannel := &keyChannel{}
	var ok bool

	value, ok := l.keys.Load(key)
	if !ok {
		return fmt.Errorf("key does not exist")
	}

	kChannel, ok = value.(*keyChannel)
	if !ok {
		return fmt.Errorf("failed to convert value, value=%v", value)
	}

	kChannel.mutex.Lock()
	defer kChannel.mutex.Unlock()

	log.Printf("retrieved key channel, key=%v, kChannel=%v", key, kChannel)

	delete(kChannel.channels, id)

	if len(kChannel.channels) == 0 {
		kChannel.cancel()
	}

	return nil
}
