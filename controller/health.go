package controller

import (
	"log"
	"net/http"
	"zydukas-rankings/utils"
)

func HealthCheckHandler(w http.ResponseWriter, _ *http.Request) {
	data := map[string]string{
		"status": "running",
	}

	if err := utils.WriteJson(w, http.StatusOK, data); err != nil {
		err = utils.WriteJsonError(w, http.StatusInternalServerError, err.Error())
		if err != nil {
			log.Println(err.Error())
		}
	}
}
