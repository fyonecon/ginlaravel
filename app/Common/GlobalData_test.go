package Common

import (
	"fmt"
	"testing"
)

func TestSetGlobalData(t *testing.T) {
	SetGlobalData("test", "2021")
}

func TestGetGlobalData(t *testing.T) {
	var test interface{} = GetGlobalData("test")
	fmt.Println(test)
}

func TestDelGlobalData(t *testing.T) {
	DelGlobalData("test")
}


