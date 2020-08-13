package templates

// ResolverExtensions ...
var ResolverExtensions = `package gen

import (
	"context"
	"time"

	"github.com/graph-gophers/dataloader"
	"github.com/99designs/gqlgen/graphql"
	"github.com/gofrs/uuid"
	"github.com/loopcontext/go-graphql-orm/events"
	"github.com/vektah/gqlparser/v2/ast"
)

{{range $ext := .Model.ObjectExtensions}}
	{{$obj := $ext.Object}}

	{{if not $ext.ExtendsLocalObject}}
		// Generated{{$obj.Name}}Resolver struct
		type Generated{{$obj.Name}}Resolver struct { *GeneratedResolver }

		{{if or $obj.HasAnyRelationships $obj.HasReadonlyColumns}}

			// {{range $col := $obj.Columns}}
			// 	{{if $col.IsReadonlyType}}
			//  // {{$col.MethodName}} ...
			// 	func (r *Generated{{$obj.Name}}Resolver) {{$col.MethodName}}(ctx context.Context, obj *{{$obj.Name}}) (res {{$col.GoType}}, err error) {
			// 		return r.Handlers.{{$obj.Name}}{{$col.MethodName}}(ctx, r, obj)
			// 	}
			//  // {{$obj.Name}}{{$col.MethodName}}Handler handler
			// 	func {{$obj.Name}}{{$col.MethodName}}Handler(ctx context.Context,r *Generated{{$obj.Name}}Resolver, obj *{{$obj.Name}}) (res {{$col.GoType}}, err error) {
			// 		return nil, fmt.Errorf("Resolver handler for {{$obj.Name}}{{$col.MethodName}} not implemented")
			// 	}
			// 	{{end}}
			// {{end}}

			// {{range $index, $rel := $obj.Relationships}}
			//  //{{$rel.MethodName}} ...
			// 	func (r *Generated{{$obj.Name}}Resolver) {{$rel.MethodName}}(ctx context.Context, obj *{{$obj.Name}}) (res {{$rel.ReturnType}}, err error) {
			// 		return r.Handlers.{{$obj.Name}}{{$rel.MethodName}}(ctx, r, obj)
			// 	}
			//  // {{$obj.Name}}{{$rel.MethodName}}Handler Handler
			// 	func {{$obj.Name}}{{$rel.MethodName}}Handler(ctx context.Context,r *Generated{{$obj.Name}}Resolver, obj *{{$obj.Name}}) (res {{$rel.ReturnType}}, err error) {
			// 		err = fmt.Errorf("not implemented")
			// 		return
			// 	}
			// {{end}}

		{{end}}
	{{end}}

{{end}}
`
