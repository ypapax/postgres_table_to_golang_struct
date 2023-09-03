package main

import (
	"fmt"
	"github.com/iancoleman/strcase"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/ypapax/logrus_conf"
	"log"
	"os"
	"regexp"
	"strings"
)

func init() {
	log.SetFlags(log.Llongfile | log.LstdFlags)
	if err := logrus_conf.PrepareFromEnv("postgres_table_to_golang_struct"); err != nil {
		log.Printf("couldn't prepare logrus_conf: %+v", err)
	}
}

func main() {
	if err := func() error {
		if len(os.Args) < 2 {
			logrus.Warnf("missing first arg: filename with postgres table creation sql")
			return nil
		}
		fileName := os.Args[1]
		logrus.Infof("reading postgres create table expression from file %+v", fileName)
		b, err := os.ReadFile(fileName)
		if err != nil {
			return errors.WithStack(err)
		}
		golangStructExr, err := postgresTableToGolangStruct(string(b))
		if err != nil {
			return errors.WithStack(err)
		}
		goStructFileName := fileName + ".golung.generated.struct"
		if err := os.WriteFile(goStructFileName, []byte(golangStructExr), 0666); err != nil {
			return errors.WithStack(err)
		}
		logrus.Infof("golangStructExr is written in file: %+v", goStructFileName)
		logrus.Infof("to copy to clipboard run: pbcopy < %+v", goStructFileName)
		return nil
	}(); err != nil {
		logrus.Errorf("%+v", err)
	}
}

const nextLine = "\n"

func postgresTableToGolangStruct(postgresTableCreatExpr string) (string, error) {
	postgresTableCreatExpr = strings.TrimSpace(postgresTableCreatExpr)
	if len(postgresTableCreatExpr) == 0 {
		return "", errors.Errorf("missing postgresTableCreatExpr expression")
	}
	lines := strings.Split(postgresTableCreatExpr, nextLine)
	tableNameExpr := lines[0] // CREATE TABLE IF NOT EXISTS income_statement
	if len(tableNameExpr) == 0 {
		return "", errors.Errorf("couldn't parse table name")
	}
	tableNameParts := strings.Split(tableNameExpr, " ")
	tableName := strings.TrimSpace(tableNameParts[len(tableNameParts)-1])
	structName := strcase.ToCamel(tableName)
	var resultLines []string
	resultLines = append(resultLines, "type "+structName+" struct {")
	resultLines = append(resultLines, getStructLine("tableName", "struct{}", fieldAttr{"-", tableName + ",discard_unknown_columns"}))
	fieldLines := lines[1:]
	for _, fl := range fieldLines {
		sl := postgresColumnLineToGolangStructField(fl)
		if len(sl) == 0 {
			continue
		}
		resultLines = append(resultLines, sl)
	}
	resultLines = append(resultLines, "}")
	return strings.Join(resultLines, nextLine), nil
}

var wordsRegex = regexp.MustCompile(`[\s,]+`)
var ignorePrefix = []string{"CONSTRAINT", "--"}

func postgresColumnLineToGolangStructField(s string) string {
	s = strings.TrimSpace(s)
	if len(s) == 0 {
		return ""
	}
	//-- calculated:
	for _, ignorePr := range ignorePrefix {
		if strings.HasPrefix(s, ignorePr) {
			return ""
		}
	}

	//interestIncome numeric not null,

	words := wordsRegex.Split(s, -1)
	if len(words) < 2 {
		return ""
	}
	postgresField := strings.TrimSpace(words[0])
	fieldType := strings.TrimSpace(words[1])
	if len(postgresField) == 0 {
		return ""
	}
	if len(fieldType) == 0 {
		return ""
	}
	fieldNameGolang := postgresColumnNameToGolangFieldName(postgresField)
	fieldTypeGolang, err := postgresColumnTypeToGolangType(fieldType)
	if err != nil {
		logrus.Warnf("couldn't convert line '%+v': %+v", s, err)
	}
	return getStructLine(fieldNameGolang, fieldTypeGolang, fieldAttr{strcase.ToLowerCamel(postgresField), postgresField})
}

type fieldAttr struct {
	Json string
	Pg   string
}

func getStructLine(fieldName, fieldType string, attrs fieldAttr) string {
	aa := []string{
		fmt.Sprintf(`json:"%+v"`, attrs.Json),
		fmt.Sprintf(`pg:"%+v"`, attrs.Pg),
	}
	return fmt.Sprintf(`    %+v %+v %+v`, fieldName, fieldType, "`"+strings.Join(aa, " ")+"`")
}

func postgresColumnNameToGolangFieldName(ps string) string {
	//r := capitalize(ps)
	r := strcase.ToCamel(ps)
	switch r {
	case "Id":
		return "ID"
	case "Url":
		return "URL"
	}

	return r
}

func postgresColumnTypeToGolangType(psType string) (string, error) {
	v, ok := typesMap[strings.ToLower(psType)]
	if !ok {
		return "", errors.Errorf("postgres type '%+v' is not supported", psType)
	}
	return v, nil
}

var typesMap = map[string]string{
	"numeric":   "float64",
	"double":    "float64",
	"bigserial": "int",
	"text":      "string",
	"int":       "int",
	"timestamp": "time.Time",
}
