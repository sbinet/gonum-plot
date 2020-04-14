// Code generated by "stringer -type Kind"; DO NOT EDIT.

package token

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Invalid-0]
	_ = x[Macro-1]
	_ = x[EmptyLine-2]
	_ = x[Comment-3]
	_ = x[Space-4]
	_ = x[Word-5]
	_ = x[Number-6]
	_ = x[Dollar-7]
	_ = x[Lbrace-8]
	_ = x[Rbrace-9]
	_ = x[Lbrack-10]
	_ = x[Rbrack-11]
	_ = x[Equal-12]
	_ = x[Underscore-13]
	_ = x[Lparen-14]
	_ = x[Rparen-15]
	_ = x[Lt-16]
	_ = x[Gt-17]
	_ = x[Hat-18]
	_ = x[Div-19]
	_ = x[Mul-20]
	_ = x[Sub-21]
	_ = x[Add-22]
	_ = x[Not-23]
	_ = x[Colon-24]
	_ = x[Backslash-25]
	_ = x[Other-26]
	_ = x[Verbatim-27]
	_ = x[EOF-28]
}

const _Kind_name = "InvalidMacroEmptyLineCommentSpaceWordNumberDollarLbraceRbraceLbrackRbrackEqualUnderscoreLparenRparenLtGtHatDivMulSubAddNotColonBackslashOtherVerbatimEOF"

var _Kind_index = [...]uint8{0, 7, 12, 21, 28, 33, 37, 43, 49, 55, 61, 67, 73, 78, 88, 94, 100, 102, 104, 107, 110, 113, 116, 119, 122, 127, 136, 141, 149, 152}

func (i Kind) String() string {
	if i < 0 || i >= Kind(len(_Kind_index)-1) {
		return "Kind(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Kind_name[_Kind_index[i]:_Kind_index[i+1]]
}
