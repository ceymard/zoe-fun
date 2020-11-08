// Code generated by a lame .js file, DO NOT EDIT.

package zoe

import "io"


func (p *Position) CreateNamespace() *Namespace {
  res := &Namespace{}
  res.ExtendPosition(p)
  return res
}

func (tk *Token) CreateNamespace() *Namespace {
  return tk.Position.CreateNamespace()
}



func (r *Namespace) Dump(w io.Writer) {

  w.Write([]byte("(namespace"))




      for _, n := range r.Children {
        w.Write([]byte(" "))
        n.Dump(w)
      }




  w.Write([]byte(")"))

}






func (r *Namespace) AddChildren(other ...Node) *Namespace {
  for _, c := range other {
    if c != nil {
      r.Children = append(r.Children, c)
      r.ExtendPosition(c)
    }
  }
  return r
}





func (p *Position) CreateFragment() *Fragment {
  res := &Fragment{}
  res.ExtendPosition(p)
  return res
}

func (tk *Token) CreateFragment() *Fragment {
  return tk.Position.CreateFragment()
}



func (r *Fragment) Dump(w io.Writer) {

  w.Write([]byte("(fragment"))




      for _, n := range r.Children {
        w.Write([]byte(" "))
        n.Dump(w)
      }




  w.Write([]byte(")"))

}






func (r *Fragment) AddChildren(other ...Node) *Fragment {
  for _, c := range other {
    if c != nil {
      r.Children = append(r.Children, c)
      r.ExtendPosition(c)
    }
  }
  return r
}





func (p *Position) CreateTypeDecl() *TypeDecl {
  res := &TypeDecl{}
  res.ExtendPosition(p)
  return res
}

func (tk *Token) CreateTypeDecl() *TypeDecl {
  return tk.Position.CreateTypeDecl()
}



func (r *TypeDecl) Dump(w io.Writer) {

  w.Write([]byte("(typedecl"))




      w.Write([]byte(" "))
      if r.Template != nil {
        r.Template.Dump(w)
      } else {
        w.Write([]byte(red("<nil>")))
      }



      w.Write([]byte(" "))
      if r.Ident != nil {
        r.Ident.Dump(w)
      } else {
        w.Write([]byte(red("<nil>")))
      }



      w.Write([]byte(" "))
      if r.Def != nil {
        r.Def.Dump(w)
      } else {
        w.Write([]byte(red("<nil>")))
      }




  w.Write([]byte(")"))

}






func (r *TypeDecl) SetTemplate(other *Template) *TypeDecl {
  r.Template = other
  if other != nil {
    r.ExtendPosition(other)
  }
  return r
}





func (r *TypeDecl) SetIdent(other *Ident) *TypeDecl {
  r.Ident = other
  if other != nil {
    r.ExtendPosition(other)
  }
  return r
}





func (r *TypeDecl) SetDef(other Node) *TypeDecl {
  r.Def = other
  if other != nil {
    r.ExtendPosition(other)
  }
  return r
}





func (p *Position) CreateUnion() *Union {
  res := &Union{}
  res.ExtendPosition(p)
  return res
}

func (tk *Token) CreateUnion() *Union {
  return tk.Position.CreateUnion()
}



func (r *Union) Dump(w io.Writer) {

  w.Write([]byte("(union"))




      for _, n := range r.TypeExprs {
        w.Write([]byte(" "))
        n.Dump(w)
      }




  w.Write([]byte(")"))

}






func (r *Union) AddTypeExprs(other ...Node) *Union {
  for _, c := range other {
    if c != nil {
      r.TypeExprs = append(r.TypeExprs, c)
      r.ExtendPosition(c)
    }
  }
  return r
}





func (p *Position) CreateImportAs() *ImportAs {
  res := &ImportAs{}
  res.ExtendPosition(p)
  return res
}

func (tk *Token) CreateImportAs() *ImportAs {
  return tk.Position.CreateImportAs()
}


func (r *ImportAs) Dump(w io.Writer) {
  w.Write([]byte(r.GetText()))
}





func (p *Position) CreateVar() *Var {
  res := &Var{}
  res.ExtendPosition(p)
  return res
}

