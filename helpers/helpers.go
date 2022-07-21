package helpers

func CheckError(err error) {
	if err != nil {
		log.fmt.Println(err.Error())
	}
}
