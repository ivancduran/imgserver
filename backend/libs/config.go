package libs

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
)

type Conf struct {
	Devel  bool
	Domain string
	Host   string
	Api    string
	Fs     bool
	Db     Db
	Mail   MailConf
}

type Db struct {
	Host         string
	Database     string
	AuthDatabase string
	AuthUsername string
	AuthPassword string
}

type MailConf struct {
	Smtp      string
	FromAlias string
	User      string
	Password  string
	Port      int
}

func GetConf() Conf {

	// file, _ := os.Open("config.json")
	// fmt.Println(file)
	// decoder := json.NewDecoder(file)

	file, _ := ioutil.ReadFile("config.json")
	decoder := json.NewDecoder(bytes.NewReader(file))

	gC := Conf{}
	err := decoder.Decode(&gC)

	if err != nil {
		panic(err)
	}

	return gC
}

var cfg = GetConf()
