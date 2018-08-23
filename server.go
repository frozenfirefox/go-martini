package main

import "github.com/go-martini/martini"

func main() {
  m := martini.Classic()
  m.Get("/", func() string {
    return "什么!"
  })
  m.Get("/test", func() string {
    return "这是测试路由！！！！";
  })
  m.Get("/admin", func() string {
    return "这是admin路由！！！！";
  })
  m.Get("/index", func() string {
    return "这是index路由！！！！";
  })
  m.Get("/user", func() string {
    return "这是user路由！！！！";
  })
  m.Get("/goods", func() string {
    return "这是goods路由！！！！";
  })
  m.Get("/pack", func() string {
    return "这是pack路由！！！！";
  })
  m.Get("/import", func() string {
    return "这是import路由！！！！";
  })
  m.Run()
}