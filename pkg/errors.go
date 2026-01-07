package pkg

import "log"

func handleFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func handleErr(err error) error {
	if err != nil {
		return err
	}
	return nil
}
