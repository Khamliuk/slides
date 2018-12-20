package main

import "log"

func main() {
	a:=make([]int,0)
	for i:=0; i<5;i++{
		a=append(a,i)
	}
	for i:=range a{
		if a[i]==3{
			log.Println(a)
			log.Println(a[i])
			a = append(a[:i], a[i+1:]...)
			log.Println(a)
			return
		}
	}

}

