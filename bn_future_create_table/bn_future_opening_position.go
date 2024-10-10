package bnfuturecreatetable

import tablemodel "xdynamodbgolang/table_model"

func NewBnFutureOpeningPosition() *tablemodel.Table {
	_attriTable := tablemodel.NewTableAttribute()
	_attriTable.SetPrimaryKey("id").SetFieldTypeString()
	_attriTable.SetAttributes("client_id").SetFieldTypeString()
	_attriTable.SetAttributes("exchange_id").SetFieldTypeNumber()
	_attriTable.SetAttributes("symbol").SetFieldTypeString()
	_attriTable.SetAttributes("type").SetFieldTypeString()
	_attriTable.SetAttributes("leverage").SetFieldTypeNumber()
	_attriTable.SetAttributes("amount").SetFieldTypeString()
	_attriTable.SetAttributes("amount_currency").SetFieldTypeString()
	_attriTable.SetAttributes("is_buy_filled").SetFieldTypeBoolean()
	_attriTable.SetAttributes("is_sell_filled").SetFieldTypeBoolean()
	_attriTable.SetAttributes("is_active").SetFieldTypeBoolean()
	_attriTable.SetAttributes("buy_order_created_at").SetFieldTypeString()
	_attriTable.SetAttributes("sell_order_created_at").SetFieldTypeString()
	_attriTable.SetAttributes("buy_updated_at").SetFieldTypeString()
	_attriTable.SetAttributes("sell_updated_at").SetFieldTypeString()
	return &tablemodel.Table{
		Table:     "bn_future_opening_position",
		Attribute: _attriTable,
	}
}
