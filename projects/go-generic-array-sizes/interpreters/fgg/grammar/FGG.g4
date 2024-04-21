grammar FGG;

program: mainPackage declaration* mainFunc;

mainPackage: 'package' 'main';
mainFunc: 'func' 'main' '(' ')' '{' '_' '=' expression '}';

declaration:
	typeDeclaration
	| methodDeclaration
	| arraySetMethodDeclaration;

typeDeclaration:
	'type' typeName typeParameterConstraints? typeLiteral;

typeParameterConstraints:
	'[' typeParameterConstraint (',' typeParameterConstraint)* ']';

methodDeclaration:
	'func' methodReceiver methodSpecification '{' 'return' expression '}';
arraySetMethodDeclaration:
	'func' methodReceiver methodName '(' methodParameter ',' methodParameter ')' type '{' variable
		'[' variable ']' '=' variable ';' 'return' variable '}';

typeLiteral: structLiteral | interfaceLiteral | arrayLiteral;

structLiteral: 'struct' '{' field* '}';
field: fieldName type;

interfaceLiteral: 'interface' '{' methodSpecification* '}';

arrayLiteral: '[' type ']' type;

methodReceiver: '(' variable typeName typeParameters? ')';
methodSpecification: methodName methodSignature;
methodSignature: '(' methodParams ')' type;
methodParams: methodParameter? (',' methodParameter)*;
methodParameter: variable type;

type:
	typeName typeArguments?	# namedType
	| integerLiteral			# intType;

typeParameters: '[' typeParameter (',' typeParameter)* ']';
typeParameterConstraint: typeParameter bound;
bound: type | 'const';
typeArguments: '[' type (',' type)* ']';

integerLiteral: decimalLiteral;
decimalLiteral: DEC_DIGITS;

expression:
	integerLiteral										# intLiteral
	| variable											# var
	| expression '.' methodName '(' expressionList ')'	# methodCall
	| type '{' expressionList '}'						# valueLiteral
	| expression '.' fieldName							# fieldSelect
	| expression '[' expression ']'						# arrIndex
	| expression '+' expression                         # add;

expressionList: expression? (',' expression)*;

variable: ID;
typeName: ID;
methodName: ID;
fieldName: ID;
typeParameter: ID;

DEC_DIGITS:
	ZERO
	| POS_DECIMAL_DIGIT ('_'? (ZERO | POS_DECIMAL_DIGIT))*;
ZERO: '0';
POS_DECIMAL_DIGIT: [1-9];

ID: LETTER (LETTER | UNICODE_DIGIT)*;
LETTER: UNICODE_LETTER | '_';
UNICODE_LETTER: [A-Za-z];
UNICODE_DIGIT: [0-9];

WS: [ \r\n\t]+ -> skip;
