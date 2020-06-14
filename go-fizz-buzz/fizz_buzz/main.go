package fizz_buzz

import "strconv"

var memoizedValue = [15]string{
	"FizzBuzz",
	"",
	"",
	"Fizz",
	"",
	"Buzz",
	"Fizz",
	"",
	"",
	"Fizz",
	"Buzz",
	"",
	"Fizz",
	"",
	"",
}
var memoizedReturnNumber = [15]bool{
	false,
	true,
	true,
	false,
	true,
	false,
	false,
	true,
	true,
	false,
	false,
	true,
	false,
	true,
	true,
}
func FizzBuzz (input uint) string {
	reminder := input % 15
	if memoizedReturnNumber[reminder] {
		return strconv.FormatUint(uint64(input), 10)
	}
	return memoizedValue[reminder]
}

func fizzBuzzSlow(input uint) string {
	if input % 15 == 0 {
		return "FizzBuzz"
	}
	if input % 3 == 0 {
		return "Fizz"
	}
	if input % 5 == 0 {
		return "Buzz"
	}
	return strconv.FormatUint(uint64(input), 10)
}


func fizzBuzzFakeSlow(input uint) string {
	if input % 15 == 0 {
		return "FizzBuzz"
	}
	if input % 3 == 0 {
		return "Fizz"
	}
	if input % 5 == 0 {
		return "Buzz"
	}
	return ""
}

func fizzBuzzFake(input uint) string {
	reminder := input % 15
	if memoizedReturnNumber[reminder] {
		return ""
	}
	return memoizedValue[reminder]
}
