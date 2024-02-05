lexer grammar GoliteLexer; 

PRINT: 'printf' ; 

STRUCT: 'struct' ;

TRUE: 'true' ; 

FALSE: 'false' ; 

NIL: 'nil' ;

NEW: 'new' ;

INT: 'int' ;  

BOOL: 'bool' ;  

SCAN: 'scan';

DELETE: 'delete';

IF: 'if';

ELSE: 'else';

FOR: 'for';

RETURN: 'return';

TYPE: 'type';

VAR: 'var';

FUNC: 'func';

NUMBER : ('+' | '-')? [0-9]+ ; 

STRING: '"' (~["])* '"';

LPAREN : '(' ; 

RPAREN : ')' ; 

LBRACE: '{' ;

RBRACE: '}' ; 

EQUALS: '=' ; 

GTN: '>' ; 

GETN: '>=' ;  

LTN: '<' ; 

LETN: '<=' ;

NOT: '!' ;

EQ: '==';

NEQ: '!=';

AND: '&&';

OR: '||';

PLUS: '+' ; 

MINUS: '-' ;

FSLASH: '/' ; 

ASTERISK: '*' ; 

COLON: ':' ;

IDENTIFIER: [a-zA-Z_][a-zA-Z0-9_]* ;

SEMICOLON: ';' ; 

DOT: '.' ;

COMMA: ',' ;

WS : [ \r\t\n]+ -> skip ;

COMMENT: '//' ~[\r\n]* -> skip;