func (tk *Token) CreateVar() *Var {
  return tk.Position.CreateVar()
}



func (r *Var) Dump(w io.Writer) {

  w.Write([]byte("(var"))




      w.Write([]byte(" "))
      if r.Ident != nil {
        r.Ident.Dump(w)
      } else {
        w.Write([]byte(red("<nil>")))
      }



      w.Write([]byte(" "))
      if r.TypeExp != nil {
        r.TypeExp.Dump(w)
      } else {
        w.Write([]byte(red("<nil>")))
      }



      w.Write([]byte(" "))
      if r.Exp != nil {
        r.Exp.Dump(w)
      } else {
        w.Write([]byte(red("<nil>")))
      }




  w.Write([]byte(")"))

}






func (r *Var) SetIdent(other *Ident) *Var {
  r.Ident = other
  if other != nil {
    r.ExtendPosition(other)
  }
  return r
}





func (r *Var) SetTypeExp(other Node) *Var {
  r.TypeExp = other
  if other != nil {
    r.ExtendPosition(other)
  }
  return r
}





func (r *Var) SetExp(other Node) *Var {
  r.Exp = other
  if other != nil {
    r.ExtendPosition(other)
  }
  return r
}





func (p *Position) CreateOperation() *Operation {
  res := &Operation{}
  res.ExtendPosition(p)
  return res
}

func (tk *Token) CreateOperation() *Operation {
  return tk.Position.CreateOperation()
}



func (r *Operation) Dump(w io.Writer) {

  w.Write([]byte("("))




      w.Write([]byte(" "))
      if r.Token != nil {
        r.Token.Dump(w)
      } else {
        w.Write([]byte(red("<nil>")))
      }



      for _, n := range r.Operands {
        w.Write([]byte(" "))
        n.Dump(w)
      }




  w.Write([]byte(")"))

}






func (r *Operation) SetToken(other *Token) *Operation {
  r.Token = other
  if other != nil {
    r.ExtendPosition(other)
  }
  return r
}





func (r *Operation) AddOperands(other ...Node) *Operation {
  for _, c := range other {
    if c != nil {
      r.Operands = append(r.Operands, c)
      r.ExtendPosition(c)
    }
  }
  return r
}





func (p *Position) CreateTemplate() *Template {
  res := &Template{}
  res.ExtendPosition(p)
  return res
}

func (tk *Token) CreateTemplate() *Template {
  return tk.Position.CreateTemplate()
}



func (r *Template) Dump(w io.Writer) {

  w.Write([]byte("(template"))




      w.Write([]byte(" "))
      if r.Args != nil {
        r.Args.Dump(w)
      } else {
        w.Write([]byte(red("<nil>")))
      }




  w.Write([]byte(")"))

}






func (r *Template) SetArgs(other *VarTuple) *Template {
  r.Args = other
  if other != nil {
    r.ExtendPosition(other)
  }
  return r
}





func (p *Position) CreateFnDef() *FnDef {
  res := &FnDef{}
  res.ExtendPosition(p)
  return res
}

func (tk *Token) CreateFnDef() *FnDef {
  return tk.Position.CreateFnDef()
}



func (r *FnDef) Dump(w io.Writer) {

  w.Write([]byte("(fndef"))




      w.Write([]byte(" "))
      if r.Template != nil {
        r.Template.Dump(w)
      } else {
        w.Write([]byte(red("<nil>")))
      }



      w.Write([]byte(" "))
      if r.Signature != nil {
        r.Signature.Dump(w)
      } else {
        w.Write([]byte(red("<nil>")))
      }



      w.Write([]byte(" "))
      if r.Definition != nil {
        r.Definition.Dump(w)
      } else {
        w.Write([]byte(red("<nil>")))
      }




  w.Write([]byte(")"))

}






func (r *FnDef) SetTemplate(other *Template) *FnDef {
  r.Template = other
  if other != nil {
    r.ExtendPosition(other)
  }
  return r
}





func (r *FnDef) SetSignature(other *Signature) *FnDef {
  r.Signature = other
  if other != nil {
    r.ExtendPosition(other)
  }
  return r
}





