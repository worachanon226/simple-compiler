package utils

import (
	"errors"
	"fmt"
	"simple-compiler/backend/config"
)

func ConnectionBuilder(stuff string, cfg *config.Config) (string, error) {
	var url string

	if stuff == "fiber" {
		url = fmt.Sprintf("%s:%s", cfg.App.Host, cfg.App.Port)
	} else {
		errMsg := fmt.Sprintf("error, connection url builder doesn't know the %s", stuff)
		return "", errors.New(errMsg)
	}

	return url, nil
}
