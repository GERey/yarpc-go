// Code generated by thriftrw v1.2.0
// @generated

// Copyright (c) 2017 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package enums

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go.uber.org/thriftrw/wire"
	"math"
	"strconv"
	"strings"
)

type EmptyEnum int32

func (v EmptyEnum) ToWire() (wire.Value, error) {
	return wire.NewValueI32(int32(v)), nil
}

func (v *EmptyEnum) FromWire(w wire.Value) error {
	*v = (EmptyEnum)(w.GetI32())
	return nil
}

func (v EmptyEnum) String() string {
	w := int32(v)
	return fmt.Sprintf("EmptyEnum(%d)", w)
}

func (v EmptyEnum) Equals(rhs EmptyEnum) bool {
	return v == rhs
}

func (v EmptyEnum) MarshalJSON() ([]byte, error) {
	return ([]byte)(strconv.FormatInt(int64(v), 10)), nil
}

func (v *EmptyEnum) UnmarshalJSON(text []byte) error {
	d := json.NewDecoder(bytes.NewReader(text))
	d.UseNumber()
	t, err := d.Token()
	if err != nil {
		return err
	}
	switch w := t.(type) {
	case json.Number:
		x, err := w.Int64()
		if err != nil {
			return err
		}
		if x > math.MaxInt32 {
			return fmt.Errorf("enum overflow from JSON %q for %q", text, "EmptyEnum")
		}
		if x < math.MinInt32 {
			return fmt.Errorf("enum underflow from JSON %q for %q", text, "EmptyEnum")
		}
		*v = (EmptyEnum)(x)
		return nil
	case string:
		switch w {
		default:
			return fmt.Errorf("unknown enum value %q for %q", w, "EmptyEnum")
		}
	default:
		return fmt.Errorf("invalid JSON value %q (%T) to unmarshal into %q", t, t, "EmptyEnum")
	}
}

type EnumDefault int32

const (
	EnumDefaultFoo EnumDefault = 0
	EnumDefaultBar EnumDefault = 1
	EnumDefaultBaz EnumDefault = 2
)

func (v EnumDefault) ToWire() (wire.Value, error) {
	return wire.NewValueI32(int32(v)), nil
}

func (v *EnumDefault) FromWire(w wire.Value) error {
	*v = (EnumDefault)(w.GetI32())
	return nil
}

func (v EnumDefault) String() string {
	w := int32(v)
	switch w {
	case 0:
		return "Foo"
	case 1:
		return "Bar"
	case 2:
		return "Baz"
	}
	return fmt.Sprintf("EnumDefault(%d)", w)
}

func (v EnumDefault) Equals(rhs EnumDefault) bool {
	return v == rhs
}

func (v EnumDefault) MarshalJSON() ([]byte, error) {
	switch int32(v) {
	case 0:
		return ([]byte)("\"Foo\""), nil
	case 1:
		return ([]byte)("\"Bar\""), nil
	case 2:
		return ([]byte)("\"Baz\""), nil
	}
	return ([]byte)(strconv.FormatInt(int64(v), 10)), nil
}

func (v *EnumDefault) UnmarshalJSON(text []byte) error {
	d := json.NewDecoder(bytes.NewReader(text))
	d.UseNumber()
	t, err := d.Token()
	if err != nil {
		return err
	}
	switch w := t.(type) {
	case json.Number:
		x, err := w.Int64()
		if err != nil {
			return err
		}
		if x > math.MaxInt32 {
			return fmt.Errorf("enum overflow from JSON %q for %q", text, "EnumDefault")
		}
		if x < math.MinInt32 {
			return fmt.Errorf("enum underflow from JSON %q for %q", text, "EnumDefault")
		}
		*v = (EnumDefault)(x)
		return nil
	case string:
		switch w {
		case "Foo":
			*v = EnumDefaultFoo
			return nil
		case "Bar":
			*v = EnumDefaultBar
			return nil
		case "Baz":
			*v = EnumDefaultBaz
			return nil
		default:
			return fmt.Errorf("unknown enum value %q for %q", w, "EnumDefault")
		}
	default:
		return fmt.Errorf("invalid JSON value %q (%T) to unmarshal into %q", t, t, "EnumDefault")
	}
}

