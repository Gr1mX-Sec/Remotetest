package main

import "fmt"

// Book 书籍实体类，对应Java中的POJO
type Book struct {
	ID     int    // 书籍ID
	Title  string // 书名
	Author string // 作者
}

// BookService 书籍服务类，对应Java中的Service层
type BookService struct {
	// 用map模拟数据库存储，key是书籍ID
	bookStore map[int]Book
}

// NewBookService 初始化书籍服务
func NewBookService() *BookService {
	return &BookService{
		bookStore: make(map[int]Book),
	}
}

// AddBook 添加书籍
func (s *BookService) AddBook(book Book) {
	s.bookStore[book.ID] = book
	fmt.Printf("[BookService] 已添加书籍：《%s》（作者：%s）\n", book.Title, book.Author)
}

// GetBook 根据ID查询书籍
func (s *BookService) GetBook(id int) (Book, bool) {
	book, exists := s.bookStore[id]
	return book, exists
}

// ListAllBooks 查询所有书籍
func (s *BookService) ListAllBooks() []Book {
	list := make([]Book, 0, len(s.bookStore))
	for _, book := range s.bookStore {
		list = append(list, book)
	}
	return list
}