package interpreter

type WorkingType interface {
	getType() string
	zeroValue() WorkingType
	operatorPlus(other WorkingType) WorkingType
	operatorMinus(other WorkingType) WorkingType
	operatorSlash(other WorkingType) WorkingType
	operatorStar(other WorkingType) WorkingType
	operatorExclamation() WorkingType
	operatorSelfminus() WorkingType
	operatorLess(other WorkingType) Bool
	operatorEqual(other WorkingType) Bool
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

func (s String) getType() string {
	return "String"
}

func (s String) zeroValue() WorkingType {
	return String{value: ""}
}

func (s String) operatorPlus(other WorkingType) WorkingType {
	if _, ok := other.(String); !ok {
		return nil
	}
	return String{value: s.value + other.(String).value}
}

func (s String) operatorMinus(other WorkingType) WorkingType {
	return nil
}

func (s String) operatorSlash(other WorkingType) WorkingType {
	return nil
}

func (s String) operatorStar(other WorkingType) WorkingType {
	return nil
}

func (s String) operatorExclamation() WorkingType {
	return nil
}

func (s String) operatorSelfminus() WorkingType {
	return nil
}

func (s String) operatorLess(other WorkingType) Bool {
	if _, ok := other.(String); !ok {
		return Bool{}
	}
	return Bool{value: s.value < other.(String).value}
}

func (s String) operatorEqual(other WorkingType) Bool {
	if _, ok := other.(String); !ok {
		return Bool{}
	}
	return Bool{value: s.value == other.(String).value}
}

func (s Int) getType() string {
	return "Int"
}

func (s Int) zeroValue() WorkingType {
	return Int{value: 0}
}

func (s Int) operatorPlus(other WorkingType) WorkingType {
	if _, ok := other.(Int); !ok {
		return nil
	}
	return Int{value: s.value + other.(Int).value}
}

func (s Int) operatorMinus(other WorkingType) WorkingType {
	if _, ok := other.(Int); !ok {
		return nil
	}
	return Int{value: s.value - other.(Int).value}
}

func (s Int) operatorSlash(other WorkingType) WorkingType {
	if _, ok := other.(Int); !ok {
		return nil
	}
	return Int{value: s.value / other.(Int).value}
}

func (s Int) operatorStar(other WorkingType) WorkingType {
	if _, ok := other.(Int); !ok {
		return nil
	}
	return Int{value: s.value * other.(Int).value}
}

func (s Int) operatorExclamation() WorkingType {
	return nil
}

func (s Int) operatorSelfminus() WorkingType {
	return Int{value: -s.value}
}

func (s Int) operatorLess(other WorkingType) Bool {
	if _, ok := other.(Int); !ok {
		return Bool{}
	}
	return Bool{value: s.value < other.(Int).value}
}

func (s Int) operatorEqual(other WorkingType) Bool {
	if _, ok := other.(Int); !ok {
		return Bool{}
	}
	return Bool{value: s.value == other.(Int).value}
}

func (s Bool) getType() string {
	return "Bool"
}

func (s Bool) zeroValue() WorkingType {
	return Bool{value: false}
}

func (s Bool) operatorPlus(other WorkingType) WorkingType {
	return nil
}

func (s Bool) operatorMinus(other WorkingType) WorkingType {
	return nil
}

func (s Bool) operatorSlash(other WorkingType) WorkingType {
	return nil
}

func (s Bool) operatorStar(other WorkingType) WorkingType {
	return nil
}

func (s Bool) operatorExclamation() WorkingType {
	return Bool{value: !s.value}
}

func (s Bool) operatorSelfminus() WorkingType {
	return nil
}

func (s Bool) operatorLess(other WorkingType) Bool {
	return Bool{}
}

func (s Bool) operatorEqual(other WorkingType) Bool {
	if _, ok := other.(Bool); !ok {
		return Bool{}
	}
	return Bool{value: s.value == other.(Bool).value}
}
