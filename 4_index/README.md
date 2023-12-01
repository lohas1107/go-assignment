# SQL Index

使用 Index B 的查詢效能會比 Index A 好。

```sql
SELECT * FROM orders WHERE user_id = ? AND created_at >= ? AND status = ?
    
index A : idx_user_id_status_created_at(user_id, status, created_at)
index B : idx_user_id_created_at_status(user_id, created_at, status)
index C : idx_user_id_created_at(user_id, created_at)
```

因為在建立複合索引 (A,B,C)，相當於創建一個單一索引 A 和兩個複合索引 A+B 與 A+B+C (不包括 B+C, C)，
所以，在這個 WHERE 的查詢條件循序，會使用到 Index B 而不會使用到 Index A。

如果 status 的資料選擇性很低時，Index B 和 Index C 的查詢效能不會有顯著的差異。