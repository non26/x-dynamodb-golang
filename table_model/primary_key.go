package tablemodel

type PrimaryField struct {
	Field     string
	FieldType string
}

func (t *PrimaryField) SetFieldTypeString() {
	t.FieldType = String
}

func (t *PrimaryField) SetFieldTypeNumber() {
	t.FieldType = Number
}

func (t *PrimaryField) SetFieldTypeBinary() {
	t.FieldType = Binary
}

func (t *PrimaryField) GetFieldType() string {
	return t.FieldType
}
