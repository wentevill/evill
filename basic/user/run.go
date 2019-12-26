package main

import (
	"evill/einit"
	"fmt"
)

func main() {
	fmt.Println(einit.Log, einit.Mysql, einit.Redis, einit.Kafka, einit.Etcd)
}
