package utils

func CheckIsOfAge(age int) (isOfAge bool) {
	if age >= 18 {
		isOfAge = true
		return
	}
	isOfAge = false
	return
}
