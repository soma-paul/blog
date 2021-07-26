package handler

import (
	"log"
	"strconv"
)

func ConvertStringtoInt(s string) int {
	integerValue, err := strconv.Atoi(s)
	CheckError("error converting a string to integer", err)
	return integerValue
}

func CheckError(msg string, err error) {
	if err != nil {
		log.Fatal(msg, err.Error())
	}
}
