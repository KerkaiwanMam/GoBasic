package main

import "fmt" //รวมเครื่องมือจัดรูปแบบ input และ output

func main() {
	math()
}

func math() {
	var num1, num2 = 10, 3
	de := 10 * 3
	fmt.Println("sum is ", num1+num2)
	fmt.Println("sub is ", num1-num2)
	fmt.Println("de is", de)
}

func tp() { //ทำงานโดยอัตโนมัติลำดับแรก
	fmt.Println("Hello Go Programming")

	name := "kerk"
	var age int = 25
	var score float32 = 98
	isPass := true
	//var isPass bool = false
	//fmt.Println("My name is ", name, age, score, isPass)
	fmt.Printf("Type name is %T\n", name) // แสดงผลชนิดของข้อมูล
	fmt.Printf("age = %T\n", age)
	fmt.Printf("score = %T\n", score)
	fmt.Printf("isPass = %T\n", isPass)

}

func con() { //นิยามค่าคงที่
	const name string = "kerkkai"

	fmt.Println("my name is ", name)
}
