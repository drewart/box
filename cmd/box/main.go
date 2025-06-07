package main

import (
	"github.com/drewart/box/cli"
)

func main() {
  cli.Execute()
  /*
    action := os.Args[1]
    fmt.Println("action:", action)

    if action == "set" {

      service := os.Args[2]
      user := os.Args[3]
      password := os.Args[4]

      // set password
      err := keyring.Set(service, user, password)
      if err != nil {
          log.Fatal(err)
      }
      fmt.Println("set password")

    } else if action == "get" {

      fmt.Println("get password")
      service := os.Args[2]
      user := os.Args[3]

      // get password
      secret, err := keyring.Get(service, user)
      if err != nil {
          log.Fatal(err)
      }
      log.Println(secret)
    }
      */

}