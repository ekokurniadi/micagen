package micagen

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"reflect"

	"gorm.io/gorm"
)

func CreateHelper(db *gorm.DB, mystruct interface{}) string {
	createFolderHelper()
	createFileHelper(mystruct)
	writeFileHelper(mystruct)
	if nameOfStruct := reflect.TypeOf(mystruct); nameOfStruct.Kind() == reflect.Ptr {
		return nameOfStruct.Elem().Name()
	} else {
		return nameOfStruct.Name()
	}
}

func createFolderHelper() error {
	path := filepath.Join("./", "helper")
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

func createFileHelper(mystruct interface{}) (string, error) {
	name := ""
	if nameOfStruct := reflect.TypeOf(mystruct); nameOfStruct.Kind() == reflect.Ptr {
		name = nameOfStruct.Elem().Name()
	} else {
		name = nameOfStruct.Name()
	}

	fmt.Println(name)

	filepath, err := filepath.Abs("./helper" + "/" + "helper" + ".go")
	if err != nil {
		log.Fatal("error")
		return filepath, err
	}

	_, err = os.Stat(filepath)

	if os.IsExist(err) {
		fmt.Println("file is exist")
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

func writeFileHelper(mystruct interface{}) (string, error) {
	filepath, err := filepath.Abs("./helper" + "/" + "helper" + ".go")
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
	_, err = file.WriteString("package helper")
	if isError(err) {
		return "", err
	}
	_, err = file.WriteString(fmt.Sprintf("%s%s", "\n", "\n"))
	if isError(err) {
		return "", err
	}
	_, err = file.WriteString("import \"github.com/go-playground/validator/v10\"\n")
	if isError(err) {
		return "", err
	}

	_, err = file.WriteString(fmt.Sprintf("type %s %s %s", "Response", "struct {", "\n"))
	if isError(err) {
		return "", err
	}

	_, err = file.WriteString("Meta Meta  `json:\"meta\"`\n")
	if isError(err) {
		return "", err
	}
	_, err = file.WriteString("Data interface{} `json:\"data\"`\n}\n")
	if isError(err) {
		return "", err
	}
	_, err = file.WriteString(fmt.Sprintf("type %s %s %s", "Meta", "struct {", "\n"))
	if isError(err) {
		return "", err
	}

	_, err = file.WriteString("Message string `json:\"message\"`\n")
	if isError(err) {
		return "", err
	}
	_, err = file.WriteString("Code int `json:\"code\"`\n")
	if isError(err) {
		return "", err
	}
	_, err = file.WriteString("Status string `json:\"status\"`\n")
	if isError(err) {
		return "", err
	}
	_, err = file.WriteString("}\n")
	if isError(err) {
		return "", err
	}
	_, err = file.WriteString("func ApiResponse(message string,code int,status string,data interface{}) Response {\n")
	if isError(err) {
		return "", err
	}

	_, err = file.WriteString("meta :=Meta{\nMessage:message,\nCode:code,\nStatus:status,\n}\n")
	if isError(err) {
		return "", err
	}
	_, err = file.WriteString("jsonResponse :=Response{\nMeta:meta,\nData:data,\n}\n")
	if isError(err) {
		return "", err
	}

	_, err = file.WriteString("return jsonResponse\n}\n")
	if isError(err) {
		return "", err
	}

	_, err = file.WriteString("func FormatValidationError(err error) []string {\n")
	if isError(err) {
		return "", err
	}
	_, err = file.WriteString("var errors []string\n")
	if isError(err) {
		return "", err
	}
	_, err = file.WriteString("for _,e := range err.(validator.Validation.Errors) {\n")
	if isError(err) {
		return "", err
	}
	_, err = file.WriteString("errors = append(errors,e.Error())\n")
	if isError(err) {
		return "", err
	}

	_, err = file.WriteString("}\n")
	if isError(err) {
		return "", err
	}

	_, err = file.WriteString("return errors\n")
	if isError(err) {
		return "", err
	}
	_, err = file.WriteString("}\n")
	if isError(err) {
		return "", err
	}
	fmt.Printf("Rewrite file is successfully \n")
	return filepath, nil
}
