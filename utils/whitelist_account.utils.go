package utils

var listID = []string{"e97cf91c-324b-4fc2-8cd5-70d64fe550ca"}

func CheckEmailExist(email string) bool {
	for _, v := range listID {
		if v == email {
			return true
		}
	}
	return false
}
