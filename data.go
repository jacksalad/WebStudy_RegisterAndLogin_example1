package main

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/crypto/bcrypt"
)

// sqlite数据库文件名
const dbFile = "RegisterAndLogin.db"

// 创建sqlite数据库文件
func DatabaseCreate() {
	// 连接数据库
	db, err := sql.Open("sqlite3", dbFile)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// 创建用户表
	createTableSQL := `
		CREATE TABLE IF NOT EXISTS users (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			username TEXT NOT NULL,
			passwordHash TEXT NOT NULL
		);
	`

	_, err = db.Exec(createTableSQL)
	if err != nil {
		log.Fatal(err)
	}

}

// 数据库中用户注册函数
func UserRegister(username, password string) error {
	// 连接数据库
	db, err := sql.Open("sqlite3", dbFile)
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer db.Close()

	// 查询用户是否存在
	query := "SELECT COUNT(*) FROM users WHERE username = ?"
	var count int
	err = db.QueryRow(query, username).Scan(&count)
	if err != nil {
		return err
	}
	if count > 0 {
		return errors.New("user has registered")
	}

	// 注册一个用户并将密码进行哈希
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
		return err
	}

	// 将用户名和哈希后的密码存入数据库
	insertUserSQL := "INSERT INTO users (username, passwordHash) VALUES (?, ?)"
	_, err = db.Exec(insertUserSQL, username, hashedPassword)
	if err != nil {
		log.Fatal(err)
		return err
	}

	fmt.Println("用户表创建成功，并成功添加一个用户！")
	return nil
}

// 数据库中用户登录查询函数
func UserCheck(user, password string) (bool, error) {
	// 连接数据库
	db, err := sql.Open("sqlite3", dbFile)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// 查询用户
	query := "SELECT passwordHash FROM users WHERE username = ?"
	var hashedPassword string
	err = db.QueryRow(query, user).Scan(&hashedPassword)
	if err == sql.ErrNoRows {
		return false, errors.New("user not found") // 用户名错误
	} else if err != nil {
		log.Fatal(err)
	}

	// 验证密码
	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return false, errors.New("password error") // 密码错误
	}

	return true, nil // 用户名和密码正确
}
