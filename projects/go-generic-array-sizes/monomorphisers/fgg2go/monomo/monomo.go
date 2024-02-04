package monomo

import (
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"

	"github.com/dawidl022/go-generic-array-sizes/interpreters/fgg/ast"
)

func Monomorphise(p ast.Program) ast.Program {
	return newVisitor().monomorphise(p).(ast.Program)
}

type visitor struct {
	// map from type name to instantiations with generic type parameters
	queue              map[ast.TypeName][][]ast.IntegerLiteral
	seenInstantiations map[ast.TypeName][][]ast.IntegerLiteral

	typeEnv           map[ast.TypeParameter]ast.IntegerLiteral
	numericalTypeArgs []ast.IntegerLiteral
}

func newVisitor() visitor {
	return visitor{
		queue:              make(map[ast.TypeName][][]ast.IntegerLiteral),
		seenInstantiations: make(map[ast.TypeName][][]ast.IntegerLiteral),
	}
}

func (v visitor) newArgVisitor(numericalTypeArgs []ast.IntegerLiteral) visitor {
	return visitor{
		queue:              v.queue,
		seenInstantiations: v.seenInstantiations,
		numericalTypeArgs:  numericalTypeArgs,
	}
}

func (v visitor) newEnvVisitor(typeEnv map[ast.TypeParameter]ast.IntegerLiteral) visitor {
	return visitor{
		queue:              v.queue,
		seenInstantiations: v.seenInstantiations,
		typeEnv:            typeEnv,
	}
}

func (v visitor) monomorphise(m ast.MapVisitable) ast.MapVisitable {
	return m.AcceptMapVisitor(v)
}

func (v visitor) enqueue(typeName ast.TypeName, numericalTypeArgs []ast.IntegerLiteral) {
	for _, instantiation := range v.seenInstantiations[typeName] {
		if cmp.Equal(instantiation, numericalTypeArgs, cmpopts.EquateEmpty()) {
			return
		}
	}
	v.queue[typeName] = append(v.queue[typeName], numericalTypeArgs)
	v.seenInstantiations[typeName] = append(v.seenInstantiations[typeName], numericalTypeArgs)
}

func (v visitor) dequeue(typeName ast.TypeName) []ast.IntegerLiteral {
	instantiations, isDeclared := v.queue[typeName]
	// TODO refactor
	if !isDeclared || len(instantiations) == 0 {
		panic("no instantiations left")
	}
	curr := instantiations[0]
	v.queue[typeName] = instantiations[1:]
	if len(instantiations) == 1 {
		delete(v.queue, typeName)
	}
	return curr
}

func (v visitor) isEmpty(typeName ast.TypeName) bool {
	_, isDeclared := v.queue[typeName]
	return !isDeclared
}

func (v visitor) isQueueEmpty() bool {
	return len(v.queue) == 0
}

func (v visitor) monoTypeName(typeName ast.TypeName, numericalTypeArgs []ast.IntegerLiteral) ast.TypeName {
	monoName := string(typeName)
	for _, arg := range numericalTypeArgs {
		monoName += "__" + arg.String()
	}
	return ast.TypeName(monoName)
}

type typeInstantiation struct {
	typeName          ast.TypeName
	numericalTypeArgs []ast.IntegerLiteral
}

func (v visitor) MapProgram(p ast.Program) ast.MapVisitable {
	monoExpr := v.monomorphise(p.Expression).(ast.Expression)

	// each declaration in original FGG program gets bucket for 0 or more
	// instantiations
	monoDecls := make([][]ast.Declaration, len(p.Declarations))

	// TODO loop while queue is not empty
	for i, decl := range p.Declarations {
		switch decl.(type) {
		case ast.TypeDeclaration:
			decl := decl.(ast.TypeDeclaration)
			for !v.isEmpty(decl.TypeName) {
				inst := v.dequeue(decl.TypeName)
				// TODO what about methods? they all need to be monomorphised too
				monoDecls[i] = append(monoDecls[i], v.newArgVisitor(inst).monomorphise(decl).(ast.Declaration))
			}
		}
	}

	flatMonoDecl := []ast.Declaration{}
	for _, bucket := range monoDecls {
		for _, decl := range bucket {
			flatMonoDecl = append(flatMonoDecl, decl)
		}
	}

	return ast.Program{
		Declarations: flatMonoDecl,
		Expression:   monoExpr,
	}
}

