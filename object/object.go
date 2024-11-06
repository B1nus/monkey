package object

import (
	"bytes"
	"fmt"
	"monkey/ast"
	"strings"
)

func NewEnclosedEnvironment(outer *Environment) *Environment {
  env := NewEnvironment()
  env.outer = outer
  return env
}

type Environment struct {
	store map[string]Object
  outer *Environment
}

func NewEnvironment() *Environment {
	s := make(map[string]Object)
  return &Environment{store: s}
}

func (e *Environment) Get(name string) (Object, bool) {
	obj, ok := e.store[name]
  if !ok && e.outer != nil {
    obj, ok = e.outer.Get(name)
  }
	return obj, ok
}

func (e *Environment) Set(name string, val Object) Object {
	e.store[name] = val
	return val
}

type ObjectType string

const (
	FUNCTION_OBJ = "FUNCTION"
	INTEGER_OBJ  = "INTEGER"
	BOOLEAN_OBJ  = "BOOLEAN"
	RETURN_OBJ   = "RETURN"
	ERROR_OBJ    = "ERROR"
	NULL_OBJ     = "NULL"
)

type Function struct {
	Parameters []*ast.Identifier
	Body       *ast.BlockStatement
	Env        *Environment
}

func (f *Function) Type() ObjectType { return FUNCTION_OBJ }
func (f *Function) Inspect() string {
	var out bytes.Buffer
	params := []string{}
	for _, p := range f.Parameters {
		params = append(params, p.String())
	}
	out.WriteString("fn")
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(") {\n")
	out.WriteString(f.Body.String())
	out.WriteString("\n}")
	return out.String()
}

type Error struct {
	Message string
}

func (e *Error) Type() ObjectType { return ERROR_OBJ }
func (e *Error) Inspect() string  { return "ERROR: " + e.Message }

type Return struct {
	Value Object
}

func (rv *Return) Type() ObjectType { return RETURN_OBJ }
func (rv *Return) Inspect() string  { return rv.Value.Inspect() }

type Null struct{}

func (n *Null) Type() ObjectType { return NULL_OBJ }
func (n *Null) Inspect() string  { return "null" }

type Boolean struct {
	Value bool
}

func (b *Boolean) Inspect() string  { return fmt.Sprintf("%t", b.Value) }
func (b *Boolean) Type() ObjectType { return BOOLEAN_OBJ }

type Integer struct {
	Value int64
}

func (i *Integer) Inspect() string  { return fmt.Sprintf("%d", i.Value) }
func (i *Integer) Type() ObjectType { return INTEGER_OBJ }

type Object interface {
	Type() ObjectType
	Inspect() string
}