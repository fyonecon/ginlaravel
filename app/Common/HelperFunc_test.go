package Common

import (
	"log"
	"testing"
)

func TestGetUrlParam(t *testing.T) {
	id := GetUrlParam("ginvel.com?date=2021&id=1949&name=牛逼", "id")
	if id != "1949"{
		t.Error("GetUrlParam()函数运行错误，test-id=", id)
	}else {
		log.Println("GetUrlParam()函数运行通过，test-id=", id)
		log.Println(ServerInfo)
	}
}