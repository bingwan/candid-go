// Do not edit. This file is auto-generated.
// Grammar: CANDID-TEST (v0.1.0) github.com/di-wu/candid-go/internal/candid_test

package test

import (
	"github.com/di-wu/parser"
	"github.com/di-wu/parser/ast"
	"github.com/di-wu/parser/op"
)

func TestData(p *ast.Parser) (*ast.Node, error) {
	return p.Expect(
		ast.Capture{
			Type:        TestDataT,
			TypeStrings: NodeTypes,
			Value: op.MinOne(
				op.Or{
					Comment,
					Test,
					EndLine,
				},
			),
		},
	)
}

func Comment(p *ast.Parser) (*ast.Node, error) {
	return p.Expect(
		op.Or{
			op.And{
				"/*",
				Ws,
				MultiComment,
				Ws,
				"*/",
			},
			op.And{
				op.And{
					"//",
					op.Optional(
						CommentText,
					),
				},
				EndLine,
			},
		},
	)
}

func CommentText(p *ast.Parser) (*ast.Node, error) {
	return p.Expect(
		ast.Capture{
			Type:        CommentTextT,
			TypeStrings: NodeTypes,
			Value: op.MinZero(
				op.And{
					op.Not{
						EndLine,
					},
					parser.CheckRuneRange(0x0000, 0x0010FFFF),
				},
			),
		},
	)
}

func MultiComment(p *ast.Parser) (*ast.Node, error) {
	return p.Expect(
		op.MinZero(
			op.And{
				op.Not{
					"*/",
				},
				op.Optional(
					CommentText,
				),
				EndLine,
			},
		),
	)
}

func Ws(p *ast.Parser) (*ast.Node, error) {
	return p.Expect(
		op.MinZero(
			op.Or{
				' ',
				0x0009,
				EndLine,
			},
		),
	)
}

func EndLine(p *ast.Parser) (*ast.Node, error) {
	return p.Expect(
		op.Or{
			0x000A,
			0x000D,
			op.And{
				0x000D,
				0x000A,
			},
		},
	)
}

func Test(p *ast.Parser) (*ast.Node, error) {
	return p.Expect(
		ast.Capture{
			Type:        TestT,
			TypeStrings: NodeTypes,
			Value: op.And{
				"assert ",
				Input,
				Ws,
				op.Or{
					TestGoodTmpl,
					TestBadTmpl,
					TestTestTmpl,
				},
				op.Optional(
					op.And{
						' ',
						Description,
					},
				),
				';',
			},
		},
	)
}

func TestGoodTmpl(p *ast.Parser) (*ast.Node, error) {
	return p.Expect(
		op.And{
			':',
			Ws,
			ValuesBr,
		},
	)
}

func TestBadTmpl(p *ast.Parser) (*ast.Node, error) {
	return p.Expect(
		op.And{
			'!',
			TestGoodTmpl,
		},
	)
}

func TestTestTmpl(p *ast.Parser) (*ast.Node, error) {
	return p.Expect(
		op.And{
			"==",
			Ws,
			Input,
			Ws,
			TestGoodTmpl,
		},
	)
}

func ValuesBr(p *ast.Parser) (*ast.Node, error) {
	return p.Expect(
		op.Or{
			"()",
			op.And{
				'(',
				Values,
				op.MinZero(
					op.And{
						", ",
						Values,
					},
				),
				')',
			},
		},
	)
}

func Values(p *ast.Parser) (*ast.Node, error) {
	return p.Expect(
		op.Or{
			Null,
			Bool,
			Nat,
			Int,
			Float,
			Text,
			Reserved,
			Empty,
			Opt,
		},
	)
}

func Null(p *ast.Parser) (*ast.Node, error) {
	return p.Expect(
		ast.Capture{
			Type:        NullT,
			TypeStrings: NodeTypes,
			Value:       "null",
		},
	)
}

