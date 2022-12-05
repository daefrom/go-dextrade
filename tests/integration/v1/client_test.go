package tests

import (
	dextrade "github.com/daefrom/go-dextrade/v1"
)

var (
	client *dextrade.Client
)

func init() {
	key := "45d251b8ad4208988b8917596f3f5795255250b4e7a4e78839a4dfee7804670c"
	secret := "f1ada910b1ba656e0e5b5413508ddd866691408a97ac2bf5ac4fee9f57067e41"
	client = dextrade.NewClient().Auth(key, secret)
}
