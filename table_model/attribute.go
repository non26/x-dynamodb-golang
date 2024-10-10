package tablemodel

type AttributeField struct {
	Field     string
	fieldType string
}

func (t *AttributeField) SetFieldTypeString() {
	t.fieldType = String
}

func (t *AttributeField) SetFieldTypeNumber() {
	t.fieldType = Number
}

func (t *AttributeField) SetFieldTypeBinary() {
	t.fieldType = Binary
}

func (t *AttributeField) SetFieldTypeStringSet() {
	t.fieldType = StringSet
}

func (t *AttributeField) SetFieldTypeNumberSet() {
	t.fieldType = NumberSet
}

func (t *AttributeField) SetFieldTypeBinarySet() {
	t.fieldType = BinarySet
}

func (t *AttributeField) SetFieldTypeMap() {
	t.fieldType = Map
}

func (t *AttributeField) SetFieldTypeList() {
	t.fieldType = List
}

func (t *AttributeField) SetFieldTypeNull() {
	t.fieldType = Null
}

func (t *AttributeField) SetFieldTypeBoolean() {
	t.fieldType = Boolean
}

func (t *AttributeField) GetFieldType() string {
	return t.fieldType
}
