package main

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
	"unicode"
)

func main() {
	task72()
}

func task1() {
	var a int
	fmt.Scan(&a)
	fmt.Println(a*2 + 100)
}

func task2() {
	var a, b int
	fmt.Scan(&a, &b)
	fmt.Println(a*a + b*b)
}

func task3() {
	var a int
	fmt.Scan(&a)
	fmt.Println(a * a)
}

func task4() {
	var a int
	fmt.Scan(&a)
	fmt.Println(a % 10)
}

func task5() {
	var a int
	fmt.Scan(&a)
	fmt.Println(a % 100 / 10)
}

func task6() {
	var degree, hours, minutes int
	fmt.Scan(&degree)
	hours = degree / 30
	minutes = degree % 30 * 2
	fmt.Printf("It is %d hours %d minutes.", hours, minutes)
}

func task7() {
	var number int
	fmt.Scan(&number)
	switch {
	case number > 0:
		fmt.Println("Число положительное")
	case number == 0:
		fmt.Println("Ноль")
	case number < 0:
		fmt.Println("Число отрицательное")
	}
}

func task8() {
	var number int
	fmt.Scan(&number)
	if number/100 != number%100/10 && number%100/10 != number%10 && number/100 != number%10 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

func task9() {
	var number int
	fmt.Scan(&number)
	switch {
	case number/10000 != 0:
		fmt.Println(number / 10000)
	case number/1000 != 0:
		fmt.Println(number % 10000 / 1000)
	case number/100 != 0:
		fmt.Println(number % 1000 / 100)
	case number/10 != 0:
		fmt.Println(number % 100 / 10)
	default:
		fmt.Println(number)
	}
}

func task10() {
	var number int
	fmt.Scan(&number)
	if number/100000+number%100000/10000+number%10000/1000 ==
		number%1000/100+number%100/10+number%10 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

func task11() {
	var year int
	fmt.Scan(&year)
	if year%400 == 0 || (year%4 == 0 && year%100 != 0) {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

func task12() {
	for i := 1; i <= 10; i++ {
		fmt.Println(i * i)
	}
}

func task13() {
	var a, b, sum int
	fmt.Scan(&a, &b)
	for ; a <= b; a++ {
		sum += a
	}
	fmt.Println(sum)
}

func task14() {
	var n, sum int
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		var temp int
		fmt.Scan(&temp)
		if temp > 9 && temp < 100 && temp%8 == 0 {
			sum += temp
		}
	}
	fmt.Println(sum)
}

func task15() {
	var number, count, max int
	for fmt.Scan(&number); number != 0; fmt.Scan(&number) {
		if number > max {
			max = number
			count = 1
		} else if number == max {
			count += 1
		}
	}
	fmt.Println(count)
}

func task16() {
	var n, c, d, answer int
	fmt.Scan(&n, &c, &d)
	for i := 1; i <= n; i++ {
		if i%c == 0 && i%d != 0 {
			answer = i
			break
		}
	}
	if answer != 0 {
		fmt.Println(answer)
	}
}

func task17() {
	var number int
	for fmt.Scan(&number); number < 100; fmt.Scan(&number) {
		if number < 10 {
			continue
		}
		fmt.Println(number)
	}
}

func task18() {
	var x, p, y, counter int
	fmt.Scan(&x, &p, &y)
	for ; x < y; x = x * (100 + p) / 100 {
		counter++
	}
	fmt.Println(counter)
}

func task19() {
	var a, b int
	fmt.Scan(&a, &b)
	for divider1 := 10000; divider1 > 0; divider1 = divider1 / 10 {
		digit1 := a / divider1
		temp := b
		for divider2 := 10000; divider2 > 0; divider2 = divider2 / 10 {
			digit2 := temp / divider2
			if digit1 == digit2 && digit1 != 0 {
				fmt.Print(digit1)
				fmt.Print(" ")
			}
			temp = temp % divider2
		}
		a = a % divider1
	}
}

func task20() {
	var number float64
	fmt.Scan(&number)
	number = number * number
	if number <= 0 {
		fmt.Printf("Число %2.2f не подходит", number)
	} else if number > 10000 {
		fmt.Printf("%e", number)
	} else {
		fmt.Printf("%4.4f", number)
	}
}

func task21() {
	var workArray [10]uint8
	for index := range workArray {
		fmt.Scan(&workArray[index])
	}
	for i := 0; i < 3; i++ {
		var i, j, temp uint8
		fmt.Scan(&i, &j)
		temp = workArray[i]
		workArray[i] = workArray[j]
		workArray[j] = temp
	}
	for _, elem := range workArray {
		fmt.Printf("%d ", elem)
	}
}

func task22() {
	var n int
	fmt.Scan(&n)
	slice := make([]int, n, n)
	for index := 0; index < n; index++ {
		fmt.Scan(&slice[index])
	}
	fmt.Println(slice[3])
}

func task23() {
	array := [5]int{}
	var a int
	for i := 0; i < 5; i++ {
		fmt.Scan(&a)
		array[i] = a
	}
	max := array[0]
	for _, elem := range array {
		if max < elem {
			max = elem
		}
	}
	fmt.Println(max)
}

func task24() {
	var n int
	fmt.Scan(&n)
	array := make([]int, n, n)
	for index := range array {
		fmt.Scan(&array[index])
	}
	for index, elem := range array {
		if index%2 == 0 {
			fmt.Printf("%d ", elem)
		}
	}
}

func task25() {
	var n, count int
	fmt.Scan(&n)
	array := make([]int, n, n)
	for index := range array {
		fmt.Scan(&array[index])
	}
	for _, elem := range array {
		if elem > 0 {
			count++
		}
	}
	fmt.Println(count)
}

func task26() {
	var n, sum int
	fmt.Scan(&n)
	for divider := 100; divider > 0; divider = divider / 10 {
		sum += n / divider
		n = n % divider
	}
	fmt.Println(sum)
}

func task27() {
	var n int
	fmt.Scan(&n)
	fmt.Println(n/100 + n/10%10*10 + n%10*100)
}

func task28() {
	var seconds int
	fmt.Scan(&seconds)
	fmt.Printf("It is %d hours %d minutes.", seconds/3600, seconds%3600/60)
}

func task29() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)
	if a*a+b*b == c*c {
		fmt.Println("Прямоугольный")
	} else {
		fmt.Println("Непрямоугольный")
	}
}

func task30() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)
	if a+b >= c && a+c >= b && b+c >= a {
		fmt.Println("Существует")
	} else {
		fmt.Println("Не существует")
	}
}

