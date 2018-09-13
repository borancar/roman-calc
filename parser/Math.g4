grammar Math;

expr
    : LPAREN expr RPAREN # Brace
    | expr POW expr      # Order
    | expr DIV expr      # Div
    | expr MUL expr      # Mul
    | expr ADD expr      # Add
    | expr SUB expr      # Sub
    | ROMAN              # Num
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
