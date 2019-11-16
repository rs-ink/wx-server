package util

import (
	"fmt"
	"testing"
)

type ActOtherOperator struct {
	Phone  string `json:"phone"`
	Remark string `json:"remark"`
}

func TestIsExistItem(t *testing.T) {
	acts := make([]ActOtherOperator, 0)
	acts = append(acts, ActOtherOperator{
		Phone: "18653256546",
	})
}

func TestShowTimes(t *testing.T) {
	fmt.Println(GetBetweenAndTimes(7))
}
