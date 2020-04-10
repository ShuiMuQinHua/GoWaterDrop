package main

type Person struct{
    firstName string
    lastName string
}

func (p Person)FirstName() string  {
    return p.firstName
}

func (p *Person) SetFirstName(newName string){
    p.firstName=newName
}

type B struct {
	thing int
}

func (b *B) change() {
	b.thing = 1
}

func (b B) write() string {
	return fmt.Sprint(b)
}


func main(){
	p1 := Person{}
	p2 := &Person{}

	/* p1 p2 可调用函数的差异*/
	fmt.Println(p1)
	fmt.Println(p2)


	var b1 B //b1是值
	b1.change()
	fmt.Println(b1.write())

	b2 := new(B) //b2是指针
	b2.change()
	fmt.Println(b2.write())
}