func (r *FnDef) SetDefinition(other *Block) *FnDef {
  r.Definition = other
  if other != nil {
    r.ExtendPosition(other)
  }
  return r
}





func (p *Position) CreateSignature() *Signature {
  res := &Signature{}
  res.ExtendPosition(p)
  return res
}

func (tk *Token) CreateSignature() *Signature {
  return tk.Position.CreateSignature()
}



func (r *Signature) Dump(w io.Writer) {

  w.Write([]byte("(signature"))




      w.Write([]byte(" "))
      if r.Args != nil {
        r.Args.Dump(w)
      } else {
        w.Write([]byte(red("<nil>")))
      }



      w.Write([]byte(" "))
      if r.ReturnTypeExp != nil {
        r.ReturnTypeExp.Dump(w)
      } else {
        w.Write([]byte(red("<nil>")))
      }




  w.Write([]byte(")"))

}






func (r *Signature) SetArgs(other *VarTuple) *Signature {
  r.Args = other
  if other != nil {
    r.ExtendPosition(other)
  }
  return r
}





func (r *Signature) SetReturnTypeExp(other Node) *Signature {
  r.ReturnTypeExp = other
  if other != nil {
    r.ExtendPosition(other)
  }
  return r
}





func (p *Position) CreateFnCall() *FnCall {
  res := &FnCall{}
  res.ExtendPosition(p)
  return res
}

func (tk *Token) CreateFnCall() *FnCall {
  return tk.Position.CreateFnCall()
}



func (r *FnCall) Dump(w io.Writer) {

  w.Write([]byte("(fncall"))




      w.Write([]byte(" "))
      if r.Left != nil {
        r.Left.Dump(w)
      } else {
        w.Write([]byte(red("<nil>")))
      }



      w.Write([]byte(" "))
      if r.Args != nil {
        r.Args.Dump(w)
      } else {
        w.Write([]byte(red("<nil>")))
      }




  w.Write([]byte(")"))

}






func (r *FnCall) SetLeft(other Node) *FnCall {
  r.Left = other
  if other != nil {
    r.ExtendPosition(other)
  }
  return r
}





func (r *FnCall) SetArgs(other *Tuple) *FnCall {
  r.Args = other
  if other != nil {
    r.ExtendPosition(other)
  }
  return r
}





func (p *Position) CreateGetIndex() *GetIndex {
  res := &GetIndex{}
  res.ExtendPosition(p)
  return res
}

func (tk *Token) CreateGetIndex() *GetIndex {
  return tk.Position.CreateGetIndex()
}



func (r *GetIndex) Dump(w io.Writer) {

  w.Write([]byte("(getindex"))




      w.Write([]byte(" "))
      if r.Left != nil {
        r.Left.Dump(w)
      } else {
        w.Write([]byte(red("<nil>")))
      }



      w.Write([]byte(" "))
      if r.Index != nil {
        r.Index.Dump(w)
      } else {
        w.Write([]byte(red("<nil>")))
      }




  w.Write([]byte(")"))

}






func (r *GetIndex) SetLeft(other Node) *GetIndex {
  r.Left = other
  if other != nil {
    r.ExtendPosition(other)
  }
  return r
}





func (r *GetIndex) SetIndex(other Node) *GetIndex {
  r.Index = other
  if other != nil {
    r.ExtendPosition(other)
  }
  return r
}





func (p *Position) CreateSetIndex() *SetIndex {
  res := &SetIndex{}
  res.ExtendPosition(p)
  return res
}

func (tk *Token) CreateSetIndex() *SetIndex {
  return tk.Position.CreateSetIndex()
}



func (r *SetIndex) Dump(w io.Writer) {

  w.Write([]byte("(setindex"))




      w.Write([]byte(" "))
      if r.Left != nil {
        r.Left.Dump(w)
      } else {
        w.Write([]byte(red("<nil>")))
      }



      w.Write([]byte(" "))
      if r.Index != nil {
        r.Index.Dump(w)
      } else {
        w.Write([]byte(red("<nil>")))
      }



      w.Write([]byte(" "))
      if r.Value != nil {
        r.Value.Dump(w)
      } else {
        w.Write([]byte(red("<nil>")))
      }




  w.Write([]byte(")"))

}






