package _struct

import (
	"fmt"
	"time"
)

//type VCard struct {
//	Name       string
//	Addr       *Address
//	Birth      string
//	Image      string
//}
//
//type Address struct {
//	AddressNum string
//	Location   string
//}
//
//func PrintMyVCard() {
//	card := &VCard{
//		Name: "Joseph",
//		Addr: &Address{
//			AddressNum: "192",
//			Location: "弘康冻住",
//		},
//		Birth: "1994-06-13",
//		Image: "url",
//	}
//
//	fmt.Printf("mycard is %v", card)
//}

type Address struct {
	Street           string
	HouseNumber      uint32
	HouseNumberAddOn string
	POBox            string
	ZipCode          string
	City             string
	Country          string
}

type VCard struct {
	FirstName string
	LastName  string
	NickName  string
	BirtDate  time.Time
	Photo     string
	Addresses map[string]*Address
}

func PrintMyVCard() {
	addr1 := &Address{"Elfenstraat", 12, "", "", "2600", "Mechelen", "België"}
	addr2 := &Address{"Heideland", 28, "", "", "2640", "Mortsel", "België"}
	addrs := make(map[string]*Address)
	addrs["youth"] = addr1
	addrs["now"] = addr2
	birthdt := time.Date(1956, 1, 17, 15, 4, 5, 0, time.Local)
	photo := "MyDocuments/MyPhotos/photo1.jpg"
	vcard := &VCard{"Ivo", "Balbaert", "", birthdt, photo, addrs}
	fmt.Printf("Here is the full VCard: %v\n", vcard)
	fmt.Printf("My Addresses are:\n %v\n %v", addr1, addr2)
}
