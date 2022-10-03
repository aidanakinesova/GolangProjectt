package main

import (
	"fmt"
	"net/http"
	"html/template"
	"github.com/gorilla/mux"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	
)

type Article struct{
	Id uint16
	Name,Anons,FullText,Image string
}
type User struct{
	Id uint16
	Name, Password string
}

var posts=[]Article{}
var showPost=Article{}

var users=[]User{}

type Customer struct{
	Id uint16
	Name, Password, Surname string
	Age int
	Sex, CardNumber string
}
var customers=[]Customer{}

func getArticles(w http.ResponseWriter, r *http.Request){
	t,err:=template.ParseFiles("templates/header.html","templates/articles.html")
	if err!=nil{
		fmt.Fprintf(w,err.Error())
	}

	posts=[]Article{}

	db,err:=sql.Open("mysql","root:@tcp(127.0.0.1:3306)/golangProject")
	if err!=nil{
		panic(err)
	}
	fmt.Println("Успешно подключились к базе данных")
	

	res,err:=db.Query("SELECT * FROM `Article`")
	if err!=nil{
		panic(err)
	}
	for res.Next(){
		var post Article
		err=res.Scan(&post.Id, &post.Name, &post.Anons,&post.FullText, &post.Image)
		if err!=nil{
			panic(err)
		}
		posts=append(posts, post)

	}
	fmt.Println("Успешно получили данных из таблицы")
	t.ExecuteTemplate(w,"articles", posts)
	defer db.Close()

}

func getArticleDetail(w http.ResponseWriter, r *http.Request){
	t,err:=template.ParseFiles("templates/header.html","templates/articleDetail.html")
	if err!=nil{
		fmt.Fprintf(w,err.Error())
	}

	posts=[]Article{}
	vars:=mux.Vars(r) //Vars returns the route variables for the current request, if any.

	db,err:=sql.Open("mysql","root:@tcp(127.0.0.1:3306)/golangProject")
	if err!=nil{
		panic(err)
	}
	fmt.Println("Успешно подключились к базе данных")
	

	res,err:=db.Query(fmt.Sprintf("SELECT * FROM `Article` WHERE `id`='%s'",vars["id"]))
	if err!=nil{
		panic(err)
	}
	for res.Next(){
		var post Article
		err=res.Scan(&post.Id, &post.Name, &post.Anons,&post.FullText, &post.Image)
		if err!=nil{
			panic(err)
		}
		showPost=post

	}
	fmt.Println("Успешно получили данных из таблицы")
	t.ExecuteTemplate(w,"articleDetail", showPost)
	defer db.Close()

}

func getUsers(w http.ResponseWriter, r *http.Request){
	t,err:=template.ParseFiles("templates/header.html","templates/users.html")
	if err!=nil{
		fmt.Fprintf(w,err.Error())
	}

	users=[]User{}

	db,err:=sql.Open("mysql","root:@tcp(127.0.0.1:3306)/golangProject")
	if err!=nil{
		panic(err)
	}
	fmt.Println("Успешно подключились к базе данных")
	

	res,err:=db.Query("SELECT * FROM `User`")
	if err!=nil{
		panic(err)
	}
	for res.Next(){
		var user User
		err=res.Scan(&user.Id, &user.Name, &user.Password)
		if err!=nil{
			panic(err)
		}
		users=append(users, user)

	}
	fmt.Println("Успешно получили данных из таблицы")
	t.ExecuteTemplate(w,"users", users)
	defer db.Close()
}

func getCustomers(w http.ResponseWriter, r *http.Request){
	t,err:=template.ParseFiles("templates/header.html","templates/customers.html")
	if err!=nil{
		fmt.Fprintf(w,err.Error())
	}

	customers=[]Customer{}

	db,err:=sql.Open("mysql","root:@tcp(127.0.0.1:3306)/golangProject")
	if err!=nil{
		panic(err)
	}
	fmt.Println("Успешно подключились к базе данных")
	

	res,err:=db.Query("SELECT * FROM `Customer`")
	if err!=nil{
		panic(err)
	}
	for res.Next(){
		var customer Customer
		err=res.Scan(&customer.Id, &customer.Name, &customer.Surname,&customer.Age, &customer.Sex, &customer.CardNumber)
		if err!=nil{
			panic(err)
		}
		customers=append(customers, customer)

	}
	fmt.Println("Успешно получили данных из таблицы")
	t.ExecuteTemplate(w,"customers", customers)
	defer db.Close()
}

func handleFunc(){
	rtr:=mux.NewRouter()
	rtr.HandleFunc("/articles",getArticles).Methods("GET")
	rtr.HandleFunc("/articles/{id:[0-9]+}",getArticleDetail).Methods("GET")
	rtr.HandleFunc("/users",getUsers).Methods("GET")
	rtr.HandleFunc("/customers",getCustomers).Methods("GET")
	// rtr.HandleFunc("/",index).Methods("GET")
	// rtr.HandleFunc("/create",create).Methods("GET")
	// rtr.HandleFunc("/save_article",save_article).Methods("POST")
	// rtr.HandleFunc("/post/{id:[0-9]+}",show_post).Methods("GET")

	http.Handle("/",rtr)

	http.Handle("/static/",http.StripPrefix("/static/",http.FileServer(http.Dir("./static/")))) //чтобы обрабатывались файлы с директории статик
	
	http.ListenAndServe(":8080",nil)
}

func main(){
	handleFunc()
}