package main

import "fmt"

// User 用户实体类
type User struct {
	ID       int    // 用户ID
	Username string // 用户名
	Email    string // 邮箱
}

// UserService 用户服务类
type UserService struct {
	userStore map[int]User
}

// NewUserService 初始化用户服务
func NewUserService() *UserService {
	return &UserService{
		userStore: make(map[int]User),
	}
}

// AddUser 添加用户
func (s *UserService) AddUser(user User) {
	s.userStore[user.ID] = user
	fmt.Printf("[UserService] 已添加用户：%s（邮箱：%s）\n", user.Username, user.Email)
}

// GetUser 根据ID查询用户
func (s *UserService) GetUser(id int) (User, bool) {
	user, exists := s.userStore[id]
	return user, exists
}

// ListAllUsers 查询所有用户
func (s *UserService) ListAllUsers() []User {
	list := make([]User, 0, len(s.userStore))
	for _, user := range s.userStore {
		list = append(list, user)
	}
	return list
}