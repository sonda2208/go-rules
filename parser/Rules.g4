grammar Rules;

rules: expression EOF;

argumentExpressionList: '(' (expression ( ',' expression)*)? ')';

numberList: NumberLiteral ( ',' NumberLiteral)*;
stringList: StringLiteral ( ',' StringLiteral)*;

valueList: '[' (numberList+ | stringList+) ']';

primaryExpression:
	Identifier							# Ident
	| BooleanLiteral					# BoolLit
	| NumberLiteral						# NumberLit
	| StringLiteral						# StringLit
	| DurationLiteral					# DurationLit
	| Identifier argumentExpressionList	# FuncExpr
	| Identifier IN valueList			# InCond
	| LPARENT expression RPARENT		# ParentExpr;

booleanExpression:
	primaryExpression													# PrimaryExpr
	| booleanExpression op = (EQ | NEQ) primaryExpression				# EqualityExpr
	| booleanExpression op = (LT | LTE | GT | GTE) primaryExpression	# RelationalExpr;

expression:
	booleanExpression								# BooleanExpr
	| expression op = (AND | OR) booleanExpression	# LogicalExpr;

AND: '&&';
OR: '||';
EQ: '==';
NEQ: '!=';
LT: '<';
LTE: '<=';
GT: '>';
GTE: '>=';
IN: 'in';

LPARENT: '(';
RPARENT: ')';

BooleanLiteral: BOOL;
NumberLiteral: INTEGER | FLOAT;
DurationLiteral: DURATION;
StringLiteral: '"' CHAR_SEQUENCE* '"';
Identifier: ID_START ID_CHAR*;

fragment BOOL: 'true' | 'false';
fragment DIGIT: [0-9];
fragment DIGIT_SEQUENCE: DIGIT+;
fragment SIGN: '+' | '-';
fragment INTEGER: SIGN? DIGIT_SEQUENCE;
fragment FLOAT: INTEGER+ ('.' DIGIT_SEQUENCE);
fragment TIME_UNIT: 'ns' | 'ms' | 's' | 'm' | 'h';
fragment DURATION: INTEGER+ TIME_UNIT;
fragment CHAR_SEQUENCE: ~["\\\r\n];
fragment ID_START: ('a' .. 'z') | ('A' .. 'Z') | '_';
fragment ID_CHAR: ID_START | ('0' .. '9');

WS: [ \r\n\t]+ -> skip;