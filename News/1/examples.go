package main

func main() {
	a:=make([]int,0)
	for i:=0; i<5;i++{
		a=append(a,i)
	}
	for i:=range a{
		if a[i]==3{
			a = append(a[:i], a[i+1:]...)
			return
		}
	}
}

