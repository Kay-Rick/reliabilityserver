package test

import (
	"rs/consul"
	"testing"
)

func TestStatusCheck(t *testing.T) {
	client := consul.ConsulClientInit()
	consul.StatusCheck(client)
}
