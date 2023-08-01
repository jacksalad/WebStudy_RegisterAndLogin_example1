package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/register", register)
	http.HandleFunc("/login", login)

	fmt.Println("服务器启动，监听在 http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", addCorsHeaders(http.DefaultServeMux)))
}

// 注册请求函数
func register(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// 获取用户名和密码
	username := r.FormValue("username")
	password := r.FormValue("password")

	// 用户注册
	err := UserRegister(username, password)
	if err != nil {
		http.Error(w, "00", http.StatusUnauthorized)
		return
	}

	// 返回注册成功的信息
	fmt.Fprintf(w, `{"username": "%s"}`, username)
}

// 登录请求函数
func login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// 获取用户名和密码
	username := r.FormValue("username")
	password := r.FormValue("password")

	isValid, err := UserCheck(username, password)

	// 登录失败，返回错误信息
	if err != nil {
		if err.Error() == "user not found" {
			http.Error(w, "01", http.StatusUnauthorized)
		} else if err.Error() == "user not found" {
			http.Error(w, "02", http.StatusUnauthorized)
		} else {
			http.Error(w, "03", http.StatusUnauthorized)
		}
		return
	}

	if !isValid {
		fmt.Fprintf(w, `{"error": "%s"}`, "登录失败")
	}

	// 登录成功，返回用户信息
	fmt.Fprintf(w, `{"username": "%s"}`, username)
}

// 添加CORS头信息的中间件
func addCorsHeaders(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 允许所有来源的请求
		w.Header().Set("Access-Control-Allow-Origin", "*")
		// 允许特定的请求头
		w.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
		// 允许特定的请求方法
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, DELETE")

		// 对于OPTIONS请求，直接返回，不进入后续处理
		if r.Method == http.MethodOptions {
			return
		}

		// 继续处理请求
		handler.ServeHTTP(w, r)
	})
}
