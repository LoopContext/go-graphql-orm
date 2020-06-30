package model

import (
	"github.com/graphql-go/graphql/language/ast"
)

// ObjectExtension ...
type ObjectExtension struct {
	Def    *ast.TypeExtensionDefinition
	Model  *Model
	Object *Object
}

// func (oe *ObjectExtension) GetObject() *Object {
// 	return &Object{
// 		Def:   oe.Def.Definition,
// 		Model: oe.Model,
// 		Extension: oe,
// 	}
// }

// IsFederatedType ...
func (oe *ObjectExtension) IsFederatedType() bool {
	return oe.Object.IsFederatedType()
}

// ExtendsLocalObject ...
func (oe *ObjectExtension) ExtendsLocalObject() bool {
	return oe.Model.HasObject(oe.Object.Name())
}
