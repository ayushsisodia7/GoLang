package math

// func Abs(a int) int {
// 	if a > 0 {
// 		return a
// 	}
// 	return -a
// }

func Contains(nums []int, target int) bool {
	for _, num := range nums {
		if num == target {
			return true
		}
	}
	return false
}
