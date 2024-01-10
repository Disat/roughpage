package model

var RootPath = "./"

type GlobalConfig struct {
	Site    siteConfig
	Server  serverConfig
	Publish publishConfig
}
type siteConfig struct {
	Title string
	logo  string
	Theme string
	Limit int
}

type serverConfig struct {
	Port int
}

type publishConfig struct {
	Output string
}