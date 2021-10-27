package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"time"

	"github.com/ekokurniadi/indodate"
	"gorm.io/gorm"
)

func CreateService(db *gorm.DB, mystruct interface{}) string {
	createFolderService()
	createFileService(mystruct)
	writeFileService(mystruct)
	if nameOfStruct := reflect.TypeOf(mystruct); nameOfStruct.Kind() == reflect.Ptr {
		return "*" + nameOfStruct.Elem().Name()
	} else {
		return nameOfStruct.Name()
	}
}

func createFolderService() error {
	path := filepath.Join("./", "service")
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

func createFileService(mystruct interface{}) (string, error) {
	name := ""
	if nameOfStruct := reflect.TypeOf(mystruct); nameOfStruct.Kind() == reflect.Ptr {
		name = nameOfStruct.Elem().Name()
	} else {
		name = nameOfStruct.Name()
	}

	fmt.Println(name)

	filepath, err := filepath.Abs("./service" + "/" + strings.ToLower(name) + ".go")
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

func writeFileService(mystruct interface{}) (string, error) {

	name := ""
	if nameOfStruct := reflect.TypeOf(mystruct); nameOfStruct.Kind() == reflect.Ptr {
		name = nameOfStruct.Elem().Name()
	} else {
		name = nameOfStruct.Name()
	}

	filepath, err := filepath.Abs("./service" + "/" + strings.ToLower(name) + ".go")
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
	_, err = file.WriteString("package service")
	if isError(err) {
		return "", err
	}
	_, err = file.WriteString(fmt.Sprintf("%s%s", "\n", "\n"))
	if isError(err) {
		return "", err
	}

	_, err = file.WriteString(fmt.Sprintf("type %s %s %s", name+"Service", "interface {", "\n"))
	if isError(err) {
		return "", err
	}

	_, err = file.WriteString(fmt.Sprintf("%sServiceGetAll() (%s %s)%s", name, "[]entity."+name, ",error", "\n"))
	if isError(err) {
		return "", err
	}
	_, err = file.WriteString(fmt.Sprintf("%sServiceGetByID(%s %s) (%s %s)%s", name, "inputID", "input.InputID"+name, "entity."+name, ",error", "\n"))
	if isError(err) {
		return "", err
	}
	_, err = file.WriteString(fmt.Sprintf("%sServiceCreate(%s %s) (%s %s)%s", name, "input", "input."+name+"Input", "entity."+name, ",error", "\n"))
	if isError(err) {
		return "", err
	}
	_, err = file.WriteString(fmt.Sprintf("%sServiceUpdate(%s %s,%s %s) (%s %s)%s", name, "inputID", "input.InputID"+name, "inputData", "input."+name+"Input", "entity."+name, ",error", "\n"))
	if isError(err) {
		return "", err
	}
	_, err = file.WriteString(fmt.Sprintf("%sServiceDeleteByID(%s %s) (%s %s)%s", name, "inputID", "input.InputID"+name, "bool", ",error", "\n"))
	if isError(err) {
		return "", err
	}

	_, err = file.WriteString("}\n")
	if isError(err) {
		return "", err
	}

	_, err = file.WriteString(fmt.Sprintf("type %s %s %s", strings.ToLower(name)+"Service", "struct {", "\n"))
	if isError(err) {
		return "", err
	}

	_, err = file.WriteString(fmt.Sprintf("repository %s", "repository."+name+"Repository\n"))
	if isError(err) {
		return "", err
	}

	_, err = file.WriteString("}\n")
	if isError(err) {
		return "", err
	}

	_, err = file.WriteString(fmt.Sprintf("func %s %s", "New"+name+"Service(repository repository."+name+"Repository)", "*"+strings.ToLower(name)+"Service {\n"))
	if isError(err) {
		return "", err
	}

	_, err = file.WriteString("return " + "&" + strings.ToLower(name) + "Service{repository}\n")
	if isError(err) {
		return "", err
	}
	_, err = file.WriteString("}\n")
	if isError(err) {
		return "", err
	}

	_, err = file.WriteString(fmt.Sprintf("func (s *%s) %s (%s", strings.ToLower(name)+"Service", name+"ServiceCreate(input input."+name+"Input)", "entity."+name+",error) {\n"))
	if isError(err) {
		return "", err
	}

	_, err = file.WriteString(strings.ToLower(name) + ":=" + "entity." + name + "{}\n")
	if isError(err) {
		return "", err
	}
	val := reflect.ValueOf(mystruct).Elem()
	for i := 0; i < val.NumField(); i++ {
		if val.Type().Field(i).Name == "ID" || val.Type().Field(i).Type.String() == "time.Time" || val.Type().Field(i).Type.String() == "time.Time" {
			_, err = file.WriteString("")
			if isError(err) {
				return "", err
			}
		} else {
			fmt.Print("Writting ... ")
			fmt.Printf("%s%s%s%s%s", strings.ToLower(name), ".", val.Type().Field(i).Name, "=input."+val.Type().Field(i).Name, "\n")
			_, err = file.WriteString(fmt.Sprintf("%s%s%s%s%s", strings.ToLower(name), ".", val.Type().Field(i).Name, "=input."+val.Type().Field(i).Name, "\n"))
			if isError(err) {
				return "", err
			}
		}
	}

	_, err = file.WriteString("new" + name + ",err :=s.repository.Save" + name + "(" + strings.ToLower(name) + ")\n")
	if isError(err) {
		return "", err
	}

	_, err = file.WriteString("if err != nil {\n")
	if isError(err) {
		return "", err
	}
	_, err = file.WriteString("return " + "new" + name + ",err" + "\n")
	if isError(err) {
		return "", err
	}
	_, err = file.WriteString("}\n")
	if isError(err) {
		return "", err
	}
	_, err = file.WriteString("return " + "new" + name + ",nil" + "\n")
	if isError(err) {
		return "", err
	}

	_, err = file.WriteString("}\n")
	if isError(err) {
		return "", err
	}

	_, err = file.WriteString(fmt.Sprintf("func (s *"+strings.ToLower(name)+"Service"+") %sServiceUpdate(%s %s,%s %s) (%s %s)%s", name, "inputID", "input.InputID"+name, "inputData", "input."+name+"Input", "entity."+name, ",error", "{\n"))
	if isError(err) {
		return "", err
	}
	_, err = file.WriteString("" + strings.ToLower(name) + ",err:=s.repository.FindByID" + name + "(inputID.ID)\n")
	if isError(err) {
		return "", err
	}

	_, err = file.WriteString("if err != nil {\n")
	if isError(err) {
		return "", err
	}
	_, err = file.WriteString("return " + strings.ToLower(name) + ",err" + "\n")
	if isError(err) {
		return "", err
	}
	_, err = file.WriteString("}\n")
	if isError(err) {
		return "", err
	}

	val = reflect.ValueOf(mystruct).Elem()
	for i := 0; i < val.NumField(); i++ {
		if val.Type().Field(i).Name == "ID" || val.Type().Field(i).Type.String() == "time.Time" || val.Type().Field(i).Type.String() == "time.Time" {
			_, err = file.WriteString("")
			if isError(err) {
				return "", err
			}
		} else {
			fmt.Print("Writting ... ")
			fmt.Printf("%s%s%s%s%s", strings.ToLower(name), ".", val.Type().Field(i).Name, "=inputData."+val.Type().Field(i).Name, "\n")
			_, err = file.WriteString(fmt.Sprintf("%s%s%s%s%s", strings.ToLower(name), ".", val.Type().Field(i).Name, "=inputData."+val.Type().Field(i).Name, "\n"))
			if isError(err) {
				return "", err
			}
		}
	}
	_, err = file.WriteString("\n")
	if isError(err) {
		return "", err
	}
	_, err = file.WriteString("updated" + name + ",err:=s.repository.Update" + name + "(" + strings.ToLower(name) + ")\n")
	if isError(err) {
		return "", err
	}
	_, err = file.WriteString("\n")
	if isError(err) {
		return "", err
	}
	_, err = file.WriteString("if err != nil {\n")
	if isError(err) {
		return "", err
	}
	_, err = file.WriteString("return " + "updated" + name + ",err" + "\n")
	if isError(err) {
		return "", err
	}

	_, err = file.WriteString("}\n")
	if isError(err) {
		return "", err
	}

	_, err = file.WriteString("return " + "updated" + name + ",nil" + "\n")
	if isError(err) {
		return "", err
	}

	_, err = file.WriteString("}\n")
	if isError(err) {
		return "", err
	}

	_, err = file.WriteString(fmt.Sprintf("func (s *"+strings.ToLower(name)+"Service"+")%sServiceGetByID(%s %s) (%s %s)%s", name, "inputID", "input.InputID"+name, "entity."+name, ",error", "{\n"))
	if isError(err) {
		return "", err
	}

	_, err = file.WriteString("" + strings.ToLower(name) + ",err:=s.repository.FindByID" + name + "(inputID.ID)\n")
	if isError(err) {
		return "", err
	}

	_, err = file.WriteString("if err != nil {\n")
	if isError(err) {
		return "", err
	}
	_, err = file.WriteString("return " + strings.ToLower(name) + ",err" + "\n")
	if isError(err) {
		return "", err
	}
	_, err = file.WriteString("}\n")
	if isError(err) {
		return "", err
	}
	_, err = file.WriteString("return " + strings.ToLower(name) + ",nil" + "\n")
	if isError(err) {
		return "", err
	}

	_, err = file.WriteString("}\n")
	if isError(err) {
		return "", err
	}

	_, err = file.WriteString(fmt.Sprintf("func (s *"+strings.ToLower(name)+"Service"+") %sServiceGetAll() (%s %s)%s", name, "[]entity."+name, ",error", "{\n"))
	if isError(err) {
		return "", err
	}

	_, err = file.WriteString("" + strings.ToLower(name) + "s, err:=s.repository.FindAll" + name + "()\n")
	if isError(err) {
		return "", err
	}

	_, err = file.WriteString("if err != nil {\n")
	if isError(err) {
		return "", err
	}
	_, err = file.WriteString("return " + strings.ToLower(name) + "s, err" + "\n")
	if isError(err) {
		return "", err
	}
	_, err = file.WriteString("}\n")
	if isError(err) {
		return "", err
	}
	_, err = file.WriteString("return " + strings.ToLower(name) + "s, nil" + "\n")
	if isError(err) {
		return "", err
	}

	_, err = file.WriteString("}\n")
	if isError(err) {
		return "", err
	}

	_, err = file.WriteString(fmt.Sprintf("func (s *"+strings.ToLower(name)+"Service"+") %sServiceDeleteByID(%s %s) (%s %s)%s", name, "inputID", "input.InputID"+name, "bool", ",error", "{\n"))
	if isError(err) {
		return "", err
	}

	_, err = file.WriteString("_,err:=s.repository.FindByID" + name + "(inputID.ID)\n")
	if isError(err) {
		return "", err
	}

	_, err = file.WriteString("if err != nil {\n")
	if isError(err) {
		return "", err
	}
	_, err = file.WriteString("return false,err" + "\n")
	if isError(err) {
		return "", err
	}
	_, err = file.WriteString("}\n")
	if isError(err) {
		return "", err
	}

	_, err = file.WriteString("_,err=s.repository.DeleteByID" + name + "(inputID.ID)\n")
	if isError(err) {
		return "", err
	}

	_, err = file.WriteString("if err != nil {\n")
	if isError(err) {
		return "", err
	}
	_, err = file.WriteString("return false,err" + "\n")
	if isError(err) {
		return "", err
	}
	_, err = file.WriteString("}\n")
	if isError(err) {
		return "", err
	}

	_, err = file.WriteString("return true,nil" + "\n")
	if isError(err) {
		return "", err
	}

	_, err = file.WriteString("}\n")
	if isError(err) {
		return "", err
	}

	_, err = file.WriteString(fmt.Sprintf("//Generated by Micagen at %v", indodate.LetterDate(time.Now())))
	if isError(err) {
		return "", err
	}

	// Save file changes.
	err = file.Sync()
	if isError(err) {
		return "", err
	}

	fmt.Printf("Rewrite file is successfully \n")
	return filepath, nil

}
