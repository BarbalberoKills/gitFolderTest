package userinput

import "fmt"

type AnyValue interface {
	int | float64
}

func GetUserValue[T AnyValue](s *string) (val T) {
	for {
		fmt.Println(*s)
		_, err := fmt.Scan(&val)
		if err == nil {
			break
		}
		fmt.Println("Value not accepted - Error: ", err)
	}
	return val
}