func task31() {
	var a, b int
	fmt.Scan(&a, &b)
	if (a+b)%2 == 0 {
		fmt.Println((a + b) / 2)
	} else {
		fmt.Printf("%.1f", (float64(a)+float64(b))/2.0)
	}
}

func task32() {
	var n, number, count int
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&number)
		if number == 0 {
			count++
		}
	}
	fmt.Println(count)
}

func task33() {
	var n, number, count, min int
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&number)
		if i == 0 || number < min {
			min = number
			count = 1
		} else if number == min {
			count++
		}
	}
	fmt.Println(count)
}

func task34() {
	var number, temp, answer int
	fmt.Scan(&number)
	for divider := 10000000; divider > 0; divider = divider / 10 {
		temp += number / divider
		number = number % divider
	}
	for divider := 100; divider > 0; divider = divider / 10 {
		answer += temp / divider
		temp = temp % divider
	}
	fmt.Println(answer)
}

func task35() {
	var a, b, i, max int
	fmt.Scan(&a, &b)
	i = a
	for ; i <= b; i++ {
		if i%7 == 0 {
			max = i
		}
	}
	if max == 0 && a > 0 {
		fmt.Println("NO")
	} else {
		fmt.Println(max)
	}
}

func task36() {
	var n, ost int
	var skl string
	fmt.Scan(&n)
	ost = n % 10
	switch {
	case n > 10 && n < 15:
		skl = "korov"
	case ost == 1:
		skl = "korova"
	case ost == 2 || ost == 3 || ost == 4:
		skl = "korovy"
	default:
		skl = "korov"
	}
	fmt.Printf("%d %s", n, skl)
}

