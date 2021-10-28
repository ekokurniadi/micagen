package micagen

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"reflect"
	"strings"

	"gorm.io/gorm"
)

func CreateHandler(db *gorm.DB, mystruct interface{}) string {
	createFolderHandler()
	createFileHandler(mystruct)
	writeFileHandler(mystruct)
	if nameOfStruct := reflect.TypeOf(mystruct); nameOfStruct.Kind() == reflect.Ptr {
		return "*" + nameOfStruct.Elem().Name()
	} else {
		return nameOfStruct.Name()
	}
}

func createFolderHandler() error {
	path := filepath.Join("./", "handler")
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

func createFileHandler(mystruct interface{}) (string, error) {
	name := ""
	if nameOfStruct := reflect.TypeOf(mystruct); nameOfStruct.Kind() == reflect.Ptr {
		name = nameOfStruct.Elem().Name()
	} else {
		name = nameOfStruct.Name()
	}

	fmt.Println(name)

	filepath, err := filepath.Abs("./handler" + "/" + strings.ToLower(name) + ".go")
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

func writeFileHandler(mystruct interface{}) (string, error) {

	name := ""
	if nameOfStruct := reflect.TypeOf(mystruct); nameOfStruct.Kind() == reflect.Ptr {
		name = nameOfStruct.Elem().Name()
	} else {
		name = nameOfStruct.Name()
	}

	filepath, err := filepath.Abs("./handler" + "/" + strings.ToLower(name) + ".go")
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
	_, err = file.WriteString("package handler")
	if isError(err) {
		return "", err
	}
	_, err = file.WriteString(fmt.Sprintf("%s%s", "\n", "\n"))
	if isError(err) {
		return "", err
	}

	_, err = file.WriteString(fmt.Sprintf("type %s %s %s", strings.ToLower(name)+"Handler", "struct {", "\n"))
	if isError(err) {
		return "", err
	}
	_, err = file.WriteString("service service." + name + "Service")
	if isError(err) {
		return "", err
	}
	_, err = file.WriteString(fmt.Sprintf("%s%s", "}", "\n"))
	if isError(err) {
		return "", err
	}

	_, err = file.WriteString(fmt.Sprintf("func %s %s", "New"+name+"Handler(service service."+name+"Service)", "*"+strings.ToLower(name)+"Handler {\n"))
	if isError(err) {
		return "", err
	}

	_, err = file.WriteString("return " + "&" + strings.ToLower(name) + "Handler{service}\n")
	if isError(err) {
		return "", err
	}
	_, err = file.WriteString("}\n")
	if isError(err) {
		return "", err
	}

	_, err = file.WriteString("func (h *" + strings.ToLower(name) + "Handler) Get" + name + "(c *gin.Context) {\n")
	if isError(err) {
		return "", err
	}
	_, err = file.WriteString("var input input.InputID" + name + "\n")
	if isError(err) {
		return "", err
	}

	_, err = file.WriteString("err := c.ShouldBindUri(&input)\n")
	if isError(err) {
		return "", err
	}

	_, err = file.WriteString("if err != nil {\n")
	if isError(err) {
		return "", err
	}

	_, err = file.WriteString("response :=helper.ApiResponse(\"Failed to get " + name + "\",http.StatusBadRequest,\"error\",nil)\n")
	if isError(err) {
		return "", err
	}

	_, err = file.WriteString("c.JSON(http.StatusBadRequest,response)\n")
	if isError(err) {
		return "", err
	}

	_, err = file.WriteString("return\n")
	if isError(err) {
		return "", err
	}

	_, err = file.WriteString("}\n")
	if isError(err) {
		return "", err
	}

	_, err = file.WriteString("}\n")
	if isError(err) {
		return "", err
	}

	// Save file changes.
	err = file.Sync()
	if isError(err) {
		return "", err
	}

	fmt.Printf("Rewrite file is successfully \n")
	return filepath + "Handler", nil
}
