# Tests

tested on:
```
Macbook Pro 14inch
M2Pro
16GB ram
512GB ssd
```

results (with sharded map):
```
1 worker, 32 shards:

1,000,000 kv pairs with 64 byte keys and 64 byte values

--- PASS: TestCMapSmallConcurrentOps/test_insert (1.35s)
--- PASS: TestCMapSmallConcurrentOps/test_retrieve (0.51s)
--- PASS: TestCMapSmallConcurrentOps/test_delete (1.20s)


10,000,000 kv pairs with 64 byte keys and 64 byte values

--- PASS: TestCMapLargeConcurrentOps/test_insert (16.97s)
--- PASS: TestCMapLargeConcurrentOps/test_retrieve (8.45s)
--- PASS: TestCMapLargeConcurrentOps/test_delete (15.87s)

1,000,000 kv pairs with 64 byte keys and 64 byte values

--- PASS: TestParallelReadWrites/test_init_key_val_pairs_in_map (0.76s)
--- PASS: TestParallelReadWrites/test_write_new_key_vals_in_map (1.56s)

3 workers, 32 shards:

1,000,000 kv pairs with 64 byte keys and 64 byte values

--- PASS: TestCMapSmallConcurrentOps/test_insert (0.76s)
--- PASS: TestCMapSmallConcurrentOps/test_retrieve (0.24s)
--- PASS: TestCMapSmallConcurrentOps/test_delete (0.76s)


10,000,000 kv pairs with 64 byte keys and 64 byte values

--- PASS: TestCMapLargeConcurrentOps/test_insert (9.69s)
--- PASS: TestCMapLargeConcurrentOps/test_retrieve (3.68s)
--- PASS: TestCMapLargeConcurrentOps/test_delete (9.67s)

1,000,000 kv pairs with 64 byte keys and 64 byte values

--- PASS: TestParallelReadWrites/test_init_key_val_pairs_in_map (0.35s)
--- PASS: TestParallelReadWrites/test_write_new_key_vals_in_map (0.80s)

1 worker, 1024 shards:

1,000,000 kv pairs with 64 byte keys and 64 byte values

--- PASS: TestCMapSmallConcurrentOps/test_insert (1.13s)
--- PASS: TestCMapSmallConcurrentOps/test_retrieve (0.45s)
--- PASS: TestCMapSmallConcurrentOps/test_delete (1.04s)


10,000,000 kv pairs with 64 byte keys and 64 byte values

--- PASS: TestCMapLargeConcurrentOps/test_insert (15.20s)
--- PASS: TestCMapLargeConcurrentOps/test_retrieve (7.70s)
--- PASS: TestCMapLargeConcurrentOps/test_delete (13.28s)

1,000,000 kv pairs with 64 byte keys and 64 byte values

--- PASS: TestParallelReadWrites/test_init_key_val_pairs_in_map (0.64s)
--- PASS: TestParallelReadWrites/test_write_new_key_vals_in_map (1.15s)

3 workers, 1024 shards:

1,000,000 kv pairs with 64 byte keys and 64 byte values

--- PASS: TestCMapSmallConcurrentOps/test_insert (0.65s)
--- PASS: TestCMapSmallConcurrentOps/test_retrieve (0.23s)
--- PASS: TestCMapSmallConcurrentOps/test_delete (0.50s)


10,000,000 kv pairs with 64 byte keys and 64 byte values

--- PASS: TestCMapLargeConcurrentOps/test_insert (8.34s)
--- PASS: TestCMapLargeConcurrentOps/test_retrieve (3.19s)
--- PASS: TestCMapLargeConcurrentOps/test_delete (7.34s)

1,000,000 kv pairs with 64 byte keys and 64 byte values

--- PASS: TestParallelReadWrites/test_init_key_val_pairs_in_map (0.29s)
--- PASS: TestParallelReadWrites/test_write_new_key_vals_in_map (0.48s)
```

comparing to `sync.Map`:

```
3 workers:

1,000,000 kv pairs with 64 byte keys and 64 byte values

--- PASS: TestSyncMapSmallConcurrentOps/test_insert (0.63s)
--- PASS: TestSyncMapSmallConcurrentOps/test_retrieve (0.29s)
--- PASS: TestSyncMapSmallConcurrentOps/test_delete (0.16s)

10,000,000 kv pairs with 64 byte keys and 64 byte values

--- PASS: TestSyncMapLargeConcurrentOps/test_insert (8.40s)
--- PASS: TestSyncMapLargeConcurrentOps/test_retrieve (3.30s)
--- PASS: TestSyncMapLargeConcurrentOps/test_delete (2.22s)

1,000,000 kv pairs with 64 byte keys and 64 byte values

--- PASS: TestParallelSyncMapReadWrites/test_write_new_key_vals_in_map (1.04s)
--- PASS: TestParallelSyncMapReadWrites/test_init_key_val_pairs_in_map (1.09s)
```
