package tablemodel

var String = "string"
var Number = "number"
var Binary = "binary"
var StringSet = "stringSet"
var NumberSet = "numberSet"
var BinarySet = "binarySet"
var Map = "map"
var List = "list"
var Null = "null"
var Boolean = "boolean"

type TableAttribute struct {
	PrimaryKey *PrimaryField
	Attributes []*AttributeField
}

func (t *TableAttribute) SetPrimaryKey(field string) *PrimaryField {
	p := &PrimaryField{
		Field: field,
	}
	t.PrimaryKey = p
	return p
}

func (t *TableAttribute) InitAttributes(no_field int) {
	t.Attributes = make([]*AttributeField, 0, no_field)
}

func (t *TableAttribute) SetAttributes(field string) *AttributeField {
	_attr := &AttributeField{
		Field: field,
	}
	t.Attributes = append(t.Attributes, _attr)
	return _attr
}

func NewTableAttribute() *TableAttribute {
	return &TableAttribute{}
}
