package bnfuturecreatetable

import tablemodel "xdynamodbgolang/table_model"

func NewBnFutureHistory() *tablemodel.Table {
	_attriTable := tablemodel.NewTableAttribute()
	_attriTable.SetPrimaryKey("id").SetFieldTypeString()
	_attriTable.SetAttributes("exchange_id").SetFieldTypeNumber()
	_attriTable.SetAttributes("client_id").SetFieldTypeString()
	_attriTable.SetAttributes("pnl").SetFieldTypeString()
	_attriTable.SetAttributes("created_at").SetFieldTypeString()
	_attriTable.SetAttributes("buy_order_created_at").SetFieldTypeString()
	_attriTable.SetAttributes("sell_order_created_at").SetFieldTypeString()
	_attriTable.SetAttributes("buy_updated_at").SetFieldTypeString()
	_attriTable.SetAttributes("sell_updated_at").SetFieldTypeString()
	return &tablemodel.Table{
		Table:     "bn_future_history",
		Attribute: _attriTable,
	}
}
