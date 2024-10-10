package bnfuturecreatetable

import (
	tablemodel "xdynamodbgolang/table_model"
)

func NewExchangeTable() *tablemodel.Table {
	_attrTable := tablemodel.NewTableAttribute()
	_attrTable.SetPrimaryKey("id").SetFieldTypeNumber()
	_attrTable.SetAttributes("exchange_name").SetFieldTypeString()
	return &tablemodel.Table{
		Table:     "exchange",
		Attribute: _attrTable,
	}
}
