package models
 
import (
   "net/http"
   "sample/config"
) 

type User struct {
   ID        string    `json:"id"`
   FirstName string `json:"first_name"`
   LastName  string `json:"last_name"`
   Avatar    string `json:"avatar"`
}


func GetUser(r *http.Request, userId string) (User, error) {
   user := User{} //Get user struct
   db   := config.DbInit() // Initialize DB
   row  := db.QueryRow("SELECT * FROM users WHERE id = $1", userId)
   err  := row.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Avatar) //Mapping of DB result to struct

   if err != nil {
      return user, err
   }
   
   return user, nil
}

func GetUsers(r *http.Request) ([]*User, error) {
   db   := config.DbInit() // Initialize DB
   rows, err  := db.Query("SELECT * FROM users")
   users := make([]*User, 0)

   if err != nil {
      return users, err
   }

   for rows.Next() {
      user := new(User)

      err  = rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Avatar) //Mapping of DB result to struct

      users = append(users, user)
   }   
   
   return users, nil
}