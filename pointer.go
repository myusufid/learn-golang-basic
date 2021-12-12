package main

import "fmt"

type Address struct{
	City, Province, Country string
}

// pointer di function
func changeCountryToIndonesia(address *Address){
	address.Country  = "Indonesia"
}

func main()  {
	address1 := Address{
		City:     "Yogyakarta",
		Province: "Yogyakarta",
		Country:  "Indonesia",
	}
	address2 := &address1
	address3 := &address1

	address2.City = "Palembang"

	// tanda bintang, untuk mengganti semua value bukan hanya satu field saja
	*address2 = Address{
		City:     "Palembang",
		Province: "Sumatera Selatan",
		Country:  "Indonesia",
	}

	fmt.Println(address1)
	fmt.Println(address2)
	fmt.Println(address3)

	address4 := new(Address)
	address4.City = "Sekayu"
	fmt.Println(address4)
	
	var alamat = Address{
		City:     "Sekayu",
		Province: "Sumatera",
		Country:  "",
	}
	changeCountryToIndonesia(&alamat)
	fmt.Println(alamat)
}