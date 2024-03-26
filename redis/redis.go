package redis

import (
	"context"
	"errors"
	"fmt"
	"github.com/Minggyyds/framework/config"
	"github.com/ghodss/yaml"
	"github.com/go-redis/redis/v8"
	"time"
)

func withClint(serviceName string, hand func(cli *redis.Client) error) error {
	content, err := config.GetConfig("DEFAULT_GROUP", serviceName)
	if err != nil {
		return err
	}
	type RedisConfig struct {
		Host string `json:"host" yaml:"host"`
		Port int    `json:"port" yaml:"port"`
	}
	var rediscfg struct {
		Redis RedisConfig `json:"Redis" yaml:"redis"`
	}
	err = yaml.Unmarshal([]byte(content), &rediscfg)
	if err != nil {
		return errors.New("转换为结构体格式失败redis" + err.Error())
	}
	cfg := rediscfg.Redis

	cli := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		DB:   0,
	})
	defer cli.Close()

	err = hand(cli)
	if err != nil {
		return err
	}

	return nil
}

func GetByKey(ctx context.Context, serviceName, key string) (string, error) {
	var data string
	var err error

	err = withClint(serviceName, func(cli *redis.Client) error {
		data, err = cli.Get(ctx, key).Result()
		return err
	})
	if err != nil {
		return "", err
	}
	return data, nil
}

func ExistKey(ctx context.Context, serviceName, key string) (bool, error) {
	var data int64
	var err error

	err = withClint(serviceName, func(cli *redis.Client) error {
		data, err = cli.Exists(ctx, key).Result()
		return err
	})
	if err != nil {
		return false, err
	}
	if data > 0 {
		return true, nil
	}
	return false, nil
}

func SetKey(ctx context.Context, serviceName, key string, val interface{}, duration time.Duration) error {
	return withClint(serviceName, func(cli *redis.Client) error {
		return cli.Set(ctx, key, val, duration).Err()
	})
}
