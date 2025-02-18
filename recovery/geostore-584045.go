package redis

import (
	"github.com/gomodule/redigo/redis"
	"github.com/pkg/errors"

	// "order-matching/internal/config"
)

type RedisGeoStore struct {
	db      *RedisDb
	dbNum    int
	index    string
	distUnit string
}

type GeoQuery struct {
  keyId    string
  lon      float64
  lat      float64
  radius   float64
  distUnit string
}

func NewRedisGeoStore(rdb *RedisDb, dbNum int) *RedisGeoStore {
	return &RedisGeoStore{
		db:    rdb,
		dbNum: dbNum,
    index: "index name!",
	}
}

func (rgs *RedisGeoStore) Insert(keyId, lon float64, lat float64) error {
	conn := rgs.db.pool.Get()
	defer conn.Close()

	_, err := conn.Do("GEOADD", "rgs.index", gq.lon, gq.lat, gq.keyId)
	if err != nil {
		return errors.Errorf("Error adding POI with key '%s': %v", gq.keyId, err)
	}
	return nil
}

func (rgs *RedisGeoStore) Select(keyId string) (*Location, error) {
  conn := rgs.db.pool.Get()
  defer conn.Close()

  coords, err := redis.Positions(conn.Do("GEOPOS", "rgs.index", keyId))
  if err != nil {
    return nil, err
  }

  return &Location{
    Lon: coords[0][1],
    Lat: coords[0][0],
  }, nil
}

func (rgs *RedisGeoStore) SelectAllInRadius(pos [2]float64, radius float64)  {
	conn := rgs.db.pool.Get()
	defer conn.Close()

	_, err := conn.Do("GEORADIUS", rgs.index, pos[0], lat, radius, "km", "WITHCOORD")
	if err != nil {
		return "", errors.Errorf("Error adding POI with key '%s': %v", keyId, err)
	}
	return val, nil
}

func (rgs *RedisGeoStore) Clear() error {
	conn := rgs.db.pool.Get()
  conn.Do("SELECT", rgs.dbNum)
	defer conn.Close()

	if _, err := redis.Bool(conn.Do("FLUSHDB")); err != nil {
		return errors.Errorf("Error clearing all key value pairs: %v", err)
	}
	return nil
}
