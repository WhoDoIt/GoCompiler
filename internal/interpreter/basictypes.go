package interpreter

type WorkingType interface {
	zeroValue() WorkingType
	operatorPlus(other WorkingType) WorkingType
	operatorMinus(other WorkingType) WorkingType
	operatorSlash(other WorkingType) WorkingType
	operatorStar(other WorkingType) WorkingType
	operatorExclamation() WorkingType
	operatorSelfminus() WorkingType
}

type String struct {
	value string
}

type Int struct {
	value int
}

type Bool struct {
	value bool
}

func (s String) zeroValue() WorkingType {
	return String{value: ""}
}

func (s String) operatorPlus(other WorkingType) WorkingType {
	if _, ok := other.(String); !ok {
		panic("not implemented")
	}
	return String{value: s.value + other.(String).value}
}

func (s String) operatorMinus(other WorkingType) WorkingType {
	panic("not implemented") // TODO: Implement
}

func (s String) operatorSlash(other WorkingType) WorkingType {
	panic("not implemented") // TODO: Implement
}

func (s String) operatorStar(other WorkingType) WorkingType {
	panic("not implemented") // TODO: Implement
}

func (s String) operatorExclamation() WorkingType {
	panic("not implemented") // TODO: Implement
}

func (s String) operatorSelfminus() WorkingType {
	panic("not implemented") // TODO: Implement
}

func (s Int) zeroValue() WorkingType {
	return Int{value: 0}
}

func (s Int) operatorPlus(other WorkingType) WorkingType {
	if _, ok := other.(Int); !ok {
		panic("not implemented")
	}
	return Int{value: s.value + other.(Int).value}
}

func (s Int) operatorMinus(other WorkingType) WorkingType {
	if _, ok := other.(Int); !ok {
		panic("not implemented")
	}
	return Int{value: s.value - other.(Int).value}
}

func (s Int) operatorSlash(other WorkingType) WorkingType {
	if _, ok := other.(Int); !ok {
		panic("not implemented")
	}
	return Int{value: s.value / other.(Int).value}
}

func (s Int) operatorStar(other WorkingType) WorkingType {
	if _, ok := other.(Int); !ok {
		panic("not implemented")
	}
	return Int{value: s.value * other.(Int).value}
}

func (s Int) operatorExclamation() WorkingType {
	panic("not implemented") // TODO: Implement
}

func (s Int) operatorSelfminus() WorkingType {
	return Int{value: -s.value}
}
