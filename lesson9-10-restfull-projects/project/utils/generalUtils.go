package utils

func checkError(err error) {
	if err != nil {
		err.Error()
	}
}