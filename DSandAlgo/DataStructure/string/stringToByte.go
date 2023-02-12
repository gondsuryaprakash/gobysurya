package string

import "fmt"

func StrToByteConversion() {

	// With []byte(str)
	str := "Hello Surya"
	// Conversion of String to Bytes
	strBytes := []byte(str)
	fmt.Println(strBytes)

	// Conversion from ByteArray to String
	newString := string(strBytes)

	fmt.Println("New String ", newString)
	if newString == str {
		fmt.Println("Strings are matched !!")
	} else {
		fmt.Println("Strings are not macthing !!")
	}
	// It will give the byte value with uint8
	for i := 0; i < len(newString); i++ {
		fmt.Printf("index: %d, value %v, type: %T\n", i, newString[i], newString[i])
	}

}

// rune alias of int32
// byte is uint8
// Conversion of String to rune

func StringToRune() {

	str := "Hello Surya "

	// convert str to rune []rune(str)
	runes := []rune(str)
	fmt.Println("Runes : = ", runes)

	// convert rune to str

	runeToStr := string(runes)
	fmt.Println("runtToStr", runeToStr)

	// Convert str to rune with Loop
	// By using the range in for loop
	for i, val := range runeToStr {
		fmt.Printf("index: %d, value %v, type: %T\n", i, val, val)
	}
	// Convert rune to Byte
	for i, val := range runeToStr {
		runtToByte := byte(val)
		fmt.Printf("index: %d, value %v, type: %T\n", i, runtToByte, runtToByte)

	}

}
