package main

import (
  "log"
  "github.com/raylicola/NFlaquiz/database"
  "github.com/raylicola/NFlaquiz/routes"
)

func main() {

  database.Connect()
  router := routes.GetRouter()

  if err := router.Run(":8888"); err != nil {
    log.Fatal("Server Run Failed.: ", err)
  }
}