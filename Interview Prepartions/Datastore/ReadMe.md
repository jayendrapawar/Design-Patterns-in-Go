Implement an in-memory relational data store
Features:
Create table [P0]
Insert row [P0]
Read row [P0]
Update row[P1]
Delete row [p2]


Create Table
Create a store for a specific table. We can assume all fields of type string. Table name and column names should be persisted as metadata either on the table itself or on a separate entity.

createTable(<table name>, [Column list]) 

Insert Row
Generate unique ID for the to-be-created row. Add the row into the table and return the newly-created-row’s ID.

insertRow([value list in the same order of table creation])

Read Row
Taking row id as input return the row from table.

readRow(id)

Update Row
Search for the row to be updated. Perform an in-place update with the columns to be updated with the new values. If a column need not be updated

updateRow(id, [map of column name to updated value])
