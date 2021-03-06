package templates

// ResolverSrcExt template
var ResolverSrcExt = `package src

import (
	"context"
	"fmt"

	"{{.Config.Package}}/gen"
	"github.com/rs/zerolog/log"
)

const (
	jwtTokenPermissionErrMsg = "You don't have permission to %s the entity %s"
)


{{range $obj := .Model.ObjectEntities}}
// {{$obj.PluralName}} method
func (r *QueryResolver) {{$obj.PluralName}}(ctx context.Context, offset *int, limit *int, q *string, sort []*gen.{{$obj.Name}}SortType, filter *gen.{{$obj.Name}}FilterType) (*gen.{{$obj.Name}}ResultType, error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if gen.HasPermission(jwtClaims, "{{$obj.TableName}}", gen.JWTPermissionConstList[:1]) {
		// TODO: Insert here any code ETL on your query/mutation for example: scope {{$obj.PluralName}} to current jwtClaims.Subject (User.ID)
		return r.GeneratedQueryResolver.{{$obj.PluralName}}(ctx, offset, limit, q, sort, filter)
	}
	return nil, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstList, "{{$obj.TableName}}")
}

// Create{{$obj.Name}} method
func (r *MutationResolver) Create{{$obj.Name}}(ctx context.Context, input map[string]interface{}) (item *gen.{{$obj.Name}}, err error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if gen.HasPermission(jwtClaims, "{{$obj.TableName}}", gen.JWTPermissionConstCreate[:1]) {
		// TODO: Insert here any code ETL on your query/mutation for example: scope {{$obj.PluralName}} to current jwtClaims.Subject (User.ID)
		return r.GeneratedMutationResolver.Create{{$obj.Name}}(ctx, input)
	}
	return nil, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstCreate, "{{$obj.TableName}}")
}

// Read{{$obj.Name}} method
func (r *QueryResolver) {{$obj.Name}}(ctx context.Context, id *string, q *string, filter *gen.{{$obj.Name}}FilterType) (*gen.{{$obj.Name}}, error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if gen.HasPermission(jwtClaims, "{{$obj.TableName}}", gen.JWTPermissionConstRead[:1]) {
		// TODO: Insert here any code ETL on your query/mutation for example: scope {{$obj.PluralName}} to current jwtClaims.Subject (User.ID)
		return r.GeneratedQueryResolver.{{$obj.Name}}(ctx, id, q, filter)
	}
	return nil, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstRead, "{{$obj.TableName}}")
}

// Update{{$obj.Name}} method
func (r *MutationResolver) Update{{$obj.Name}}(ctx context.Context, id string, input map[string]interface{}) (item *gen.{{$obj.Name}}, err error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if gen.HasPermission(jwtClaims, "{{$obj.TableName}}", gen.JWTPermissionConstUpdate[:1]) {
		// TODO: Insert here any code ETL on your query/mutation for example: scope {{$obj.PluralName}} to current jwtClaims.Subject (User.ID)
		return r.GeneratedMutationResolver.Update{{$obj.Name}}(ctx, id, input)
	}
	return nil, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstUpdate, "{{$obj.TableName}}")
}

// Delete{{$obj.Name}} method
func (r *MutationResolver) Delete{{$obj.Name}}(ctx context.Context, id string) (item *gen.{{$obj.Name}}, err error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if gen.HasPermission(jwtClaims, "{{$obj.TableName}}", gen.JWTPermissionConstDelete[:1]) {
		// TODO: Insert here any code ETL on your query/mutation for example: scope {{$obj.PluralName}} to current jwtClaims.Subject (User.ID)
		return r.GeneratedMutationResolver.Delete{{$obj.Name}}(ctx, id)
	}
	return nil, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstDelete, "{{$obj.TableName}}")
}

// DeleteAll{{$obj.PluralName}} method
func (r *MutationResolver) DeleteAll{{$obj.PluralName}}(ctx context.Context) (ok bool, err error) {
	jwtClaims := gen.GetJWTClaimsFromContext(ctx)
	if gen.HasRole(jwtClaims, "admin") && gen.HasPermission(jwtClaims, "{{$obj.TableName}}", gen.JWTPermissionConstDelete[:1]) {
		 // TODO: Insert here any code ETL on your query/mutation for example: scope {{$obj.PluralName}} to current jwtClaims.Subject (User.ID)
		 return r.GeneratedMutationResolver.DeleteAll{{$obj.PluralName}}(ctx)
	}
	return false, fmt.Errorf(jwtTokenPermissionErrMsg, gen.JWTPermissionConstDelete, "{{$obj.TableName}}")
}

{{end}}
`
