package main

import "fmt"

func main(){
    p("rocky")
    wc("girls")
}

/* block comment */
func p(name string){
    fmt.Println("Hello," + name)
}

func wc(name string){
    fmt.Println("Hello, fuck you" + name)
}




// single line comment
// func p(){
    // fmt.Println("Hello world!")
// }
