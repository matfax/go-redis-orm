package model

//! conf.redis
import (
	"fmt"
	"strings"

	"github.com/ezbuy/redis-orm/orm"
)

var (
	_redis_store *orm.RedisStore
)

const (
	PAIR = "pair"
	HASH = "hash"
	SET  = "set"
	ZSET = "zset"
	GEO  = "geo"
	LIST = "list"
)

type Object interface {
	GetClassName() string
	GetStoreType() string
	GetPrimaryName() string
	GetIndexes() []string
}

type RedisConfig struct {
	Host     string
	Port     int
	Password string
}

func RedisSetUp(cf *RedisConfig) {
	store, err := orm.NewRedisStore(cf.Host, cf.Port, cf.Password, 0)
	if err != nil {
		panic(err)
	}
	_redis_store = store
}

//! util functions
func keyOfObject(obj Object, keys ...string) string {
	suffix := strings.Join(keys, ":")
	if suffix != "" {
		return fmt.Sprintf("%s:%s:object:%s", obj.GetStoreType(), obj.GetClassName(), suffix)
	}
	return keyOfClass(obj)
}

func keyOfClass(obj Object, keys ...string) string {
	suffix := strings.Join(keys, ":")
	if suffix != "" {
		return fmt.Sprintf("%s:%s:%s", obj.GetStoreType(), obj.GetClassName(), suffix)
	}
	return fmt.Sprintf("%s:%s", obj.GetStoreType(), obj.GetClassName())
}

func pairOfClass(obj Object, keys ...string) string {
	suffix := strings.Join(keys, ":")
	if suffix != "" {
		return fmt.Sprintf("%s:%s:%s", PAIR, obj.GetClassName(), suffix)
	}
	return fmt.Sprintf("%s:%s", PAIR, obj.GetClassName())
}

func setOfClass(obj Object, keys ...string) string {
	suffix := strings.Join(keys, ":")
	if suffix != "" {
		return fmt.Sprintf("%s:%s:%s", SET, obj.GetClassName(), suffix)
	}
	return fmt.Sprintf("%s:%s", SET, obj.GetClassName())
}

func zsetOfClass(obj Object, keys ...string) string {
	suffix := strings.Join(keys, ":")
	if suffix != "" {
		return fmt.Sprintf("%s:%s:%s", ZSET, obj.GetClassName(), suffix)
	}
	return fmt.Sprintf("%s:%s", ZSET, obj.GetClassName())
}

func geoOfClass(obj Object, keys ...string) string {
	suffix := strings.Join(keys, ":")
	if suffix != "" {
		return fmt.Sprintf("%s:%s:%s", GEO, obj.GetClassName(), suffix)
	}
	return fmt.Sprintf("%s:%s", GEO, obj.GetClassName())
}

func listOfClass(obj Object, keys ...string) string {
	suffix := strings.Join(keys, ":")
	if suffix != "" {
		return fmt.Sprintf("%s:%s:%s", LIST, obj.GetClassName(), suffix)
	}
	return fmt.Sprintf("%s:%s", LIST, obj.GetClassName())
}
