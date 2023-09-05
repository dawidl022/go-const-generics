grammar FG;

program: mainPackage declaration* mainFunc;

mainPackage: 'package' 'main';
mainFunc: 'func' 'main' '(' ')' '{' '_' '=' expression '}';

declaration:
	typeDeclaration
	| methodDeclaration
	| arraySetMethodDeclaration;

typeDeclaration: 'type' typeName typeLiteral;
methodDeclaration:
	'func' methodReceiver methodSpecification '{' 'return' expression '}';
arraySetMethodDeclaration:
	'func' methodReceiver methodName '(' methodParameter ',' methodParameter ')' typeName '{'
		variable '[' variable ']' '=' variable ';' 'return' variable '}';

typeLiteral: structLiteral | interfaceLiteral | arrayLiteral;

structLiteral: 'struct' '{' field* '}';
field: fieldName typeName;

interfaceLiteral: 'interface' '{' methodSpecification* '}';

methodReceiver: '(' methodParameter ')';
methodSpecification: methodName methodSignature;
methodSignature: '(' methodParams ')' typeName;
methodParams: methodParameter? (',' methodParameter)*;
methodParameter: variable typeName;

arrayLiteral: '[' integerLiteral ']' typeName;
integerLiteral: decimalLiteral;
decimalLiteral: ZERO | POS_DECIMAL_DIGIT ('_'? (ZERO | POS_DECIMAL_DIGIT))*;

expression:
	integerLiteral										# intLiteral
	| variable											# var
	| expression '.' methodName '(' expressionList ')'	# methodCall
	| typeName '{' expressionList '}'					# valueLiteral
	| expression '.' fieldName							# fieldSelect
	| expression '[' expression ']'						# arrIndex;

expressionList: expression? (',' expression)*;

variable: ID;
typeName: ID;
methodName: ID;
fieldName: ID;

ZERO: '0';
POS_DECIMAL_DIGIT: [1-9];

ID: LETTER (LETTER | UNICODE_DIGIT)*;
LETTER: UNICODE_LETTER | '_';
UNICODE_LETTER: [A-Za-z];
UNICODE_DIGIT: [0-9];

WS: [ \r\n\t]+ -> skip;
