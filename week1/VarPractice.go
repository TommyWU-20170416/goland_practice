package main

import "fmt"

/**
 * 變數 種類
 */
func varType() {
	var a int = 1
	var b string = "Hello"
	var c bool = true

	fmt.Print(a, b, c)
}

/**
 * 變數 省略寫法
 */
func varInit() {
	// 省略寫法
	var s1 = "Hello String" // 省略 string 宣告
	var s2 string           // 省略賦值

	fmt.Print(s, s1, s2)
}

/**
 * 變數 初始值
 */
func varIfNoInit() {
	var a int     // 0
	var b float64 // 0.0
	var c string  // ""
	var d bool    // false
	var e *int    // nil

	fmt.Print(a, b, c, d, e)
	// 型別			初始值
	// 布林值		false
	// 數值			0
	// 字串			空字串
	// 集合與函數	 nil
}

/**
 * 變數 宣告方式
 */
var aGlobal int = 10

func varDeclare() {
	var a, b, c int = 1, 2, 3
	var x, y = 10, "Go"

	var (
		a1 int    = 10
		b1 string = "Hello"
		c1 bool   = true
	)

	// 在函數內部，可以使用 := 來聲明並初始化變數（這是 var 的簡化形式）：
	a2 := 10
	b2 := "Hello"

	fmt.Print(a, b, c, x, y, a1, b1, c1, a2, b2)
}
