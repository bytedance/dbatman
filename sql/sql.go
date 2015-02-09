//line sql.y:6
package sqlparser

import __yyfmt__ "fmt"

//line sql.y:6
import "bytes"

func SetParseTree(yylex interface{}, stmt Statement) {
	yylex.(*Tokenizer).ParseTree = stmt
}

func SetAllowComments(yylex interface{}, allow bool) {
	yylex.(*Tokenizer).AllowComments = allow
}

func ForceEOF(yylex interface{}) {
	yylex.(*Tokenizer).ForceEOF = true
}

var (
	SHARE          = []byte("share")
	MODE           = []byte("mode")
	IF_BYTES       = []byte("if")
	VALUES_BYTES   = []byte("values")
	DATABASE_BYTES = []byte("database")
)

//line sql.y:32
type yySymType struct {
	yys         int
	empty       struct{}
	statement   Statement
	selStmt     SelectStatement
	byt         byte
	bytes       []byte
	bytes2      [][]byte
	str         string
	selectExprs SelectExprs
	selectExpr  SelectExpr
	columns     Columns
	colName     *ColName
	tableExprs  TableExprs
	tableExpr   TableExpr
	smTableExpr SimpleTableExpr
	tableName   *TableName
	indexHints  *IndexHints
	expr        Expr
	boolExpr    BoolExpr
	valExpr     ValExpr
	colTuple    ColTuple
	valExprs    ValExprs
	values      Values
	rowTuple    RowTuple
	subquery    *Subquery
	caseExpr    *CaseExpr
	whens       []*When
	when        *When
	orderBy     OrderBy
	order       *Order
	limit       *Limit
	insRows     InsertRows
	updateExprs UpdateExprs
	updateExpr  *UpdateExpr
}

const LEX_ERROR = 57346
const SELECT = 57347
const INSERT = 57348
const UPDATE = 57349
const DELETE = 57350
const FROM = 57351
const WHERE = 57352
const GROUP = 57353
const HAVING = 57354
const ORDER = 57355
const BY = 57356
const LIMIT = 57357
const FOR = 57358
const ALL = 57359
const DISTINCT = 57360
const AS = 57361
const EXISTS = 57362
const IN = 57363
const IS = 57364
const LIKE = 57365
const BETWEEN = 57366
const NULL = 57367
const ASC = 57368
const DESC = 57369
const VALUES = 57370
const INTO = 57371
const DUPLICATE = 57372
const KEY = 57373
const DEFAULT = 57374
const SET = 57375
const LOCK = 57376
const FUNCTION = 57377
const PROCEDURE = 57378
const TEMPORARY = 57379
const FULLTEXT = 57380
const PRIMARY = 57381
const AUTO_INCREMENT = 57382
const INDEXES = 57383
const KEYS = 57384
const VALUE = 57385
const ID = 57386
const STRING = 57387
const NUMBER = 57388
const VALUE_ARG = 57389
const LIST_ARG = 57390
const COMMENT = 57391
const GLOBAL = 57392
const SESSION = 57393
const LE = 57394
const GE = 57395
const NE = 57396
const NULL_SAFE_EQUAL = 57397
const UNION = 57398
const MINUS = 57399
const EXCEPT = 57400
const INTERSECT = 57401
const JOIN = 57402
const STRAIGHT_JOIN = 57403
const LEFT = 57404
const RIGHT = 57405
const INNER = 57406
const OUTER = 57407
const CROSS = 57408
const NATURAL = 57409
const USE = 57410
const FORCE = 57411
const ON = 57412
const OR = 57413
const AND = 57414
const NOT = 57415
const UNARY = 57416
const CASE = 57417
const WHEN = 57418
const THEN = 57419
const ELSE = 57420
const END = 57421
const BEGIN = 57422
const COMMIT = 57423
const ROLLBACK = 57424
const NAMES = 57425
const REPLACE = 57426
const ADMIN = 57427
const DATABASE = 57428
const DATABASES = 57429
const TABLES = 57430
const PROXY = 57431
const COLUMNS = 57432
const VARIABLES = 57433
const CREATE = 57434
const ALTER = 57435
const DROP = 57436
const RENAME = 57437
const ANALYZE = 57438
const TABLE = 57439
const INDEX = 57440
const VIEW = 57441
const TO = 57442
const IGNORE = 57443
const IF = 57444
const UNIQUE = 57445
const USING = 57446
const SHOW = 57447
const DESCRIBE = 57448
const EXPLAIN = 57449
const STATUS = 57450
const WARNINGS = 57451
const ERRORS = 57452
const BIT = 57453
const TINYINT = 57454
const SMALLINT = 57455
const MEDIUMINT = 57456
const INT = 57457
const INTEGER = 57458
const BIGINT = 57459
const REAL = 57460
const DOUBLE = 57461
const FLOAT = 57462
const DECIMAL = 57463
const DATE = 57464
const TIME = 57465
const TIMESTAMP = 57466
const DATETIME = 57467
const YEAR = 57468
const CHAR = 57469
const VARCHAR = 57470
const BINARY = 57471
const VARBINARY = 57472
const TINYBLOB = 57473
const BLOB = 57474
const MEDIUMBLOB = 57475
const LONGBLOB = 57476
const TINYTEXT = 57477
const TEXT = 57478
const MEDIUMTEXT = 57479
const LONGTEXT = 57480
const LONG = 57481
const NUMERIC = 57482
const BOOL = 57483
const ENUM = 57484
const CHARACTER = 57485
const ZEROFILL = 57486
const COLLATE = 57487
const UNSIGNED = 57488
const SIGNED = 57489
const PRECISION = 57490
const EXTENDED = 57491
const PARTITIONS = 57492

var yyToknames = []string{
	"LEX_ERROR",
	"SELECT",
	"INSERT",
	"UPDATE",
	"DELETE",
	"FROM",
	"WHERE",
	"GROUP",
	"HAVING",
	"ORDER",
	"BY",
	"LIMIT",
	"FOR",
	"ALL",
	"DISTINCT",
	"AS",
	"EXISTS",
	"IN",
	"IS",
	"LIKE",
	"BETWEEN",
	"NULL",
	"ASC",
	"DESC",
	"VALUES",
	"INTO",
	"DUPLICATE",
	"KEY",
	"DEFAULT",
	"SET",
	"LOCK",
	"FUNCTION",
	"PROCEDURE",
	"TEMPORARY",
	"FULLTEXT",
	"PRIMARY",
	"AUTO_INCREMENT",
	"INDEXES",
	"KEYS",
	"VALUE",
	"ID",
	"STRING",
	"NUMBER",
	"VALUE_ARG",
	"LIST_ARG",
	"COMMENT",
	"GLOBAL",
	"SESSION",
	"LE",
	"GE",
	"NE",
	"NULL_SAFE_EQUAL",
	" (",
	" =",
	" <",
	" >",
	" ~",
	"UNION",
	"MINUS",
	"EXCEPT",
	"INTERSECT",
	" ,",
	"JOIN",
	"STRAIGHT_JOIN",
	"LEFT",
	"RIGHT",
	"INNER",
	"OUTER",
	"CROSS",
	"NATURAL",
	"USE",
	"FORCE",
	"ON",
	"OR",
	"AND",
	"NOT",
	" &",
	" |",
	" ^",
	" +",
	" -",
	" *",
	" /",
	" %",
	" .",
	"UNARY",
	"CASE",
	"WHEN",
	"THEN",
	"ELSE",
	"END",
	"BEGIN",
	"COMMIT",
	"ROLLBACK",
	"NAMES",
	"REPLACE",
	"ADMIN",
	"DATABASE",
	"DATABASES",
	"TABLES",
	"PROXY",
	"COLUMNS",
	"VARIABLES",
	"CREATE",
	"ALTER",
	"DROP",
	"RENAME",
	"ANALYZE",
	"TABLE",
	"INDEX",
	"VIEW",
	"TO",
	"IGNORE",
	"IF",
	"UNIQUE",
	"USING",
	"SHOW",
	"DESCRIBE",
	"EXPLAIN",
	"STATUS",
	"WARNINGS",
	"ERRORS",
	"BIT",
	"TINYINT",
	"SMALLINT",
	"MEDIUMINT",
	"INT",
	"INTEGER",
	"BIGINT",
	"REAL",
	"DOUBLE",
	"FLOAT",
	"DECIMAL",
	"DATE",
	"TIME",
	"TIMESTAMP",
	"DATETIME",
	"YEAR",
	"CHAR",
	"VARCHAR",
	"BINARY",
	"VARBINARY",
	"TINYBLOB",
	"BLOB",
	"MEDIUMBLOB",
	"LONGBLOB",
	"TINYTEXT",
	"TEXT",
	"MEDIUMTEXT",
	"LONGTEXT",
	"LONG",
	"NUMERIC",
	"BOOL",
	"ENUM",
	"CHARACTER",
	"ZEROFILL",
	"COLLATE",
	"UNSIGNED",
	"SIGNED",
	"PRECISION",
	"EXTENDED",
	"PARTITIONS",
}
var yyStatenames = []string{}

const yyEofCode = 1
const yyErrCode = 2
const yyMaxDepth = 200

//line yacctab:1
var yyExca = []int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 26,
	113, 344,
	-2, 333,
	-1, 415,
	1, 349,
	-2, 348,
	-1, 537,
	142, 140,
	160, 140,
	-2, 136,
}

const yyNprod = 350
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 1125