func Bool(p *ast.Parser) (*ast.Node, error) {
	return p.Expect(
		ast.Capture{
			Type:        BoolT,
			TypeStrings: NodeTypes,
			Value:       "bool",
		},
	)
}

func Nat(p *ast.Parser) (*ast.Node, error) {
	return p.Expect(
		ast.Capture{
			Type:        NatT,
			TypeStrings: NodeTypes,
			Value: op.And{
				"nat",
				op.Optional(
					Base,
				),
			},
		},
	)
}

func Int(p *ast.Parser) (*ast.Node, error) {
	return p.Expect(
		ast.Capture{
			Type:        IntT,
			TypeStrings: NodeTypes,
			Value: op.And{
				"int",
				op.Optional(
					Base,
				),
			},
		},
	)
}

func Float(p *ast.Parser) (*ast.Node, error) {
	return p.Expect(
		ast.Capture{
			Type:        FloatT,
			TypeStrings: NodeTypes,
			Value: op.And{
				"float",
				Base,
			},
		},
	)
}

func Base(p *ast.Parser) (*ast.Node, error) {
	return p.Expect(
		ast.Capture{
			Type:        BaseT,
			TypeStrings: NodeTypes,
			Value: op.MinOne(
				Digit,
			),
		},
	)
}

func Text(p *ast.Parser) (*ast.Node, error) {
	return p.Expect(
		ast.Capture{
			Type:        TextT,
			TypeStrings: NodeTypes,
			Value:       "text",
		},
	)
}

func Reserved(p *ast.Parser) (*ast.Node, error) {
	return p.Expect(
		ast.Capture{
			Type:        ReservedT,
			TypeStrings: NodeTypes,
			Value:       "reserved",
		},
	)
}

func Empty(p *ast.Parser) (*ast.Node, error) {
	return p.Expect(
		ast.Capture{
			Type:        EmptyT,
			TypeStrings: NodeTypes,
			Value:       "empty",
		},
	)
}

func Opt(p *ast.Parser) (*ast.Node, error) {
	return p.Expect(
		ast.Capture{
			Type:        OptT,
			TypeStrings: NodeTypes,
			Value: op.And{
				"opt ",
				Values,
			},
		},
	)
}

func Input(p *ast.Parser) (*ast.Node, error) {
	return p.Expect(
		op.Or{
			BlobInput,
			TextInput,
		},
	)
}

func TextInput(p *ast.Parser) (*ast.Node, error) {
	return p.Expect(
		ast.Capture{
			Type:        TextInputT,
			TypeStrings: NodeTypes,
			Value:       QuotedString,
		},
	)
}

func BlobInput(p *ast.Parser) (*ast.Node, error) {
	return p.Expect(
		ast.Capture{
			Type:        BlobInputT,
			TypeStrings: NodeTypes,
			Value: op.And{
				"blob ",
				QuotedString,
			},
		},
	)
}

func Description(p *ast.Parser) (*ast.Node, error) {
	return p.Expect(
		ast.Capture{
			Type:        DescriptionT,
			TypeStrings: NodeTypes,
			Value:       QuotedString,
		},
	)
}

func QuotedString(p *ast.Parser) (*ast.Node, error) {
	return p.Expect(
		op.And{
			'"',
			String,
			'"',
		},
	)
}

func String(p *ast.Parser) (*ast.Node, error) {
	return p.Expect(
		op.MinZero(
			Char,
		),
	)
}

func Char(p *ast.Parser) (*ast.Node, error) {
	return p.Expect(
		op.Or{
			Utf,
			op.And{
				ESC,
				op.Repeat(2,
					Hex,
				),
			},
			op.And{
				ESC,
				Escape,
			},
			op.And{
				"\\u{",
				HexNum,
				'}',
			},
		},
	)
}

func HexNum(p *ast.Parser) (*ast.Node, error) {
	return p.Expect(
		op.And{
			Hex,
			op.MinZero(
				op.And{
					op.Optional(
						'_',
					),
					Hex,
				},
			),
		},
	)
}

