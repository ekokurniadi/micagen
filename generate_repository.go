package micagen

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

func CreateRepository(db *gorm.DB, mystruct interface{}) string {
	createFolderRepository()
	createFileRepository(mystruct)
	writeFileRepository(mystruct)
	if nameOfStruct := reflect.TypeOf(mystruct); nameOfStruct.Kind() == reflect.Ptr {
		return "*" + nameOfStruct.Elem().Name()
	} else {
		return nameOfStruct.Name()
	}

}

func writeFileRepository(mystruct interface{}) (string, error) {

	name := ""
	if nameOfStruct := reflect.TypeOf(mystruct); nameOfStruct.Kind() == reflect.Ptr {
		name = nameOfStruct.Elem().Name()
	} else {
		name = nameOfStruct.Name()
	}

	filepath, err := filepath.Abs("./repository" + "/" + strings.ToLower(name) + ".go")
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
	_, err = file.WriteString("package repository")
	if isError(err) {
		return "", err
	}
	_, err = file.WriteString(fmt.Sprintf("%s%s", "\n", "\n"))
	if isError(err) {
		return "", err
	}

	_, err = file.WriteString(fmt.Sprintf("type %s %s %s", name+"Repository", "interface {", "\n"))
	if isError(err) {
		return "", err
	}

	_, err = file.WriteString(fmt.Sprintf("Save%s(%s %s) (%s %s)%s", name, strings.ToLower(name), "entity."+name, "entity."+name, ",error", "\n"))
	if isError(err) {
		return "", err
	}
	_, err = file.WriteString(fmt.Sprintf("Update%s(%s %s) (%s %s)%s", name, strings.ToLower(name), "entity."+name, "entity."+name, ",error", "\n"))
	if isError(err) {
		return "", err
	}
	_, err = file.WriteString(fmt.Sprintf("FindByID%s(%s %s) (%s %s)%s", name, "ID", "int", "entity."+name, ",error", "\n"))
	if isError(err) {
		return "", err
	}

	_, err = file.WriteString(fmt.Sprintf("FindAll%s() (%s %s)%s", name, "[]"+"entity."+name, ",error", "\n"))
	if isError(err) {
		return "", err
	}

	_, err = file.WriteString(fmt.Sprintf("DeleteByID%s(%s %s) (%s %s)", name, "ID", "int", "string", "error"))
	if isError(err) {
		return "", err
	}

	_, err = file.WriteString("}\n\n")
	if isError(err) {
		return "", err
	}

	_, err = file.WriteString(fmt.Sprintf("type %s %s %s", strings.ToLower(name)+"Repository", "struct {", "\n"))
	if isError(err) {
		return "", err
	}
	_, err = file.WriteString("db *gorm.DB \n")
	if isError(err) {
		return "", err
	}
	_, err = file.WriteString("} \n\n")
	if isError(err) {
		return "", err
	}

	_, err = file.WriteString(fmt.Sprintf("func %s %s %s", "New"+name+"Repository(db *gorm.DB)", "*"+strings.ToLower(name)+"Repository"+"{", "\n"))
	if isError(err) {
		return "", err
	}
	_, err = file.WriteString("return &" + strings.ToLower(name) + "Repository" + "{db}")
	if isError(err) {
		return "", err
	}
	_, err = file.WriteString("}\n\n")
	if isError(err) {
		return "", err
	}

	_, err = file.WriteString("func (r *" + strings.ToLower(name) + "Repository" + ") Save" + name + "(" + strings.ToLower(name) + " entity." + name + ") (entity." + name + ",error) {\n")
	if isError(err) {
		return "", err
	}

	_, err = file.WriteString("err:= r.db.Create(&" + strings.ToLower(name) + ").Error\n")
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
	_, err = file.WriteString("\n")
	if isError(err) {
		return "", err
	}

	_, err = file.WriteString("}\n")
	if isError(err) {
		return "", err
	}

	_, err = file.WriteString("func (r *" + strings.ToLower(name) + "Repository" + ") FindByID" + name + "(ID int) (entity." + name + ",error) {\n")
	if isError(err) {
		return "", err
	}
	_, err = file.WriteString("var " + strings.ToLower(name) + " entity." + name + "\n")
	if isError(err) {
		return "", err
	}
	_, err = file.WriteString("err:= r.db.Where(\"id = ? \",ID).Find(&" + strings.ToLower(name) + ").Error\n")
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
	_, err = file.WriteString("\n")
	if isError(err) {
		return "", err
	}

	_, err = file.WriteString("}\n")
	if isError(err) {
		return "", err
	}

	_, err = file.WriteString("func (r *" + strings.ToLower(name) + "Repository" + ") Update" + name + "(" + strings.ToLower(name) + " entity." + name + ") (entity." + name + ",error) {\n")
	if isError(err) {
		return "", err
	}

	_, err = file.WriteString("err:= r.db.Save(&" + strings.ToLower(name) + ").Error\n")
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
	_, err = file.WriteString("\n")
	if isError(err) {
		return "", err
	}

	_, err = file.WriteString("}\n")
	if isError(err) {
		return "", err
	}

	_, err = file.WriteString("func (r *" + strings.ToLower(name) + "Repository" + ") FindAll" + name + "() ([]entity." + name + ",error) {\n")
	if isError(err) {
		return "", err
	}
	_, err = file.WriteString("var " + strings.ToLower(name) + "s  []entity." + name + "\n")
	if isError(err) {
		return "", err
	}
	_, err = file.WriteString("err:= r.db.Find(&" + strings.ToLower(name) + "s).Error\n")
	if isError(err) {
		return "", err
	}
	_, err = file.WriteString("if err != nil {\n")
	if isError(err) {
		return "", err
	}
	_, err = file.WriteString("return " + strings.ToLower(name) + "s ,err" + "\n")
	if isError(err) {
		return "", err
	}
	_, err = file.WriteString("}\n")
	if isError(err) {
		return "", err
	}
	_, err = file.WriteString("return " + strings.ToLower(name) + "s ,nil" + "\n")
	if isError(err) {
		return "", err
	}
	_, err = file.WriteString("\n")
	if isError(err) {
		return "", err
	}

	_, err = file.WriteString("}\n")
	if isError(err) {
		return "", err
	}

	_, err = file.WriteString("func (r *" + strings.ToLower(name) + "Repository" + ") DeleteByID" + name + "(ID int) (entity." + name + ",error) {\n")
	if isError(err) {
		return "", err
	}
	_, err = file.WriteString("var " + strings.ToLower(name) + " entity." + name + "\n")
	if isError(err) {
		return "", err
	}
	_, err = file.WriteString("err:= r.db.Where(\"id = ? \",ID).Delete(&" + strings.ToLower(name) + ").Error\n")
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
	_, err = file.WriteString("\n")
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

func createFileRepository(mystruct interface{}) (string, error) {
	name := ""
	if nameOfStruct := reflect.TypeOf(mystruct); nameOfStruct.Kind() == reflect.Ptr {
		name = nameOfStruct.Elem().Name()
	} else {
		name = nameOfStruct.Name()
	}

	fmt.Println(name)

	filepath, err := filepath.Abs("./repository" + "/" + strings.ToLower(name) + ".go")
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
	return filepath + "Repository", nil
}

func createFolderRepository() error {
	path := filepath.Join("./", "repository")
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
