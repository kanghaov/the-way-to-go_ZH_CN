package main

import (
	"fmt"
	"time"
)

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

func main() {
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

/* Output:
Here is the full VCard: &{Ivo Balbaert  Sun Jan 17 15:04:05 +0000 1956 MyDocuments/MyPhotos/photo1.jpg map[now:0x126d57c0 youth:0x126d5500]}
My Addresses are:
 &{Elfenstraat 12   2600 Mechelen België}
 &{Heideland 28   2640 Mortsel België}
*/


//下面是我的作业

package main

import (
	"fmt"
)

type Address struct {
	addressCode int
}

type VCard struct {
	name string
	addressCode *Address
	birthday string
	image string
}

func main() {
	Nemo := &VCard{"Nemo",&Address{111},"1994-01-26","xxx"}
	fmt.Printf("Nemo's name:%s,addressCode:%d,birthday:%s,image:%s",Nemo.name,Nemo.addressCode.addressCode,Nemo.birthday,Nemo.image)
}

//或者
package main

import (
	"fmt"
	"time"
)

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
	Address   map[string]*Address
}

func NewAddress(name string, Street string, HouseNumber uint32, HouseNumberAddOn string, POBox string, ZipCode string, City string, Country string) *Address {
	name := Address{
		Street: Street,
		HouseNumber: HouseNumber,
		HouseNumberAddOn: HouseNumberAddOn,
		POBox: POBox,
		ZipCode: ZipCode,
		City:City,
		Country: Country
	}
	return &name
}



func main() {
	addr1 := NewAddress("youth","Elfenstraat", 12, "", "", "2600", "Mechelen", "België")
	addr2 := NewAddress{"now","Heideland", 28, "", "", "2640", "Mortsel", "België"}
	addrs := make(map[string]*Address, 2)
	addrs["youth"] = addr1
	addrs["now"] = addr2
	birthDate := time.Date(1956,1,17,15,4,5,0,time.Local) 
	photo := "MyDocuments/MyPhotos/photo1.jpg"
	vcard := &VCard{"Ivo", "Balbaert", "", birthdt, photo, addrs}
	fmt.Printf("Here is the full VCard: %v\n", vcard)
	fmt.Printf("My Addresses are:\n %v\n %v", addr1, addr2)
}
