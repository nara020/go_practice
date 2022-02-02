package main

import (
	"fmt"
	"strings"
)

type person struct {
	name    string
	age     int
	favFood []string
}

func multiply(a, b int) int {
	return a * b
}

func lenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

func lenAndUpper2(name string) (lenght int, upperCase string) {
	defer fmt.Println("I'm done") //함수가 끝나면 실행시킴
	lenght = len(name)
	upperCase = strings.ToUpper(name)
	return
}

func repeatMe(words ...string) {

}

func superAdd(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	for i := 0; i < len(numbers); i++ {
		fmt.Println(numbers[i], "포")
	}

	return total
}

func canIDrink(age int) bool {

	if koreanAge := age + 2; koreanAge <= 18 {
		return false
	}
	return true

}

func canIDrink2(age int) bool {
	switch koreanAge := age + 2; {
	case koreanAge <= 18:
		return false
	case koreanAge > 18:
		return true
	}
	return true
}

func main() {

	total := superAdd(1, 2, 3, 4, 5, 6)
	fmt.Println(total)
	fmt.Println(canIDrink(16))
	fmt.Println(canIDrink2(16))

	favFood := []string{"kimchi", "ramen"}
	jinStruct := person{name: "jinhyeok", age: 27, favFood: favFood}
	fmt.Println(jinStruct)

	//포인터
	a := 2
	b := &a
	a = 5
	fmt.Println(b, *b)
	*b = 20 //b가 보고있는 a의 값을 변환
	fmt.Println(a)

	//

	names := []string{"kkk", "jjj", "hhh"}
	// names[3] = "h"
	// names[4] = "i"
	//맵
	jinHyeok := map[string]string{"name": "jinHyeok", "age": "27"}
	fmt.Println(jinHyeok) //앞이 키, 뒤가 값

	for _, value := range jinHyeok {
		fmt.Print(value)
	}
	//

	fmt.Println(names)
	names = append(names, "m2m") //얘는 names에 바꾸는게 아닌, 하나 추가한 새로운 slice를 반환
	fmt.Println(names)

	const name string = " jh" //const : 상수
	var name2 string = " 진혁"
	name3 := " 진혁" //고가 알아서 타입을 찾아줌 변경은 안됨
	//하지만 이건 func 안에서만 쓸 수 있는 방법
	name2 = " x"

	fmt.Println(name + name2 + name3)

	totalLength, upperName := lenAndUpper("JinHyeok")
	fmt.Println(totalLength, upperName)

	totalLength2, upperName2 := lenAndUpper2("JinHyeokE")
	fmt.Println(totalLength2, upperName2)

	repeatMe("oh", "this", "is", "good")

}

//go 문서에 가면 지정할 수 있는 많은 타입이 설명되어 있다. (string부터 bool, 많은 범위의 int 등)