func Utf(p *ast.Parser) (*ast.Node, error) {
	return p.Expect(
		op.Or{
			Ascii,
			UtfEnc,
		},
	)
}

func UtfEnc(p *ast.Parser) (*ast.Node, error) {
	return p.Expect(
		op.Or{
			op.And{
				parser.CheckRuneRange(0x00C2, 0x00DF),
				Utfcont,
			},
			op.And{
				0x00E0,
				parser.CheckRuneRange(0x00A0, 0x00BF),
				Utfcont,
			},
			op.And{
				0x00ED,
				parser.CheckRuneRange(0x0080, 0x009F),
				Utfcont,
			},
			op.And{
				parser.CheckRuneRange(0x00E1, 0x00EC),
				op.Repeat(2,
					Utfcont,
				),
			},
			op.And{
				parser.CheckRuneRange(0x00EE, 0x00EF),
				op.Repeat(2,
					Utfcont,
				),
			},
			op.And{
				0x00F0,
				parser.CheckRuneRange(0x0090, 0x00BF),
				op.Repeat(2,
					Utfcont,
				),
			},
			op.And{
				0x00F4,
				parser.CheckRuneRange(0x0080, 0x008F),
				op.Repeat(2,
					Utfcont,
				),
			},
			op.And{
				parser.CheckRuneRange(0x00F1, 0x00F3),
				op.Repeat(3,
					Utfcont,
				),
			},
		},
	)
}

func Utfcont(p *parser.Parser) (*parser.Cursor, bool) {
	return p.Check(parser.CheckRuneRange(0x0080, 0x00BF))
}

func Ascii(p *parser.Parser) (*parser.Cursor, bool) {
	return p.Check(op.Or{
		parser.CheckRuneRange(0x0020, 0x0021),
		parser.CheckRuneRange(0x0023, 0x005B),
		parser.CheckRuneRange(0x005D, 0x007E),
	})
}

func Escape(p *parser.Parser) (*parser.Cursor, bool) {
	return p.Check(op.Or{
		'n',
		'r',
		't',
		ESC,
		0x0022,
		0x0027,
	})
}

func Letter(p *parser.Parser) (*parser.Cursor, bool) {
	return p.Check(op.Or{
		parser.CheckRuneRange('A', 'Z'),
		parser.CheckRuneRange('a', 'z'),
	})
}

func Digit(p *parser.Parser) (*parser.Cursor, bool) {
	return p.Check(parser.CheckRuneRange('0', '9'))
}

func Hex(p *parser.Parser) (*parser.Cursor, bool) {
	return p.Check(op.Or{
		Digit,
		parser.CheckRuneRange('A', 'F'),
		parser.CheckRuneRange('a', 'f'),
	})
}

// Token Definitions
const (
	// CANDID-TEST (github.com/di-wu/candid-go/internal/candid_test)

	ESC = 0x005C // \
)

// Node Types
const (
	Unknown = iota

	// CANDID-TEST (github.com/di-wu/candid-go/internal/candid_test)

	TestDataT    // 001
	CommentTextT // 002
	TestT        // 003
	NullT        // 004
	BoolT        // 005
	NatT         // 006
	IntT         // 007
	FloatT       // 008
	BaseT        // 009
	TextT        // 010
	ReservedT    // 011
	EmptyT       // 012
	OptT         // 013
	TextInputT   // 014
	BlobInputT   // 015
	DescriptionT // 016
)

var NodeTypes = []string{
	"UNKNOWN",

	// CANDID-TEST (github.com/di-wu/candid-go/internal/candid_test)

	"TestData",
	"CommentText",
	"Test",
	"Null",
	"Bool",
	"Nat",
	"Int",
	"Float",
	"Base",
	"Text",
	"Reserved",
	"Empty",
	"Opt",
	"TextInput",
	"BlobInput",
	"Description",
}