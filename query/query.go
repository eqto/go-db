package query

func TableOf(stmt interface{}) []Table {
	switch stmt := stmt.(type) {
	case *SelectStmt:
		return stmt.tables
	case *InsertStmt:
		if stmt.table != nil {
			return []Table{*stmt.table}
		}
	}
	return nil
}

func WhereOf(stmt interface{}) []string {
	if stmt, ok := stmt.(*SelectStmt); ok {
		return stmt.wheres
	}
	return nil
}

func OrderByOf(stmt interface{}) []string {
	if stmt, ok := stmt.(*SelectStmt); ok {
		return stmt.orders
	}
	return nil
}

func LimitOf(stmt interface{}) (int, int) {
	if stmt, ok := stmt.(*SelectStmt); ok {
		return stmt.offset, stmt.count
	}
	return 0, 0
}

func ValueOf(stmt interface{}) []string {
	if stmt, ok := stmt.(*InsertStmt); ok {
		return stmt.values
	}
	return nil
}