package main
import "fmt"

type Human struct {
	name string
	age int
	phone string
}

type Student struct {
	Human //匿名フィールド
	school string
	loan float32
}

type Employee struct {
	Human //匿名フィールド
	company string
	money float32
}

//Humanでmethodを定義
func (h Human) SayHi() {
	fmt.Printf("Hi, I am %s you can call me on %s\n", h.name, h.phone)
}

//HumanにSingメソッドを実装
func (h Human) Sing(lyrics string) {
	fmt.Println("La la la la...", lyrics)
}

//EmployeeのmethodでHumanのmethodを書き直す。
func (e Employee) SayHi() {
	fmt.Printf("Hi, I am %s, I work at %s. Call me on %s\n", e.name, e.company, e.phone)
}

// Interface MenはHuman, StudentおよびEmployeeに実装される
// この3つの型はこの2つのメソッドを実装するから
type Men interface {
	SayHi()
	Sing(lyrics string)
}

func main() {
	mike := Student{Human{"Mark", 25, "222-222-YYYY"}, "MIT", 0.00}
	paul := Student{Human{"Paul", 26, "111-222-CCCC"}, "Harvard", 100}
	sam := Employee{Human{"Sam", 45, "111-888-XXXX"}, "Golang Inc", 1000}
	tom := Employee{Human{"Tom", 37, "222-555-EEEE"}, "Things Ltd.", 5000}

	//Men型の変数iを定義
	var i Men

	//iにはStudentを保存できる
	i = mike
	fmt.Println("This is Mike, a Student:")
	i.SayHi()
	i.Sing("November rain")

	//iにはEmployeeを保存することもできる
	i = tom
	fmt.Println("This is Tom, an Employee:")
	i.SayHi()
	i.Sing("Born to be wild")

	//sliceのMenを定義
	fmt.Println("Let's use a slice of Men and see what happens")
	x := make([]Men, 3)
	//この3つはどれも異なる要素ですが、同じインターフェースを実装
	x[0], x[1], x[2] = paul, sam, mike

	for _, value := range x {
		value.SayHi()
	}
}