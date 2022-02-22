package config

import "os"

var ACCESS_TOKEN string = "Bot " + os.Getenv("ACCESS_TOKEN")
var CLIENT_SECRET string = os.Getenv("CLIENT_SECRET")
var CLIENT_ID string = os.Getenv("CLIENT_ID")
