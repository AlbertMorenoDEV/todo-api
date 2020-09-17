package create

//func TodoCreate(w http.ResponseWriter, r *http.Request) {
//	var todo Todo
//	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
//	if err != nil {
//		panic(err)
//	}
//	if err := r.Body.Close(); err != nil {
//		panic(err)
//	}
//	if err := json.Unmarshal(body, &todo); err != nil {
//		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
//		w.WriteHeader(422) // unprocessable entity
//		if err := json.NewEncoder(w).Encode(err); err != nil {
//			panic(err)
//		}
//	}
//
//	t := RepoCreateTodo(todo)
//	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
//	w.WriteHeader(http.StatusCreated)
//	if err := json.NewEncoder(w).Encode(t); err != nil {
//		panic(err)
//	}
//}
//
//type addGopherRequest struct {
//	ID    string `json:"ID"`
//	Name  string `json:"name"`
//	Image string `json:"image"`
//	Age   int    `json:"age"`
//}
//
//// AddGopher save a gopher
//func (s *server) AddGopher(w http.ResponseWriter, r *http.Request) {
//	decoder := json.NewDecoder(r.Body)
//
//	var g addGopherRequest
//	err := decoder.Decode(&g)
//
//	w.Header().Set("Content-Type", "application/json")
//	if err != nil {
//		w.WriteHeader(http.StatusInternalServerError)
//		_ = json.NewEncoder(w).Encode("Error unmarshalling request body")
//		return
//	}
//	if err := s.adding.AddGopher(r.Context(), g.ID, g.Name, g.Image, g.Age); err != nil {
//		w.WriteHeader(http.StatusInternalServerError)
//		_ = json.NewEncoder(w).Encode("Can't create a gopher")
//		return
//	}
//
//	w.WriteHeader(http.StatusCreated)
//}