var yyAct = []int{

	158, 592, 608, 178, 549, 587, 353, 459, 622, 356,
	253, 417, 156, 402, 346, 144, 155, 411, 404, 302,
	316, 166, 149, 267, 179, 145, 263, 183, 89, 43,
	44, 45, 46, 227, 226, 348, 86, 252, 3, 376,
	377, 378, 379, 380, 126, 381, 382, 666, 342, 641,
	161, 654, 654, 654, 409, 165, 110, 111, 171, 640,
	550, 409, 114, 409, 409, 645, 118, 532, 643, 633,
	409, 409, 409, 125, 148, 162, 163, 164, 91, 102,
	95, 210, 637, 409, 464, 221, 153, 300, 409, 221,
	169, 221, 137, 300, 300, 636, 242, 635, 634, 129,
	638, 555, 556, 557, 554, 181, 123, 87, 309, 152,
	190, 276, 108, 167, 168, 146, 189, 367, 117, 450,
	173, 180, 337, 196, 239, 240, 241, 518, 141, 175,
	517, 172, 177, 107, 338, 347, 271, 531, 510, 444,
	516, 387, 237, 238, 239, 240, 241, 170, 223, 217,
	209, 203, 656, 655, 653, 652, 84, 85, 242, 213,
	214, 215, 648, 109, 647, 646, 644, 130, 204, 642,
	632, 619, 618, 615, 541, 250, 251, 181, 242, 150,
	266, 511, 216, 259, 614, 463, 449, 113, 446, 408,
	397, 112, 395, 180, 339, 299, 340, 88, 92, 93,
	310, 56, 283, 59, 105, 420, 205, 61, 63, 62,
	211, 225, 422, 419, 142, 272, 440, 442, 182, 226,
	218, 281, 277, 227, 226, 308, 312, 67, 68, 504,
	291, 325, 207, 284, 87, 255, 212, 269, 452, 256,
	227, 226, 306, 336, 403, 361, 193, 307, 403, 313,
	314, 318, 293, 199, 505, 262, 269, 294, 441, 296,
	297, 347, 515, 400, 106, 181, 181, 53, 352, 514,
	181, 432, 312, 341, 343, 311, 433, 362, 430, 513,
	55, 180, 354, 431, 58, 326, 180, 421, 218, 512,
	355, 358, 423, 181, 359, 435, 369, 280, 282, 279,
	434, 351, 300, 150, 64, 65, 66, 524, 454, 180,
	315, 503, 268, 323, 324, 364, 327, 328, 330, 331,
	332, 333, 334, 335, 386, 373, 306, 368, 388, 351,
	526, 527, 436, 270, 437, 438, 499, 268, 532, 150,
	150, 389, 318, 574, 234, 235, 236, 237, 238, 239,
	240, 241, 360, 116, 43, 44, 45, 46, 396, 220,
	399, 21, 405, 405, 181, 410, 407, 374, 319, 406,
	401, 376, 377, 378, 379, 380, 317, 381, 382, 573,
	418, 572, 372, 242, 234, 235, 236, 237, 238, 239,
	240, 241, 269, 420, 428, 429, 306, 306, 571, 218,
	422, 419, 305, 390, 391, 564, 415, 82, 83, 563,
	181, 394, 363, 560, 304, 221, 84, 85, 657, 559,
	119, 120, 121, 242, 150, 558, 455, 548, 456, 467,
	470, 471, 472, 469, 473, 474, 475, 476, 477, 478,
	480, 481, 482, 483, 484, 485, 486, 487, 488, 489,
	490, 491, 492, 493, 494, 495, 496, 497, 479, 468,
	498, 292, 264, 546, 206, 181, 506, 522, 70, 71,
	72, 76, 73, 75, 448, 421, 529, 21, 74, 81,
	423, 418, 451, 502, 265, 265, 501, 260, 258, 257,
	78, 79, 138, 174, 457, 460, 628, 629, 186, 551,
	552, 553, 181, 181, 80, 609, 181, 181, 561, 562,
	184, 185, 565, 566, 602, 601, 305, 567, 354, 354,
	569, 586, 354, 354, 667, 610, 662, 659, 304, 519,
	612, 568, 577, 570, 520, 385, 87, 224, 182, 135,
	588, 589, 590, 591, 593, 594, 595, 600, 597, 598,
	599, 603, 445, 604, 605, 606, 443, 104, 425, 414,
	384, 181, 87, 21, 181, 181, 287, 613, 285, 91,
	616, 617, 181, 623, 623, 623, 181, 354, 621, 626,
	354, 354, 624, 625, 219, 620, 349, 202, 354, 201,
	200, 218, 180, 197, 194, 188, 139, 115, 94, 658,
	650, 350, 639, 187, 176, 576, 460, 533, 534, 535,
	536, 537, 538, 539, 540, 542, 543, 544, 651, 136,
	631, 579, 630, 545, 521, 547, 500, 583, 580, 140,
	453, 122, 161, 101, 143, 582, 584, 165, 627, 393,
	171, 320, 660, 321, 322, 585, 661, 127, 664, 133,
	191, 192, 274, 198, 195, 665, 148, 162, 163, 164,
	128, 134, 98, 96, 412, 668, 208, 509, 153, 669,
	21, 413, 169, 357, 508, 578, 124, 427, 268, 596,
	131, 47, 103, 663, 575, 161, 21, 48, 17, 15,
	165, 152, 14, 171, 13, 167, 168, 146, 12, 611,
	607, 465, 173, 530, 49, 50, 51, 52, 466, 182,
	162, 163, 164, 172, 581, 69, 416, 90, 100, 275,
	54, 153, 366, 278, 60, 169, 649, 132, 273, 170,
	77, 37, 57, 525, 21, 458, 507, 286, 426, 398,
	288, 289, 290, 261, 152, 345, 160, 157, 167, 168,
	159, 154, 228, 295, 165, 173, 298, 171, 21, 22,
	23, 24, 151, 439, 303, 375, 172, 301, 147, 383,
	222, 97, 42, 182, 162, 163, 164, 99, 344, 523,
	40, 20, 170, 19, 16, 206, 25, 18, 11, 169,
	10, 9, 8, 7, 234, 235, 236, 237, 238, 239,
	240, 241, 447, 6, 234, 235, 236, 237, 238, 239,
	240, 241, 167, 168, 5, 4, 2, 1, 365, 173,
	0, 0, 0, 0, 0, 0, 0, 38, 0, 0,
	172, 254, 0, 242, 0, 0, 370, 371, 0, 0,
	0, 0, 0, 242, 0, 0, 170, 0, 30, 31,
	32, 0, 33, 35, 229, 233, 231, 232, 0, 0,
	26, 27, 29, 28, 36, 0, 0, 0, 0, 0,
	0, 0, 0, 34, 41, 39, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 246, 247, 248, 249, 0,
	243, 244, 245, 0, 0, 254, 0, 161, 0, 0,
	0, 0, 165, 0, 424, 171, 0, 0, 0, 0,
	0, 0, 230, 234, 235, 236, 237, 238, 239, 240,
	241, 148, 162, 163, 164, 0, 0, 0, 0, 0,
	0, 0, 0, 153, 0, 0, 0, 169, 0, 0,
	0, 0, 0, 0, 0, 161, 0, 0, 0, 0,
	165, 0, 242, 171, 461, 462, 152, 0, 0, 21,
	167, 168, 146, 0, 0, 0, 0, 173, 0, 182,
	162, 163, 164, 0, 0, 0, 0, 0, 172, 165,
	0, 153, 171, 0, 392, 169, 234, 235, 236, 237,
	238, 239, 240, 241, 170, 0, 0, 0, 182, 162,
	163, 164, 0, 528, 152, 0, 0, 0, 167, 168,
	206, 0, 0, 0, 169, 173, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 242, 172, 165, 0, 0,
	171, 0, 165, 0, 0, 171, 0, 167, 168, 0,
	0, 0, 170, 0, 173, 0, 182, 162, 163, 164,
	0, 182, 162, 163, 164, 172, 0, 0, 206, 0,
	0, 0, 169, 206, 0, 0, 0, 169, 0, 0,
	0, 170, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 329, 0, 167, 168, 0, 0, 0,
	167, 168, 173, 0, 0, 0, 0, 173, 0, 0,
	0, 0, 0, 172, 0, 0, 0, 0, 172, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 170,
	0, 0, 0, 0, 170,
}
var yyPact = []int{

	753, -1000, -1000, 293, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, 166, 93, 96, 192,
	-1000, -1000, -1000, -1000, 366, 492, 85, 34, 554, -1000,
	-1000, -1000, 681, 646, -1000, -1000, -1000, 644, -1000, 604,
	525, 673, 106, -5, 50, 492, 492, 79, -1000, -1000,
	75, 492, -1000, 553, 1, 492, 1, 1, 1, 602,
	-1000, 667, 492, 637, -24, 55, 671, 640, -1000, -1000,
	-31, -1000, -1000, -1000, -1000, -1000, 436, -1000, 552, -1000,
	681, 126, -1000, -1000, -1000, -1000, -1000, 877, -1000, 444,
	525, -1000, 571, 525, 494, 466, 570, 551, 37, 492,
	-1000, -1000, -5, 550, -1000, 8, 549, 633, 177, 546,
	545, 543, 525, 637, 1007, 667, -1000, 925, 1007, 667,
	525, 525, 525, -1000, -1000, -1000, -1000, 637, 1007, -1000,
	-1000, 293, 540, -1000, 350, -1000, -1000, 518, 123, 163,
	833, -1000, 925, 665, -1000, -1000, -1000, 1007, 433, 432,
	-1000, 431, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, 1007, -1000, 429, 494, 668, 191, -1000,
	276, -1000, 48, -1000, -1000, -1000, -1000, 466, -1000, 632,
	-8, -1000, -1000, 525, 189, -1000, 524, -1000, -1000, 522,
	-1000, -1000, -1000, 428, -1000, 264, 729, 637, -1000, 163,
	833, 264, 637, -1000, 637, 637, -1000, 29, 264, -1000,
	358, 877, -1000, -1000, 63, 190, 925, 925, 1007, 320,
	620, 1007, 1007, 206, 1007, 1002, 1007, 1007, 1007, 1007,
	1007, 1007, 492, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -44, -32, 28, -1000, 833, -23, 30, 612, -1000,
	681, 44, 264, 558, 494, 494, 327, 660, 925, 494,
	1007, 492, -1000, -1000, -1000, 169, 492, 356, -1000, 2,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, 558, 494, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	1007, 302, 305, 516, 472, 53, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, 141, 264, -1000, 954, -1000, -1000,
	320, 1007, 1007, 264, 906, -1000, 614, 59, 59, 1007,
	59, 39, 39, -23, -23, -23, -1000, -1000, -1000, -1000,
	-1000, 26, 877, 24, -1000, 170, -1000, 925, 168, 408,
	408, 293, 172, 23, -1000, 660, 649, 657, 163, -1000,
	264, 515, -1000, 362, 293, -1000, 514, -1000, -1000, 191,
	-1000, -1000, 264, 666, 358, 358, -1000, -1000, 212, 205,
	234, 229, 266, 142, -1000, 512, -27, 508, 22, -1000,
	264, 724, 1007, -1000, 59, -1000, 20, -1000, 25, -1000,
	1007, 146, -1000, 600, 243, -1000, 243, -1000, -1000, 494,
	649, -1000, 1007, 1007, -1000, 48, 19, -1000, 303, 595,
	430, 427, 198, 410, -1000, -1000, 662, 653, 305, 62,
	-1000, 223, -1000, 213, -1000, -1000, -1000, 203, 196, -1000,
	27, 17, 14, -1000, -1000, -1000, -1000, 1007, 264, -1000,
	-1000, 264, 1007, 593, 408, -1000, -1000, 714, 242, -1000,
	304, -1000, -1000, -1000, 174, -1000, -1000, 282, -1000, 282,
	282, 282, 282, 282, 282, 282, 11, 282, 282, 282,
	-1000, -1000, -1000, -1000, -1000, 282, 407, 282, 371, -1000,
	-1000, -1000, -1000, -84, -84, -84, -84, -41, 369, 363,
	357, 494, 494, 353, 349, 494, 494, 660, 925, 1007,
	925, 342, -1000, -1000, -1000, -1000, 325, 323, 287, 264,
	264, 677, -1000, 1007, 1007, -1000, -1000, -1000, -1000, -1000,
	596, -1000, 475, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, 282, -1000, -1000, -1000, -1000, 469, -1000, 468, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 480, 485,
	494, 18, 7, 494, 494, 6, 5, 649, 163, 237,
	163, 494, 492, 492, 492, 494, 264, -1000, 613, -1000,
	451, 591, 589, -1000, -1000, -1000, 4, -64, -64, -64,
	-64, -64, -60, -64, -64, -64, -1000, -64, -64, -64,
	-60, -107, -117, -60, -60, -60, -60, 3, -1000, -1000,
	-1000, 0, -1000, -1, -1000, -1000, -2, -4, -1000, -1000,
	584, -11, -12, -1000, -13, -14, 191, -1000, -1000, -1000,
	-1000, -1000, -1000, 372, -1000, -1000, -1000, 566, 482, -64,
	-1000, -1000, -1000, 480, -1000, 481, -1000, -1000, -1000, -1000,
	676, 627, -1000, -1000, 492, -1000, -1000, -119, 479, -1000,
	-60, -1000, -1000, -1000, 492, -1000, -1000, -1000, 492, -1000,
}
var yyPgo = []int{

	0, 817, 816, 37, 815, 814, 803, 793, 792, 791,
	790, 788, 787, 784, 783, 781, 681, 777, 772, 771,
	15, 25, 770, 769, 504, 27, 768, 767, 19, 765,
	764, 28, 763, 8, 23, 22, 762, 752, 35, 751,
	81, 20, 10, 18, 12, 750, 21, 747, 16, 746,
	745, 14, 743, 739, 738, 736, 9, 735, 7, 733,
	353, 133, 732, 137, 5, 1, 4, 731, 730, 727,
	17, 726, 26, 6, 13, 3, 24, 724, 723, 722,
	720, 719, 718, 717, 0, 539, 716, 11, 708, 703,
	701, 700, 2, 699, 698, 694, 692, 689, 688, 106,
	44, 687,
}
var yyR1 = []int{

	0, 1, 2, 2, 2, 2, 2, 2, 2, 2,
	2, 2, 2, 2, 2, 2, 2, 2, 2, 2,
	3, 3, 3, 4, 4, 97, 97, 5, 6, 7,
	7, 7, 25, 25, 25, 24, 24, 24, 94, 95,
	96, 98, 13, 13, 13, 13, 13, 13, 13, 13,
	13, 13, 13, 68, 68, 68, 69, 69, 8, 8,
	8, 8, 8, 8, 8, 86, 86, 87, 87, 87,
	87, 87, 87, 87, 87, 90, 89, 89, 89, 89,
	89, 89, 89, 89, 89, 89, 89, 88, 88, 88,
	88, 88, 88, 88, 88, 88, 88, 88, 88, 88,
	88, 88, 88, 88, 88, 88, 88, 88, 88, 88,
	88, 88, 88, 88, 88, 88, 88, 88, 88, 88,
	88, 88, 88, 88, 88, 91, 91, 92, 92, 93,
	93, 63, 63, 63, 66, 66, 64, 64, 64, 64,
	65, 65, 65, 9, 9, 9, 10, 11, 11, 11,
	11, 11, 12, 15, 14, 14, 83, 83, 83, 67,
	67, 67, 101, 16, 17, 17, 18, 18, 18, 18,
	18, 19, 19, 20, 20, 21, 21, 21, 26, 26,
	22, 22, 22, 22, 22, 27, 27, 28, 28, 28,
	28, 28, 23, 23, 23, 29, 29, 29, 29, 29,
	29, 29, 29, 29, 29, 29, 30, 30, 30, 31,
	31, 32, 32, 32, 32, 33, 33, 34, 34, 100,
	100, 100, 99, 99, 35, 35, 35, 35, 35, 36,
	36, 36, 36, 36, 36, 36, 36, 36, 36, 37,
	37, 37, 37, 37, 37, 37, 41, 41, 41, 46,
	42, 42, 40, 40, 40, 40, 40, 40, 40, 40,
	40, 40, 40, 40, 40, 40, 40, 40, 40, 40,
	40, 40, 45, 45, 45, 47, 47, 47, 49, 52,
	52, 50, 50, 51, 53, 53, 48, 48, 39, 39,
	39, 39, 54, 54, 55, 55, 56, 56, 57, 57,
	58, 59, 59, 59, 70, 70, 70, 71, 71, 71,
	72, 72, 73, 73, 74, 74, 38, 38, 38, 43,
	43, 44, 44, 44, 75, 75, 76, 82, 82, 60,
	60, 61, 61, 62, 62, 77, 77, 78, 78, 78,
	78, 78, 79, 79, 80, 80, 81, 81, 84, 85,
}
var yyR2 = []int{

	0, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	4, 12, 3, 7, 7, 6, 6, 8, 7, 4,
	4, 5, 1, 1, 1, 0, 1, 1, 1, 1,
	1, 5, 2, 4, 5, 4, 5, 5, 6, 6,
	3, 3, 5, 1, 1, 1, 1, 1, 5, 8,
	4, 4, 8, 9, 7, 1, 3, 2, 5, 4,
	4, 5, 5, 4, 4, 2, 0, 3, 2, 3,
	3, 2, 3, 3, 2, 2, 2, 2, 1, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 4, 3,
	3, 3, 1, 1, 1, 1, 1, 3, 5, 2,
	4, 1, 1, 1, 1, 3, 3, 3, 3, 2,
	2, 2, 2, 4, 4, 1, 3, 1, 1, 1,
	3, 0, 3, 5, 0, 1, 0, 2, 2, 2,
	0, 4, 3, 6, 7, 4, 5, 5, 5, 5,
	5, 5, 3, 3, 3, 3, 0, 1, 1, 1,
	1, 1, 0, 2, 0, 2, 1, 2, 1, 1,
	1, 0, 1, 1, 3, 1, 2, 3, 1, 1,
	0, 1, 2, 2, 2, 1, 3, 3, 3, 3,
	5, 7, 0, 1, 2, 1, 1, 2, 3, 2,
	3, 2, 2, 2, 3, 3, 1, 3, 1, 1,
	3, 0, 5, 5, 5, 1, 3, 0, 2, 0,
	2, 2, 0, 2, 1, 3, 3, 2, 3, 3,
	3, 4, 3, 4, 5, 6, 3, 4, 2, 1,
	1, 1, 1, 1, 1, 1, 3, 1, 1, 3,
	1, 3, 1, 1, 1, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 4, 2, 3, 4, 5, 4,
	3, 1, 1, 1, 1, 1, 1, 1, 5, 0,
	1, 1, 2, 4, 0, 2, 1, 3, 1, 1,
	1, 1, 0, 3, 0, 2, 0, 3, 1, 3,
	2, 0, 1, 1, 0, 2, 4, 0, 2, 4,
	0, 3, 1, 3, 0, 5, 2, 2, 1, 1,
	3, 3, 2, 1, 1, 3, 3, 0, 1, 0,
	2, 0, 3, 0, 1, 0, 1, 1, 1, 1,
	1, 1, 0, 1, 0, 1, 0, 2, 1, 0,
}
var yyChk = []int{

	-1000, -1, -2, -3, -4, -5, -6, -7, -8, -9,
	-10, -11, -94, -95, -96, -97, -13, -98, -12, -14,
	-15, 5, 6, 7, 8, 33, 107, 108, 110, 109,
	95, 96, 97, 99, 120, 100, 111, -67, 74, 122,
	27, 121, -18, 61, 62, 63, 64, -16, -101, -16,
	-16, -16, -16, 101, -80, 114, 35, -62, 118, 37,
	-77, 114, 116, 112, 112, 113, 114, 35, 36, -16,
	102, 103, 104, 106, 112, 107, 105, -68, 124, 125,
	-24, 113, 41, 42, 50, 51, -84, 44, 112, -31,
	-83, 44, 164, 165, 44, -3, 17, -19, 18, -17,
	-82, 29, -31, 9, -24, 98, 158, -61, 117, 113,
	-84, -84, 112, 112, -84, 44, -60, 117, -84, -60,
	-60, -60, 29, -99, 9, -84, -100, 10, 23, 123,
	112, 9, -69, 9, 21, -85, -85, 123, 56, 44,
	-85, -3, 88, -85, -20, -21, 85, -26, 44, -35,
	-40, -36, 79, 56, -39, -48, -44, -47, -84, -45,
	-49, 20, 45, 46, 47, 25, -46, 83, 84, 60,
	117, 28, 101, 90, 49, -31, 33, -31, -75, -76,
	-48, -84, 44, -25, 44, 45, 32, 33, 44, 79,
	-84, -85, -85, -61, 44, -85, 115, 44, 20, 76,
	44, 44, 44, -31, -100, -40, 56, -99, -85, -35,
	-40, -40, -99, -31, -31, -31, -100, -42, -40, 44,
	9, 65, -22, -84, 19, 88, 78, 77, -37, 21,
	79, 23, 24, 22, 80, 81, 82, 83, 84, 85,
	86, 87, 119, 57, 58, 59, 52, 53, 54, 55,
	-35, -35, -3, -42, 166, -40, -40, 56, 56, -46,
	56, -52, -40, -72, 33, 56, -75, -34, 10, 65,
	57, 88, -25, -85, 20, -81, 119, -31, -78, 110,
	108, 32, 109, 13, 44, 44, -85, 44, -85, -85,
	-85, -72, 33, -100, -100, -85, -100, -100, -85, 166,
	65, -27, -28, -30, 56, 44, -46, -21, -84, 45,
	137, 85, -84, -35, -35, -40, -41, 56, -46, 48,
	21, 23, 24, -40, -40, 25, 79, -40, -40, 81,
	-40, -40, -40, -40, -40, -40, -84, 166, 166, 166,
	166, -20, 18, -20, 166, -50, -51, 91, -38, 28,
	43, -3, -75, -73, -48, -34, -56, 13, -35, -76,
	-40, 76, -84, 56, -3, -85, -79, 115, -38, -75,
	-85, -85, -40, -34, 65, -29, 66, 67, 68, 69,
	70, 72, 73, -23, 44, 19, -28, 88, -42, -41,
	-40, -40, 78, 25, -40, 166, -20, 166, -53, -51,
	93, -35, -74, 76, -43, -44, -43, -74, 166, 65,
	-56, -70, 15, 14, 44, 44, -86, -87, -48, 39,
	31, 113, 38, 118, -85, 44, -54, 11, -28, -28,
	66, 71, 66, 71, 66, 66, 66, 68, 69, -32,
	74, 116, 75, 44, 166, 44, 166, 78, -40, 166,
	94, -40, 92, 30, 65, -48, -70, -40, -57, -58,
	-40, -85, -85, 166, 65, -90, -88, 126, 156, 130,
	127, 128, 129, 131, 132, 133, 134, 135, 136, 155,
	137, 138, 139, 140, 141, 142, 143, 144, 145, 146,
	147, 148, 149, 150, 151, 152, 153, 154, 157, 33,
	31, 56, 56, 113, 31, 56, 56, -55, 12, 14,
	76, 119, 66, 66, 66, 66, 113, 113, 113, -40,
	-40, 31, -44, 65, 65, -59, 26, 27, -85, -87,
	-89, -63, 56, -63, -63, -63, -63, -63, -63, -63,
	-63, 163, -63, -63, -63, -63, 56, -63, 56, -66,
	144, -66, -66, -66, 145, 142, 143, 144, 56, 56,
	56, -73, -73, 56, 56, -73, -73, -56, -35, -42,
	-35, 56, 56, 56, 56, 7, -40, -58, 79, 25,
	32, 118, 39, 31, 40, 49, 46, -64, -64, -64,
	-64, -64, -65, -64, -64, -64, -63, -64, -64, -64,
	-65, 46, 46, -65, -65, -65, -65, -91, -92, 25,
	45, -93, 45, -73, 166, 166, -73, -73, 166, 166,
	-70, -73, -33, -84, -33, -33, -75, 25, 45, 46,
	31, 31, 166, 65, 162, 161, 159, 142, 160, -64,
	166, 166, 166, 65, 166, 65, 166, 166, 166, -71,
	16, 34, 166, 166, 65, 166, 166, 46, 33, 45,
	-65, -92, 45, 7, 21, -84, 166, 45, -84, -84,
}
var yyDef = []int{

	0, -2, 1, 2, 3, 4, 5, 6, 7, 8,
	9, 10, 11, 12, 13, 14, 15, 16, 17, 18,
	19, 162, 162, 162, 162, 162, -2, 335, 0, 0,
	38, 39, 40, 162, 35, 0, 0, 156, 0, 159,
	160, 161, 0, 166, 168, 169, 170, 171, 164, 327,
	0, 0, 35, 331, 0, 0, 0, 0, 345, 334,
	0, 0, 336, 0, 329, 0, 329, 329, 329, 0,
	42, 222, 0, 219, 0, 0, 0, 0, 349, 349,
	0, 53, 54, 55, 36, 37, 0, 348, 0, 349,
	0, 209, 157, 158, 349, 22, 167, 0, 172, 163,
	0, 328, 0, 0, 0, 0, 0, 0, 0, 0,
	349, 349, 331, 0, 349, 0, 0, 0, 0, 0,
	0, 0, 0, 219, 0, 222, 349, 0, 0, 222,
	0, 0, 0, 56, 57, 50, 51, 219, 0, 152,
	154, 155, 0, 153, 20, 173, 175, 180, 348, 178,
	179, 224, 0, 0, 252, 253, 254, 0, 286, 0,
	271, 0, 288, 289, 290, 291, 323, 275, 276, 277,
	272, 273, 274, 279, 165, 310, 0, 217, 29, 324,
	0, 286, 348, 30, 32, 33, 34, 0, 349, 0,
	346, 60, 61, 0, 0, 145, 0, 349, 330, 0,
	349, 349, 349, 310, 43, 223, 0, 219, 45, 220,
	0, 221, 219, 349, 219, 219, 349, 0, 250, 210,
	0, 0, 176, 181, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 239, 240, 241, 242, 243, 244, 245,
	227, 0, 0, 0, 322, 250, 265, 0, 0, 238,
	0, 0, 280, 0, 0, 0, 217, 296, 0, 0,
	0, 0, 31, 58, 332, 0, 0, 0, 349, 342,
	337, 338, 339, 340, 341, 146, 147, 148, 149, 150,
	151, 0, 0, 44, 46, 47, 349, 349, 52, 41,
	0, 217, 185, 192, 0, 206, 208, 174, 182, 183,
	184, 177, 287, 225, 226, 229, 230, 0, 247, 248,
	0, 0, 0, 232, 0, 236, 0, 255, 256, 0,
	257, 258, 259, 260, 261, 262, 263, 228, 249, 321,
	266, 0, 0, 0, 270, 284, 281, 0, 314, 0,
	0, 318, 314, 0, 312, 296, 304, 0, 218, 325,
	326, 0, 347, 0, 349, 143, 0, 343, 25, 26,
	48, 49, 251, 292, 0, 0, 195, 196, 0, 0,
	0, 0, 0, 211, 193, 0, 0, 0, 0, 231,
	233, 0, 0, 237, 264, 267, 0, 269, 0, 282,
	0, 0, 23, 0, 316, 319, 317, 24, 311, 0,
	304, 28, 0, 0, 349, -2, 0, 65, 0, 0,
	0, 0, 0, 0, 64, 144, 294, 0, 186, 189,
	197, 0, 199, 0, 201, 202, 203, 0, 0, 187,
	0, 0, 0, 194, 188, 207, 246, 0, 234, 268,
	278, 285, 0, 0, 0, 313, 27, 305, 297, 298,
	301, 59, 62, 349, 0, 67, 76, 131, 88, 131,
	131, 131, 131, 131, 131, 131, 131, 131, 131, 131,
	102, 103, 104, 105, 106, 131, 0, 131, 0, 111,
	112, 113, 114, 134, 134, 134, 134, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 296, 0, 0,
	0, 0, 198, 200, 204, 205, 0, 0, 0, 235,
	283, 0, 320, 0, 0, 300, 302, 303, 63, 66,
	75, 87, 0, 136, 136, 136, 136, -2, 136, 136,
	136, 131, 136, 136, 136, 140, 0, 109, 0, 140,
	135, 140, 140, 140, 119, 120, 121, 122, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 304, 295, 293,
	190, 0, 0, 0, 0, 0, 306, 299, 0, 78,
	0, 81, 0, 84, 85, 86, 0, 89, 90, 91,
	92, 93, 94, 95, 96, 97, 136, 99, 100, 101,
	107, 0, 0, 115, 116, 117, 118, 0, 125, 127,
	128, 0, 129, 0, 69, 70, 0, 0, 73, 74,
	307, 0, 0, 215, 0, 0, 315, 77, 79, 80,
	82, 83, 132, 0, 137, 138, 139, 0, 0, 98,
	140, 110, 123, 0, 124, 0, 68, 71, 72, 21,
	0, 0, 191, 212, 0, 213, 214, 0, 0, 142,
	108, 126, 130, 308, 0, 216, 133, 141, 0, 309,
}
var yyTok1 = []int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 87, 80, 3,
	56, 166, 85, 83, 65, 84, 88, 86, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	58, 57, 59, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 82, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 81, 3, 60,
}
var yyTok2 = []int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 37, 38, 39, 40, 41,
	42, 43, 44, 45, 46, 47, 48, 49, 50, 51,
	52, 53, 54, 55, 61, 62, 63, 64, 66, 67,
	68, 69, 70, 71, 72, 73, 74, 75, 76, 77,
	78, 79, 89, 90, 91, 92, 93, 94, 95, 96,
	97, 98, 99, 100, 101, 102, 103, 104, 105, 106,
	107, 108, 109, 110, 111, 112, 113, 114, 115, 116,
	117, 118, 119, 120, 121, 122, 123, 124, 125, 126,
	127, 128, 129, 130, 131, 132, 133, 134, 135, 136,
	137, 138, 139, 140, 141, 142, 143, 144, 145, 146,
	147, 148, 149, 150, 151, 152, 153, 154, 155, 156,
	157, 158, 159, 160, 161, 162, 163, 164, 165,
}
var yyTok3 = []int{
	0,
}

