package main

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

}
