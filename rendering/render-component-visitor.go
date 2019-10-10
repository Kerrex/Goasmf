package rendering

import (
	"context"
	"fmt"
	"github.com/aymerick/raymond/ast"
	"goasmf/component"
	"goasmf/global"
	"goasmf/helpers"
)

type RenderComponentVisitor struct {
	RenderingContext     context.Context
}

func NewRenderComponentVisitor(renderingContext context.Context) *RenderComponentVisitor {
	return &RenderComponentVisitor{RenderingContext: renderingContext}
}

func (r *RenderComponentVisitor) VisitProgram(node *ast.Program) interface{} {
	for _, n := range node.Body {
		n.Accept(r)
	}

	return nil
}

func (r *RenderComponentVisitor) VisitMustache(node *ast.MustacheStatement) interface{} {
	node.Expression.Accept(r)
	return nil
}

func (r *RenderComponentVisitor) VisitBlock(node *ast.BlockStatement) interface{} {
	node.Expression.Accept(r)

	if node.Program != nil {
		node.Program.Accept(r)
	}

	if node.Inverse != nil {
		node.Inverse.Accept(r)
	}
	return nil
}

func (r *RenderComponentVisitor) VisitPartial(node *ast.PartialStatement) interface{} {
	node.Name.Accept(r)

	if len(node.Params) > 0 {
		node.Params[0].Accept(r)
	}

	// hash
	if node.Hash != nil {
		node.Hash.Accept(r)
	}

	return nil
}

func (r *RenderComponentVisitor) VisitContent(*ast.ContentStatement) interface{} {
	return nil
}

func (r *RenderComponentVisitor) VisitComment(*ast.CommentStatement) interface{} {
	return nil
}

func (r *RenderComponentVisitor) VisitExpression(node *ast.Expression) interface{} {
	// path
	node.Path.Accept(r)
	if node.HelperName() == "renderComponent" {
		r.handleRenderComponentExpression(node)
	}

	// params
	for _, n := range node.Params {
		n.Accept(r)
	}

	// hash
	if node.Hash != nil {
		node.Hash.Accept(r)
	}

	return nil
}

func (r *RenderComponentVisitor) handleRenderComponentExpression(node *ast.Expression) {
	if node.Hash == nil {
		node.Hash = ast.NewHash(0, 0)
	}
	if len(node.Params) < 1 {
		panic("renderComponent must be followed by component name!")
	}

	generatedInstanceId := helpers.RandomString(64)

	r.addInstanceIdToHash(generatedInstanceId, node)
	r.generateNewComponentInstance(node, generatedInstanceId)
}

func (r *RenderComponentVisitor) generateNewComponentInstance(node *ast.Expression, generatedInstanceId string) {
	componentName := node.Params[0].(*ast.StringLiteral).Value
	componentInstanceFactory := global.componentFactories[componentName]
	if componentInstanceFactory == nil {
		panic(fmt.Sprintf("component %s does not exists", componentName))
	}
	componentInstance := componentInstanceFactory(r.RenderingContext)
	componentInstance.SetInstanceId(generatedInstanceId)

	r.addComponentInstanceToGlobalContext(generatedInstanceId, componentInstance)
}

func (r *RenderComponentVisitor) addComponentInstanceToGlobalContext(generatedInstanceId string, componentInstance component.Component) {
	currentlyRenderedComponentInstances := global.GetGlobalContext().Value(global.CurrentlyRenderedComponentInstances).(map[string]component.Component)
	currentlyRenderedComponentInstances[generatedInstanceId] = componentInstance
	global.SetGlobalContext(context.WithValue(global.GetGlobalContext(), global.CurrentlyRenderedComponentInstances, currentlyRenderedComponentInstances))
}

func (r *RenderComponentVisitor) addInstanceIdToHash(generatedInstanceId string, node *ast.Expression) {
	newHashPair := ast.NewHashPair(0, 0)
	newHashPair.Key = "componentInstanceId"
	newHashPair.Val = ast.NewStringLiteral(0, 0, generatedInstanceId)
	node.Hash.Pairs = append(node.Hash.Pairs, newHashPair)
}

func (r *RenderComponentVisitor) VisitSubExpression(node *ast.SubExpression) interface{} {
	node.Expression.Accept(r)

	return nil
}

func (r *RenderComponentVisitor) VisitPath(*ast.PathExpression) interface{} {
	return nil
}

func (r *RenderComponentVisitor) VisitString(*ast.StringLiteral) interface{} {
	return nil
}

func (r *RenderComponentVisitor) VisitBoolean(*ast.BooleanLiteral) interface{} {
	return nil
}

func (r *RenderComponentVisitor) VisitNumber(*ast.NumberLiteral) interface{} {
	return nil
}

func (r *RenderComponentVisitor) VisitHash(*ast.Hash) interface{} {
	return nil
}

func (r *RenderComponentVisitor) VisitHashPair(*ast.HashPair) interface{} {
	return nil
}
