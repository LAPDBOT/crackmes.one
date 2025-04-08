package controller

import (
    "github.com/xushneg6/crackmes.one/app/model"
    "github.com/xushneg6/crackmes.one/app/shared/view"
    "log"
    "net/http"
)

// IndexGET displays the home page
func IndexGET(w http.ResponseWriter, r *http.Request) {
	// Display the view
	v := view.New(r)
	v.Name = "index/index"
	var nbusers, nbcrackmes, nbsolutions int
	var err error

	nbusers, err = model.CountUsers()
	if err != nil {
		log.Println(err)
		Error500(w, r)
		return
	}

	nbcrackmes, err = model.CountCrackmes()
	if err != nil {
		log.Println(err)
		Error500(w, r)
		return
	}

	nbsolutions, err = model.CountSolutions()
	if err != nil {
		log.Println(err)
		Error500(w, r)
		return
	}

	v.Vars["nbusers"] = nbusers
	v.Vars["nbsolutions"] = nbsolutions
	v.Vars["nbcrackmes"] = nbcrackmes
	v.Render(w)
}
