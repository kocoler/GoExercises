package main

//void HelloGo();
//#include <stdio.h>
import "C"

func main(){
    // fmt.Println("Hello, World!")
    C.HelloGo()
    //C.GoString()
    C.printf()
    C.puts(C.CString("Hello, World!"))
}

//export HelloC
func HelloC() {

}