func task37() {
	var n int
	fmt.Scan(&n)
	for i := 1; i <= n; i *= 2 {
		fmt.Printf("%d ", i)
	}
}

func task38() {
	n1 := 1
	n2 := 2
	count := 3
	var isFibonacciNumber bool
	var a int
	fmt.Scan(&a)
	for n2 <= a {
		if n2 == a {
			isFibonacciNumber = true
			break
		}
		temp := n2
		n2 = n2 + n1
		n1 = temp
		count++
	}
	if isFibonacciNumber {
		fmt.Println(count)
	} else {
		fmt.Println(-1)
	}
}

func task39() {
	var number int
	fmt.Scan(&number)
	fmt.Printf("%b", number)
}

func task40() {
	var number int
	var digit string
	fmt.Scan(&number, &digit)
	s := fmt.Sprintf("%v", number)
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == digit[0] {
			s = s[0:i] + s[i+1:len(s)]
		}
	}
	fmt.Println(s)
}

func task41(text string) {
	fmt.Println(text)
}

func task42() int {
	var a, b, c, d int
	fmt.Scan(&a, &b, &c, &d)
	return task42min(task42min(a, b), task42min(c, d))
}

func task42min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func task43(x, y, z int) int {
	if x+y+z >= 2 {
		return 1
	} else {
		return 0
	}
}

func task44(n int) int {
	if n > 0 && n < 3 {
		return 1
	}
	n1 := 1
	n2 := 1
	for i := 3; i <= n; i++ {
		temp := n2
		n2 += n1
		n1 = temp
	}
	return n2
}

func task45(numbers ...int) (int, int) {
	var sum int
	for _, elem := range numbers {
		sum += elem
	}
	return len(numbers), sum
}

func task46(x1, x2 *int) {
	fmt.Println(*x1 * *x2)
}

func task47(x1, x2 *int) {
	temp := *x1
	*x1 = *x2
	*x2 = temp
	fmt.Println(*x1, *x2)
}

type myStruct struct {
	On    bool
	Ammo  int
	Power int
}

func (ts *myStruct) Shoot() bool {
	if !ts.On || ts.Ammo == 0 {
		return false
	} else {
		ts.Ammo--
		return true
	}
}

func (ts *myStruct) RideBike() bool {
	if !ts.On || ts.Power == 0 {
		return false
	} else {
		ts.Power--
		return true
	}
}

func task48() {
	myStruct := new(myStruct)
	myStruct.Shoot()
	myStruct.RideBike()
}

func task49() {
	input, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	runes := []rune(input)
	if unicode.IsUpper(runes[0]) && runes[len(runes)-1] == '.' {
		fmt.Println("Right")
	} else {
		fmt.Println("Wrong")
	}
}

func task50() {
	var input, part1, part2 string
	fmt.Scan(&input)
	runes := []rune(input)
	if len(runes)%2 == 0 {
		part1 = string(runes[0 : len(runes)/2])
	} else {
		part1 = string(runes[0 : len(runes)/2+1])
	}
	part2 = string(task50reverse(runes[len(runes)/2 : len(runes)]))
	fmt.Println(part1, part2)
	if part1 == part2 {
		fmt.Println("Палиндром")
	} else {
		fmt.Println("Нет")
	}
}

func task50reverse(rns []rune) []rune {
	for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {
		rns[i], rns[j] = rns[j], rns[i]
	}
	return rns
}

func task51() {
	var input, subString string
	fmt.Scan(&input, &subString)
	fmt.Println(strings.Index(input, subString))
}

func task52() {
	var input, answer string
	fmt.Scan(&input)
	runes := []rune(input)
	for index, elem := range runes {
		if index%2 != 0 {
			answer += string(elem)
		}
	}
	fmt.Println(answer)
}

