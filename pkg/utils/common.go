package utils

import (
	"fmt"

	"gorm.io/gorm"
)

func IF(bool2 bool, a interface{}, b interface{}) interface{} {
	if bool2 {
		return a
	}
	return b
}

func WrapDbErr(err error) error {
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil
		}
		fmt.Printf("db err:%v\n", err)
		return fmt.Errorf("db err")
	}
	return err
}

func ErrToPanic(err error) {
	if err != nil {
		panic(err)
	}
}
