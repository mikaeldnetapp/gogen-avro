// Code generated by github.com/actgardner/gogen-avro/v7. DO NOT EDIT.
/*
 * SOURCE:
 *     recursive.avsc
 */
package avro

import (
	"github.com/actgardner/gogen-avro/v7/compiler"
	"github.com/actgardner/gogen-avro/v7/vm"
	"github.com/actgardner/gogen-avro/v7/vm/types"
	"io"
)

type RecursiveUnionTestRecord struct {
	RecursiveField *RecursiveFieldUnion `json:"RecursiveField"`
}

const RecursiveUnionTestRecordAvroCRC64Fingerprint = "\xc6U)C\v\x8a\xa6\x89"

func NewRecursiveUnionTestRecord() *RecursiveUnionTestRecord {
	return &RecursiveUnionTestRecord{}
}

func DeserializeRecursiveUnionTestRecord(r io.Reader) (*RecursiveUnionTestRecord, error) {
	t := NewRecursiveUnionTestRecord()
	deser, err := compiler.CompileSchemaBytes([]byte(t.Schema()), []byte(t.Schema()))
	if err != nil {
		return nil, err
	}

	err = vm.Eval(r, deser, t)
	if err != nil {
		return nil, err
	}
	return t, err
}

func DeserializeRecursiveUnionTestRecordFromSchema(r io.Reader, schema string) (*RecursiveUnionTestRecord, error) {
	t := NewRecursiveUnionTestRecord()

	deser, err := compiler.CompileSchemaBytes([]byte(schema), []byte(t.Schema()))
	if err != nil {
		return nil, err
	}

	err = vm.Eval(r, deser, t)
	if err != nil {
		return nil, err
	}
	return t, err
}

func writeRecursiveUnionTestRecord(r *RecursiveUnionTestRecord, w io.Writer) error {
	var err error
	err = writeRecursiveFieldUnion(r.RecursiveField, w)
	if err != nil {
		return err
	}
	return err
}

func (r *RecursiveUnionTestRecord) Serialize(w io.Writer) error {
	return writeRecursiveUnionTestRecord(r, w)
}

func (r *RecursiveUnionTestRecord) Schema() string {
	return "{\"fields\":[{\"name\":\"RecursiveField\",\"type\":[\"null\",\"RecursiveUnionTestRecord\"]}],\"name\":\"RecursiveUnionTestRecord\",\"type\":\"record\"}"
}

func (r *RecursiveUnionTestRecord) SchemaName() string {
	return "RecursiveUnionTestRecord"
}

func (_ *RecursiveUnionTestRecord) SetBoolean(v bool)    { panic("Unsupported operation") }
func (_ *RecursiveUnionTestRecord) SetInt(v int32)       { panic("Unsupported operation") }
func (_ *RecursiveUnionTestRecord) SetLong(v int64)      { panic("Unsupported operation") }
func (_ *RecursiveUnionTestRecord) SetFloat(v float32)   { panic("Unsupported operation") }
func (_ *RecursiveUnionTestRecord) SetDouble(v float64)  { panic("Unsupported operation") }
func (_ *RecursiveUnionTestRecord) SetBytes(v []byte)    { panic("Unsupported operation") }
func (_ *RecursiveUnionTestRecord) SetString(v string)   { panic("Unsupported operation") }
func (_ *RecursiveUnionTestRecord) SetUnionElem(v int64) { panic("Unsupported operation") }

func (r *RecursiveUnionTestRecord) Get(i int) types.Field {
	switch i {
	case 0:
		r.RecursiveField = NewRecursiveFieldUnion()

		return r.RecursiveField
	}
	panic("Unknown field index")
}

func (r *RecursiveUnionTestRecord) SetDefault(i int) {
	switch i {
	}
	panic("Unknown field index")
}

func (r *RecursiveUnionTestRecord) NullField(i int) {
	switch i {
	case 0:
		r.RecursiveField = nil
		return
	}
	panic("Not a nullable field index")
}

func (_ *RecursiveUnionTestRecord) AppendMap(key string) types.Field { panic("Unsupported operation") }
func (_ *RecursiveUnionTestRecord) AppendArray() types.Field         { panic("Unsupported operation") }
func (_ *RecursiveUnionTestRecord) Finalize()                        {}

func (_ *RecursiveUnionTestRecord) AvroCRC64Fingerprint() []byte {
	return []byte(RecursiveUnionTestRecordAvroCRC64Fingerprint)
}