func task53() {
	var input string
	fmt.Scan(&input)
	runes := []rune(input)
	for _, elem := range runes {
		if strings.Count(input, string(elem)) > 1 {
			input = strings.Replace(input, string(elem), "", -1)
		}
	}
	fmt.Println(input)
}

func task54() {
	var input string
	checker := true
	fmt.Scan(&input)
	runes := []rune(input)
	for _, elem := range runes {
		if (!unicode.IsLetter(elem) || !unicode.Is(unicode.Latin, elem)) && !unicode.IsDigit(elem) || len(runes) < 5 {
			checker = false
			break
		}
	}
	if checker {
		fmt.Println("Ok")
	} else {
		fmt.Println("Wrong password")
	}
}

func task55() {
	var a, b float64
	fmt.Scan(&a, &b)
	fmt.Println(math.Sqrt(a*a + b*b))
}

func task56() {
	var input string
	fmt.Scan(&input)
	array := strings.Split(input, "")
	fmt.Println(strings.Join(array, "*"))
}

func task57() {
	var input string
	fmt.Scan(&input)
	runes := []rune(input)
	max := runes[0]
	for _, elem := range runes {
		if elem > max {
			max = elem
		}
	}
	fmt.Println(string(max))
}

func task58() {
	var numbers, answer string
	fmt.Scan(&numbers)
	runes := []rune(numbers)
	for _, number := range runes {
		intNum, _ := strconv.Atoi(string(number))
		answer += fmt.Sprintf("%v", intNum*intNum)
	}
	fmt.Println(answer)
}

func task59() {
	myMap := make(map[int]int)
	var array [10]int
	for i := 0; i < 10; i++ {
		var number int
		fmt.Scan(&number)
		array[i] = number
		if _, inMap := myMap[number]; inMap {
			continue
		}
		myMap[number] = work(number)
	}
	for _, value := range array {
		fmt.Printf("%d ", myMap[value])
	}
}

func work(x int) int {
	return x
}

func task60() {
	groupCity := map[int][]string{
		10:   []string{},
		100:  []string{},
		1000: []string{},
	}
	cityPopulation := make(map[string]int)
	for key := range cityPopulation {
		var isIn bool
		for _, value := range groupCity[100] {
			if value == key {
				isIn = true
			}
		}
		if isIn {
			continue
		}
		delete(cityPopulation, key)
	}
}

func task61(myInt64 int64) uint16 {
	return uint16(myInt64)
}

func task62(a, b string) int64 {
	runesA := []rune(a)
	runesB := []rune(b)
	for index, value := range runesA {
		if !unicode.IsDigit(value) {
			runesA[index] = ' '
		}
	}
	for index, value := range runesB {
		if !unicode.IsDigit(value) {
			runesB[index] = ' '
		}
	}
	a = strings.Replace(string(runesA), " ", "", -1)
	b = strings.Replace(string(runesB), " ", "", -1)
	number1, _ := strconv.ParseInt(a, 10, 64)
	number2, _ := strconv.ParseInt(a, 10, 64)
	return number1 + number2
}

func task63() {
	fmt.Println([]rune("0")[0] == '0')
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()
	myStrings := strings.Split(input, ";")
	fmt.Println(myStrings)
	for index := range myStrings {
		myStrings[index] = strings.Replace(myStrings[index], " ", "", -1)
		myStrings[index] = strings.Replace(myStrings[index], ",", ".", -1)
	}
	number1, _ := strconv.ParseFloat(myStrings[0], 64)
	number2, _ := strconv.ParseFloat(myStrings[1], 64)
	fmt.Printf("%.4f", number1/number2)
}

