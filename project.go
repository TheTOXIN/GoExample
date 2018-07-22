package main

import "fmt"

func main() {
	// --Hello, World!!!--
	fmt.Println("Hello, Go!!!")
	
	// --Varible--
	var age int = 19
	var token = 1234

	const pi float32 = 3.14
	var e = 2.7

	var login = "Toxin"
	var password string = "qwerty"

	var size = len(login)
	var isWeak bool = true

	fmt.Printf("%f \n", pi)
	fmt.Printf("%T \n", e)

	// --Conditions--
	if age >= 18 {
		fmt.Println("Hmmm")
	} else if age == 17 {
		fmt.Println("So close") 
	} else if (age < 16) && (age > 7) {
		fmt.Println("Go to school")
	} else {
		fmt.Println("Are you fucking kiding me&")
	}

	// --Cyicles--
	var i = 1
	for i <= 10 {
		fmt.Println(i)
		i++
	}

	for j := 10; j > 0; j-- {
		fmt.Println(j)
	} 

	// --Arrays--
	var arr[3] int
	arr[0] = 1
	arr[1] = 2
	arr[2] = 3

	fmt.Println(arr[1])

	nums := [3]float64 {3.33, 4.44, 5.55}
	for k, value := range nums {
		fmt.Println(value, k)
	}

	fmt.Println(token, password, size, isWeak)

	// --Maps--
	web := make(map[string]int)

	web["vk"] = 0
	web["youtube"] = 1
	web["google"] = 2

	fmt.Println(web["google"])

	// --Functions--
	var a = 2
	var b = 3

    var res = summ(a, b)
	fmt.Println(res)

	a, b = swap(a, b)
	fmt.Println(a, b)

	// --Zamikaniya???--
	var num = 3
	multi := func() int {
		num *= 2
		return num
	}
	fmt.Println(multi())

	// --Otklasivaniya???--
	defer two()
	one()

	// --Links--
	var x = 0
	pointer(&x)
	fmt.Println(x)

	// --Structures--
	bob := Cats{"Bob", 7, 0.87}
	fmt.Println("Cat name = " + bob.name)
	bob.toString()
}

type Cats struct {
	name string
	age int
	happiness float64
}

func (cat *Cats) toString() {
	fmt.Println("name = ", cat.name);
	fmt.Println("age = ", cat.age);
	fmt.Println("happiness = ", cat.happiness );
}

func summ(num_1 int, num_2 int) int {
	var res int
	res = num_1 + num_2
	return res
}

func swap(num_1 int, num_2 int) (int, int) {
	var tmp int
	tmp = num_1
	num_1 = num_2
	num_2 = tmp

	return num_1, num_2
}

func one() {
	fmt.Println("1")
}

func two() {
	fmt.Println("2")
}

func pointer(x *int) {
	*x = 2
}