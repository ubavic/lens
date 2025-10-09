package main

import (
	"fmt"

	"github.com/ubavic/lens/server"
	"github.com/ubavic/lens/view"
	"github.com/ubavic/lens/view/component"
	"github.com/ubavic/lens/view/form"
	"github.com/ubavic/lens/view/layout"
	"github.com/ubavic/lens/view/table"
)

func main() {

	table := table.Table{
		Columns: []table.Column{
			{
				Name:        "Name",
				Description: "bla bla bla das",
				SortedBy:    -1,
				Filter:      "sda",
			},
			{
				Name:        "Address",
				Description: "bla bla bla das",
			},
			{
				Name: "City",
			},
			{
				Name: "Balance",
			},
			{
				Name: "Status",
			},
		},
	}

	group1 := layout.Group{
		Name: "grupa1",
		Columns: []layout.Column{
			{
				Name: "column",
			},
		},
	}

	group2 := layout.Group{
		Name: "grupa2",
		Columns: []layout.Column{
			{
				Name: "column1",
			},
			{
				Name: "column2",
			},
		},
	}

	group3 := layout.Group{
		Name: "grupa2",
		Columns: []layout.Column{
			{
				Name: "column1",
			},
			{
				Name: "column2",
			},
			{
				Name: "column3",
			},
		},
	}

	pageCustomers := view.Page{
		Id:   "customer",
		Name: "Customers",
		Components: []component.Component{
			&table,
			&group1,
			&group2,
			&group3,
		},
	}

	group4 := layout.Group{
		Name: "Grupa2",
		Columns: []layout.Column{
			{
				Name: "column1",
				Elements: []component.Component{
					&form.Field{
						Name:        "Field1",
						Type:        form.FieldText,
						Description: "Name2 of the field",
						Validate: func(s string) error {
							if s != "a" {
								return fmt.Errorf("not a")
							}

							return nil
						},
					},
					&form.Field{
						Name:        "Field2",
						Type:        form.FieldText,
						Description: "Name2 of the field",
					},
					&form.Field{
						Name:        "Field3",
						Type:        form.FieldText,
						Description: "Name2 of the field",
					},
				},
			},
			{
				Name: "column2",
				Elements: []component.Component{
					&form.Field{

						Name:        "Field4",
						Type:        form.FieldText,
						Description: "Name2 of the field",
					},
					&form.Field{

						Name:        "Field5",
						Type:        form.FieldText,
						Description: "Name2 of the field",
					},
				},
			},
		},
	}

	pageSettings := view.Page{
		Id:         "settings",
		Name:       "Settings",
		Components: []component.Component{&group4},
	}

	config := server.ServerConfig{
		Port: "8081",
		Pages: []view.Page{
			pageCustomers,
			pageSettings,
		},
	}

	server.StartServer(config)
}
