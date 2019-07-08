package main

import "fmt"

//Person is struct to store personal data//
type Person struct {
	Name     string
	Age      int
	Location string
}

func main() {
	/////Lesson 1 hellowolrd/////
	//fmt.Println("Hello World")

	/////Lesson 2 print function/////
	//fmt.Println(len("Hello World"))
	//fmt.Println("Hello Wolrd"[1])
	//fmt.Println("Hello" + "World")

	/////Lesson 3 variable/////
	///normal var///
	//var word1 = "Rawit"
	//var word2 = "Supathanakorn"
	///constant var///
	//const num1 = 22
	//fmt.Println("Name: " + word1 + " " + word2)

	/////Lesson 4 loop/////
	///format 1///
	//i := 1
	//for i <= 10 {
	//	fmt.Println(i)
	//	i++
	//}
	///format 2///
	//for i := 11; i <= 20; i++ {
	//	fmt.Println(i)
	//}

	/////Lesson 5 if-else/////
	// for i := 1; i <= 100; i++ {
	// 	if i % 2 == 0 {
	// 		fmt.Println(i, " is even")
	// 	} else {
	// 		fmt.Println(i, " is odd")
	// 	}
	// }

	/////Lesson 6 switch case/////
	// for i := 1; i <= 5; i++ {
	// 	switch i {
	// 	case 1:
	// 		fmt.Println("One")
	// 	case 2:
	// 		fmt.Println("Two")
	// 	case 3:
	// 		fmt.Println("Three")
	// 	case 4:
	// 		fmt.Println("Four")
	// 	case 5:
	// 		fmt.Println("Five")
	// 	default:
	// 		fmt.Println("Unknow number")

	// 	}
	// }

	/////Lesson 7 array/////

	// var sumScore float32
	// var avgScore float32
	// var score [5]float32
	///Create Array shortcut///
	// score := [5]float32{98, 93, 77, 82, 83}
	///or///
	// score := [5]float32{98,
	// 	93,
	// 	77,
	// 	82,
	// 	83}

	// score[0] = 98
	// score[1] = 93
	// score[2] = 77
	// score[3] = 82
	// score[4] = 83

	///format 1///
	// for i := 0; i < len(score); i++ {
	// 	sumScore += score[i]
	// }
	// avgScore = sumScore / float32(len(score))
	// fmt.Println("Average score = ", avgScore)

	///format 2///
	// for _, value := range score {
	// 	sumScore += value
	// }
	// avgScore = sumScore / float32(len(score))
	// fmt.Println("Average score = ", avgScore)

	/////Lesson 8 slice/////
	///create slice///
	// slice := make([]float32, 5)
	// array := [5]float32{1, 2, 3, 4, 5}
	///select all of array value///
	// slice = array[0:]
	///select 1 - 3 of array value///
	// slice = array[0:3]
	// fmt.Println(slice)

	///using function of slice///
	// slice1 := []int{1, 2, 3}
	///add 10 20 in back of slice 1 array///
	// slice2 := append(slice1, 10, 20)
	// fmt.Println(slice1, slice2)

	// slice3 := make([]int, 5)
	///copy data from slice2 into slice3///
	// copy(slice3, slice2)
	// fmt.Println(slice3)

	/////Lesson 9 map/////
	// nameList := make(map[int]string)
	// nameList[1] = "Ham"
	// nameList[2] = "Bas"
	// nameList[3] = "Bank"
	// nameList[4] = "Pree"
	// nameList[5] = "kee"
	// nameList[6] = "Touch"

	// fmt.Println(nameList)

	///can create nested map///
	// teamData := map[string]map[string]string{
	// 	"team1": map[string]string{
	// 		"teamname": "VNP",
	// 		"leader":   "Rawit",
	// 	},
	// }
	// fmt.Println(teamData["team1"])

	///Find min and max number on array///
	// x := []int{
	// 	48, 96, 86, 68,
	// 	57, 82, 63, 70,
	// 	37, 34, 83, 27,
	// 	19, 97, 9, 17,
	// }
	///set for avoid value empty slice is 0 not corrct///
	// minValue := x[0]
	// maxValue := x[0]

	// for _, e := range x {
	// 	if e < minValue {
	// 		minValue = e
	// 	} else if e > maxValue {
	// 		maxValue = e
	// 	}
	// }

	// fmt.Println("Min Value is : ", minValue)
	// fmt.Println("Max Value is : ", maxValue)

	///call function in lesson 10///
	// fmt.Println(f())
	// fmt.Println(r())
	// i(20)
	// t(1, 2, 3)
	// a(1, 2, 3, 4, 5, 6, 7, 8, 9)

	/////Lesson 11 closure/////
	// localVar := 0
	// add := func(x, y int) int {
	// 	sum := x + y
	// 	localVar = sum
	// 	return sum
	// }
	// fmt.Println(add(1, 6))
	///Can using variable in same function///
	// fmt.Println(localVar)

	///call function in lesson 12///
	// fmt.Println(plus(50))

	///call function in lesson 13///
	///if function have defer they with work at the last of main work///
	// defer two()
	// one()

	/////Lesson 14 panic and recover/////
	///if not have recover when program meet panic program with exit///
	// 	defer func() {
	// 		str := recover()
	// 		fmt.Println(str)
	// 	}()
	// 	panic("Something went wrong!!")

	///call function from pratice///
	// ar := []int{2, 4, 6, 8, 10}
	// fmt.Println(signature(ar))

	// fmt.Println(half(1))

	// fmt.Println(highestFunc(1, 2, 3, 4, 5, 6, 7))

	// for i := 1; i < 20; i++ {
	// 	fmt.Print(fibonacci(i), " ")
	// }

	///call function from lesson 15///
	///use pointer///
	// 	x := 5
	// 	mutiplyFunc(&x)
	// 	fmt.Println(x)

	// 	///use new///
	// 	input := new(int)
	// 	*input = 20
	// 	mutiplyFunc(input)
	// 	fmt.Println(*input)

	/////lesson 16 struct/////
	// type Dimensions struct {
	// 	height,
	// 	width,
	// 	thickness float32
	// }
	// ///initialization///
	// d := Dimensions{100.4, 50.2, 21.1}
	// fmt.Println(d)
	// ///access fields///
	// fmt.Println(d.height)
	// volume := func(d Dimensions) float32 {
	// 	return d.height * d.width * d.thickness
	// }
	// fmt.Println(volume(d))

	///call function from lesson 17///
	p := Person{"Ham", 32, "Chonburi"}
	fmt.Println(p.Detail())

}

