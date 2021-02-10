package main

import "fmt"

func main() {
	// creating array
	arr := [6]int{1, 4, 3, 5, 8, 2}

	sliceFromArray(arr)

	//this is slice in Go its just like ArrayList in java have grovable nature...
	// originalSlc1 := []int{1, 4, 3, 5, 8, 2, 7}
	// slc2 := []int{10, 12, 13, 17, 20, 14}
	multiDimnsnSlc1 := [][]int{{1, 2, 3}, {4, 7, 6}, {8, 5, 9}} // multi-dimensional slice

	var slc1, slc2, slc3, slc4 = diffrentWaysToCreateSlices()
	sliceFromSlice(slc1)
	sliceUsingMakeFunc()
	iterateSliceByForLoop(slc1)
	slicesareReferanceType(arr) //proff
	compareSlices(slc1, slc2)
	multiDimensionSlices(multiDimnsnSlc1)
	copySlice(slc1, slc2, slc3, slc4)

}

func copySlice(slc1, slc2, slc3, slc4 []int) {

	// Before copying
	fmt.Println("Slice_1:", slc1)
	fmt.Println("Slice_2:", slc2)
	fmt.Println("Slice_3:", slc3)
	fmt.Println("Slice_4:", slc4)

	// Copying the slices
	copy1 := copy(slc4, slc1)
	fmt.Println("No of elements copied into Slice_4 : ", copy1)
	fmt.Println("Slice_4 after copy: ", slc4) //nothing is coppied because slc4 is nill

	copy2 := copy(slc3, slc1)
	fmt.Println("No of elements copied into Slice_3 : ", copy2)
	fmt.Println("Slice_3 after copy: ", slc3)

	copy3 := copy(slc2, slc1)
	fmt.Println("No of elements copied into Slice_2 : ", copy3)
	fmt.Println("Slice_2 after copy: ", slc2) //All slc2 valus are overriden

}

func diffrentWaysToCreateSlices() ([]int, []int, []int, []int) {
	// Creating slices
	slc1 := []int{58, 69, 40, 45, 11, 56, 67, 21, 65}
	slc2 := []int{78, 50, 67, 77}
	slc3 := make([]int, 5)
	var slc4 []int

	fmt.Println("Slice_1:", slc1)
	fmt.Println("Slice_2:", slc2)
	fmt.Println("Slice_3:", slc3)
	fmt.Println("Slice_4:", slc4)

	return slc1, slc2, slc3, slc4

}

func multiDimensionSlices(multiSlice [][]int) {
	fmt.Println("MultiSlice: ", multiSlice)
}

func compareSlices(slice1, slice2 []int) {
	//you can only use == operator to check the given slice is nill or not.
	fmt.Println(slice1 == nil)

	//If you try to compare two slices with the help of == operator then it will give you an error
	// fmt.Println(slice1 == slice2)// this line gives error

	//Note: If you want to compare two slices, then use range for loop to match each element or you can use DeepEqual function.
}

func slicesareReferanceType(arr [6]int) {
	slc := arr[0:4]

	// Before modifying
	fmt.Println("Original_Array: ", arr)
	fmt.Println("Original_Slice: ", slc)

	// After modification
	// these changes will effect on array aswell because slice is a reference type
	slc[0] = 100
	slc[2] = 200
	slc[1] = 300

	fmt.Println("\nNew_Array: ", arr)
	fmt.Println("New_Slice: ", slc)
}

func iterateSliceByForLoop(originalSlice []int) {
	// Iterate using for loop
	for i := 0; i < len(originalSlice); i++ {
		fmt.Println(originalSlice[i])
	}

	// Iterate slice
	// using range in for loop
	for index, element := range originalSlice {
		fmt.Printf("Index = %d and element = %s\n", index, element)
	}

	// Iterate slice
	// using range in for loop
	// without index
	for _, ele := range originalSlice {
		fmt.Printf("Element = %s\n", ele)
	}
}

func sliceUsingMakeFunc() {
	// Creating an array of size 7
	// and slice this array  till 4
	// and return the reference of the slice
	// Using make function
	var mySlice1 = make([]int, 4, 7)
	fmt.Printf("Slice 1 = %v, \nlength = %d, \ncapacity = %d\n",
		mySlice1, len(mySlice1), cap(mySlice1))

	// Creating another array of size 7
	// and return the reference of the slice
	// Using make function
	var mySlice2 = make([]int, 7)
	fmt.Printf("Slice 2 = %v, \nlength = %d, \ncapacity = %d\n",
		mySlice2, len(mySlice2), cap(mySlice2))
}

func sliceFromSlice(originalSlice []int) {
	// Creating slices from the given slice
	var mySlice1 = originalSlice[1:5]
	mySlice2 := originalSlice[0:]
	mySlice3 := originalSlice[:6]
	mySlice4 := originalSlice[:]
	mySlice5 := mySlice3[2:4]
	// Display the result
	fmt.Println("Original Slice:", originalSlice)
	fmt.Println("New Slice 1:", mySlice1)
	fmt.Println("New Slice 2:", mySlice2)
	fmt.Println("New Slice 3:", mySlice3)
	fmt.Println("New Slice 4:", mySlice4)
	fmt.Println("New Slice 5:", mySlice5)
}

func sliceFromArray(arr [6]int) {
	// Creating slices from the given array
	var mySlice1 = arr[1:2]
	mySlice2 := arr[0:]
	mySlice3 := arr[:2]
	mySlice4 := arr[:]

	// Display the result
	fmt.Println("My Array: ", arr)
	fmt.Println("My Slice 1: ", mySlice1)
	fmt.Println("My Slice 2: ", mySlice2)
	fmt.Println("My Slice 3: ", mySlice3)
	fmt.Println("My Slice 4: ", mySlice4)
}
