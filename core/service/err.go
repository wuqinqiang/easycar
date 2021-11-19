package service

import "fmt"

var (
	BeginTransactionErr = fmt.Errorf("begin transaction err")
	NotFindTransaction  = fmt.Errorf("not found transaction")
)
