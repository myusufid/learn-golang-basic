package main

import "fmt"

func main()  {
	var months = [...] string{
		"Januari", // 0
		"Februai",// 1
		"Maret", // 2
		"April", // 3
		"Mei", // 4
		"Juni", // 5
		"Juli", // 6
		"Agustus", // 7
		"September", // 8
		"Oktober", // 9
		"November", // 10
		"Desember", // 11
	}

	var slice1 = months[4:7]
	fmt.Println(slice1)
	fmt.Println(len(slice1)) // 3 total data
	fmt.Println(cap(slice1)) // kapasitas, karna dimulai dari mei ke desember

	// slice diubah, maka array juga berubah
	slice1[0] = "MeiDiubah"
	fmt.Println(months)

	/**
		append : untuk membuat slice baru dengan menambahkan data ke posisi terakhir slice,
		jika kapasitas udah penuh, maka akan membuat array baru
	 */

	var slice2 = months[10:]
	fmt.Println(slice2)

	var slice3 = append(slice2, "Yusuf") // karna udah penuh di arrays months, maka jadi slice baru
	fmt.Println(slice3)
	slice3[1] = "DesemberEdit"
	fmt.Println(slice3)

	// disini tidak ada perubahan, masih data awal
	fmt.Println(slice2)
	fmt.Println(months)


	// kita coba ketika slice masih bisa menampung data

	var slice4 = months[2:4]
	fmt.Println(slice4)

	var slice5 = append(slice4, "Yusuf")
	fmt.Println(slice5)
	fmt.Println(months)


	// make slice with make function
	newSlice := make([]string, 2, 5)
	newSlice[0] = "make: M"
	newSlice[1] = "Yusuf"
	//newSlice[2] = "Coba" // ga bisa, karna len nya cuman 2

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	copySlice := make([]string, len(newSlice), cap(newSlice))
	copy(copySlice, newSlice)
	fmt.Println(copySlice)


	/**
		Array vs Slice
	 */

	iniArray := [...]int{1,2,3,4,5}
	iniSlice := []int{1,2,3,4,5}
	fmt.Println(iniArray)
	fmt.Println(iniSlice)

}
