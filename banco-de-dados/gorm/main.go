package main

import (
	"gorm.io/driver/mysql"
	_ "gorm.io/driver/mysql"
	"gorm.io/gorm"
	_ "gorm.io/gorm"
)

type Product struct {
	gorm.Model
	ID           int
	Name         string
	Price        float64
	CategoryID   int
	Category     Category
	SerialNumber SerialNumber
}

type Category struct {
	gorm.Model
	ID       int
	Name     string
	Products []Product
}

type SerialNumber struct {
	gorm.Model
	ID        int
	Number    string
	ProductID uint
}

func main() {

	dsn := "root:root@tcp(localhost:3306)/banco?parseTime=true"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Product{}, &Category{}, &SerialNumber{})

	categories := []Category{
		{Name: "Eletronicos"},
		{Name: "Livros"},
		{Name: "Programacao"},
		{Name: "Perifericos"},
		{Name: "Moveis"},
		{Name: "Games"},
		{Name: "Audio"},
		{Name: "Smart Home"},
	}

	db.Create(&categories)

	products := []Product{

		{Name: "Notebook Dell", Price: 3500, CategoryID: 1, SerialNumber: SerialNumber{Number: "SN1001"}},
		{Name: "Smartphone Samsung", Price: 2200, CategoryID: 1, SerialNumber: SerialNumber{Number: "SN1002"}},
		{Name: "Tablet Lenovo", Price: 900, CategoryID: 1, SerialNumber: SerialNumber{Number: "SN1003"}},
		{Name: "Smart TV LG", Price: 2800, CategoryID: 1, SerialNumber: SerialNumber{Number: "SN1004"}},

		{Name: "Clean Code", Price: 120, CategoryID: 2, SerialNumber: SerialNumber{Number: "SN1005"}},
		{Name: "Entendendo Algoritmos", Price: 95, CategoryID: 2, SerialNumber: SerialNumber{Number: "SN1006"}},
		{Name: "Design Patterns", Price: 150, CategoryID: 2, SerialNumber: SerialNumber{Number: "SN1007"}},

		{Name: "Java Completo", Price: 200, CategoryID: 3, SerialNumber: SerialNumber{Number: "SN1008"}},
		{Name: "Go Programming", Price: 180, CategoryID: 3, SerialNumber: SerialNumber{Number: "SN1009"}},
		{Name: "Python Avancado", Price: 170, CategoryID: 3, SerialNumber: SerialNumber{Number: "SN1010"}},
		{Name: "Rust Essentials", Price: 160, CategoryID: 3, SerialNumber: SerialNumber{Number: "SN1011"}},

		{Name: "Mouse Gamer", Price: 80, CategoryID: 4, SerialNumber: SerialNumber{Number: "SN1012"}},
		{Name: "Teclado Mecânico", Price: 300, CategoryID: 4, SerialNumber: SerialNumber{Number: "SN1013"}},
		{Name: "Mousepad RGB", Price: 90, CategoryID: 4, SerialNumber: SerialNumber{Number: "SN1014"}},

		{Name: "Cadeira Gamer", Price: 1200, CategoryID: 5, SerialNumber: SerialNumber{Number: "SN1015"}},
		{Name: "Mesa Escritório", Price: 750, CategoryID: 5, SerialNumber: SerialNumber{Number: "SN1016"}},
		{Name: "Estante Livros", Price: 400, CategoryID: 5, SerialNumber: SerialNumber{Number: "SN1017"}},

		{Name: "PlayStation 5", Price: 3200, CategoryID: 6, SerialNumber: SerialNumber{Number: "SN1018"}},
		{Name: "Xbox Series X", Price: 3100, CategoryID: 6, SerialNumber: SerialNumber{Number: "SN1019"}},
		{Name: "Controle Xbox", Price: 350, CategoryID: 6, SerialNumber: SerialNumber{Number: "SN1020"}},

		{Name: "Headset Gamer", Price: 450, CategoryID: 7, SerialNumber: SerialNumber{Number: "SN1021"}},
		{Name: "Microfone USB", Price: 380, CategoryID: 7, SerialNumber: SerialNumber{Number: "SN1022"}},
		{Name: "Caixa Som JBL", Price: 600, CategoryID: 7, SerialNumber: SerialNumber{Number: "SN1023"}},

		{Name: "Alexa Echo", Price: 500, CategoryID: 8, SerialNumber: SerialNumber{Number: "SN1024"}},
		{Name: "Lampada Smart", Price: 120, CategoryID: 8, SerialNumber: SerialNumber{Number: "SN1025"}},
		{Name: "Tomada Inteligente", Price: 150, CategoryID: 8, SerialNumber: SerialNumber{Number: "SN1026"}},

		{Name: "Router WiFi 6", Price: 700, CategoryID: 1, SerialNumber: SerialNumber{Number: "SN1027"}},
		{Name: "HD Externo 2TB", Price: 550, CategoryID: 1, SerialNumber: SerialNumber{Number: "SN1028"}},
		{Name: "SSD NVMe 1TB", Price: 620, CategoryID: 1, SerialNumber: SerialNumber{Number: "SN1029"}},
		{Name: "Webcam HD", Price: 300, CategoryID: 4, SerialNumber: SerialNumber{Number: "SN1030"}},
	}

	db.Create(&products)

	// var product Product
	// db.First(&product, 1)
	// fmt.Println(product)
	// var product2 Product
	// db.First(&product2, "name = ? ", "Teclado")
	// fmt.Println(product2)

	// var category Category
	// db.First(&category, 1)
	// fmt.Println(category)

	// var products []Product
	// db.Find(&products)
	// for _, p := range products {
	// 	fmt.Println(p)
	// }

}
