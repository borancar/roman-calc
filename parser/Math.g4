grammar Math;

expr
    : LPAREN expr RPAREN     # Brace
    | expr POW expr          # Order
    | expr op=(DIV|MUL) expr # DivMul
    | expr op=(ADD|SUB) expr # AddSub
    | ROMAN                  # Num
    ;

ADD : '+' ;
SUB : '-' ;
MUL : '*' ;
DIV : '/' ;
POW : '^' ;
ROMAN : [IVXLCDM]+ ;
LPAREN : '(' ;
RPAREN : ')' ;

WS : [ \t\n\r]+ -> skip ;
