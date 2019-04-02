package main

import "fmt"

//func main() {
//	a := time.Now().Unix()
//	b:=time.Unix(a,0)
//	fmt.Println(a)
//	fmt.Println(b)
//}

func main()  {
	m:=make(map[string]interface{})
	m["id"]= 666
	fmt.Println(m)
}
