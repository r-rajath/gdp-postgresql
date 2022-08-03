package main

import (
  "context"
  "database/sql"
  "fmt"
  "log"

  _ "github.com/lib/pq"
)

const (
  host     = "localhost"
  port     = 5432
  user     = "rajath"
  password = "password"
  dbname   = "postgres"
)


type User struct {
	ID          int
	Username    string
	Email       string
}


func main() {
  psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
    "password=%s dbname=%s sslmode=disable",
    host, port, user, password, dbname)
  db, err := sql.Open("postgres", psqlInfo)
  if err != nil {
    panic(err)
  }
  defer db.Close()


  ctx:= context.Background()
  tx, err := db.BeginTx(ctx, nil)
  if err != nil {
    log.Fatal(err)
  }
  _, err = tx.Exec("UPDATE rapido SET salary = (salary+25000) WHERE name='Ashish'")
  if err != nil {
    // Incase we find any error in the query execution, rollback the transaction
    tx.Rollback()
    fmt.Println("\n", (err), "\n ....Transaction rollback!\n")
    return
  }
  
  result, err := tx.Exec("UPDATE rapido SET salary = (salary-15000) WHERE name='Ardra'")
  count,err := result.RowsAffected()
  if count == 0 {
    tx.Rollback()
    fmt.Println("\n", (err), "\n ....Transaction rollback!\n")
    return
  }
  if err != nil {
    // Incase we find any error in the query execution, rollback the transaction
    tx.Rollback()
    fmt.Println("\n", (err), "\n ....Transaction rollback!\n")
    return
  }
  // If no error in the query execution, commit the transaction
  tx.Commit()
  if err != nil {
    log.Fatal(err)
  } else {
    fmt.Println("\n ....Transaction committed!\n")
  }


}
