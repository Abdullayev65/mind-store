package hash

import (
	"errors"
	"reflect"
)

type Int int

var hashInt HashInteger[int] = NewHashInt()

func (i Int) Int() int {
	return int(i)
}
func (i Int) I64() int64 {
	return int64(i)
}
func (i Int) I32() int32 {
	return int32(i)
}
func (i Int) HashToStr() string {
	encode, err := hashInt.Encode(int(i))
	if err != nil {
		return ""
	}
	return string(encode)
}
func (i *Int) UnhashStr(hashStr string) error {
	if i == nil {
		return errors.New("hash.Int: i can't be nil")
	}

	decode, err := hashInt.Decode([]byte(hashStr))
	if err != nil {
		return err
	}

	*i = Int(decode)
	return nil
}

func (i *Int) GoMarshal() ([]byte, error) {
	if i == nil {
		return nil, errors.New("hash.Int can't be <nil>")
	}

	encode, err := hashInt.Encode(int(*i))
	if err != nil {
		return nil, err
	}

	encode = append([]byte{'"'}, encode...)
	encode = append(encode, '"')

	return encode, nil
}
func (i *Int) GoUnmarshal(data []byte) error {
	if isIn2QuotationMark(data) {
		data = data[1 : len(data)-1]
	}
	decode, err := hashInt.Decode(data)
	if err != nil {
		return err
	}

	if i == nil {
		return errors.New("hash.Int can't set value, because pointer is nil")
	}

	*i = Int(decode)

	return nil
}

func (i *Int) MarshalJSON() ([]byte, error) {
	return i.GoMarshal()
}

func (i *Int) UnmarshalJSON(data []byte) error {
	return i.GoUnmarshal(data)
}

func (i Int) UnmarshalSetter(data []byte, setter reflect.Value) error {
	err := i.GoUnmarshal(data)
	if err != nil {
		return err
	}

	setter.SetInt(int64(i))
	return nil
}

func isIn2QuotationMark(data []byte) bool {
	if len(data) < 2 {
		return false
	}

	return data[0] == '"' && data[len(data)-1] == '"'
}
