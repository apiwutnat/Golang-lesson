package main

import "fmt"

func main() {
	//ตัวแปรตัวแรก สามารถใช้ : ได้
	// GO ไม่ชอบไห้ ประกาศตัวแปรที่ไม่ได้ใช้
	msg := "Hello"
	age := 55
	price, check := 22.52, true

	fmt.Println("msg", msg)
	fmt.Println("age", age)
	fmt.Println("price", price)
	fmt.Println("check", check)
}
