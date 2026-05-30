package main

import "fmt"

// Student 学生实体类
type Student struct {
	ID   int    // 学号
	Name string // 姓名
	Age  int    // 年龄
}

// StudentService 学生服务类
type StudentService struct {
	studentStore map[int]Student
}

// NewStudentService 初始化学生服务
func NewStudentService() *StudentService {
	return &StudentService{
		studentStore: make(map[int]Student),
	}
}

// AddStudent 添加学生
func (s *StudentService) AddStudent(student Student) {
	s.studentStore[student.ID] = student
	fmt.Printf("[StudentService] 已添加学生：%s（学号：%d）\n", student.Name, student.ID)
}

// GetStudent 根据学号查询学生
func (s *StudentService) GetStudent(id int) (Student, bool) {
	student, exists := s.studentStore[id]
	return student, exists
}

// ListAllStudents 查询所有学生
func (s *StudentService) ListAllStudents() []Student {
	list := make([]Student, 0, len(s.studentStore))
	for _, student := range s.studentStore {
		list = append(list, student)
	}
	return list
}