package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var faultStr = []string{"oms127.0.0.1:9876,127.0.0.1:9875/bj01",
	"omsrocketmq://adminpwd@127.0.0.1:9876,127.0.0.1:9875/bj01",
	"oms:rocketmq://adminpwd127.0.0.1:9876,127.0.0.1:9875#bj01#ddds",
	"oms8888:rocketmq://admin#pwd127.0.0.1:9876,127.0.0.1:9875"}

var (
	Str       = "oms:rocketmq://admin#pwd@127.0.0.1:9876,127.0.0.1:9875/bj01"
	Driver    = "rocketmq"
	address   = "127.0.0.1:9876,127.0.0.1:9875"
	accessKey = "admin"
	region    = "bj01"
)

func TestAccessPointURI(t *testing.T) {
	for _, v := range faultStr {
		_, Err := AccessPointURI(v)
		assert.Errorf(t, Err, " ", Err)
	}
	data, err := AccessPointURI(Str)
	assert.NoError(t, err, err)
	assert.Equal(t, data.Prefix, Prefix, data.Prefix)
	assert.Equal(t, data.Driver, Driver, data.Driver)
	assert.Equal(t, data.Address, address, data.Address)
	assert.Equal(t, data.AccessKey, accessKey, data.AccessKey)
	assert.Equal(t, data.Region, region, data.Region)

}
