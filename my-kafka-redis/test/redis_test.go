package test

import (
	"testing"
	"strings"
)

func Test_Redis_Get(t *testing.T) {
	redis := NewRedisClient()
	defer redis.Close()

	t.Log("[golang]: ", redis.Get("golang"))
}

func Test_Redis_Get_Pipeline(t *testing.T) {
	redis := NewRedisClient()
	defer redis.Close()
	pipeline := redis.Pipeline()
	defer pipeline.Close()

	v1 := pipeline.IncrByFloat("f_key", 0.1)
	v2 := pipeline.HIncrByFloat("f_h_key", "f_key", 0.01)
	v3 := pipeline.PFAdd("databases", "MySQL", "MongoDB", "Redis", "Oracle")
	v4 := pipeline.PFCount("databases")

	cmd, err := pipeline.Exec()
	if err == nil {
		for index := range cmd {
			r := cmd[index]
			v := strings.Split(r.String(), " ")
			t.Log(r.Name(), "\t|\t", v[len(v)-1], "\t|\t", r.Args(), "\t|\t", r.Err())
		}
	}

	t.Log(v1.Val(), "\t|\t", v2.Val(), "\t|\t", v3.Val(), "\t|\t", v4.Val())
}
