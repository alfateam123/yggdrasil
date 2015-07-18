package main_test

import (
	yggdrasil "github.com/alfateam123/yggdrasil"
	"os"
	"testing"
)

func TheOpening(filepath string) (conffile *os.File) {
	conffile, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}
	return conffile
}

func TestGetConfigDoesNotFailOnCorrectConf(test *testing.T) {
	conffile := TheOpening("./test_data/correct_conf.json")

	_, err := yggdrasil.GetConfig(conffile)
	//test.Log(conf); //golang pls go, you are drunk.
	if err != nil {
		test.Fail() //"GetConfig must not fail when the file exists and it's correct")
	}
}

func TestGetConfigFailsOnEmptyConf(test *testing.T) {
	conffile := TheOpening("./test_data/empty_conf.json")

	_, err := yggdrasil.GetConfig(conffile)
	//test.Log(conf);
	if err == nil {
		//test.Log(err);
		test.Fail()
	}
}

func TestGetConfigFailsOnNonCorrectConf(test *testing.T) {
	conffile := TheOpening("./test_data/notcorrect_conf.json")

	_, err := yggdrasil.GetConfig(conffile)
	if err == nil {
		test.Fail()
	}
}