func (v visitor) MapTypeDeclaration(t ast.TypeDeclaration) ast.MapVisitable {
	// TODO handle recursive type parameter bounds
	numericalTypeParams := make(map[ast.TypeParameter]ast.IntegerLiteral)
	var monoTypeParams []ast.TypeParameterConstraint

	numIndex := 0

	for _, typeParam := range t.TypeParameters {
		if _, isConstParam := typeParam.Bound.(ast.ConstType); isConstParam {
			numericalTypeParams[typeParam.TypeParameter] = v.numericalTypeArgs[numIndex]
			numIndex++
		} else {
			monoTypeParams = append(monoTypeParams, typeParam)
		}
	}

	return ast.TypeDeclaration{
		TypeName:       v.monoTypeName(t.TypeName, v.numericalTypeArgs),
		TypeParameters: monoTypeParams,
		TypeLiteral:    v.newEnvVisitor(numericalTypeParams).monomorphise(t.TypeLiteral).(ast.TypeLiteral),
	}
}

func (v visitor) MapMethodDeclaration(m ast.MethodDeclaration) ast.MapVisitable {
	//TODO implement me
	panic("implement me")
}

func (v visitor) MapArraySetMethodDeclaration(a ast.ArraySetMethodDeclaration) ast.MapVisitable {
	//TODO implement me
	panic("implement me")
}

func (v visitor) MapTypeParameterConstraint(t ast.TypeParameterConstraint) ast.MapVisitable {
	//TODO implement me
	panic("implement me")
}

func (v visitor) MapStructTypeLiteral(s ast.StructTypeLiteral) ast.MapVisitable {
	//TODO implement me
	panic("implement me")
}

func (v visitor) MapInterfaceTypeLiteral(i ast.InterfaceTypeLiteral) ast.MapVisitable {
	//TODO implement me
	panic("implement me")
}

func (v visitor) MapArrayTypeLiteral(a ast.ArrayTypeLiteral) ast.MapVisitable {
	return ast.ArrayTypeLiteral{
		Length:      v.monomorphise(a.Length).(ast.IntegerLiteral),
		ElementType: v.monomorphise(a.ElementType).(ast.Type),
	}
}

func (v visitor) MapMethodSpecification(m ast.MethodSpecification) ast.MapVisitable {
	//TODO implement me
	panic("implement me")
}

func (v visitor) MapIntegerLiteral(i ast.IntegerLiteral) ast.MapVisitable {
	return i
}

func (v visitor) MapVariable(variable ast.Variable) ast.MapVisitable {
	//TODO implement me
	panic("implement me")
}

func (v visitor) MapMethodCall(m ast.MethodCall) ast.MapVisitable {
	//TODO implement me
	panic("implement me")
}

func (v visitor) MapValueLiteral(value ast.ValueLiteral) ast.MapVisitable {
	monoType := v.monomorphise(value.Type).(ast.Type)
	monoValues := make([]ast.Expression, 0, len(value.Values))
	for _, val := range value.Values {
		monoValues = append(monoValues, v.monomorphise(val).(ast.Expression))
	}
	return ast.ValueLiteral{
		Type:   monoType,
		Values: monoValues,
	}
}

func (v visitor) MapSelect(s ast.Select) ast.MapVisitable {
	//TODO implement me
	panic("implement me")
}

func (v visitor) MapArrayIndex(a ast.ArrayIndex) ast.MapVisitable {
	//TODO implement me
	panic("implement me")
}

func (v visitor) MapMethodParameter(p ast.MethodParameter) ast.MapVisitable {
	//TODO implement me
	panic("implement me")
}

func (v visitor) MapConstType(c ast.ConstType) ast.MapVisitable {
	//TODO implement me
	panic("implement me")
}

func (v visitor) MapNamedType(n ast.NamedType) ast.MapVisitable {
	// instantiate any numerical type args
	var numericalTypeArgs []ast.IntegerLiteral
	var monoTypeArgs []ast.Type

	for _, arg := range n.TypeArguments {
		// TODO recursively monomorphise
		if numArg, isNumArg := arg.(ast.IntegerLiteral); isNumArg {
			numericalTypeArgs = append(numericalTypeArgs, numArg)
		} else {
			monoTypeArgs = append(monoTypeArgs, arg)
		}
	}
	// enqueue the instantiation
	v.enqueue(n.TypeName, numericalTypeArgs)

	// monomorphise the instantation
	return ast.NamedType{
		TypeName:      v.monoTypeName(n.TypeName, numericalTypeArgs),
		TypeArguments: monoTypeArgs,
	}
}

func (v visitor) MapTypeParameter(t ast.TypeParameter) ast.MapVisitable {
	if typeArg, shouldMonomorphise := v.typeEnv[t]; shouldMonomorphise {
		return typeArg
	}
	return t
}

func (v visitor) MapField(f ast.Field) ast.MapVisitable {
	//TODO implement me
	panic("implement me")
}

func (v visitor) MapMethodSignature(m ast.MethodSignature) ast.MapVisitable {
	//TODO implement me
	panic("implement me")
}