func task64() {
	fn := func(myUint uint) uint {
		myStr := strconv.FormatUint(uint64(myUint), 10)
		runes := []rune(myStr)
		count := 1
		for index := range runes {
			number, _ := strconv.ParseInt(string(runes[index]), 10, 64)
			if number%2 == 1 || runes[index] == '0' {
				runes = append(runes[0:index-count], runes[index-count:len(runes)]...)
				count++
			}
		}
		answer, _ := strconv.ParseUint(string(runes), 10, 64)
		return uint(answer)
	}
	fmt.Println(fn)
}

func task65() {
	var value1, value2, operation interface{} //:= readTask()
	var answer float64
	switch v := value1.(type) {
	case float64:
	default:
		errors.New(fmt.Sprintf("value=%v: %T", value1, v))
		os.Exit(0)
	}
	switch v := value2.(type) {
	case float64:
	default:
		errors.New(fmt.Sprintf("value=%v: %T", value2, v))
		os.Exit(0)
	}
	switch operation.(type) {
	case string:
		switch operation {
		case "/":
			answer = value1.(float64) / value2.(float64)
		case "+":
			answer = value1.(float64) + value2.(float64)
		case "-":
			answer = value1.(float64) - value2.(float64)
		case "*":
			answer = value1.(float64) * value2.(float64)
		}
	default:
		errors.New("неизвестная операция")
		os.Exit(0)
	}
	fmt.Println(answer)
}

type Battery struct {
	charge string
}

func (a Battery) String() string {
	count := strings.Count(a.charge, "1")
	return strings.Repeat(" ", 10-count) + strings.Repeat("X", count)
}

func task66() {
	var charge string
	fmt.Scan(&charge)
	batteryForTest := Battery{
		charge: charge,
	}
	fmt.Println(batteryForTest)
}

func task67() {
	csvFile, _ := os.Open("../task.data")
	rd := bufio.NewReader(csvFile)
	reader := csv.NewReader(rd)
	reader.Comma = ';'
	for {
		record, _ := reader.Read()
		for num, item := range record {
			if item == "0" {
				fmt.Println(num + 1)
				break
			}
		}
	}
}

type (
	Student struct {
		Rating []int
	}

	Group struct {
		Students []Student
	}

	Rating struct {
		Average float32
	}
)

func task68() {
	data, _ := ioutil.ReadAll(os.Stdin)
	var group Group
	var count, rates int
	if err := json.Unmarshal(data, &group); err != nil {
		fmt.Println(err)
		return
	}
	count = len(group.Students)
	for _, student := range group.Students {
		rates += len(student.Rating)
	}
	rating := Rating{Average: float32(rates) / float32(count)}
	data1, _ := json.MarshalIndent(rating, "", "\t")
	io.WriteString(os.Stdout, string(data1))
}

type dataStruct struct {
	GlobalId int64 `json:"global_id"`
}

type rawData struct {
	units []dataStruct
}

type User []struct {
	GlobalId int `json:"global_id"`
}

func task69() {
	jsonFile, err := os.Open("C:\\Users\\sahar\\GolandProjects\\learnGo\\data.json")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened users.json")
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var users User
	json.Unmarshal(byteValue, &users)
	var sum int
	for i := 0; i < len(users); i++ {
		sum += users[i].GlobalId
	}
	fmt.Println(sum)
}

func task70() {
	var input string
	fmt.Scan(&input)
	var year, month, day, hour, minutes, seconds, utcHours, utcMinutes int
	year, _ = strconv.Atoi(input[0:4])
	month, _ = strconv.Atoi(input[5:7])
	day, _ = strconv.Atoi(input[8:10])
	hour, _ = strconv.Atoi(input[11:13])
	minutes, _ = strconv.Atoi(input[14:16])
	seconds, _ = strconv.Atoi(input[17:19])
	utcHours, _ = strconv.Atoi(input[20:22])
	utcMinutes, _ = strconv.Atoi(input[23:25])
	offset := time.FixedZone("offset", (utcHours*60+utcMinutes)*60)
	time := time.Date(year, time.Month(month), day, hour, minutes, seconds, 0, offset)
	fmt.Println(time.Format("Mon Jan _2 15:04:05 -0700 2006"))
}