type EnumWithDuplicateName int32

const (
	EnumWithDuplicateNameA EnumWithDuplicateName = 0
	EnumWithDuplicateNameB EnumWithDuplicateName = 1
	EnumWithDuplicateNameC EnumWithDuplicateName = 2
	EnumWithDuplicateNameP EnumWithDuplicateName = 3
	EnumWithDuplicateNameQ EnumWithDuplicateName = 4
	EnumWithDuplicateNameR EnumWithDuplicateName = 5
	EnumWithDuplicateNameX EnumWithDuplicateName = 6
	EnumWithDuplicateNameY EnumWithDuplicateName = 7
	EnumWithDuplicateNameZ EnumWithDuplicateName = 8
)

func (v EnumWithDuplicateName) ToWire() (wire.Value, error) {
	return wire.NewValueI32(int32(v)), nil
}

func (v *EnumWithDuplicateName) FromWire(w wire.Value) error {
	*v = (EnumWithDuplicateName)(w.GetI32())
	return nil
}

func (v EnumWithDuplicateName) String() string {
	w := int32(v)
	switch w {
	case 0:
		return "A"
	case 1:
		return "B"
	case 2:
		return "C"
	case 3:
		return "P"
	case 4:
		return "Q"
	case 5:
		return "R"
	case 6:
		return "X"
	case 7:
		return "Y"
	case 8:
		return "Z"
	}
	return fmt.Sprintf("EnumWithDuplicateName(%d)", w)
}

func (v EnumWithDuplicateName) Equals(rhs EnumWithDuplicateName) bool {
	return v == rhs
}

func (v EnumWithDuplicateName) MarshalJSON() ([]byte, error) {
	switch int32(v) {
	case 0:
		return ([]byte)("\"A\""), nil
	case 1:
		return ([]byte)("\"B\""), nil
	case 2:
		return ([]byte)("\"C\""), nil
	case 3:
		return ([]byte)("\"P\""), nil
	case 4:
		return ([]byte)("\"Q\""), nil
	case 5:
		return ([]byte)("\"R\""), nil
	case 6:
		return ([]byte)("\"X\""), nil
	case 7:
		return ([]byte)("\"Y\""), nil
	case 8:
		return ([]byte)("\"Z\""), nil
	}
	return ([]byte)(strconv.FormatInt(int64(v), 10)), nil
}

func (v *EnumWithDuplicateName) UnmarshalJSON(text []byte) error {
	d := json.NewDecoder(bytes.NewReader(text))
	d.UseNumber()
	t, err := d.Token()
	if err != nil {
		return err
	}
	switch w := t.(type) {
	case json.Number:
		x, err := w.Int64()
		if err != nil {
			return err
		}
		if x > math.MaxInt32 {
			return fmt.Errorf("enum overflow from JSON %q for %q", text, "EnumWithDuplicateName")
		}
		if x < math.MinInt32 {
			return fmt.Errorf("enum underflow from JSON %q for %q", text, "EnumWithDuplicateName")
		}
		*v = (EnumWithDuplicateName)(x)
		return nil
	case string:
		switch w {
		case "A":
			*v = EnumWithDuplicateNameA
			return nil
		case "B":
			*v = EnumWithDuplicateNameB
			return nil
		case "C":
			*v = EnumWithDuplicateNameC
			return nil
		case "P":
			*v = EnumWithDuplicateNameP
			return nil
		case "Q":
			*v = EnumWithDuplicateNameQ
			return nil
		case "R":
			*v = EnumWithDuplicateNameR
			return nil
		case "X":
			*v = EnumWithDuplicateNameX
			return nil
		case "Y":
			*v = EnumWithDuplicateNameY
			return nil
		case "Z":
			*v = EnumWithDuplicateNameZ
			return nil
		default:
			return fmt.Errorf("unknown enum value %q for %q", w, "EnumWithDuplicateName")
		}
	default:
		return fmt.Errorf("invalid JSON value %q (%T) to unmarshal into %q", t, t, "EnumWithDuplicateName")
	}
}

