package main

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Username string
	Email    string
}

var tmpl = template.Must(template.New("form").Parse(`
<!DOCTYPE html>
<html>
<head>
	<title>User Form</title>
	<style>
		body {
			font-family: Arial, sans-serif;
			background-color: #f4f4f4;
			padding: 30px;
		}
		form {
			background: white;
			padding: 20px;
			border-radius: 10px;
			box-shadow: 0px 0px 10px rgba(0,0,0,0.1);
			width: 300px;
			margin-bottom: 20px;
		}
		input {
			width: 100%;
			margin: 10px 0;
			padding: 8px;
		}
		table {
			width: 100%;
			border-collapse: collapse;
			background: white;
			box-shadow: 0px 0px 10px rgba(0,0,0,0.1);
			border-radius: 10px;
			overflow: hidden;
		}
		th, td {
			padding: 12px;
			text-align: left;
			border-bottom: 1px solid #ddd;
		}
	</style>
</head>
<body>

	<h2>Submit Your Info</h2>
	<form method="POST">
		<input name="username" placeholder="Username" required>
		<input name="email" placeholder="Email" type="email" required>
		<input type="submit" value="Submit">
	</form>

	{{if .Success}}
		<p><strong>Data saved!</strong></p>
	{{end}}

	{{if .Users}}
	<h3>Submitted Users:</h3>
	<table>
		<tr><th>Username</th><th>Email</th></tr>
		{{range .Users}}
			<tr>
				<td>{{.Username}}</td>
				<td>{{.Email}}</td>
			</tr>
		{{end}}
	</table>
	{{end}}

</body>
</html>
`))

func main() {
	db, err := sql.Open("mysql", "root:rootpassword@tcp(db:3306)/userdb")
	if err != nil {
		log.Fatal("DB connection failed:", err)
	}
	defer db.Close()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")

		data := struct {
			Users   []User
			Success bool
		}{}

		if r.Method == http.MethodPost {
			username := r.FormValue("username")
			email := r.FormValue("email")

			_, err := db.Exec("INSERT INTO users (username, email) VALUES (?, ?)", username, email)
			if err != nil {
				http.Error(w, "DB insert failed", 500)
				log.Println(err)
				return
			}
			data.Success = true
		}

		rows, err := db.Query("SELECT username, email FROM users")
		if err != nil {
			http.Error(w, "DB select failed", 500)
			log.Println(err)
			return
		}
		defer rows.Close()

		for rows.Next() {
			var u User
			rows.Scan(&u.Username, &u.Email)
			data.Users = append(data.Users, u)
		}

		tmpl.Execute(w, data)
	})

	log.Println("Server running at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
