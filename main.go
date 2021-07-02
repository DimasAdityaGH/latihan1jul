package main

import "fmt"

func main () {
	//helloworld
	fmt.Println("hello, world!")
	//string
	var helo = "hai" //string menggunakan dua kutip  satu
	fmt.Println(helo)
	//variabel
	var example = "variabel menggunakan kata var lalu diikuti nama variabelnya"
	fmt.Println(example)
	//type declaration
	type str string
	type inr int8

	var nama str = "jordi"
	var nomor inr = 123
	fmt.Println(nama)
	fmt.Println(nomor)
	//type number
	var nilai int = 3264
	var nilaii32 int32 = 2147483647
	var nilaii64 int64 = 9223372036854775807
	fmt.Println(nilai)
	fmt.Println(nilaii32)
	fmt.Println(nilaii64)
	//type boolean
	var nikah = false
	fmt.Println(nikah)//boolean menghasilkan nilai true dan false

	//const
	const harga = 100
	fmt.Println(harga)

	const (
		jenisK = "laki"

	)
	fmt.Println(jenisK)
	//operasi matematika
	var i = 10
	i++
	fmt.Println(i)

	var j = 20
	j += 20
	fmt.Println(j)

	var l = 2
	l -= 1
	fmt.Println(l)

	var result = 10 + 1
	fmt.Println(result)
	//conversion
	var bahasa = "inggris"
	var b = bahasa[0]
	var bBahasa = string(b)
	fmt.Println(bahasa)
	fmt.Println(bBahasa)
	//perbandingan

	var x = "dimas"
	var y = "fifi"
	var z = x == y
	fmt.Println(z)
}