type EnumWithDuplicateValues int32

const (
	EnumWithDuplicateValuesP EnumWithDuplicateValues = 0
	EnumWithDuplicateValuesQ EnumWithDuplicateValues = -1
	EnumWithDuplicateValuesR EnumWithDuplicateValues = 0
)

func (v EnumWithDuplicateValues) ToWire() (wire.Value, error) {
	return wire.NewValueI32(int32(v)), nil
}

func (v *EnumWithDuplicateValues) FromWire(w wire.Value) error {
	*v = (EnumWithDuplicateValues)(w.GetI32())
	return nil
}

func (v EnumWithDuplicateValues) String() string {
	w := int32(v)
	switch w {
	case 0:
		return "P"
	case -1:
		return "Q"
	}
	return fmt.Sprintf("EnumWithDuplicateValues(%d)", w)
}

func (v EnumWithDuplicateValues) Equals(rhs EnumWithDuplicateValues) bool {
	return v == rhs
}

func (v EnumWithDuplicateValues) MarshalJSON() ([]byte, error) {
	switch int32(v) {
	case 0:
		return ([]byte)("\"P\""), nil
	case -1:
		return ([]byte)("\"Q\""), nil
	}
	return ([]byte)(strconv.FormatInt(int64(v), 10)), nil
}

func (v *EnumWithDuplicateValues) UnmarshalJSON(text []byte) error {
	d := json.NewDecoder(bytes.NewReader(text))
	d.UseNumber()
	t, err := d.Token()
	if err != nil {
		return err
	}
	switch w := t.(type) {
	case json.Number:
		x, err := w.Int64()
		if err != nil {
			return err
		}
		if x > math.MaxInt32 {
			return fmt.Errorf("enum overflow from JSON %q for %q", text, "EnumWithDuplicateValues")
		}
		if x < math.MinInt32 {
			return fmt.Errorf("enum underflow from JSON %q for %q", text, "EnumWithDuplicateValues")
		}
		*v = (EnumWithDuplicateValues)(x)
		return nil
	case string:
		switch w {
		case "P":
			*v = EnumWithDuplicateValuesP
			return nil
		case "Q":
			*v = EnumWithDuplicateValuesQ
			return nil
		case "R":
			*v = EnumWithDuplicateValuesR
			return nil
		default:
			return fmt.Errorf("unknown enum value %q for %q", w, "EnumWithDuplicateValues")
		}
	default:
		return fmt.Errorf("invalid JSON value %q (%T) to unmarshal into %q", t, t, "EnumWithDuplicateValues")
	}
}

type EnumWithValues int32

const (
	EnumWithValuesX EnumWithValues = 123
	EnumWithValuesY EnumWithValues = 456
	EnumWithValuesZ EnumWithValues = 789
)

func (v EnumWithValues) ToWire() (wire.Value, error) {
	return wire.NewValueI32(int32(v)), nil
}

func (v *EnumWithValues) FromWire(w wire.Value) error {
	*v = (EnumWithValues)(w.GetI32())
	return nil
}

func (v EnumWithValues) String() string {
	w := int32(v)
	switch w {
	case 123:
		return "X"
	case 456:
		return "Y"
	case 789:
		return "Z"
	}
	return fmt.Sprintf("EnumWithValues(%d)", w)
}

func (v EnumWithValues) Equals(rhs EnumWithValues) bool {
	return v == rhs
}

