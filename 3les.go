package main 

import ("fmt") 

func main() { 


//6 MASALA
a:=234

c:=a%100
b:=(a-c)/100

	

	fmt.Println("yuzlar xonasidagi son = ",b)


	//7 MASALA

	k:=463
	f:=k%100
	e:=k%10
	y:=(f-e)/10
	d:=0
	d=(k-e)/100


	fmt.Println("7 - MASALA")

	fmt.Println("birlilar xonasidagi son = ",e)
	fmt.Println("o'nliklar xonasidagi son = ",y)



     j:=e+y+d
     
	 fmt.Println("8- MASALA")

	 fmt.Println("uch xonali raqamlari yig'gindisi = ",j)


	fmt.Println("9 - MASALA")
	fmt.Println("o'nliklar xonasidagi son = ",e,y,d)
	fmt.Println("9 - MASALA")
	fmt.Print("uch xonlai soning teskari  soni = ")
	fmt.Print(e)
	fmt.Print(y)
	fmt.Println(d)
}