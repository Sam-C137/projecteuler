package number_letter_counts

import (
	"strconv"
)

var dict = map[string]string{
	"1":    "one",
	"2":    "two",
	"3":    "three",
	"4":    "four",
	"5":    "five",
	"6":    "six",
	"7":    "seven",
	"8":    "eight",
	"9":    "nine",
	"10":   "ten",
	"11":   "eleven",
	"12":   "twelve",
	"13":   "thirteen",
	"14":   "fourteen",
	"15":   "fifteen",
	"16":   "sixteen",
	"17":   "seventeen",
	"18":   "eighteen",
	"19":   "nineteen",
	"20":   "twenty",
	"30":   "thirty",
	"40":   "forty",
	"50":   "fifty",
	"60":   "sixty",
	"70":   "seventy",
	"80":   "eighty",
	"90":   "ninety",
	"100":  "onehundred",
	"200":  "twohundred",
	"300":  "threehundred",
	"400":  "fourhundred",
	"500":  "fivehundred",
	"600":  "sixhundred",
	"700":  "sevenhundred",
	"800":  "eighthundred",
	"900":  "ninehundred",
	"1000": "onethousand",
}

func Run(target int) int {
	var acc int
	for n := 1; n <= target; n++ {
		str := strconv.Itoa(n)

		switch {
		case n <= 20:
			acc += len(dict[str])
		case n < 100:
			tens := str[0:1] + "0"
			units := str[1:2]
			acc += len(dict[tens])
			if units != "0" {
				acc += len(dict[units])
			}
		case n < 1000:
			hundreds := str[0:1] + "00"
			remainder := n % 100
			acc += len(dict[hundreds])
			if remainder != 0 {
				acc += len("and")
				if remainder <= 20 {
					acc += len(dict[strconv.Itoa(remainder)])
				} else {
					tens := strconv.Itoa(remainder)[0:1] + "0"
					units := strconv.Itoa(remainder)[1:2]
					acc += len(dict[tens])
					if units != "0" {
						acc += len(dict[units])
					}
				}
			}
		default:
			acc += len(dict[str])
		}
	}

	return acc
}