func (v EnumWithValues) MarshalJSON() ([]byte, error) {
	switch int32(v) {
	case 123:
		return ([]byte)("\"X\""), nil
	case 456:
		return ([]byte)("\"Y\""), nil
	case 789:
		return ([]byte)("\"Z\""), nil
	}
	return ([]byte)(strconv.FormatInt(int64(v), 10)), nil
}

func (v *EnumWithValues) UnmarshalJSON(text []byte) error {
	d := json.NewDecoder(bytes.NewReader(text))
	d.UseNumber()
	t, err := d.Token()
	if err != nil {
		return err
	}
	switch w := t.(type) {
	case json.Number:
		x, err := w.Int64()
		if err != nil {
			return err
		}
		if x > math.MaxInt32 {
			return fmt.Errorf("enum overflow from JSON %q for %q", text, "EnumWithValues")
		}
		if x < math.MinInt32 {
			return fmt.Errorf("enum underflow from JSON %q for %q", text, "EnumWithValues")
		}
		*v = (EnumWithValues)(x)
		return nil
	case string:
		switch w {
		case "X":
			*v = EnumWithValuesX
			return nil
		case "Y":
			*v = EnumWithValuesY
			return nil
		case "Z":
			*v = EnumWithValuesZ
			return nil
		default:
			return fmt.Errorf("unknown enum value %q for %q", w, "EnumWithValues")
		}
	default:
		return fmt.Errorf("invalid JSON value %q (%T) to unmarshal into %q", t, t, "EnumWithValues")
	}
}

type RecordType int32

const (
	RecordTypeName        RecordType = 0
	RecordTypeHomeAddress RecordType = 1
	RecordTypeWorkAddress RecordType = 2
)

func (v RecordType) ToWire() (wire.Value, error) {
	return wire.NewValueI32(int32(v)), nil
}

func (v *RecordType) FromWire(w wire.Value) error {
	*v = (RecordType)(w.GetI32())
	return nil
}

func (v RecordType) String() string {
	w := int32(v)
	switch w {
	case 0:
		return "NAME"
	case 1:
		return "HOME_ADDRESS"
	case 2:
		return "WORK_ADDRESS"
	}
	return fmt.Sprintf("RecordType(%d)", w)
}

func (v RecordType) Equals(rhs RecordType) bool {
	return v == rhs
}

func (v RecordType) MarshalJSON() ([]byte, error) {
	switch int32(v) {
	case 0:
		return ([]byte)("\"NAME\""), nil
	case 1:
		return ([]byte)("\"HOME_ADDRESS\""), nil
	case 2:
		return ([]byte)("\"WORK_ADDRESS\""), nil
	}
	return ([]byte)(strconv.FormatInt(int64(v), 10)), nil
}

func (v *RecordType) UnmarshalJSON(text []byte) error {
	d := json.NewDecoder(bytes.NewReader(text))
	d.UseNumber()
	t, err := d.Token()
	if err != nil {
		return err
	}
	switch w := t.(type) {
	case json.Number:
		x, err := w.Int64()
		if err != nil {
			return err
		}
		if x > math.MaxInt32 {
			return fmt.Errorf("enum overflow from JSON %q for %q", text, "RecordType")
		}
		if x < math.MinInt32 {
			return fmt.Errorf("enum underflow from JSON %q for %q", text, "RecordType")
		}
		*v = (RecordType)(x)
		return nil
	case string:
		switch w {
		case "NAME":
			*v = RecordTypeName
			return nil
		case "HOME_ADDRESS":
			*v = RecordTypeHomeAddress
			return nil
		case "WORK_ADDRESS":
			*v = RecordTypeWorkAddress
			return nil
		default:
			return fmt.Errorf("unknown enum value %q for %q", w, "RecordType")
		}
	default:
		return fmt.Errorf("invalid JSON value %q (%T) to unmarshal into %q", t, t, "RecordType")
	}
}

type StructWithOptionalEnum struct {
	E *EnumDefault `json:"e,omitempty"`
}

