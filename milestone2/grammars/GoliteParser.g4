parser grammar GoliteParser; 

options {
    tokenVocab = GoliteLexer; 
}
 

/*Program = Types Declarations Functions 'eof'                                                      ;
Types = {TypeDeclaration}                                                                         ;
TypeDeclaration = 'type' 'id' 'struct' '{' Fields '}' ';'                                         ;
Fields = Decl ';' {Decl ';'}                                                                      ;
Decl = 'id' Type                                                                                  ;
Type = 'int' | 'bool' | '*' 'id'                                                                  ;
Declarations = {Declaration}                                                                      ;
Declaration = 'var' Ids Type ';'                                                                  ;
Ids = 'id' {',' 'id'}                                                                             ;
Functions = {Function}                                                                            ;
Function = 'func' 'id' Parameters [ReturnType] '{' Declarations Statements '}'                    ;
Parameters = '(' [ Decl {',' Decl}] ')'                                                           ;
ReturnType = type                                                                                 ;
Statements = {Statement}                                                                          ;
Statement = Assignment | Print | Read |  Delete | Conditional | Loop | Return | Invocation        ;
Block = '{' Statements '}'                                                                        ;
Delete = 'delete' Expression ';'                                                                  ;
Read = 'scan' LValue  ';'                                                                         ;
Assignment = LValue '=' Expression  ';'                                                           ;
Print = 'printf' '(' 'string' { ',' Expression} ')'  ';'                                          ;
Conditional = 'if' '(' Expression ')' Block ['else' Block]                                        ;
Loop = 'for' '(' Expression ')' Block                                                             ;
Return = 'return' [Expression] ';'                                                                ;
Invocation = 'id' Arguments ';'                                                                   ;
Arguments = '(' [Expression {',' Expression}] ')'                                                 ;
LValue = 'id' {'.' 'id'}                                                                          ;
Expression = BoolTerm {'||' BoolTerm}                                                             ;
BoolTerm = EqualTerm {'&&' EqualTerm}                                                             ;
EqualTerm =  RelationTerm {('=='| '!=') RelationTerm}                                             ;
RelationTerm = SimpleTerm {('>'| '<' | '<=' | '>=') SimpleTerm}                                   ;
SimpleTerm = Term {('+'| '-') Term}                                                               ;
Term = UnaryTerm {('*'| '/') UnaryTerm}                                                           ;
UnaryTerm = '!' SelectorTerm | '-' SelectorTerm | SelectorTerm                                    ;
SelectorTerm = Factor {'.' 'id'}                                                                  ;
Factor = '(' Expression ')' | 'id' [Arguments] | 'number' | 'new' id | 'true' | 'false' | 'nil'   ;*/




program: types declarations functions EOF; 
types: typeDeclaration*;
typeDeclaration: TYPE IDENTIFIER STRUCT LBRACE fields RBRACE SEMICOLON;
fields: decl SEMICOLON (decl SEMICOLON)*;
decl: IDENTIFIER type;
type: INT | BOOL | ASTERISK IDENTIFIER;
declarations: declaration*;
declaration: VAR ids type SEMICOLON;
ids: IDENTIFIER (COMMA IDENTIFIER)* ;
functions: function*;  
function: FUNC IDENTIFIER parameters returnType? LBRACE declarations statements RBRACE;
parameters: LPAREN (decl (COMMA decl)*)? RPAREN;
returnType: type;
statements: statement*;
statement: assignment
         | print
         | read
         | delete
         | conditional
         | loop
         | return
         | invocation;

assignment: lValue EQUALS expression SEMICOLON;
print: PRINT LPAREN STRING (COMMA expression)* RPAREN SEMICOLON;
read: SCAN lValue SEMICOLON;
delete: DELETE expression SEMICOLON;
conditional: IF LPAREN expression RPAREN block (ELSE block)?;
loop: FOR LPAREN expression RPAREN block;
return: RETURN expression? SEMICOLON;
invocation: IDENTIFIER arguments SEMICOLON;
arguments:  LPAREN (expression (COMMA expression)*)? RPAREN;
block: LBRACE statements RBRACE;
lValue: IDENTIFIER (DOT IDENTIFIER)?;
expression: boolTerm (OR boolTerm)*;
boolTerm: equalTerm (AND equalTerm)*;
equalTerm: relationTerm ((EQ | NEQ) relationTerm)?;
relationTerm: simpleTerm ((LTN | GTN | LETN | GETN) simpleTerm)?;
simpleTerm: term ((PLUS | MINUS) term)*;
term: unaryTerm ((ASTERISK | FSLASH) unaryTerm)*;
unaryTerm: (NOT | MINUS | PLUS)? selectorTerm;
selectorTerm: factor (DOT IDENTIFIER)?;
factor: LPAREN expression RPAREN | IDENTIFIER arguments? | NUMBER | NEW IDENTIFIER | TRUE | FALSE | NIL;
