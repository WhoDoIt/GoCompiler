package interpreter

import (
	"github.com/WhoDoIt/GoCompiler/internal/tokenizer"
)

type RoseType interface {
	getType() string
	zeroValue() RoseType
	operatorBinary(operator tokenizer.TokenType, other RoseType) RoseType
	operatorUnary(operator tokenizer.TokenType) RoseType
	operatorCall(args []RoseType) RoseType
}

type RoseString struct {
	value string
}

type RoseInt struct {
	value int
}

type RoseBool struct {
	value bool
}

type RuntimeError struct {
	value string
}

func tryDifferentTypesError(a RoseType, b RoseType) RuntimeError {
	if val, ok := a.(RuntimeError); ok {
		return val
	}
	if val, ok := b.(RuntimeError); ok {
		return val
	}
	return RuntimeError{value: "unsupported operation of (" + a.getType() + " and " + b.getType() + ")"}
}

func (s RoseString) getType() string {
	return "String"
}

func (s RoseString) zeroValue() RoseType {
	return RoseString{value: ""}
}

func (s RoseString) operatorBinary(operator tokenizer.TokenType, other RoseType) RoseType {
	if other.getType() != s.getType() {
		return tryDifferentTypesError(s, other)
	}
	switch operator {
	case tokenizer.PLUS:
		return RoseString{value: s.value + other.(RoseString).value}
	case tokenizer.EQUAL_EQUAL:
		return RoseBool{value: s.value == other.(RoseString).value}
	case tokenizer.LESS:
		return RoseBool{value: s.value < other.(RoseString).value}
	}
	return tryDifferentTypesError(s, s)
}

func (s RoseString) operatorUnary(operator tokenizer.TokenType) RoseType {
	return tryDifferentTypesError(s, s)
}

func (s RoseString) operatorCall(args []RoseType) RoseType {
	return tryDifferentTypesError(s, s)
}

func (s RoseInt) getType() string {
	return "Int"
}

func (s RoseInt) zeroValue() RoseType {
	return RoseInt{value: 0}
}

func (s RoseInt) operatorBinary(operator tokenizer.TokenType, other RoseType) RoseType {
	if other.getType() != s.getType() {
		return tryDifferentTypesError(s, other)
	}
	switch operator {
	case tokenizer.PLUS:
		return RoseInt{value: s.value + other.(RoseInt).value}
	case tokenizer.MINUS:
		return RoseInt{value: s.value - other.(RoseInt).value}
	case tokenizer.STAR:
		return RoseInt{value: s.value * other.(RoseInt).value}
	case tokenizer.SLASH:
		return RoseInt{value: s.value / other.(RoseInt).value}
	case tokenizer.PIPE:
		return RoseInt{value: s.value & other.(RoseInt).value}
	case tokenizer.AMPERSAND:
		return RoseInt{value: s.value + other.(RoseInt).value}
	case tokenizer.EQUAL_EQUAL:
		return RoseBool{value: s.value == other.(RoseInt).value}
	case tokenizer.LESS:
		return RoseBool{value: s.value < other.(RoseInt).value}
	}
	return tryDifferentTypesError(s, s)
}

func (s RoseInt) operatorUnary(operator tokenizer.TokenType) RoseType {
	if operator == tokenizer.MINUS {
		return RoseInt{value: -s.value}
	}
	return tryDifferentTypesError(s, s)
}

func (s RoseInt) operatorCall(args []RoseType) RoseType {
	return tryDifferentTypesError(s, s)
}

func (s RoseBool) getType() string {
	return "Bool"
}

func (s RoseBool) zeroValue() RoseType {
	return RoseBool{value: false}
}

func (s RoseBool) operatorBinary(operator tokenizer.TokenType, other RoseType) RoseType {
	return tryDifferentTypesError(s, s)
}

func (s RoseBool) operatorUnary(operator tokenizer.TokenType) RoseType {
	if operator == tokenizer.EXCLAMATION {
		return RoseBool{value: !s.value}
	}
	return tryDifferentTypesError(s, s)
}

func (s RoseBool) operatorCall(args []RoseType) RoseType {
	return tryDifferentTypesError(s, s)
}

func (s RuntimeError) getType() string {
	return "RuntimeError"
}

func (s RuntimeError) zeroValue() RoseType {
	return s
}
func (s RuntimeError) operatorBinary(operator tokenizer.TokenType, other RoseType) RoseType {
	return s
}

func (s RuntimeError) operatorUnary(operator tokenizer.TokenType) RoseType {
	return s
}

func (s RuntimeError) operatorCall(args []RoseType) RoseType {
	return s
}
