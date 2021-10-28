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

	filename, err := os.Create(filepath)

	if err != nil {
		log.Fatal("Cannot create a file please check your directory again")
		return filename.Name(), err
	}

	fmt.Printf("Create %s is successfully \n", filename.Name())
	return filepath, nil
}
