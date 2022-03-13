package main

import (
	"net/http"

	"timesavvy/app/controllers"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {
	// creating DB object
	// db, err = sql.Open("mysql", "root:@/my_app_db")
	// if err != nil {
	// 	panic(err.Error())
	// }
	// defer db.Close()

	// init router with mux
	router := mux.NewRouter()

	router.Handle("/", http.FileServer(http.Dir("./views/")))
	// creating endpoints
	router.HandleFunc("/auth/register", controllers.RegisterHandler).Methods("POST")
	// router.HandleFunc("/posts", getPosts).Methods("GET")
	// router.HandleFunc("/posts", createPost).Methods("POST")
	// router.HandleFunc("/posts/{id}", getPost).Methods("GET")
	// router.HandleFunc("/posts/{id}", updatePost).Methods("PUT")
	// router.HandleFunc("/posts/{id}", deletePost).Methods("DELETE")

	// run our server on port 8000
	http.ListenAndServe(":8000", router)
}

// handle get req
// func getPosts(w http.ResponseWriter, r *http.Request) {
// 	// setting the header “Content-Type” to “application/json” cos its get request
// 	w.Header().Set("Content-Type", "application/json")
// 	// define array of posts called posts
// 	var posts []Post
// 	// sql query
// 	result, err := db.Query("SELECT id, title from posts")
// 	// check for errors
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	// closing connection everytime we done
// 	defer result.Close()
// 	// loop over the result and for every iteration we create a new Post instance
// 	for result.Next() {
// 		var post Post
// 		err := result.Scan(&post.ID, &post.Title)
// 		if err != nil {
// 			panic(err.Error())
// 		}
// 		// if all good append our post object into post array
// 		posts = append(posts, post)
// 	}
// 	// encode our posts to JSON
// 	json.NewEncoder(w).Encode(posts)
// }

// // handle create req
// func createPost(w http.ResponseWriter, r *http.Request) {
// 	stmt, err := db.Prepare("INSERT INTO posts(title) VALUES(?)")
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	body, err := ioutil.ReadAll(r.Body)
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	keyVal := make(map[string]string)
// 	json.Unmarshal(body, &keyVal)
// 	title := keyVal["title"]
// 	_, err = stmt.Exec(title)
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	fmt.Fprintf(w, "New post was created")
// }

// // handle get one post req
// func getPost(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	params := mux.Vars(r)
// 	result, err := db.Query("SELECT id, title FROM posts WHERE id = ?", params["id"])
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	defer result.Close()
// 	var post Post
// 	for result.Next() {
// 		err := result.Scan(&post.ID, &post.Title)
// 		if err != nil {
// 			panic(err.Error())
// 		}
// 	}
// 	json.NewEncoder(w).Encode(post)
// }

// // handle update post req
// func updatePost(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	params := mux.Vars(r)
// 	stmt, err := db.Prepare("UPDATE posts SET title = ? WHERE id = ?")
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	body, err := ioutil.ReadAll(r.Body)
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	keyVal := make(map[string]string)
// 	json.Unmarshal(body, &keyVal)
// 	newTitle := keyVal["title"]
// 	_, err = stmt.Exec(newTitle, params["id"])
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	fmt.Fprintf(w, "Post with ID = %s was updated", params["id"])
// }

// // handle delete post req
// func deletePost(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	params := mux.Vars(r)
// 	stmt, err := db.Prepare("DELETE FROM posts WHERE id = ?")
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	_, err = stmt.Exec(params["id"])
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	fmt.Fprintf(w, "Post with ID = %s was deleted", params["id"])
// }
