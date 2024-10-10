package bnfuturecreatetable

import tablemodel "xdynamodbgolang/table_model"

func NewBnFutureQouteUSDT() *tablemodel.Table {
	_attrTable := tablemodel.NewTableAttribute()
	_attrTable.SetPrimaryKey("id").SetFieldTypeString()
	_attrTable.SetAttributes("exchange_id").SetFieldTypeNumber()
	_attrTable.SetAttributes("symbol").SetFieldTypeString()
	_attrTable.SetAttributes("counting_symbol").SetFieldTypeNumber()
	return &tablemodel.Table{
		Table:     "bn_future_qoute_usdt",
		Attribute: _attrTable,
	}
}
