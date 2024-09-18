---
title: Defer behavior
tags:
    - golang
    - defer    
---

A interesting form to implement defer closing or cleaning method of a process is using the example above.

The code in defer will be evaluated and executed before the Close method. This allow a initialization of the process and the Close method will be executed when the method ends.
Check this out.

```golang
package main

import "fmt"

type (
	Service struct {
		name string
	}
)

func NewService(name string) *Service {
	fmt.Printf("Creating new service %s\n", name)
	return &Service{name: name}
}

func (s *Service) DoSomething() {
	fmt.Printf("Doing something on service %s\n", s.name)
}

func (s *Service) Close() {
	fmt.Printf("Closing service %s\n", s.name)
}

func main() {
	defer NewService("test_defer").Close()

	fmt.Println("START MAIN")
	fmt.Println("END MAIN")
}
```

```bash
❯ go run defer_t.go
Creating new service test_defer
START MAIN
END MAIN
Closing service test_defer
```