//line yaccpar:1

/*	parser for yacc output	*/

var yyDebug = 0

type yyLexer interface {
	Lex(lval *yySymType) int
	Error(s string)
}

const yyFlag = -1000

func yyTokname(c int) string {
	// 4 is TOKSTART above
	if c >= 4 && c-4 < len(yyToknames) {
		if yyToknames[c-4] != "" {
			return yyToknames[c-4]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func yyStatname(s int) string {
	if s >= 0 && s < len(yyStatenames) {
		if yyStatenames[s] != "" {
			return yyStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func yylex1(lex yyLexer, lval *yySymType) int {
	c := 0
	char := lex.Lex(lval)
	if char <= 0 {
		c = yyTok1[0]
		goto out
	}
	if char < len(yyTok1) {
		c = yyTok1[char]
		goto out
	}
	if char >= yyPrivate {
		if char < yyPrivate+len(yyTok2) {
			c = yyTok2[char-yyPrivate]
			goto out
		}
	}
	for i := 0; i < len(yyTok3); i += 2 {
		c = yyTok3[i+0]
		if c == char {
			c = yyTok3[i+1]
			goto out
		}
	}

out:
	if c == 0 {
		c = yyTok2[1] /* unknown char */
	}
	if yyDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", yyTokname(c), uint(char))
	}
	return c
}

func yyParse(yylex yyLexer) int {
	var yyn int
	var yylval yySymType
	var yyVAL yySymType
	yyS := make([]yySymType, yyMaxDepth)

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yystate := 0
	yychar := -1
	yyp := -1
	goto yystack

ret0:
	return 0

ret1:
	return 1

yystack:
	/* put a state and value onto the stack */
	if yyDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", yyTokname(yychar), yyStatname(yystate))
	}

	yyp++
	if yyp >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyS[yyp] = yyVAL
	yyS[yyp].yys = yystate

yynewstate:
	yyn = yyPact[yystate]
	if yyn <= yyFlag {
		goto yydefault /* simple state */
	}
	if yychar < 0 {
		yychar = yylex1(yylex, &yylval)
	}
	yyn += yychar
	if yyn < 0 || yyn >= yyLast {
		goto yydefault
	}
	yyn = yyAct[yyn]
	if yyChk[yyn] == yychar { /* valid shift */
		yychar = -1
		yyVAL = yylval
		yystate = yyn
		if Errflag > 0 {
			Errflag--
		}
		goto yystack
	}

yydefault:
	/* default state action */
	yyn = yyDef[yystate]
	if yyn == -2 {
		if yychar < 0 {
			yychar = yylex1(yylex, &yylval)
		}

		/* look through exception table */
		xi := 0
		for {
			if yyExca[xi+0] == -1 && yyExca[xi+1] == yystate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			yyn = yyExca[xi+0]
			if yyn < 0 || yyn == yychar {
				break
			}
		}
		yyn = yyExca[xi+1]
		if yyn < 0 {
			goto ret0
		}
	}
	if yyn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			yylex.Error("syntax error")
			Nerrs++
			if yyDebug >= 1 {
				__yyfmt__.Printf("%s", yyStatname(yystate))
				__yyfmt__.Printf(" saw %s\n", yyTokname(yychar))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for yyp >= 0 {
				yyn = yyPact[yyS[yyp].yys] + yyErrCode
				if yyn >= 0 && yyn < yyLast {
					yystate = yyAct[yyn] /* simulate a shift of "error" */
					if yyChk[yystate] == yyErrCode {
						goto yystack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if yyDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", yyS[yyp].yys)
				}
				yyp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if yyDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", yyTokname(yychar))
			}
			if yychar == yyEofCode {
				goto ret1
			}
			yychar = -1
			goto yynewstate /* try again in the same state */
		}
	}

	/* reduction by production yyn */
	if yyDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", yyn, yyStatname(yystate))
	}

	yynt := yyn
	yypt := yyp
	_ = yypt // guard against "declared and not used"

	yyp -= yyR2[yyn]
	yyVAL = yyS[yyp+1]

	/* consult goto table to find next state */
	yyn = yyR1[yyn]
	yyg := yyPgo[yyn]
	yyj := yyg + yyS[yyp].yys + 1

	if yyj >= yyLast {
		yystate = yyAct[yyg]
	} else {
		yystate = yyAct[yyj]
		if yyChk[yystate] != -yyn {
			yystate = yyAct[yyg]
		}
	}
	// dummy call; replaced with literal code
	switch yynt {

	case 1:
		//line sql.y:180
		{
			SetParseTree(yylex, yyS[yypt-0].statement)
		}
	case 2:
		//line sql.y:186
		{
			yyVAL.statement = yyS[yypt-0].selStmt
		}
	case 3:
		yyVAL.statement = yyS[yypt-0].statement
	case 4:
		yyVAL.statement = yyS[yypt-0].statement
	case 5:
		yyVAL.statement = yyS[yypt-0].statement
	case 6:
		yyVAL.statement = yyS[yypt-0].statement
	case 7:
		yyVAL.statement = yyS[yypt-0].statement
	case 8:
		yyVAL.statement = yyS[yypt-0].statement
	case 9:
		yyVAL.statement = yyS[yypt-0].statement
	case 10:
		yyVAL.statement = yyS[yypt-0].statement
	case 11:
		yyVAL.statement = yyS[yypt-0].statement
	case 12:
		yyVAL.statement = yyS[yypt-0].statement
	case 13:
		yyVAL.statement = yyS[yypt-0].statement
	case 14:
		yyVAL.statement = yyS[yypt-0].statement
	case 15:
		yyVAL.statement = yyS[yypt-0].statement
	case 16:
		yyVAL.statement = yyS[yypt-0].statement
	case 17:
		yyVAL.statement = yyS[yypt-0].statement
	case 18:
		yyVAL.statement = yyS[yypt-0].statement
	case 19:
		yyVAL.statement = yyS[yypt-0].statement
	case 20:
		//line sql.y:209
		{
			yyVAL.selStmt = &Select{Comments: Comments(yyS[yypt-2].bytes2), Distinct: yyS[yypt-1].str, SelectExprs: yyS[yypt-0].selectExprs}
		}
	case 21:
		//line sql.y:213
		{
			yyVAL.selStmt = &Select{Comments: Comments(yyS[yypt-10].bytes2), Distinct: yyS[yypt-9].str, SelectExprs: yyS[yypt-8].selectExprs, From: yyS[yypt-6].tableExprs, Where: NewWhere(AST_WHERE, yyS[yypt-5].boolExpr), GroupBy: GroupBy(yyS[yypt-4].valExprs), Having: NewWhere(AST_HAVING, yyS[yypt-3].boolExpr), OrderBy: yyS[yypt-2].orderBy, Limit: yyS[yypt-1].limit, Lock: yyS[yypt-0].str}
		}
	case 22:
		//line sql.y:217
		{
			yyVAL.selStmt = &Union{Type: yyS[yypt-1].str, Left: yyS[yypt-2].selStmt, Right: yyS[yypt-0].selStmt}
		}
	case 23:
		//line sql.y:224
		{
			yyVAL.statement = &Insert{Comments: Comments(yyS[yypt-5].bytes2), Table: yyS[yypt-3].tableName, Columns: yyS[yypt-2].columns, Rows: yyS[yypt-1].insRows, OnDup: OnDup(yyS[yypt-0].updateExprs)}
		}
	case 24:
		//line sql.y:228
		{
			cols := make(Columns, 0, len(yyS[yypt-1].updateExprs))
			vals := make(ValTuple, 0, len(yyS[yypt-1].updateExprs))
			for _, col := range yyS[yypt-1].updateExprs {
				cols = append(cols, &NonStarExpr{Expr: col.Name})
				vals = append(vals, col.Expr)
			}
			yyVAL.statement = &Insert{Comments: Comments(yyS[yypt-5].bytes2), Table: yyS[yypt-3].tableName, Columns: cols, Rows: Values{vals}, OnDup: OnDup(yyS[yypt-0].updateExprs)}
		}
	case 25:
		//line sql.y:240
		{
			yyVAL.statement = &Replace{Comments: Comments(yyS[yypt-4].bytes2), Table: yyS[yypt-2].tableName, Columns: yyS[yypt-1].columns, Rows: yyS[yypt-0].insRows}
		}
	case 26:
		//line sql.y:244
		{
			cols := make(Columns, 0, len(yyS[yypt-0].updateExprs))
			vals := make(ValTuple, 0, len(yyS[yypt-0].updateExprs))
			for _, col := range yyS[yypt-0].updateExprs {
				cols = append(cols, &NonStarExpr{Expr: col.Name})
				vals = append(vals, col.Expr)
			}
			yyVAL.statement = &Replace{Comments: Comments(yyS[yypt-4].bytes2), Table: yyS[yypt-2].tableName, Columns: cols, Rows: Values{vals}}
		}
	case 27:
		//line sql.y:257
		{
			yyVAL.statement = &Update{Comments: Comments(yyS[yypt-6].bytes2), Table: yyS[yypt-5].tableName, Exprs: yyS[yypt-3].updateExprs, Where: NewWhere(AST_WHERE, yyS[yypt-2].boolExpr), OrderBy: yyS[yypt-1].orderBy, Limit: yyS[yypt-0].limit}
		}
	case 28:
		//line sql.y:263
		{
			yyVAL.statement = &Delete{Comments: Comments(yyS[yypt-5].bytes2), Table: yyS[yypt-3].tableName, Where: NewWhere(AST_WHERE, yyS[yypt-2].boolExpr), OrderBy: yyS[yypt-1].orderBy, Limit: yyS[yypt-0].limit}
		}
	case 29:
		//line sql.y:269
		{
			yyVAL.statement = &Set{Comments: Comments(yyS[yypt-2].bytes2), Exprs: yyS[yypt-0].updateExprs}
		}
	case 30:
		//line sql.y:273
		{
			yyVAL.statement = &Set{Comments: Comments(yyS[yypt-2].bytes2), Exprs: UpdateExprs{&UpdateExpr{Name: &ColName{Name: []byte("names")}, Expr: StrVal(yyS[yypt-0].bytes)}}}
		}
	case 31:
		//line sql.y:277
		{
			yyVAL.statement = &Set{Comments: Comments(yyS[yypt-3].bytes2), Exprs: UpdateExprs{&UpdateExpr{Name: &ColName{Name: []byte("names")}, Expr: StrVal(yyS[yypt-0].bytes)}}}
		}
	case 32:
		//line sql.y:283
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 33:
		//line sql.y:285
		{
			yyVAL.bytes = []byte(yyS[yypt-0].bytes)
		}
	case 34:
		//line sql.y:287
		{
			yyVAL.bytes = []byte("default")
		}
	case 35:
		//line sql.y:290
		{
		}
	case 36:
		//line sql.y:292
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 37:
		//line sql.y:294
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 38:
		//line sql.y:298
		{
			yyVAL.statement = &Begin{}
		}
	case 39:
		//line sql.y:304
		{
			yyVAL.statement = &Commit{}
		}
	case 40:
		//line sql.y:310
		{
			yyVAL.statement = &Rollback{}
		}
	case 41:
		//line sql.y:316
		{
			yyVAL.statement = &Admin{Name: yyS[yypt-3].bytes, Values: yyS[yypt-1].valExprs}
		}
	case 42:
		//line sql.y:322
		{
			yyVAL.statement = &Show{Section: "databases"}
		}
	case 43:
		//line sql.y:326
		{
			yyVAL.statement = &Show{Section: "tables", From: yyS[yypt-1].valExpr, LikeOrWhere: yyS[yypt-0].expr}
		}
	case 44:
		//line sql.y:330
		{
			yyVAL.statement = &Show{Section: "proxy", Key: string(yyS[yypt-2].bytes), From: yyS[yypt-1].valExpr, LikeOrWhere: yyS[yypt-0].expr}
		}
	case 45:
		//line sql.y:334
		{
			yyVAL.statement = &Show{Section: "variables", LikeOrWhere: yyS[yypt-1].expr}
		}
	case 46:
		//line sql.y:338
		{
			yyVAL.statement = &Show{Section: "table", Key: string("status"), From: yyS[yypt-1].valExpr, LikeOrWhere: yyS[yypt-0].expr}
		}
	case 47:
		//line sql.y:341
		{
			yyVAL.statement = &Show{Section: "create table", Table: yyS[yypt-1].tableName}
		}
	case 48:
		//line sql.y:344
		{
			yyVAL.statement = &Show{Section: "columns", Table: yyS[yypt-2].tableName, LikeOrWhere: yyS[yypt-1].expr}
		}
	case 49:
		//line sql.y:348
		{
			yyVAL.statement = &Show{Section: yyS[yypt-4].str, Table: yyS[yypt-2].tableName}
		}
	case 50:
		//line sql.y:352
		{
			yyVAL.statement = &Show{Section: AST_WARNINGS}
		}
	case 51:
		//line sql.y:356
		{
			yyVAL.statement = &Show{Section: AST_ERRORS}
		}
	case 52:
		//line sql.y:360
		{
			yyVAL.statement = &Show{Section: AST_STATUS}
		}
	case 53:
		//line sql.y:366
		{
			yyVAL.str = AST_INDEX
		}
	case 54:
		//line sql.y:368
		{
			yyVAL.str = AST_INDEXES
		}
	case 55:
		//line sql.y:370
		{
			yyVAL.str = AST_KEYS
		}
	case 56:
		//line sql.y:374
		{
			yyVAL.str = AST_FROM
		}
	case 57:
		//line sql.y:376
		{
			yyVAL.str = AST_IN
		}
	case 58:
		//line sql.y:380
		{
			yyVAL.statement = &Create{Obj: AST_DATABASE, NotExistsOpt: yyS[yypt-2].str, ID: yyS[yypt-1].bytes}
		}
	case 59:
		//line sql.y:384
		{
			// Change this to an alter statement
			yyVAL.statement = &DDL{Action: AST_ALTER, Table: yyS[yypt-1].bytes, NewName: yyS[yypt-1].bytes}
		}
	case 60:
		//line sql.y:389
		{
			yyVAL.statement = &Create{Obj: AST_VIEW, ID: yyS[yypt-1].bytes}
		}
	case 61:
		//line sql.y:393
		{
			yyVAL.statement = &Create{Obj: AST_FUNCTION, ID: yyS[yypt-1].bytes}
		}
	case 62:
		//line sql.y:397
		{
			yyVAL.statement = &CreateTable{Table: yyS[yypt-3].tableName, NotExistsOpt: yyS[yypt-4].str}
		}
	case 63:
		//line sql.y:401
		{
			yyVAL.statement = &CreateTable{Table: yyS[yypt-4].tableName, NotExistsOpt: yyS[yypt-5].str}
		}
	case 64:
		//line sql.y:405
		{
			yyVAL.statement = &CreateTable{Table: yyS[yypt-2].tableName, NotExistsOpt: yyS[yypt-3].str}
		}
	case 65:
		//line sql.y:417
		{
			yyVAL.empty = struct{}{}
		}
	case 66:
		//line sql.y:421
		{
			yyVAL.empty = struct{}{}
		}
	case 67:
		//line sql.y:427
		{
			yyVAL.empty = struct{}{}
		}
	case 68:
		//line sql.y:432
		{
			yyVAL.empty = struct{}{}
		}
	case 69:
		//line sql.y:436
		{
			yyVAL.empty = struct{}{}
		}
	case 70:
		//line sql.y:440
		{
			yyVAL.empty = struct{}{}
		}
	case 71:
		//line sql.y:444
		{
			yyVAL.empty = struct{}{}
		}
	case 72:
		//line sql.y:448
		{
			yyVAL.empty = struct{}{}
		}
	case 73:
		//line sql.y:452
		{
			yyVAL.empty = struct{}{}
		}
	case 74:
		//line sql.y:456
		{
			yyVAL.empty = struct{}{}
		}
	case 75:
		//line sql.y:462
		{
			yyVAL.empty = struct{}{}
		}
	case 76:
		//line sql.y:465
		{
			yyVAL.empty = struct{}{}
		}
	case 77:
		//line sql.y:467
		{
			yyVAL.empty = struct{}{}
		}
	case 78:
		//line sql.y:469
		{
			yyVAL.empty = struct{}{}
		}
	case 79:
		//line sql.y:471
		{
			yyVAL.empty = struct{}{}
		}
	case 80:
		//line sql.y:473
		{
			yyVAL.empty = struct{}{}
		}
	case 81:
		//line sql.y:475
		{
			yyVAL.empty = struct{}{}
		}
	case 82:
		//line sql.y:477
		{
			yyVAL.empty = struct{}{}
		}
	case 83:
		//line sql.y:479
		{
			yyVAL.empty = struct{}{}
		}
	case 84:
		//line sql.y:481
		{
			yyVAL.empty = struct{}{}
		}
	case 85:
		//line sql.y:483
		{
			yyVAL.empty = struct{}{}
		}
	case 86:
		//line sql.y:485
		{
			yyVAL.empty = struct{}{}
		}
	case 87:
		//line sql.y:490
		{
			yyVAL.empty = struct{}{}
		}
	case 88:
		//line sql.y:494
		{
			yyVAL.empty = struct{}{}
		}
	case 89:
		//line sql.y:498
		{
			yyVAL.empty = struct{}{}
		}
	case 90:
		//line sql.y:502
		{
			yyVAL.empty = struct{}{}
		}
	case 91:
		//line sql.y:506
		{
			yyVAL.empty = struct{}{}
		}
	case 92:
		//line sql.y:510
		{
			yyVAL.empty = struct{}{}
		}
	case 93:
		//line sql.y:514
		{
			yyVAL.empty = struct{}{}
		}
	case 94:
		//line sql.y:516
		{
			yyVAL.empty = struct{}{}
		}
	case 95:
		//line sql.y:518
		{
			yyVAL.empty = struct{}{}
		}
	case 96:
		//line sql.y:520
		{
			yyVAL.empty = struct{}{}
		}
	case 97:
		//line sql.y:522
		{
			yyVAL.empty = struct{}{}
		}
	case 98:
		//line sql.y:524
		{
			yyVAL.empty = struct{}{}
		}
	case 99:
		//line sql.y:526
		{
			yyVAL.empty = struct{}{}
		}
	case 100:
		//line sql.y:528
		{
			yyVAL.empty = struct{}{}
		}
	case 101:
		//line sql.y:530
		{
			yyVAL.empty = struct{}{}
		}
	case 102:
		//line sql.y:532
		{
			yyVAL.empty = struct{}{}
		}
	case 103:
		//line sql.y:534
		{
			yyVAL.empty = struct{}{}
		}
	case 104:
		//line sql.y:536
		{
			yyVAL.empty = struct{}{}
		}
	case 105:
		//line sql.y:538
		{
			yyVAL.empty = struct{}{}
		}
	case 106:
		//line sql.y:540
		{
			yyVAL.empty = struct{}{}
		}
	case 107:
		//line sql.y:542
		{
			yyVAL.empty = struct{}{}
		}
	case 108:
		//line sql.y:544
		{
			yyVAL.empty = struct{}{}
		}
	case 109:
		//line sql.y:546
		{
			yyVAL.empty = struct{}{}
		}
	case 110:
		//line sql.y:548
		{
			yyVAL.empty = struct{}{}
		}
	case 111:
		//line sql.y:550
		{
			yyVAL.empty = struct{}{}
		}
	case 112:
		//line sql.y:552
		{
			yyVAL.empty = struct{}{}
		}
	case 113:
		//line sql.y:554
		{
			yyVAL.empty = struct{}{}
		}
	case 114:
		//line sql.y:556
		{
			yyVAL.empty = struct{}{}
		}
	case 115:
		//line sql.y:558
		{
			yyVAL.empty = struct{}{}
		}
	case 116:
		//line sql.y:560
		{
			yyVAL.empty = struct{}{}
		}
	case 117:
		//line sql.y:562
		{
			yyVAL.empty = struct{}{}
		}
	case 118:
		//line sql.y:564
		{
			yyVAL.empty = struct{}{}
		}
	case 119:
		//line sql.y:566
		{
			yyVAL.empty = struct{}{}
		}
	case 120:
		//line sql.y:568
		{
			yyVAL.empty = struct{}{}
		}
	case 121:
		//line sql.y:570
		{
			yyVAL.empty = struct{}{}
		}
	case 122:
		//line sql.y:572
		{
			yyVAL.empty = struct{}{}
		}
	case 123:
		//line sql.y:574
		{
			yyVAL.empty = struct{}{}
		}
	case 124:
		//line sql.y:576
		{
			yyVAL.empty = struct{}{}
		}
	case 125:
		//line sql.y:581
		{
			yyVAL.empty = struct{}{}
		}
	case 126:
		//line sql.y:583
		{
			yyVAL.empty = struct{}{}
		}
	case 127:
		//line sql.y:587
		{
			yyVAL.empty = struct{}{}
		}
	case 128:
		//line sql.y:589
		{
			yyVAL.empty = struct{}{}
		}
	case 129:
		//line sql.y:593
		{
			yyVAL.empty = struct{}{}
		}
	case 130:
		//line sql.y:595
		{
			yyVAL.empty = struct{}{}
		}
	case 131:
		//line sql.y:598
		{
			yyVAL.str = ""
		}
	case 132:
		//line sql.y:602
		{
			yyVAL.str = __yyfmt__.Sprint("(", yyS[yypt-1].bytes, ")")
		}
	case 133:
		//line sql.y:606
		{
			yyVAL.str = __yyfmt__.Sprint("(", yyS[yypt-3].bytes, ",", yyS[yypt-1].bytes, ")")
		}
	case 134:
		//line sql.y:611
		{
			yyVAL.str = ""
		}
	case 135:
		//line sql.y:613
		{
			yyVAL.str = "binary"
		}
	case 136:
		//line sql.y:616
		{
			yyVAL.str = ""
		}
	case 137:
		//line sql.y:618
		{
			yyVAL.str = yyS[yypt-1].str + " signed"
		}
	case 138:
		//line sql.y:620
		{
			yyVAL.str = yyS[yypt-1].str + " unsigned"
		}
	case 139:
		//line sql.y:622
		{
			yyVAL.str = yyS[yypt-1].str + " zerofill"
		}
	case 140:
		//line sql.y:625
		{
			yyVAL.str = ""
		}
	case 141:
		//line sql.y:627
		{
			yyVAL.str = __yyfmt__.Sprintf("COLCHARSET %s", yyS[yypt-0].bytes)
		}
	case 142:
		//line sql.y:629
		{
			yyVAL.str = __yyfmt__.Sprintf("COLCHARSET %s", yyS[yypt-0].bytes)
		}
	case 143:
		//line sql.y:634
		{
			yyVAL.statement = &DDL{Action: AST_ALTER, Table: yyS[yypt-2].bytes, NewName: yyS[yypt-2].bytes}
		}
	case 144:
		//line sql.y:638
		{
			// Change this to a rename statement
			yyVAL.statement = &DDL{Action: AST_RENAME, Table: yyS[yypt-3].bytes, NewName: yyS[yypt-0].bytes}
		}
	case 145:
		//line sql.y:643
		{
			yyVAL.statement = &DDL{Action: AST_ALTER, Table: yyS[yypt-1].bytes, NewName: yyS[yypt-1].bytes}
		}
	case 146:
		//line sql.y:649
		{
			yyVAL.statement = &DDL{Action: AST_RENAME, Table: yyS[yypt-2].bytes, NewName: yyS[yypt-0].bytes}
		}
	case 147:
		//line sql.y:655
		{
			yyVAL.statement = &Drop{Key: AST_TABLE, ExistsOpt: yyS[yypt-2].str, ID: yyS[yypt-1].bytes}
		}
	case 148:
		//line sql.y:659
		{
			// Change this to an alter statement
			yyVAL.statement = &DDL{Action: AST_ALTER, Table: yyS[yypt-0].bytes, NewName: yyS[yypt-0].bytes}
		}
	case 149:
		//line sql.y:664
		{
			yyVAL.statement = &Drop{Key: AST_VIEW, ExistsOpt: yyS[yypt-2].str, ID: yyS[yypt-1].bytes}
		}
	case 150:
		//line sql.y:668
		{
			yyVAL.statement = &Drop{Key: AST_FUNCTION, ExistsOpt: yyS[yypt-2].str, ID: yyS[yypt-1].bytes}
		}
	case 151:
		//line sql.y:672
		{
			yyVAL.statement = &Drop{Key: AST_PROCEDURE, ExistsOpt: yyS[yypt-2].str, ID: yyS[yypt-1].bytes}
		}
	case 152:
		//line sql.y:678
		{
			yyVAL.statement = &DDL{Action: AST_ALTER, Table: yyS[yypt-0].bytes, NewName: yyS[yypt-0].bytes}
		}
	case 153:
		//line sql.y:684
		{
			yyVAL.statement = &Use{Action: AST_USE, ID: yyS[yypt-1].bytes}
		}
	case 154:
		//line sql.y:690
		{
			yyVAL.statement = &Explain{Table: yyS[yypt-1].tableName}
		}
	case 155:
		//line sql.y:694
		{
			yyVAL.statement = &Explain{}
		}
	case 156:
		//line sql.y:699
		{
		}
	case 157:
		//line sql.y:701
		{
		}
	case 158:
		//line sql.y:703
		{
		}
	case 159:
		//line sql.y:707
		{
			yyVAL.str = AST_EXPLAIN
		}
	case 160:
		//line sql.y:710
		{
			yyVAL.str = AST_EXPLAIN
		}
	case 161:
		//line sql.y:713
		{
			yyVAL.str = AST_EXPLAIN
		}
	case 162:
		//line sql.y:716
		{
			SetAllowComments(yylex, true)
		}
	case 163:
		//line sql.y:720
		{
			yyVAL.bytes2 = yyS[yypt-0].bytes2
			SetAllowComments(yylex, false)
		}
	case 164:
		//line sql.y:726
		{
			yyVAL.bytes2 = nil
		}
	case 165:
		//line sql.y:730
		{
			yyVAL.bytes2 = append(yyS[yypt-1].bytes2, yyS[yypt-0].bytes)
		}
	case 166:
		//line sql.y:736
		{
			yyVAL.str = AST_UNION
		}
	case 167:
		//line sql.y:740
		{
			yyVAL.str = AST_UNION_ALL
		}
	case 168:
		//line sql.y:744
		{
			yyVAL.str = AST_SET_MINUS
		}
	case 169:
		//line sql.y:748
		{
			yyVAL.str = AST_EXCEPT
		}
	case 170:
		//line sql.y:752
		{
			yyVAL.str = AST_INTERSECT
		}
	case 171:
		//line sql.y:757
		{
			yyVAL.str = ""
		}
	case 172:
		//line sql.y:761
		{
			yyVAL.str = AST_DISTINCT
		}
	case 173:
		//line sql.y:767
		{
			yyVAL.selectExprs = SelectExprs{yyS[yypt-0].selectExpr}
		}
	case 174:
		//line sql.y:771
		{
			yyVAL.selectExprs = append(yyVAL.selectExprs, yyS[yypt-0].selectExpr)
		}
	case 175:
		//line sql.y:777
		{
			yyVAL.selectExpr = &StarExpr{}
		}
	case 176:
		//line sql.y:781
		{
			yyVAL.selectExpr = &NonStarExpr{Expr: yyS[yypt-1].expr, As: yyS[yypt-0].bytes}
		}
	case 177:
		//line sql.y:785
		{
			yyVAL.selectExpr = &StarExpr{TableName: yyS[yypt-2].bytes}
		}
	case 178:
		//line sql.y:791
		{
			yyVAL.expr = yyS[yypt-0].boolExpr
		}
	case 179:
		//line sql.y:795
		{
			yyVAL.expr = yyS[yypt-0].valExpr
		}
	case 180:
		//line sql.y:800
		{
			yyVAL.bytes = nil
		}
	case 181:
		//line sql.y:804
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 182:
		//line sql.y:808
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 183:
		//line sql.y:812
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 184:
		//line sql.y:816
		{
			yyVAL.bytes = bytes.NewBufferString("date").Bytes()
		}
	case 185:
		//line sql.y:822
		{
			yyVAL.tableExprs = TableExprs{yyS[yypt-0].tableExpr}
		}
	case 186:
		//line sql.y:826
		{
			yyVAL.tableExprs = append(yyVAL.tableExprs, yyS[yypt-0].tableExpr)
		}
	case 187:
		//line sql.y:832
		{
			yyVAL.tableExpr = &AliasedTableExpr{Expr: yyS[yypt-2].smTableExpr, As: yyS[yypt-1].bytes, Hints: yyS[yypt-0].indexHints}
		}
	case 188:
		//line sql.y:836
		{
			yyVAL.tableExpr = &ParenTableExpr{Expr: yyS[yypt-1].tableExpr}
		}
	case 189:
		//line sql.y:840
		{
			yyVAL.tableExpr = &JoinTableExpr{LeftExpr: yyS[yypt-2].tableExpr, Join: yyS[yypt-1].str, RightExpr: yyS[yypt-0].tableExpr}
		}
	case 190:
		//line sql.y:844
		{
			yyVAL.tableExpr = &JoinTableExpr{LeftExpr: yyS[yypt-4].tableExpr, Join: yyS[yypt-3].str, RightExpr: yyS[yypt-2].tableExpr, On: yyS[yypt-0].boolExpr}
		}
	case 191:
		//line sql.y:848
		{
			yyVAL.tableExpr = &JoinTableExpr{LeftExpr: yyS[yypt-6].tableExpr, Join: yyS[yypt-5].str, RightExpr: yyS[yypt-4].tableExpr}
		}
	case 192:
		//line sql.y:853
		{
			yyVAL.bytes = nil
		}
	case 193:
		//line sql.y:857
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 194:
		//line sql.y:861
		{
			yyVAL.bytes = yyS[yypt-0].bytes
		}
	case 195:
		//line sql.y:867
		{
			yyVAL.str = AST_JOIN
		}
	case 196:
		//line sql.y:871
		{
			yyVAL.str = AST_STRAIGHT_JOIN
		}
	case 197:
		//line sql.y:875
		{
			yyVAL.str = AST_LEFT_JOIN
		}
	case 198:
		//line sql.y:879
		{
			yyVAL.str = AST_LEFT_JOIN
		}
	case 199:
		//line sql.y:883
		{
			yyVAL.str = AST_RIGHT_JOIN
		}
	case 200:
		//line sql.y:887
		{
			yyVAL.str = AST_RIGHT_JOIN
		}
	case 201:
		//line sql.y:891
		{
			yyVAL.str = AST_JOIN
		}
	case 202:
		//line sql.y:895
		{
			yyVAL.str = AST_CROSS_JOIN
		}
	case 203:
		//line sql.y:899
		{
			yyVAL.str = AST_NATURAL_JOIN
		}
	case 204:
		//line sql.y:903
		{
			yyVAL.str = "natural left join"
		}
	case 205:
		//line sql.y:907
		{
			yyVAL.str = "natural right join"
		}
	case 206:
		//line sql.y:913
		{
			yyVAL.smTableExpr = &TableName{Name: yyS[yypt-0].bytes}
		}
	case 207:
		//line sql.y:917
		{
			yyVAL.smTableExpr = &TableName{Qualifier: yyS[yypt-2].bytes, Name: yyS[yypt-0].bytes}
		}
	case 208:
		//line sql.y:921
		{
			yyVAL.smTableExpr = yyS[yypt-0].subquery
		}
	case 209:
		//line sql.y:927
		{
			yyVAL.tableName = &TableName{Name: yyS[yypt-0].bytes}
		}
	case 210:
		//line sql.y:931
		{
			yyVAL.tableName = &TableName{Qualifier: yyS[yypt-2].bytes, Name: yyS[yypt-0].bytes}
		}
	case 211:
		//line sql.y:936
		{
			yyVAL.indexHints = nil
		}
	case 212:
		//line sql.y:940
		{
			yyVAL.indexHints = &IndexHints{Type: AST_USE, Indexes: yyS[yypt-1].bytes2}
		}
	case 213:
		//line sql.y:944
		{
			yyVAL.indexHints = &IndexHints{Type: AST_IGNORE, Indexes: yyS[yypt-1].bytes2}
		}
	case 214:
		//line sql.y:948
		{
			yyVAL.indexHints = &IndexHints{Type: AST_FORCE, Indexes: yyS[yypt-1].bytes2}
		}
	case 215:
		//line sql.y:954
		{
			yyVAL.bytes2 = [][]byte{yyS[yypt-0].bytes}
		}
	case 216:
		//line sql.y:958
		{
			yyVAL.bytes2 = append(yyS[yypt-2].bytes2, yyS[yypt-0].bytes)
		}
	case 217:
		//line sql.y:963
		{
			yyVAL.boolExpr = nil
		}
	case 218:
		//line sql.y:967
		{
			yyVAL.boolExpr = yyS[yypt-0].boolExpr
		}
	case 219:
		//line sql.y:972
		{
			yyVAL.expr = nil
		}
	case 220:
		//line sql.y:976
		{
			yyVAL.expr = yyS[yypt-0].boolExpr
		}
	case 221:
		//line sql.y:980
		{
			yyVAL.expr = yyS[yypt-0].valExpr
		}
	case 222:
		//line sql.y:985
		{
			yyVAL.valExpr = nil
		}
	case 223:
		//line sql.y:989
		{
			yyVAL.valExpr = yyS[yypt-0].valExpr
		}
	case 224:
		yyVAL.boolExpr = yyS[yypt-0].boolExpr
	case 225:
		//line sql.y:996
		{
			yyVAL.boolExpr = &AndExpr{Left: yyS[yypt-2].boolExpr, Right: yyS[yypt-0].boolExpr}
		}
	case 226:
		//line sql.y:1000
		{
			yyVAL.boolExpr = &OrExpr{Left: yyS[yypt-2].boolExpr, Right: yyS[yypt-0].boolExpr}
		}
	case 227:
		//line sql.y:1004
		{
			yyVAL.boolExpr = &NotExpr{Expr: yyS[yypt-0].boolExpr}
		}
	case 228:
		//line sql.y:1008
		{
			yyVAL.boolExpr = &ParenBoolExpr{Expr: yyS[yypt-1].boolExpr}
		}
	case 229:
		//line sql.y:1014
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyS[yypt-2].valExpr, Operator: yyS[yypt-1].str, Right: yyS[yypt-0].valExpr}
		}
	case 230:
		//line sql.y:1018
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyS[yypt-2].valExpr, Operator: AST_IN, Right: yyS[yypt-0].colTuple}
		}
	case 231:
		//line sql.y:1022
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyS[yypt-3].valExpr, Operator: AST_NOT_IN, Right: yyS[yypt-0].colTuple}
		}
	case 232:
		//line sql.y:1026
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyS[yypt-2].valExpr, Operator: AST_LIKE, Right: yyS[yypt-0].valExpr}
		}
	case 233:
		//line sql.y:1030
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyS[yypt-3].valExpr, Operator: AST_NOT_LIKE, Right: yyS[yypt-0].valExpr}
		}
	case 234:
		//line sql.y:1034
		{
			yyVAL.boolExpr = &RangeCond{Left: yyS[yypt-4].valExpr, Operator: AST_BETWEEN, From: yyS[yypt-2].valExpr, To: yyS[yypt-0].valExpr}
		}
	case 235:
		//line sql.y:1038
		{
			yyVAL.boolExpr = &RangeCond{Left: yyS[yypt-5].valExpr, Operator: AST_NOT_BETWEEN, From: yyS[yypt-2].valExpr, To: yyS[yypt-0].valExpr}
		}
	case 236:
		//line sql.y:1042
		{
			yyVAL.boolExpr = &NullCheck{Operator: AST_IS_NULL, Expr: yyS[yypt-2].valExpr}
		}
	case 237:
		//line sql.y:1046
		{
			yyVAL.boolExpr = &NullCheck{Operator: AST_IS_NOT_NULL, Expr: yyS[yypt-3].valExpr}
		}
	case 238:
		//line sql.y:1050
		{
			yyVAL.boolExpr = &ExistsExpr{Subquery: yyS[yypt-0].subquery}
		}
	case 239:
		//line sql.y:1056
		{
			yyVAL.str = AST_EQ
		}
	case 240:
		//line sql.y:1060
		{
			yyVAL.str = AST_LT
		}
	case 241:
		//line sql.y:1064
		{
			yyVAL.str = AST_GT
		}
	case 242:
		//line sql.y:1068
		{
			yyVAL.str = AST_LE
		}
	case 243:
		//line sql.y:1072
		{
			yyVAL.str = AST_GE
		}
	case 244:
		//line sql.y:1076
		{
			yyVAL.str = AST_NE
		}
	case 245:
		//line sql.y:1080
		{
			yyVAL.str = AST_NSE
		}
	case 246:
		//line sql.y:1086
		{
			yyVAL.colTuple = ValTuple(yyS[yypt-1].valExprs)
		}
	case 247:
		//line sql.y:1090
		{
			yyVAL.colTuple = yyS[yypt-0].subquery
		}
	case 248:
		//line sql.y:1094
		{
			yyVAL.colTuple = ListArg(yyS[yypt-0].bytes)
		}
	case 249:
		//line sql.y:1100
		{
			yyVAL.subquery = &Subquery{yyS[yypt-1].selStmt}
		}
	case 250:
		//line sql.y:1106
		{
			yyVAL.valExprs = ValExprs{yyS[yypt-0].valExpr}
		}
	case 251:
		//line sql.y:1110
		{
			yyVAL.valExprs = append(yyS[yypt-2].valExprs, yyS[yypt-0].valExpr)
		}
	case 252:
		//line sql.y:1116
		{
			yyVAL.valExpr = yyS[yypt-0].valExpr
		}
	case 253:
		//line sql.y:1120
		{
			yyVAL.valExpr = yyS[yypt-0].colName
		}
	case 254:
		//line sql.y:1124
		{
			yyVAL.valExpr = yyS[yypt-0].rowTuple
		}
	case 255:
		//line sql.y:1128
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyS[yypt-2].valExpr, Operator: AST_BITAND, Right: yyS[yypt-0].valExpr}
		}
	case 256:
		//line sql.y:1132
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyS[yypt-2].valExpr, Operator: AST_BITOR, Right: yyS[yypt-0].valExpr}
		}
	case 257:
		//line sql.y:1136
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyS[yypt-2].valExpr, Operator: AST_BITXOR, Right: yyS[yypt-0].valExpr}
		}
	case 258:
		//line sql.y:1140
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyS[yypt-2].valExpr, Operator: AST_PLUS, Right: yyS[yypt-0].valExpr}
		}
	case 259:
		//line sql.y:1144
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyS[yypt-2].valExpr, Operator: AST_MINUS, Right: yyS[yypt-0].valExpr}
		}
	case 260:
		//line sql.y:1148
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyS[yypt-2].valExpr, Operator: AST_MULT, Right: yyS[yypt-0].valExpr}
		}
	case 261:
		//line sql.y:1152
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyS[yypt-2].valExpr, Operator: AST_DIV, Right: yyS[yypt-0].valExpr}
		}
	case 262:
		//line sql.y:1156
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyS[yypt-2].valExpr, Operator: AST_MOD, Right: yyS[yypt-0].valExpr}
		}
	case 263:
		//line sql.y:1160
		{
		}
	case 264:
		//line sql.y:1162
		{
		}
	case 265:
		//line sql.y:1164
		{
			if num, ok := yyS[yypt-0].valExpr.(NumVal); ok {
				switch yyS[yypt-1].byt {
				case '-':
					yyVAL.valExpr = append(NumVal("-"), num...)
				case '+':
					yyVAL.valExpr = num
				default:
					yyVAL.valExpr = &UnaryExpr{Operator: yyS[yypt-1].byt, Expr: yyS[yypt-0].valExpr}
				}
			} else {
				yyVAL.valExpr = &UnaryExpr{Operator: yyS[yypt-1].byt, Expr: yyS[yypt-0].valExpr}
			}
		}
	case 266:
		//line sql.y:1179
		{
			yyVAL.valExpr = &FuncExpr{Name: yyS[yypt-2].bytes}
		}
	case 267:
		//line sql.y:1183
		{
			yyVAL.valExpr = &FuncExpr{Name: yyS[yypt-3].bytes, Exprs: yyS[yypt-1].selectExprs}
		}
	case 268:
		//line sql.y:1187
		{
			yyVAL.valExpr = &FuncExpr{Name: yyS[yypt-4].bytes, Distinct: true, Exprs: yyS[yypt-1].selectExprs}
		}
	case 269:
		//line sql.y:1191
		{
			yyVAL.valExpr = &FuncExpr{Name: yyS[yypt-3].bytes, Exprs: yyS[yypt-1].selectExprs}
		}
	case 270:
		//line sql.y:1195
		{
			yyVAL.valExpr = &FuncExpr{Name: yyS[yypt-2].bytes}
		}
	case 271:
		//line sql.y:1199
		{
			yyVAL.valExpr = yyS[yypt-0].caseExpr
		}
	case 272:
		//line sql.y:1205
		{
			yyVAL.bytes = IF_BYTES
		}
	case 273:
		//line sql.y:1209
		{
			yyVAL.bytes = VALUES_BYTES
		}
	case 274:
		//line sql.y:1213
		{
			yyVAL.bytes = DATABASE_BYTES
		}
	case 275:
		//line sql.y:1219
		{
			yyVAL.byt = AST_UPLUS
		}
	case 276:
		//line sql.y:1223
		{
			yyVAL.byt = AST_UMINUS
		}
	case 277:
		//line sql.y:1227
		{
			yyVAL.byt = AST_TILDA
		}
	case 278:
		//line sql.y:1233
		{
			yyVAL.caseExpr = &CaseExpr{Expr: yyS[yypt-3].valExpr, Whens: yyS[yypt-2].whens, Else: yyS[yypt-1].valExpr}
		}
	case 279:
		//line sql.y:1238
		{
			yyVAL.valExpr = nil
		}
	case 280:
		//line sql.y:1242
		{
			yyVAL.valExpr = yyS[yypt-0].valExpr
		}
	case 281:
		//line sql.y:1248
		{
			yyVAL.whens = []*When{yyS[yypt-0].when}
		}
	case 282:
		//line sql.y:1252
		{
			yyVAL.whens = append(yyS[yypt-1].whens, yyS[yypt-0].when)
		}
	case 283:
		//line sql.y:1258
		{
			yyVAL.when = &When{Cond: yyS[yypt-2].boolExpr, Val: yyS[yypt-0].valExpr}
		}
	case 284:
		//line sql.y:1263
		{
			yyVAL.valExpr = nil
		}
	case 285:
		//line sql.y:1267
		{
			yyVAL.valExpr = yyS[yypt-0].valExpr
		}
	case 286:
		//line sql.y:1273
		{
			yyVAL.colName = &ColName{Name: yyS[yypt-0].bytes}
		}
	case 287:
		//line sql.y:1277
		{
			yyVAL.colName = &ColName{Qualifier: yyS[yypt-2].bytes, Name: yyS[yypt-0].bytes}
		}
	case 288:
		//line sql.y:1283
		{
			yyVAL.valExpr = StrVal(yyS[yypt-0].bytes)
		}
	case 289:
		//line sql.y:1287
		{
			yyVAL.valExpr = NumVal(yyS[yypt-0].bytes)
		}
	case 290:
		//line sql.y:1291
		{
			yyVAL.valExpr = ValArg(yyS[yypt-0].bytes)
		}
	case 291:
		//line sql.y:1295
		{
			yyVAL.valExpr = &NullVal{}
		}
	case 292:
		//line sql.y:1300
		{
			yyVAL.valExprs = nil
		}
	case 293:
		//line sql.y:1304
		{
			yyVAL.valExprs = yyS[yypt-0].valExprs
		}
	case 294:
		//line sql.y:1309
		{
			yyVAL.boolExpr = nil
		}
	case 295:
		//line sql.y:1313
		{
			yyVAL.boolExpr = yyS[yypt-0].boolExpr
		}
	case 296:
		//line sql.y:1318
		{
			yyVAL.orderBy = nil
		}
	case 297:
		//line sql.y:1322
		{
			yyVAL.orderBy = yyS[yypt-0].orderBy
		}
	case 298:
		//line sql.y:1328
		{
			yyVAL.orderBy = OrderBy{yyS[yypt-0].order}
		}
	case 299:
		//line sql.y:1332
		{
			yyVAL.orderBy = append(yyS[yypt-2].orderBy, yyS[yypt-0].order)
		}
	case 300:
		//line sql.y:1338
		{
			yyVAL.order = &Order{Expr: yyS[yypt-1].valExpr, Direction: yyS[yypt-0].str}
		}
	case 301:
		//line sql.y:1343
		{
			yyVAL.str = AST_ASC
		}
	case 302:
		//line sql.y:1347
		{
			yyVAL.str = AST_ASC
		}
	case 303:
		//line sql.y:1351
		{
			yyVAL.str = AST_DESC
		}
	case 304:
		//line sql.y:1356
		{
			yyVAL.limit = nil
		}
	case 305:
		//line sql.y:1360
		{
			yyVAL.limit = &Limit{Rowcount: yyS[yypt-0].valExpr}
		}
	case 306:
		//line sql.y:1364
		{
			yyVAL.limit = &Limit{Offset: yyS[yypt-2].valExpr, Rowcount: yyS[yypt-0].valExpr}
		}
	case 307:
		//line sql.y:1369
		{
			yyVAL.str = ""
		}
	case 308:
		//line sql.y:1373
		{
			yyVAL.str = AST_FOR_UPDATE
		}
	case 309:
		//line sql.y:1377
		{
			if !bytes.Equal(yyS[yypt-1].bytes, SHARE) {
				yylex.Error("expecting share")
				return 1
			}
			if !bytes.Equal(yyS[yypt-0].bytes, MODE) {
				yylex.Error("expecting mode")
				return 1
			}
			yyVAL.str = AST_SHARE_MODE
		}
	case 310:
		//line sql.y:1390
		{
			yyVAL.columns = nil
		}
	case 311:
		//line sql.y:1394
		{
			yyVAL.columns = yyS[yypt-1].columns
		}
	case 312:
		//line sql.y:1400
		{
			yyVAL.columns = Columns{&NonStarExpr{Expr: yyS[yypt-0].colName}}
		}
	case 313:
		//line sql.y:1404
		{
			yyVAL.columns = append(yyVAL.columns, &NonStarExpr{Expr: yyS[yypt-0].colName})
		}
	case 314:
		//line sql.y:1409
		{
			yyVAL.updateExprs = nil
		}
	case 315:
		//line sql.y:1413
		{
			yyVAL.updateExprs = yyS[yypt-0].updateExprs
		}
	case 316:
		//line sql.y:1419
		{
			yyVAL.insRows = yyS[yypt-0].values
		}
	case 317:
		//line sql.y:1423
		{
			yyVAL.insRows = yyS[yypt-0].values
		}
	case 318:
		//line sql.y:1427
		{
			yyVAL.insRows = yyS[yypt-0].selStmt
		}
	case 319:
		//line sql.y:1433
		{
			yyVAL.values = Values{yyS[yypt-0].rowTuple}
		}
	case 320:
		//line sql.y:1437
		{
			yyVAL.values = append(yyS[yypt-2].values, yyS[yypt-0].rowTuple)
		}
	case 321:
		//line sql.y:1443
		{
			yyVAL.rowTuple = ValTuple(yyS[yypt-1].valExprs)
		}
	case 322:
		//line sql.y:1447
		{
			yyVAL.rowTuple = ValTuple(ValExprs{&NullVal{}})
		}
	case 323:
		//line sql.y:1451
		{
			yyVAL.rowTuple = yyS[yypt-0].subquery
		}
	case 324:
		//line sql.y:1457
		{
			yyVAL.updateExprs = UpdateExprs{yyS[yypt-0].updateExpr}
		}
	case 325:
		//line sql.y:1461
		{
			yyVAL.updateExprs = append(yyS[yypt-2].updateExprs, yyS[yypt-0].updateExpr)
		}
	case 326:
		//line sql.y:1467
		{
			yyVAL.updateExpr = &UpdateExpr{Name: yyS[yypt-2].colName, Expr: yyS[yypt-0].valExpr}
		}
	case 327:
		//line sql.y:1472
		{
			yyVAL.empty = struct{}{}
		}
	case 328:
		//line sql.y:1474
		{
			yyVAL.empty = struct{}{}
		}
	case 329:
		//line sql.y:1477
		{
			yyVAL.str = ""
		}
	case 330:
		//line sql.y:1479
		{
			yyVAL.str = AST_EXISTS
		}
	case 331:
		//line sql.y:1482
		{
			yyVAL.str = ""
		}
	case 332:
		//line sql.y:1484
		{
			yyVAL.str = AST_NOT_EXISTS
		}
	case 333:
		//line sql.y:1487
		{
			yyVAL.str = ""
		}
	case 334:
		//line sql.y:1489
		{
			yyVAL.str = AST_TEMPORARY
		}
	case 335:
		//line sql.y:1492
		{
			yyVAL.empty = struct{}{}
		}
	case 336:
		//line sql.y:1494
		{
			yyVAL.empty = struct{}{}
		}
	case 337:
		//line sql.y:1498
		{
			yyVAL.empty = struct{}{}
		}
	case 338:
		//line sql.y:1500
		{
			yyVAL.empty = struct{}{}
		}
	case 339:
		//line sql.y:1502
		{
			yyVAL.empty = struct{}{}
		}
	case 340:
		//line sql.y:1504
		{
			yyVAL.empty = struct{}{}
		}
	case 341:
		//line sql.y:1506
		{
			yyVAL.empty = struct{}{}
		}
	case 342:
		//line sql.y:1509
		{
			yyVAL.empty = struct{}{}
		}
	case 343:
		//line sql.y:1511
		{
			yyVAL.empty = struct{}{}
		}
	case 344:
		//line sql.y:1514
		{
			yyVAL.empty = struct{}{}
		}
	case 345:
		//line sql.y:1516
		{
			yyVAL.empty = struct{}{}
		}
	case 346:
		//line sql.y:1519
		{
			yyVAL.empty = struct{}{}
		}
	case 347:
		//line sql.y:1521
		{
			yyVAL.empty = struct{}{}
		}
	case 348:
		//line sql.y:1525
		{
			yyVAL.bytes = bytes.ToLower(yyS[yypt-0].bytes)
		}
	case 349:
		//line sql.y:1530
		{
			ForceEOF(yylex)
		}
	}
	goto yystack /* stack new state and value */
}
