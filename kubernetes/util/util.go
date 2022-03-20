package util

func Must(e interface{}) {
	if e != nil {
		panic(e)
	}
}
