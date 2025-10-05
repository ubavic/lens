package main

import (
	"github.com/ubavic/lens/server"
	"github.com/ubavic/lens/view"
	"github.com/ubavic/lens/view/form"
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

	pageCustomers := view.Page{
		Id:   "customer",
		Name: "Customers",
		Components: []view.Component{
			&table,
		},
	}

	form := form.Form{
		Fields: []form.Field{
			{
				ID:          "name",
				Name:        "Name",
				Type:        form.FieldText,
				Required:    true,
				Description: "Name of the field",
			},
			{
				ID:          "name",
				Name:        "Name2",
				Type:        form.FieldText,
				Required:    true,
				Description: "Name2 of the field",
				Error:       "Errored",
			},
		},
	}

	pageSettings := view.Page{
		Id:         "settings",
		Name:       "Settings",
		Components: []view.Component{&form},
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
