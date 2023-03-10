// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// CharactersColumns holds the columns for the "characters" table.
	CharactersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "name", Type: field.TypeString},
		{Name: "combat", Type: field.TypeInt},
		{Name: "connaissance", Type: field.TypeInt},
		{Name: "discretion", Type: field.TypeInt},
		{Name: "endurance", Type: field.TypeInt},
		{Name: "force", Type: field.TypeInt},
		{Name: "habilite", Type: field.TypeInt},
		{Name: "mouvement", Type: field.TypeInt},
		{Name: "perception", Type: field.TypeInt},
		{Name: "sociabilite", Type: field.TypeInt},
		{Name: "survie", Type: field.TypeInt},
		{Name: "tir", Type: field.TypeInt},
		{Name: "volonte", Type: field.TypeInt},
		{Name: "exp", Type: field.TypeInt},
		{Name: "armes_hast", Type: field.TypeInt},
		{Name: "armes_moine", Type: field.TypeInt},
		{Name: "armes_doubles", Type: field.TypeInt},
		{Name: "armes_naturelles", Type: field.TypeInt},
		{Name: "batons", Type: field.TypeInt},
		{Name: "cimeterres", Type: field.TypeInt},
		{Name: "fleaux", Type: field.TypeInt},
		{Name: "fouets", Type: field.TypeInt},
		{Name: "haches", Type: field.TypeInt},
		{Name: "katanas", Type: field.TypeInt},
		{Name: "lames_legeres", Type: field.TypeInt},
		{Name: "lames_lourdes", Type: field.TypeInt},
		{Name: "lances", Type: field.TypeInt},
		{Name: "marteaux", Type: field.TypeInt},
		{Name: "mains_nues", Type: field.TypeInt},
		{Name: "mysteres", Type: field.TypeInt},
		{Name: "exploration_souterraine", Type: field.TypeInt},
		{Name: "ingenierie", Type: field.TypeInt},
		{Name: "geographie", Type: field.TypeInt},
		{Name: "histoire", Type: field.TypeInt},
		{Name: "folklore", Type: field.TypeInt},
		{Name: "nature", Type: field.TypeInt},
		{Name: "noblesse", Type: field.TypeInt},
		{Name: "plans", Type: field.TypeInt},
		{Name: "religon", Type: field.TypeInt},
		{Name: "anatomie", Type: field.TypeInt},
		{Name: "magie_theorique", Type: field.TypeInt},
		{Name: "economie", Type: field.TypeInt},
		{Name: "linguistique", Type: field.TypeInt},
	}
	// CharactersTable holds the schema information for the "characters" table.
	CharactersTable = &schema.Table{
		Name:       "characters",
		Columns:    CharactersColumns,
		PrimaryKey: []*schema.Column{CharactersColumns[0]},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "username", Type: field.TypeString, Unique: true},
		{Name: "password", Type: field.TypeString},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		CharactersTable,
		UsersTable,
	}
)

func init() {
}
