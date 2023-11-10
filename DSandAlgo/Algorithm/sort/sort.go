package sort

import (
	"fmt"
	"sort"
)

func SortingSlice() {
	// ints
	numbers := []int{2, 3, 53, 1, 33, 222}
	sort.Ints(numbers)
	fmt.Println("Sorting the numbers in Ascending Orders :", numbers)

	// String
	str := []string{"a", "b", "d", "c"}
	sort.Strings(str)
	fmt.Println("Sorting the strings in Ascending Orders :", str)

	// Floats
	flt := []float64{1.23, 2.2, 1.11, 3.5}
	sort.Float64s(flt)
	fmt.Println("Sorting the  in Ascending Orders :", flt)

}

func CustomSorting() {

	type Person struct {
		Name string
		Age  int
	}

	persons := []Person{
		{
			Name: "ABC",
			Age:  27,
		}, {Name: "BCA",
			Age: 28},

		{
			Name: "CAB",
			Age:  27,
		},
		{
			Name: "DAC",
			Age:  29,
		},
	}

	// Sort on the Basis of the Age
	sort.Slice(persons, func(i, j int) bool {
		return persons[i].Age < persons[j].Age
	})

	fmt.Println("Sorted persons based on Age: ", persons)

	// Sort on the Basis of the Name
	sort.Slice(persons, func(i, j int) bool {
		return persons[i].Name < persons[j].Name
	})
	fmt.Println("Sorted persons based on Name: ", persons)
}
