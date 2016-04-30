package kittysnap

import (
	"github.com/garyburd/redigo/redis"
)

type RedisUploadQueue struct {
	conn        redis.Conn
	keyCreated  string
	keyUploaded string
	keyDeleted  string
}

func OpenRedisUploadQueue() (*RedisUploadQueue, error) {
	c, err := redis.Dial("tcp", ":6379")
	if err != nil {
		return nil, err
	}

	return &RedisUploadQueue{
		conn:        c,
		keyCreated:  "created",
		keyUploaded: "uploaded",
		keyDeleted:  "deleted",
	}, nil
}

func (q *RedisUploadQueue) CloseRedisUploadQueue() {
	q.conn.Close()
}

func (q *RedisUploadQueue) AppendCreatedFile(path string) {
	q.conn.Send("RPUSH", q.keyCreated, path)
	q.conn.Flush()
}

func RedisTest() {

	// Take created file and move them to uploaded.

	// Take uploaded

	log.Println("Connected!")
}
