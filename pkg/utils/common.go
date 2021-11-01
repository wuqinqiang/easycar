package utils

import (
	"fmt"
	"gorm.io/gorm"
)

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