func task71() {
	var input1, input2 string
	fmt.Scan(&input1, &input2)
	input := input1 + " " + input2
	var year, month, day, hour, minutes, seconds int
	year, _ = strconv.Atoi(input[0:4])
	month, _ = strconv.Atoi(input[5:7])
	day, _ = strconv.Atoi(input[8:10])
	hour, _ = strconv.Atoi(input[11:13])
	minutes, _ = strconv.Atoi(input[14:16])
	seconds, _ = strconv.Atoi(input[17:19])
	if hour > 13 {
		day = day + 1
	}
	time := time.Date(year, time.Month(month), day, hour, minutes, seconds, 0, time.Local)
	fmt.Println(time.Format("2006-01-02 15:04:05"))
}

func task72() {
	var input1, input2, input3 string
	fmt.Scan(&input1, &input2, &input3)
	input := input1 + " " + input2 + " " + input3
	times := strings.Split(input, ",")
	var myTimes [2]time.Time
	for index, input := range times {
		var year, month, day, hour, minutes, seconds int
		year, _ = strconv.Atoi(input[6:10])
		month, _ = strconv.Atoi(input[3:5])
		day, _ = strconv.Atoi(input[0:2])
		hour, _ = strconv.Atoi(input[11:13])
		minutes, _ = strconv.Atoi(input[14:16])
		seconds, _ = strconv.Atoi(input[17:19])
		if hour > 13 {
			day = day + 1
		}
		myTimes[index] = time.Date(year, time.Month(month), day, hour, minutes, seconds, 0, time.Local)
	}
	var duration time.Duration
	if myTimes[0].Before(myTimes[1]) {
		duration = time.Since(myTimes[0]).Round(time.Second) - time.Since(myTimes[1]).Round(time.Second)
	} else {
		duration = time.Since(myTimes[1]).Round(time.Second) - time.Since(myTimes[0]).Round(time.Second)
	}
	fmt.Println(duration)
}

func task73() {
	const now = 1589570165
	var min, sec int64
	fmt.Scanf("%d мин. %d", &min, &sec)
	answer := time.Unix(now+min*60+sec, 0).Add(0)
	fmt.Println(answer.Format(time.UnixDate))
}

func task74(chanel chan int, N int) {
	chanel <- N + 1
}

func task75(chanel chan string, s string) {
	chanel <- strings.Repeat(s, 5) + " "
}

func task76(inputStream, outputStream chan string) {
	var lastAdded string
	for elem := range inputStream {
		if lastAdded != elem {
			outputStream <- elem
			lastAdded = elem
		}
	}
	close(inputStream)
	close(outputStream)
}

func task77(firstChan <-chan int, secondChan <-chan int, stopChan <-chan struct{}) <-chan int {
	var output chan int
	go func() {
		defer close(output)
		select {
		case x := <-firstChan:
			output <- x * x
		case x := <-secondChan:
			output <- x * 3
		case <-stopChan:
			return
		default:
		}
	}()
	return output
}

func task78(arguments <-chan int, done <-chan struct{}) <-chan int {
	output := make(chan int)
	var sum int
	go func() {
		defer close(output)
		for {
			select {
			case x := <-arguments:
				sum += x
			case <-done:
				output <- sum
			}
		}
	}()
	return output
}

func task79(f func(int) int, in1 <-chan int, in2 <-chan int, out chan<- int, n int) {
	idx := make(chan int)
	val := make(chan int)

	go func() {
		for i := 0; i < n; i++ {
			x1 := <-in1
			x2 := <-in2

			done := make(chan int)
			go func() {
				done <- f(x1)
			}()
			go func() {
				done <- f(x2)
			}()

			go func(i int) {
				idx <- i
				val <- (<-done) + (<-done)
			}(i)
		}
	}()

	go func() {
		m := make(map[int]int)
		for i := 0; i < n; i++ {
			m[<-idx] = <-val
		}

		for i := 0; i < n; i++ {
			out <- m[i]
		}
	}()
}