func (r *SetIndex) SetLeft(other Node) *SetIndex {
  r.Left = other
  if other != nil {
    r.ExtendPosition(other)
  }
  return r
}





func (r *SetIndex) SetIndex(other Node) *SetIndex {
  r.Index = other
  if other != nil {
    r.ExtendPosition(other)
  }
  return r
}





func (r *SetIndex) SetValue(other Node) *SetIndex {
  r.Value = other
  if other != nil {
    r.ExtendPosition(other)
  }
  return r
}





func (p *Position) CreateIf() *If {
  res := &If{}
  res.ExtendPosition(p)
  return res
}

func (tk *Token) CreateIf() *If {
  return tk.Position.CreateIf()
}



func (r *If) Dump(w io.Writer) {

  w.Write([]byte("(if"))




      w.Write([]byte(" "))
      if r.Cond != nil {
        r.Cond.Dump(w)
      } else {
        w.Write([]byte(red("<nil>")))
      }



      w.Write([]byte(" "))
      if r.Then != nil {
        r.Then.Dump(w)
      } else {
        w.Write([]byte(red("<nil>")))
      }



      w.Write([]byte(" "))
      if r.Else != nil {
        r.Else.Dump(w)
      } else {
        w.Write([]byte(red("<nil>")))
      }




  w.Write([]byte(")"))

}






func (r *If) SetCond(other Node) *If {
  r.Cond = other
  if other != nil {
    r.ExtendPosition(other)
  }
  return r
}





func (r *If) SetThen(other Node) *If {
  r.Then = other
  if other != nil {
    r.ExtendPosition(other)
  }
  return r
}





func (r *If) SetElse(other Node) *If {
  r.Else = other
  if other != nil {
    r.ExtendPosition(other)
  }
  return r
}





func (p *Position) CreateFnDecl() *FnDecl {
  res := &FnDecl{}
  res.ExtendPosition(p)
  return res
}

func (tk *Token) CreateFnDecl() *FnDecl {
  return tk.Position.CreateFnDecl()
}



func (r *FnDecl) Dump(w io.Writer) {

  w.Write([]byte("(fndecl"))




      w.Write([]byte(" "))
      if r.Ident != nil {
        r.Ident.Dump(w)
      } else {
        w.Write([]byte(red("<nil>")))
      }



      w.Write([]byte(" "))
      if r.FnDef != nil {
        r.FnDef.Dump(w)
      } else {
        w.Write([]byte(red("<nil>")))
      }




  w.Write([]byte(")"))

}






func (r *FnDecl) SetIdent(other *Ident) *FnDecl {
  r.Ident = other
  if other != nil {
    r.ExtendPosition(other)
  }
  return r
}





func (r *FnDecl) SetFnDef(other *FnDef) *FnDecl {
  r.FnDef = other
  if other != nil {
    r.ExtendPosition(other)
  }
  return r
}





func (p *Position) CreateTuple() *Tuple {
  res := &Tuple{}
  res.ExtendPosition(p)
  return res
}

func (tk *Token) CreateTuple() *Tuple {
  return tk.Position.CreateTuple()
}



func (r *Tuple) Dump(w io.Writer) {

  w.Write([]byte("["))




      for _, n := range r.Children {
        w.Write([]byte(" "))
        n.Dump(w)
      }




  w.Write([]byte("]"))

}






func (r *Tuple) AddChildren(other ...Node) *Tuple {
  for _, c := range other {
    if c != nil {
      r.Children = append(r.Children, c)
      r.ExtendPosition(c)
    }
  }
  return r
}





func (p *Position) CreateVarTuple() *VarTuple {
  res := &VarTuple{}
  res.ExtendPosition(p)
  return res
}

func (tk *Token) CreateVarTuple() *VarTuple {
  return tk.Position.CreateVarTuple()
}



func (r *VarTuple) Dump(w io.Writer) {

  w.Write([]byte("["))




      for _, n := range r.Vars {
        w.Write([]byte(" "))
        n.Dump(w)
      }




  w.Write([]byte("]"))

}






