package main

import (
	"fmt"
	"time"

	"github.com/hailocab/gocassa"
)

type Sale struct {
	Id         string
	CustomerId string
	SellerId   string
	Price      int
	Created    time.Time
}

func main() {
	keySpace, err := gocassa.ConnectToKeySpace("test", []string{"127.0.0.1"}, "", "")
	if err != nil {
		panic(err)
	}
	salesTable := keySpace.Table("sale", &Sale{}, gocassa.Keys{
		PartitionKeys: []string{"Id"},
	})

	err = salesTable.CreateIfNotExist()
	if err != nil {
		panic(err)
	}

	err = salesTable.Set(Sale{
		Id:         "sale-1",
		CustomerId: "customer-1",
		SellerId:   "seller-1",
		Price:      42,
		Created:    time.Now(),
	}).Run()
	if err != nil {
		panic(err)
	}

	result := Sale{}
	if err := salesTable.Where(gocassa.Eq("Id", "sale-1")).ReadOne(&result).Run(); err != nil {
		panic(err)
	}
	fmt.Println(result)
}
