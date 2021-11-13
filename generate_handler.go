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

	_, err = file.WriteString("" + strings.ToLower(name) + "Detail,err:= h.service." + name + "ServiceGetByID(input)\n")
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

	_, err = file.WriteString("response :=helper.ApiResponse(\"Detail " + name + "\",http.StatusOK,\"success\",formatter.Format" + name + "(" + strings.ToLower(name) + "Detail))\n")
	if isError(err) {
		return "", err
	}

	_, err = file.WriteString("c.JSON(http.StatusOK,response)\n")
	if isError(err) {
		return "", err
	}

	_, err = file.WriteString("}\n")
	if isError(err) {
		return "", err
	}
	_, err = file.WriteString("\n")
	if isError(err) {
		return "", err
	}

	_, err = file.WriteString("func (h *" + strings.ToLower(name) + "Handler) Get" + name + "s" + "(c *gin.Context) {\n")
	if isError(err) {
		return "", err
	}

	_, err = file.WriteString(strings.ToLower(name) + "s,err :=h.service." + name + "ServiceGetAll()\n")
	if isError(err) {
		return "", err
	}
	_, err = file.WriteString("if err !=nil {\n")
	if isError(err) {
		return "", err
	}

	_, err = file.WriteString("response :=helper.ApiResponse(\"Failed to get " + name + "s\",http.StatusBadRequest,\"error\",nil)\n")
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

	_, err = file.WriteString("response :=helper.ApiResponse(\"Detail " + name + "s\",http.StatusOK,\"success\",formatter.Format" + name + "s(" + strings.ToLower(name) + "s" + "))\n")
	if isError(err) {
		return "", err
	}

	_, err = file.WriteString("c.JSON(http.StatusOK,response)\n")
	if isError(err) {
		return "", err
	}

	_, err = file.WriteString("}\n")
	if isError(err) {
		return "", err
	}
	_, err = file.WriteString("\n")
	if isError(err) {
		return "", err
	}

	_, err = file.WriteString("func (h *" + strings.ToLower(name) + "Handler) Create" + name + "(c *gin.Context) {\n")
	if isError(err) {
		return "", err
	}
	_, err = file.WriteString("var input input." + name + "Input\n")
	if isError(err) {
		return "", err
	}

	_, err = file.WriteString("err :=c.ShouldBindJSON(&input)\n")
	if isError(err) {
		return "", err
	}

	_, err = file.WriteString("if err !=nil{\n")
	if isError(err) {
		return "", err
	}
	_, err = file.WriteString("errors:=helper.FormatValidationError(err)\n")
	if isError(err) {
		return "", err
	}
	_, err = file.WriteString("errorMessage :=gin.H{\"errors\":errors}\n")
	if isError(err) {
		return "", err
	}
	_, err = file.WriteString("response :=helper.ApiResponse(\"Create " + name + "failed\",http.StatusBadRequest,\"error\",errorMessage)\n")
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

	_, err = file.WriteString("new" + name + ",err :=h.service." + name + "ServiceCreate(input)\n")
	if isError(err) {
		return "", err
	}

	_, err = file.WriteString("if err !=nil{\n")
	if isError(err) {
		return "", err
	}

	_, err = file.WriteString("response :=helper.ApiResponse(\"Create " + name + "failed\",http.StatusBadRequest,\"error\",nil)\n")
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

	_, err = file.WriteString("response :=helper.ApiResponse(\"Create " + name + "\",http.StatusOK,\"success\",formatter.Format" + name + "(new" + name + "))\n")
	if isError(err) {
		return "", err
	}

	_, err = file.WriteString("c.JSON(http.StatusOK,response)\n")
	if isError(err) {
		return "", err
	}

	_, err = file.WriteString("}\n")
	if isError(err) {
		return "", err
	}

	_, err = file.WriteString("func (h *" + strings.ToLower(name) + "Handler) Update" + name + "(c *gin.Context) {\n")
	if isError(err) {
		return "", err
	}
	_, err = file.WriteString("var inputID input.InputID" + name + "\n")
	if isError(err) {
		return "", err
	}

	_, err = file.WriteString("err := c.ShouldBindUri(&inputID)\n")
	if isError(err) {
		return "", err
	}

	_, err = file.WriteString("if err !=nil {\n")
	if isError(err) {
		return "", err
	}

	_, err = file.WriteString("response :=helper.ApiResponse(\"Failed to get " + name + "s\",http.StatusBadRequest,\"error\",nil)\n")
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

	_, err = file.WriteString("var inputData input." + name + "Input\n")
	if isError(err) {
		return "", err
	}
	_, err = file.WriteString("err = c.ShouldBindJSON(&inputData)\n")
	if isError(err) {
		return "", err
	}
	_, err = file.WriteString("if err !=nil{\n")
	if isError(err) {
		return "", err
	}
	_, err = file.WriteString("errors:=helper.FormatValidationError(err)\n")
	if isError(err) {
		return "", err
	}
	_, err = file.WriteString("errorMessage :=gin.H{\"errors\":errors}\n")
	if isError(err) {
		return "", err
	}
	_, err = file.WriteString("response :=helper.ApiResponse(\"Update " + name + "failed\",http.StatusBadRequest,\"error\",errorMessage)\n")
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

	_, err = file.WriteString("updated" + name + ",err := h.service." + name + "ServiceUpdate(inputID,inputData)\n")
	if isError(err) {
		return "", err
	}

	_, err = file.WriteString("if err !=nil {\n")
	if isError(err) {
		return "", err
	}

	_, err = file.WriteString("response :=helper.ApiResponse(\"Failed to get " + name + "s\",http.StatusBadRequest,\"error\",nil)\n")
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
	_, err = file.WriteString("response :=helper.ApiResponse(\"Update " + name + "\",http.StatusOK,\"success\",formatter.Format" + name + "(updated" + name + "))\n")
	if isError(err) {
		return "", err
	}

	_, err = file.WriteString("c.JSON(http.StatusOK,response)\n")
	if isError(err) {
		return "", err
	}

	_, err = file.WriteString("}\n")
	if isError(err) {
		return "", err
	}

	_, err = file.WriteString("func (h *" + strings.ToLower(name) + "Handler) Delete" + name + "(c *gin.Context) {\n")
	if isError(err) {
		return "", err
	}
	_, err = file.WriteString("param := c.Param(\"id\")\n")
	if isError(err) {
		return "", err
	}
	_, err = file.WriteString("id,_ := strconv.Atoi(param)\n")
	if isError(err) {
		return "", err
	}
	_, err = file.WriteString("_,err := h.service." + name + "ServiceGetByID(id)\n")
	if isError(err) {
		return "", err
	}

	_, err = file.WriteString("if err !=nil {\n")
	if isError(err) {
		return "", err
	}

	_, err = file.WriteString("response :=helper.ApiResponse(\"Failed to get " + name + "s\",http.StatusBadRequest,\"error\",nil)\n")
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

	_, err = file.WriteString("_,err = h.service." + name + "ServiceDeleteByID(id)\n")
	if isError(err) {
		return "", err
	}

	_, err = file.WriteString("if err !=nil {\n")
	if isError(err) {
		return "", err
	}

	_, err = file.WriteString("response :=helper.ApiResponse(\"Failed to get " + name + "s\",http.StatusBadRequest,\"error\",nil)\n")
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
	_, err = file.WriteString("response :=helper.ApiResponse(\"Delete " + name + "\",http.StatusOK,\"success\",nil)\n")
	if isError(err) {
		return "", err
	}

	_, err = file.WriteString("c.JSON(http.StatusOK,response)\n")
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
