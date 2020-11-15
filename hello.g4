// hello.g4
grammar hello;

// Tokens
WHITESPACE: [ \t]+ -> skip;
NEWLINE: [\r\n]+;
NODE_EDGE_SEP: '<->';   
ID: [0-9]+;
EDGES_SEP: ',';

// Rules
start : expressions EOF;

edgeExpr
   : id
   | id EDGES_SEP edgeExpr
   ;

id: ID;

expression
   : id NODE_EDGE_SEP edgeExpr
   ;

expressions
   : expression NEWLINE expressions
   | expression
   ;
