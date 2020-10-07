package templates

// Migrations ...
var Migrations = `package gen

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/jinzhu/gorm"
	"github.com/rs/zerolog/log"
	"gopkg.in/gormigrate.v1"
)

// Migrate runs migrations
func Migrate(db *gorm.DB, options *gormigrate.Options, migrations []*gormigrate.Migration) error {
	m := gormigrate.New(db, options, migrations)

	// // it's possible to use this, but in case of any specific keys or columns are created in migrations, they will not be generated by automigrate
	// m.InitSchema(func(tx *gorm.DB) error {
	// 	return AutoMigrate(db)
	// })

	return m.Migrate();
}

func AutoMigrate(db *gorm.DB) (err error) {
	_db := db.AutoMigrate({{range $obj := .Model.ObjectEntities}}
		{{.Name}}{},{{end}}
	)
	if _db.Error != nil {
		log.Err(_db.Error).Send()
		return _db.Error
	}
	if(_db.Dialect().GetName() != "sqlite3"){
		{{range $obj := .Model.ObjectEntities}}
			{{range $rel := $obj.Relationships}}
				{{if $rel.IsToOne}}
					err = _db.Model({{$obj.Name}}{}).RemoveForeignKey("{{$rel.Name}}Id",TableName("{{$rel.Target.TableName}}")+"({{$rel.ForeignKeyDestinationColumn}})").Error
					if err != nil {
						log.Err(err).Send()
					}
					err = _db.Model({{$obj.Name}}{}).AddForeignKey("{{$rel.Name}}Id",TableName("{{$rel.Target.TableName}}")+"({{$rel.ForeignKeyDestinationColumn}})", "{{$rel.OnDelete "SET NULL"}}", "{{$rel.OnUpdate "SET NULL"}}").Error
					if err != nil {
						log.Err(err).Send()
					}
				{{else if $rel.IsManyToMany}}
					err = _db.Model({{$rel.ManyToManyObjectNameCC}}{}).RemoveForeignKey("{{$rel.ForeignKeyDestinationColumn}}",TableName("{{$rel.Obj.TableName}}")+"(id)").Error
					if err != nil {
						log.Err(err).Send()
					}
					err = _db.Model({{$rel.ManyToManyObjectNameCC}}{}).AddForeignKey("{{$rel.ForeignKeyDestinationColumn}}",TableName("{{$rel.Obj.TableName}}")+"(id)", "{{$rel.OnDelete "CASCADE"}}", "{{$rel.OnUpdate "CASCADE"}}").Error
					if err != nil {
						log.Err(err).Send()
					}
				{{end}}
			{{end}}
		{{end}}
		if _db.Error != nil {
			log.Err(_db.Error).Send()
		}
	}
	return nil // _db.Error
}
`
