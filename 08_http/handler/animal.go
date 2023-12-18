package handler

// func AnimalGetPost(w http.ResponseWriter, r *http.Request) {
// 	if !helper.CheckMethod(w, r, http.MethodGet) {
// 		return
// 	}
// 	o, err := service.GetAllAnimalCsv()
// 	if err != nil {
// 		helper.InternalError(w, err)
// 		return
// 	}
// 	err = helper.AsJson(w, o)
// 	if err != nil {
// 		helper.InternalError(w, err)
// 		return
// 	}
// }

// func Animal(w http.ResponseWriter, r *http.Request) {

// 	id, err := helper.GetIntId(r.URL.Path)

// 	if err != nil {
// 		util.InternalError(w, err)
// 		return
// 	}

// 	switch r.Method {
// 	case http.MethodGet:
// 		a, err := service.GetAnimalCsv(id)
// 		if err != nil {
// 			if errors.Is(err, service.ErrNotFound) {
// 				util.NotFound(w)
// 			} else {
// 				util.InternalError(w, err.Error())
// 			}
// 			return
// 		}
// 		util.AsJson(w, a)
// 		return
// 	case http.MethodPut:
// 		return
// 	case http.MethodDelete:
// 		return
// 	default:
// 		m := []string{
// 			http.MethodGet,
// 			http.MethodPut,
// 			http.MethodDelete,
// 		}
// 		helper.NotAllowed(w, m)
// 	}
// }