func (v *StructWithOptionalEnum) ToWire() (wire.Value, error) {
	var (
		fields [1]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)
	if v.E != nil {
		w, err = v.E.ToWire()
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 1, Value: w}
		i++
	}
	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

func _EnumDefault_Read(w wire.Value) (EnumDefault, error) {
	var v EnumDefault
	err := v.FromWire(w)
	return v, err
}

func (v *StructWithOptionalEnum) FromWire(w wire.Value) error {
	var err error
	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 1:
			if field.Value.Type() == wire.TI32 {
				var x EnumDefault
				x, err = _EnumDefault_Read(field.Value)
				v.E = &x
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func (v *StructWithOptionalEnum) String() string {
	if v == nil {
		return "<nil>"
	}
	var fields [1]string
	i := 0
	if v.E != nil {
		fields[i] = fmt.Sprintf("E: %v", *(v.E))
		i++
	}
	return fmt.Sprintf("StructWithOptionalEnum{%v}", strings.Join(fields[:i], ", "))
}

func _EnumDefault_EqualsPtr(lhs, rhs *EnumDefault) bool {
	if lhs != nil && rhs != nil {
		x := *lhs
		y := *rhs
		return x.Equals(y)
	}
	return lhs == nil && rhs == nil
}

func (v *StructWithOptionalEnum) Equals(rhs *StructWithOptionalEnum) bool {
	if !_EnumDefault_EqualsPtr(v.E, rhs.E) {
		return false
	}
	return true
}

type LowerCaseEnum int32

const (
	LowerCaseEnumContaining LowerCaseEnum = 0
	LowerCaseEnumLowerCase  LowerCaseEnum = 1
	LowerCaseEnumItems      LowerCaseEnum = 2
)

func (v LowerCaseEnum) ToWire() (wire.Value, error) {
	return wire.NewValueI32(int32(v)), nil
}

func (v *LowerCaseEnum) FromWire(w wire.Value) error {
	*v = (LowerCaseEnum)(w.GetI32())
	return nil
}

func (v LowerCaseEnum) String() string {
	w := int32(v)
	switch w {
	case 0:
		return "containing"
	case 1:
		return "lower_case"
	case 2:
		return "items"
	}
	return fmt.Sprintf("LowerCaseEnum(%d)", w)
}

func (v LowerCaseEnum) Equals(rhs LowerCaseEnum) bool {
	return v == rhs
}

func (v LowerCaseEnum) MarshalJSON() ([]byte, error) {
	switch int32(v) {
	case 0:
		return ([]byte)("\"containing\""), nil
	case 1:
		return ([]byte)("\"lower_case\""), nil
	case 2:
		return ([]byte)("\"items\""), nil
	}
	return ([]byte)(strconv.FormatInt(int64(v), 10)), nil
}

func (v *LowerCaseEnum) UnmarshalJSON(text []byte) error {
	d := json.NewDecoder(bytes.NewReader(text))
	d.UseNumber()
	t, err := d.Token()
	if err != nil {
		return err
	}
	switch w := t.(type) {
	case json.Number:
		x, err := w.Int64()
		if err != nil {
			return err
		}
		if x > math.MaxInt32 {
			return fmt.Errorf("enum overflow from JSON %q for %q", text, "LowerCaseEnum")
		}
		if x < math.MinInt32 {
			return fmt.Errorf("enum underflow from JSON %q for %q", text, "LowerCaseEnum")
		}
		*v = (LowerCaseEnum)(x)
		return nil
	case string:
		switch w {
		case "containing":
			*v = LowerCaseEnumContaining
			return nil
		case "lower_case":
			*v = LowerCaseEnumLowerCase
			return nil
		case "items":
			*v = LowerCaseEnumItems
			return nil
		default:
			return fmt.Errorf("unknown enum value %q for %q", w, "LowerCaseEnum")
		}
	default:
		return fmt.Errorf("invalid JSON value %q (%T) to unmarshal into %q", t, t, "LowerCaseEnum")
	}
}