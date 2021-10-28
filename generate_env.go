package micagen

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"reflect"

	"gorm.io/gorm"
)

func CreateEnv(db *gorm.DB, mystruct interface{}) string {
	createFileEnv(mystruct)
	writeFileEnv(mystruct)
	if nameOfStruct := reflect.TypeOf(mystruct); nameOfStruct.Kind() == reflect.Ptr {
		return nameOfStruct.Elem().Name()
	} else {
		return nameOfStruct.Name()
	}
}
func createFileEnv(mystruct interface{}) (string, error) {
	name := ""
	if nameOfStruct := reflect.TypeOf(mystruct); nameOfStruct.Kind() == reflect.Ptr {
		name = nameOfStruct.Elem().Name()
	} else {
		name = nameOfStruct.Name()
	}

	fmt.Println(name)

	filepath, err := filepath.Abs("./" + ".env")
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

func writeFileEnv(mystruct interface{}) (string, error) {

	filepath, err := filepath.Abs("./" + ".env")
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
	_, err = file.WriteString("SECRET_KEY = ABCDEFGHIJKLMNOPQRSTUVWXYZ\n")
	if isError(err) {
		return "", err
	}
	_, err = file.WriteString(fmt.Sprintf("%s%s", "\n", "\n"))
	if isError(err) {
		return "", err
	}

	fmt.Printf("Create %s is successfully \n", filepath)
	return filepath, nil
}
