package main

import "fmt"


func main (){

a:=258

var c,b,f,d int

c=a%100
b=a%10
d=(c-b)/10
f=(a-c+b)/100

fmt.Print("10-masal chapdan birchi raqamni olish almashtrish = ")
fmt.Print(d)
fmt.Print(f)
fmt.Println(b,"\n")

fmt.Print("11-masal o'ngdan birchi raqamni olish almashtrish = ")

fmt.Print(f)
fmt.Print(b)
fmt.Println(d)


fmt.Print("12-masal o'ngdan birchi raqamni olish almashtrish = ")
fmt.Printf("%d%d%d\n",d,f,b)

//sovol yana bita ozgaruvchi ochib MISOL uchun t=d,f,c*/
}