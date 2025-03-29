package main

import "fmt"

// 同一包 package main 只能有一個 main() 函數
// 所以 /week1 底下的任一程式都不能有 main() 函數
func main() {
	// Hello, World!%，由於沒有換行會出現 ％，終端機中的 % 並不是程式實際印出的內容，而是你的終端機提示符號（prompt）。
	fmt.Print("Hello, World!")
	fmt.Println("Hello, World!")
}

// 執行程式
// go run ./HelloWorld.go
// 這樣會先編譯 然後再執行

// 打包程式
// go build ./HelloWorld.go
// 這樣會直接打包成執行檔 .exe(windows) .app(mac)
// 這樣就可以直接執行程式 ./HelloWorld

// { 這符號不能單獨一行，否則會有格式上錯誤
