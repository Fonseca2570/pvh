package globals

import (
	"github.com/BurntSushi/toml"
	"io/ioutil"
)

var App Application

func Init(app *Application) {
	conf := Config{}
	c, err := ioutil.ReadFile("conf.toml")
	if err != nil {
		panic(err)
	}


	if _, err = toml.Decode(string(c), &conf); err != nil {
		panic(err)
	}

	app.Config = conf
}