func (r *VarTuple) AddVars(other ...*Var) *VarTuple {
  for _, c := range other {
    if c != nil {
      r.Vars = append(r.Vars, c)
      r.ExtendPosition(c)
    }
  }
  return r
}





func (p *Position) CreateBlock() *Block {
  res := &Block{}
  res.ExtendPosition(p)
  return res
}

func (tk *Token) CreateBlock() *Block {
  return tk.Position.CreateBlock()
}



func (r *Block) Dump(w io.Writer) {

  w.Write([]byte("{"))




      for _, n := range r.Children {
        w.Write([]byte(" "))
        n.Dump(w)
      }




  w.Write([]byte("}"))

}






func (r *Block) AddChildren(other ...Node) *Block {
  for _, c := range other {
    if c != nil {
      r.Children = append(r.Children, c)
      r.ExtendPosition(c)
    }
  }
  return r
}





func (p *Position) CreateReturn() *Return {
  res := &Return{}
  res.ExtendPosition(p)
  return res
}

func (tk *Token) CreateReturn() *Return {
  return tk.Position.CreateReturn()
}



func (r *Return) Dump(w io.Writer) {

  w.Write([]byte("(return"))




      w.Write([]byte(" "))
      if r.Expr != nil {
        r.Expr.Dump(w)
      } else {
        w.Write([]byte(red("<nil>")))
      }




  w.Write([]byte(")"))

}






func (r *Return) SetExpr(other Node) *Return {
  r.Expr = other
  if other != nil {
    r.ExtendPosition(other)
  }
  return r
}





func (p *Position) CreateIdent() *Ident {
  res := &Ident{}
  res.ExtendPosition(p)
  return res
}

func (tk *Token) CreateIdent() *Ident {
  return tk.Position.CreateIdent()
}


func (r *Ident) Dump(w io.Writer) {
  w.Write([]byte(r.GetText()))
}





func (p *Position) CreateNull() *Null {
  res := &Null{}
  res.ExtendPosition(p)
  return res
}

func (tk *Token) CreateNull() *Null {
  return tk.Position.CreateNull()
}


func (r *Null) Dump(w io.Writer) {
  w.Write([]byte(r.GetText()))
}





func (p *Position) CreateFalse() *False {
  res := &False{}
  res.ExtendPosition(p)
  return res
}

func (tk *Token) CreateFalse() *False {
  return tk.Position.CreateFalse()
}


func (r *False) Dump(w io.Writer) {
  w.Write([]byte(r.GetText()))
}





func (p *Position) CreateTrue() *True {
  res := &True{}
  res.ExtendPosition(p)
  return res
}

func (tk *Token) CreateTrue() *True {
  return tk.Position.CreateTrue()
}


func (r *True) Dump(w io.Writer) {
  w.Write([]byte(r.GetText()))
}





func (p *Position) CreateString() *String {
  res := &String{}
  res.ExtendPosition(p)
  return res
}

func (tk *Token) CreateString() *String {
  return tk.Position.CreateString()
}


func (r *String) Dump(w io.Writer) {
  w.Write([]byte(r.GetText()))
}





func (p *Position) CreateInteger() *Integer {
  res := &Integer{}
  res.ExtendPosition(p)
  return res
}

func (tk *Token) CreateInteger() *Integer {
  return tk.Position.CreateInteger()
}


func (r *Integer) Dump(w io.Writer) {
  w.Write([]byte(r.GetText()))
}





func (p *Position) CreateFloat() *Float {
  res := &Float{}
  res.ExtendPosition(p)
  return res
}

func (tk *Token) CreateFloat() *Float {
  return tk.Position.CreateFloat()
}


func (r *Float) Dump(w io.Writer) {
  w.Write([]byte(r.GetText()))
}





func (p *Position) CreateEof() *Eof {
  res := &Eof{}
  res.ExtendPosition(p)
  return res
}

func (tk *Token) CreateEof() *Eof {
  return tk.Position.CreateEof()
}


func (r *Eof) Dump(w io.Writer) {
  w.Write([]byte(r.GetText()))
}




