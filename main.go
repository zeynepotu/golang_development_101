package main

import (
	"fmt"
	"golesson/project"
)

func main() {

	//conditionals.Demo1()
	//enum.PrintBrand(enum.GOOGLE)
	//slices.Demo1()
	//numTerms, sum := functions.Topla(3, 5, 6)
	//fmt.Println("added:", numTerms, "terms for a total of", sum)
	//sayilar := []int{4, 5, 6}
	//sonuc := functions.ToplaVariadic(sayilar...)
	//fmt.Println(sonuc)
	//for_range.Demo2()

	//ciftSayiCn := make(chan int)
	//tekSayiCn := make(chan int)
	/*go channels.CiftSayilar(ciftSayiCn)
	  go channels.TekSayilar(tekSayiCn)

	  //ciftSayiToplam, tekSayiToplam := <-ciftSayiCn, <-tekSayiCn

	  //carpim := ciftSayiToplam * tekSayiToplam
	  //fmt.Println("Çarpım :", carpim)
	  //interfaces.Demo1()
	  //interfaces.Demo2()
	  string_functions.Demo1()
	  string_functions.Demo2()*/

	//array1 := []int{1, 2, 3, 4, 5}
	//array2 := []int{6, 7, 8, 9, 0}
	//result := test.Demo3(array1, array2)
	//fmt.Print(result)
	project.AddProduct()
	products, _ := project.GetAllProducts()

	for i := 0; i < len(products); i++ {
		fmt.Println(products[i].ProductName)
	}

}
