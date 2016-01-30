package main

import(
  "github.com/loosecannon93/chittyrc/lib/chilog"
)

func main() {
  chilog.Init(chilog.INFO)
  chilog.Log(chilog.WARNING, "wow such logging")
}
