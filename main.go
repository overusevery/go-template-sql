package main

import (
	"bytes"
	"fmt"
	"log"
	"text/template"

	_ "embed"
)

//go:embed sql/template.sql
var tmpl string

type StringValueInCondition struct {
	Value string
	Skip  bool
}
type Data struct {
	Id    string
	Name  StringValueInCondition
	Color []string
	Price int
}

func main() {
	t, err := template.New("select_statement").Parse(tmpl)
	if err != nil {
		log.Fatal(err)
	}
	var b bytes.Buffer
	if err = t.Execute(&b, &Data{
		Id:    "x",
		Name:  StringValueInCondition{"pen", false},
		Color: []string{"black", "red"},
		Price: 100,
	}); err != nil {
		log.Fatal(err)
	}
	queryString := b.String()
	fmt.Print(queryString)
	/*
		--output--
			SELECT
		        id,
		        name,
		        color,
		        price,
			FROM table
			WHERE
				1=1
				AND name = 'pen'
				AND color IN ('black','red')
			;
	*/

}
