package alenar

import "log"

func checkError(err error) {
	if err != nil {
		//fmt.Println(err)
		log.Println(err)
	}
}
