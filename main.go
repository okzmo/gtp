package main

import (
	"errors"
	"flag"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"regexp"
	"strings"
)

var types = make(map[string]string)

func generateMapToCastTypes() {
	types["int"] = "number"
	types["int8"] = "number"
	types["int16"] = "number"
	types["int32"] = "number"
	types["int64"] = "number"
	types["uint"] = "number"
	types["uint8"] = "number"
	types["uint16"] = "number"
	types["uint32"] = "number"
	types["uint64"] = "number"
	types["string"] = "string"
	types["bool"] = "boolean"
}

func main() {
	var input, output, namespace string

	flag.StringVar(&input, "in", "", "Where is your model.go file ? Or any file that contain models.")
	flag.StringVar(&output, "out", "", "Where do you want to output the types.d.ts file.")
	flag.StringVar(&namespace, "namespace", "", "Have all your types in a namespace, great if using JSDoc.")

	flag.Parse()

	if !strings.HasSuffix(input, ".go") {
		fmt.Println("Your input file needs to be a .go file.")
		os.Exit(1)
	}

	var fileNameOutput = "types.d.ts"
	if strings.HasSuffix(output, ".d.ts") {
		fileNameOutput, output = getFileName(&output)
		output = strings.TrimSuffix(output, "/")
	}

	if _, err := os.Stat(output); errors.Is(err, os.ErrNotExist) {
		err := os.MkdirAll(output, os.ModePerm)
		if err != nil {
			panic(err)
		}
	}

	f, err := os.OpenFile(output+"/"+fileNameOutput, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	generateMapToCastTypes()

	if _, err := f.WriteString("// AUTO-GENERATED FILE FROM GTP DO NOT EDIT\n\n"); err != nil {
		panic(err)
	}

	fset := token.NewFileSet()

	node, err := parser.ParseFile(fset, input, nil, parser.ParseComments)
	if err != nil {
		panic(err)
	}

	indentGlobal := ""
	indentTypeField := ""
	if namespace != "" {
		declareNamespace := fmt.Sprintf("declare namespace %s {\n", namespace)
		if _, err := f.WriteString(declareNamespace); err != nil {
			panic(err)
		}
		indentGlobal = "	"
		indentTypeField = "		"
	}

	ast.Inspect(node, func(n ast.Node) bool {

		ts, ok := n.(*ast.TypeSpec)
		if !ok {
			return true
		}

		st, ok := ts.Type.(*ast.StructType)
		if ok {

			// Write the opening of the type
			typeStart := fmt.Sprintf(indentGlobal+"export type %s = {\n", ts.Name.Name)
			if _, err := f.WriteString(typeStart); err != nil {
				panic(err)
			}

			// Generate the fields
			generateFields(f, st, indentTypeField)

			// Close the type
			if _, err := f.WriteString(indentGlobal + "}\n"); err != nil {
				panic(err)
			}

		}

		return true
	})

	if namespace != "" {
		if _, err := f.WriteString("}"); err != nil {
			panic(err)
		}
	}
}

func generateFields(file *os.File, structType *ast.StructType, indent string) {
	for _, field := range structType.Fields.List {
		for _, fieldName := range field.Names {
			var fieldType string

			switch t := field.Type.(type) {
			case *ast.ArrayType:
				fieldType = fmt.Sprintf("%s[]", t.Elt)
			default:
				fieldType = fmt.Sprintf("%s", field.Type)
			}

			cast := fieldType
			if _, ok := types[fieldType]; ok {
				cast = types[fieldType]
			}

			fieldName := strings.ToLower(fieldName.Name) + strings.Replace(field.Comment.Text(), "\n", "", -1)

			if _, err := file.WriteString(indent + fieldName + ": " + cast + "\n"); err != nil {
				panic(err)
			}
		}
	}
}

func getFileName(output *string) (string, string) {
	re := regexp.MustCompile(`[^/]+\.d\.ts$`)
	matches := re.FindStringSubmatch(*output)

	out, _ := strings.CutSuffix(*output, matches[0])

	return matches[0], out
}
