package main

/**
 * @Author: LFM
 * @Date: 2023/5/28 23:25
 * @Since: 1.0.0
 * @Desc: TODO
 */

import "fmt"

// UserService 单一职责原则（SRP）示例
type UserService struct{}

func (s *UserService) RegisterUser(username, password string) {
	// 执行用户注册逻辑
	fmt.Printf("用户：%s 注册成功\n", username)
}

func (s *UserService) AuthenticateUser(username, password string) bool {
	// 执行用户认证逻辑
	fmt.Printf("用户：%s 认证成功\n", username)
	return true
}

// Shape 开放封闭原则（OCP）示例
type Shape interface {
	Area() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

func CalculateArea(shape Shape) float64 {
	return shape.Area()
}

// Database 依赖倒置原则（DIP）示例
type Database interface {
	Save(data string)
}

type MySQL struct{}

func (m MySQL) Save(data string) {
	fmt.Println("Saving data to MySQL database:", data)
}

type MongoDB struct{}

func (m MongoDB) Save(data string) {
	fmt.Println("Saving data to MongoDB database:", data)
}

type DataManager struct {
	Database Database
}

func (dm DataManager) SaveData(data string) {
	dm.Database.Save(data)
}

// Printer 接口隔离原则（ISP）示例
type Printer interface {
	Print()
}

type Scanner interface {
	Scan()
}

type MultifunctionDevice interface {
	Printer
	Scanner
}

type SimplePrinter struct{}

func (sp SimplePrinter) Print() {
	fmt.Println("Printing document...")
}

type SimpleScanner struct{}

func (ss SimpleScanner) Scan() {
	fmt.Println("Scanning document...")
}

// Vehicle 里氏替换原则（LSP）示例
type Vehicle interface {
	Drive()
}

type Car struct{}

func (c Car) Drive() {
	fmt.Println("Driving a car...")
}

type ElectricCar struct{}

func (ec ElectricCar) Drive() {
	fmt.Println("Driving an electric car...")
}

// Reader 最小接口原则（POLA）示例
type Reader interface {
	Read()
}

type Writer interface {
	Write()
}

type ReaderWriter interface {
	Reader
	Writer
}

type FileReader struct{}

func (fr FileReader) Read() {
	fmt.Println("Reading file...")
}

func (fr FileReader) Write() {
	fmt.Println("Writing file...")
}

func main() {
	// 单一职责原则（SRP）示例
	userService := &UserService{}
	userService.RegisterUser("john", "123456")
	userService.AuthenticateUser("john", "123456")

	// 开放封闭原则（OCP）示例
	rectangle := Rectangle{Width: 5, Height: 3}
	circle := Circle{Radius: 2}
	fmt.Println("Rectangle area:", CalculateArea(rectangle))
	fmt.Println("Circle area:", CalculateArea(circle))

	// 依赖倒置原则（DIP）示例
	mySQL := MySQL{}
	mongoDB := MongoDB{}
	dataManagerMySQL := DataManager{Database: mySQL}
	dataManagerMongoDB := DataManager{Database: mongoDB}
	dataManagerMySQL.SaveData("Data to be saved in MySQL database")
	dataManagerMongoDB.SaveData("Data to be saved in MongoDB database")

	// 接口隔离原则（ISP）示例
	printer := SimplePrinter{}
	printer.Print()
	scanner := SimpleScanner{}
	scanner.Scan()

	// 里氏替换原则（LSP）示例
	car := Car{}
	electricCar := ElectricCar{}
	car.Drive()
	electricCar.Drive()

	// 最小接口原则（POLA）示例
	fileReader := FileReader{}
	fileReader.Read()
	fileReader.Write()
}