//////////////////////////////////////////////////////////////////////////////////////
/////Lesson 10 function/////

///function with one return///
// func f() int {
// 	return 100
// }

///function with more than one return///
// func r() (int, int, int) {
// 	return 10, 20, 30
// }

///function with one input///
// func i(x int) {
// 	fmt.Println(x)
// }

///function with more than one input///
// func t(x int, y int, z int) {
// 	fmt.Println(x, y, z)
// }

///function with infinity input///
// func a(args ...int) {
// 	total := 0
// 	for _, v := range args {
// 		total += v
// 	}
// 	fmt.Println(total)
// }

/////Lesson 12 function call itself/////
// func plus(x uint) uint {
///Note:if not put that code with give run time error///
// if x == 0 {
// 	return 1
// }
///plus until x = 0///
// 	return x + plus(x-1)
// }

/////Lesson 13 defer/////
// func one() {
// 	fmt.Println(1)
// }

// func two() {
// 	fmt.Println(2)
// }

/////pratice/////
// func signature(arr []int) int {
// 	sum := 0
// 	for i := 0; i < len(arr); i++ {
// 		sum += arr[i]
// 	}
// 	return sum
// }

// func half(input int) (int, bool) {
// 	if input%2 == 0 {
// 		return (input / 2), true
// 	} else if input%2 == 1 {
// 		return (input / 2), false
// 	}
// 	return 3, false
// }

// func highestFunc(input ...int) int {
// 	highestValue := 0
// 	for _, v := range input {
// 		if highestValue < v {
// 			highestValue = v
// 		}
// 	}
// 	return highestValue
// }

// func fibonacci(input int) int {
// 	if input == 1 {
// 		return 1
// 	} else if input == 2 {
// 		return 2
// 	} else {
// 		return fibonacci(input-1) + fibonacci(input-2)
// 	}

// }

/////Lesson 15 pointer and new/////
// func mutiplyFunc(input *int) {
// 	*input = *input * 2
// }

/////lesson 17 method/////

//Detail //
func (p *Person) Detail() string {
	return p.Name + " " + p.Location
}
