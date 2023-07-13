package configurer

import "github.com/sirupsen/logrus"

func Logger(loglevel string) error {

	logger := logrus.New()

	level, err := logrus.ParseLevel(loglevel)
	if err != nil {
		return err
	}

	logger.SetLevel(level)
	return nil
}
