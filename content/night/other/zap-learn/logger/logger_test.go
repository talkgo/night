package logger

import (
	"testing"
)

func TestFinfo(t *testing.T) {
	FInfo("~~~prefix", "key", "value", "key2", "value2", "key3", "value3", "key4", "value4")
}

func init() {
	//if err := config.InitConfig("../assets/config.yaml"); err != nil {
	//	panic(err)
	//}

	// init logger
	if err := Setup(); err != nil {
		panic(err)
	}
}
