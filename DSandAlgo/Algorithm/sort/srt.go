package sort

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func SortArrayOnBasisOfModOf3(arr []string) []string {

	var lenght []string
	for _, val := range arr {
		mod := len(val) % 3

		fmt.Println(mod)
		lenght = append(lenght, val+"_"+strconv.Itoa(mod))
	}
	fmt.Println(lenght)
	var arr_2 []string
	var arr_1 []string
	var arr_0 []string
	for _, val := range lenght {
		if strings.Contains(val, "_2") {
			arr_2 = append(arr_2, val[:len(val)-2])
		}
		if strings.Contains(val, "_1") {
			arr_1 = append(arr_1, val[:len(val)-2])
		}
		if strings.Contains(val, "_0") {
			arr_0 = append(arr_0, val[:len(val)-2])
		}
	}

	sort.Strings(arr_0)
	sort.Strings(arr_1)
	sort.Strings(arr_2)

	arr_2 = append(arr_2, arr_1...)
	arr_2 = append(arr_2, arr_0...)
	return arr_2
}
