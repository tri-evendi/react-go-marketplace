package database

import (
	"backend/config"
	"backend/models"
	"log"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func PopulateUsers() {
	const usercount int = 4
	var firstnames 	 = [usercount]string{"Admin", "Operator", "Jamie", "User"}
	var lastnames 	= [usercount]string{"Market", "Market", "Jones", "Tester"}
	var roles 		= [usercount]string{"admin", "operator", "user", "user"}
	var emails 		= [usercount]string{"admin@admin.com", "operator@operator.com", "jones@gmail.com", "user@user.com"}
	var basePassword = "password"
	var users 		= []models.User{}

	for i := 0; i < usercount; i++ {
		password := []byte(basePassword)
		hash, _ := bcrypt.GenerateFromPassword(password, config.PasswordCost)
		users = append(users, models.User{
			FirstName: 	firstnames[i],
			LastName:  	lastnames[i],
			Email:     	emails[i],
			Password:  	hash,
			Role:  		roles[i],
		})
	}
	result := DB.Create(&users)
	log.Println(result.Error, result.RowsAffected)
}

func PopulateProducts() {
	const productcount int = 6
	var names 	 		= [productcount]string{"Voucher 1", "Voucher 2", "Voucher 3", "Voucher 4", "Voucher 5", "Voucher 6"}
	var descriptions 	= [productcount]string{
		"Some quick example text to build on the card title and make up the bulk of the card's content.", 
		"Some quick example text to build on the card title and make up the bulk of the card's content.",
		"Some quick example text to build on the card title and make up the bulk of the card's content.", 
		"Some quick example text to build on the card title and make up the bulk ofthe card's content.",
		"Some quick example text to build on the card title and make up the bulk of the card's content.",
		"Some quick example text to build on the card title and make up the bulk of the card's content.",}
	var category_ids 	= [productcount]uint{1, 2, 3, 4, 1, 2}
	var prices 			= [productcount]float64{50000, 40000, 100000, 30000, 50000, 40000}
	var images 			= [productcount]string{"misc/bgimage.jpg", "misc/bgimage.jpg", "misc/bgimage.jpg", "misc/bgimage.jpg", "misc/bgimage.jpg", "misc/bgimage.jpg"}
	var products 		= []models.Product{}

	for i := 0; i < productcount; i++ {
		products = append(products, models.Product{
			Name: 			names[i],
			Description:  	descriptions[i],
			CategoryID: 	category_ids[i],
			Price:  		prices[i],
			ImagePath:  	images[i],
		})
	}
	result := DB.Create(&products)
	log.Println(result.Error, result.RowsAffected)
}

func PopulateCategories() {
	const categoriescount int = 4
	var names 	 		= [categoriescount]string{"Categori 1", "Categori 2", "Categori 3", "Categori 4"}
	var descriptions 	= [categoriescount]string{"Lorem ipsum dummy...", "Lorem ipsum dummy...", "Lorem ipsum dummy...", "Lorem ipsum dummy..."}
	var categories 		= []models.Category{}

	for i := 0; i < categoriescount; i++ {
		categories = append(categories, models.Category{
			Name: 			names[i],
			Description:  	descriptions[i],
		})
	}
	result := DB.Create(&categories)
	log.Println(result.Error, result.RowsAffected)
}

func PopulateCodes() {
	timeWithoutNanoseconds := time.Date(2022, 8, 1, 12, 13, 14, 0, time.UTC)
	const codescount int = 8
	var codenumbers 	= [codescount]string{"XYZ-ABCD-HIJK", "ABC-WXYZ-LMNO", "DEF-OPQR-ABCD", "HIJ-LMNO-PQRS", "PQR-STU-VWXYZ", "STU-VWXY-LMNO", "VWX-YZ-ABCD", "YZ-ABCD-HIJK"}
	var product_ids 	= [codescount]uint{1, 2, 3, 4, 2, 2, 3, 4}
	var expireds 		= [codescount]time.Time{timeWithoutNanoseconds, timeWithoutNanoseconds, timeWithoutNanoseconds, timeWithoutNanoseconds, timeWithoutNanoseconds, timeWithoutNanoseconds, timeWithoutNanoseconds, timeWithoutNanoseconds}
	var availabilities 	= [codescount]bool{true, true, true,true, true, true, true, true}
	var codes 			= []models.Code{}

	for i := 0; i < codescount; i++ {
		codes = append(codes, models.Code{
			Code: 			codenumbers[i],
			ProductID: 		product_ids[i],
			DateExpired:  	expireds[i],
			IsAvailable:  	availabilities[i],
		})
	}
	result := DB.Create(&codes)
	log.Println(result.Error, result.RowsAffected)
}



func PopulateDB() {
	// To be run only when database doesn't exists.
	PopulateUsers()
	PopulateCategories()
	PopulateProducts()
	PopulateCodes()
}
