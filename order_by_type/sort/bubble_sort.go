package sort

/**

默认升序

j := 0; j < n; j++
i := 0; i < n-1-j; i++ 先-1，再-j


*/
func bubbleSort(array []int) []int {
	if len(array) <= 1 {
		return array
	}

	n := len(array)
	for j := 0; j < n; j++ {
		for i := 0; i < n-j-1; i++ { //n-1代表最后一位，-j代表后几位，因为没有=号，相当于少了一位，所以不用怕溢出的问题

			//如果大于后面的，就和后面的一个交换
			if array[i] > array[i+1] {
				array[i], array[i+1] = array[i+1], array[i]
			}
		}

	}

	return array

}

/**

优化版本
1.加个flag，外循环flag,内循环里如果有序，则break
 */


func bubbleSort2(array []int) []int {
	if len(array) <= 1 {
		return array
	}

	n := len(array)
	for j := 0; j < n; j++ {

		flag := true

		for i := 0; i < n-j-1; i++ { //n-1代表最后一位，-j代表后几位，因为没有=号，相当于少了一位，所以不用怕溢出的问题

			//如果大于后面的，就和后面的一个交换
			if array[i] > array[i+1] {
				flag = false
				array[i], array[i+1] = array[i+1], array[i]
			}
		}

		if flag {	//如果有序，就break
			break
		}

	}

	return array

}
