package main

import (
	ab "addressbook"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/golang/protobuf/proto"
)

func write() {
	p1 := &ab.Person{
		Id:   1,
		Name: "小张",
		Phones: []*ab.Person_Phone{
			{"111-111-1111", ab.Person_MOBILE},
			{"111-111-1112", ab.Person_HOME},
		},
	}
	p2 := &ab.Person{
		Id:   2,
		Name: "小王",
		Phones: []*ab.Person_Phone{
			{"111-111-1113", ab.Person_HOME},
			{"111-111-1114", ab.Person_WORK},
		},
	}

	//创建地址簿
	book := &ab.AddressBook{}
	book.Persons = append(book.Persons, p1)
	book.Persons = append(book.Persons, p2)

	//编码数据
	data, _ := proto.Marshal(book)
	//把数据写入文件
	ioutil.WriteFile("./addressbook.txt", data, os.ModePerm)
}

func read() {
	//读取文件数据
	data, _ := ioutil.ReadFile("./addressbook.txt")
	book := &ab.AddressBook{}
	//解码数据
	proto.Unmarshal(data, book)
	for _, v := range book.Persons {
		fmt.Println(v.Id, v.Name)
		for _, vv := range v.Phones {
			fmt.Println(vv.Type, vv.Number)
		}
	}
}

func main() {
	write()
	read()
}
