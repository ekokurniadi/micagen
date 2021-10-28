package micagen

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"reflect"

	"gorm.io/gorm"
)

func CreateAuth(db *gorm.DB, mystruct interface{}) string {
	createFolderAuth()
	createFileAuth(mystruct)
	writeFileAuth(mystruct)
	if nameOfStruct := reflect.TypeOf(mystruct); nameOfStruct.Kind() == reflect.Ptr {
		return "*" + nameOfStruct.Elem().Name()
	} else {
		return nameOfStruct.Name()
	}
}

func createFolderAuth() error {
	path := filepath.Join("./", "auth")
	_, err := os.Stat(path)

	if os.IsExist(err) {
		fmt.Println("your directory is already exist but it's ok")
		return err
	}

	err = os.Mkdir(path, 0755)
	if err != nil {
		fmt.Println("your directory is already exist but it's ok")
		return err
	}

	return nil
}

func createFileAuth(mystruct interface{}) (string, error) {
	name := ""
	if nameOfStruct := reflect.TypeOf(mystruct); nameOfStruct.Kind() == reflect.Ptr {
		name = nameOfStruct.Elem().Name()
	} else {
		name = nameOfStruct.Name()
	}

	fmt.Println(name)

	filepath, err := filepath.Abs("./auth" + "/" + "service" + ".go")
	if err != nil {
		log.Fatal("error")
		return filepath, err
	}

	_, err = os.Stat(filepath)
	if os.IsExist(err) {
		fmt.Println("File is exist")
		return filepath, err
	}

	filename, err := os.Create(filepath)

	if err != nil {
		log.Fatal("Cannot create a file please check your directory again")
		return filename.Name(), err
	}

	fmt.Printf("Create %s is successfully \n", filename.Name())
	return filepath, nil
}

func writeFileAuth(mystruct interface{}) (string, error) {
	filepath, err := filepath.Abs("./auth" + "/" + "service" + ".go")
	if err != nil {
		log.Fatal("error")
		return filepath, err
	}

	file, err := os.OpenFile(filepath, os.O_RDWR, 0644)
	if isError(err) {
		return "", err
	}
	defer file.Close()

	//Write some text line-by-line to file.
	_, err = file.WriteString("package auth")
	if isError(err) {
		return "", err
	}
	_, err = file.WriteString(fmt.Sprintf("%s%s", "\n", "\n"))
	if isError(err) {
		return "", err
	}
	_, err = file.WriteString("type Service interface {\n")
	if isError(err) {
		return "", err
	}
	_, err = file.WriteString("GenerateToken(userID int) (string,error)\n")
	if isError(err) {
		return "", err
	}
	_, err = file.WriteString("ValidateToken(token string) (*jwt.Token,error)\n")
	if isError(err) {
		return "", err
	}
	_, err = file.WriteString("}\n")
	if isError(err) {
		return "", err
	}
	_, err = file.WriteString("type jwtService struct {\n")
	if isError(err) {
		return "", err
	}
	_, err = file.WriteString("}\n")
	if isError(err) {
		return "", err
	}
	_, err = file.WriteString("s := godotenv.Load()\nif s != nil { \nfmt.Println(s)\n}\n")
	if isError(err) {
		return "", err
	}
	_, err = file.WriteString("secret := os.Getenv(\"SECRET_KEY\")\n")
	if isError(err) {
		return "", err
	}

	_, err = file.WriteString("var SECRET_KEY =[]byte(secret)\n")
	if isError(err) {
		return "", err
	}

	_, err = file.WriteString("func NewService() *jwtService {\n return &jwtService{}\n}\n")
	if isError(err) {
		return "", err
	}
	_, err = file.WriteString("func (s *jwtService) GenerateToken(userID int) (string,error) {\n")
	if isError(err) {
		return "", err
	}
	_, err = file.WriteString("claim := jwt.MapClaims{}\n")
	if isError(err) {
		return "", err
	}
	_, err = file.WriteString("claim[\"user_id\"] = userID\n")
	if isError(err) {
		return "", err
	}
	_, err = file.WriteString("token := jwt.NewWithClaims(jwt.SigninMethodHS256,claim)\n")
	if isError(err) {
		return "", err
	}
	_, err = file.WriteString("signToken,err := token.SignedString(SECRET_KEY)\n")
	if isError(err) {
		return "", err
	}
	_, err = file.WriteString("if err != nil {\n return signToken,nil\n}\nreturn signToken,nil\n}\n")
	if isError(err) {
		return "", err
	}
	_, err = file.WriteString("func (s *jwtService) ValidateToken(encodedToken string) (*jwt.Token,error){\n")
	if isError(err) {
		return "", err
	}
	_, err = file.WriteString("token,err := jwt.Parse(encodedToken,func(token *jwt.Token) (interface{},error){\n")
	if isError(err) {
		return "", err
	}
	_, err = file.WriteString("_,ok := token.Method.(*jwt.SigningMethodHMAC)\n")
	if isError(err) {
		return "", err
	}
	_, err = file.WriteString("if !ok {\n")
	if isError(err) {
		return "", err
	}
	_, err = file.WriteString("return nil, errors.New(\"invalid token\")\n}\n")
	if isError(err) {
		return "", err
	}
	_, err = file.WriteString("return []byte(SECRET_KEY),nil\n")
	if isError(err) {
		return "", err
	}
	_, err = file.WriteString("})\n")
	if isError(err) {
		return "", err
	}
	_, err = file.WriteString("if err !=nil{\n return token,err \n}\n return token,nil\n}\n")
	if isError(err) {
		return "", err
	}

	fmt.Printf("Rewrite file is successfully \n")
	return filepath, nil
}
