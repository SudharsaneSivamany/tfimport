package contains

func Contains(list []interface{}, str string) bool {
	for _, value := range list {
		if value == str {
			return true
		}
	}
	return false
}
