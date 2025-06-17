# tests

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

  --- PASS: TestCMapSmallConcurrentOps/test_insert (1.33s)
  --- PASS: TestCMapSmallConcurrentOps/test_retrieve (0.52s)
  --- PASS: TestCMapSmallConcurrentOps/test_delete (1.19s)

10,000,000 kv pairs with 64 byte keys and 64 byte values

  --- PASS: TestCMapLargeConcurrentOps/test_insert (15.99s)
  --- PASS: TestCMapLargeConcurrentOps/test_retrieve (7.87s)
  --- PASS: TestCMapLargeConcurrentOps/test_delete (14.97s)

1,000,000 kv pairs with 64 byte keys and 64 byte values

  --- PASS: TestParallelReadWrites/test_init_key_val_pairs_in_map (0.77s)
  --- PASS: TestParallelReadWrites/test_write_new_key_vals_in_map (1.49s)

3 workers, 32 shards:

1,000,000 kv pairs with 64 byte keys and 64 byte values

  --- PASS: TestCMapSmallConcurrentOps/test_insert (0.76s)
  --- PASS: TestCMapSmallConcurrentOps/test_retrieve (0.31s)
  --- PASS: TestCMapSmallConcurrentOps/test_delete (0.66s)

10,000,000 kv pairs with 64 byte keys and 64 byte values

  --- PASS: TestCMapLargeConcurrentOps/test_insert (10.73s)
  --- PASS: TestCMapLargeConcurrentOps/test_retrieve (3.49s)
  --- PASS: TestCMapLargeConcurrentOps/test_delete (10.34s)

1,000,000 kv pairs with 64 byte keys and 64 byte values

  --- PASS: TestParallelReadWrites/test_init_key_val_pairs_in_map (0.37s)
  --- PASS: TestParallelReadWrites/test_write_new_key_vals_in_map (0.65s)

1 worker, 1024 shards:

1,000,000 kv pairs with 64 byte keys and 64 byte values

  --- PASS: TestCMapSmallConcurrentOps/test_insert (1.12s)
  --- PASS: TestCMapSmallConcurrentOps/test_retrieve (0.48s)
  --- PASS: TestCMapSmallConcurrentOps/test_delete (1.07s)

10,000,000 kv pairs with 64 byte keys and 64 byte values

  --- PASS: TestCMapLargeConcurrentOps/test_insert (15.12s)
  --- PASS: TestCMapLargeConcurrentOps/test_retrieve (8.11s)
  --- PASS: TestCMapLargeConcurrentOps/test_delete (12.92s)

1,000,000 kv pairs with 64 byte keys and 64 byte values

  --- PASS: TestParallelReadWrites/test_init_key_val_pairs_in_map (0.77s)
  --- PASS: TestParallelReadWrites/test_write_new_key_vals_in_map (1.57s)

3 workers, 1024 shards:

1,000,000 kv pairs with 64 byte keys and 64 byte values

  --- PASS: TestCMapSmallConcurrentOps/test_insert (0.65s)
  --- PASS: TestCMapSmallConcurrentOps/test_retrieve (0.24s)
  --- PASS: TestCMapSmallConcurrentOps/test_delete (0.51s)

10,000,000 kv pairs with 64 byte keys and 64 byte values

  --- PASS: TestCMapLargeConcurrentOps/test_insert (7.91s)
  --- PASS: TestCMapLargeConcurrentOps/test_retrieve (3.32s)
  --- PASS: TestCMapLargeConcurrentOps/test_delete (7.65s)

1,000,000 kv pairs with 64 byte keys and 64 byte values

  --- PASS: TestParallelReadWrites/test_init_key_val_pairs_in_map (0.30s)
  --- PASS: TestParallelReadWrites/test_write_new_key_vals_in_map (0.50s)

3 workers, 4096 shards:

1,000,000 kv pairs with 64 byte keys and 64 byte values
    --- PASS: TestCMapSmallConcurrentOps/test_insert (0.52s)
    --- PASS: TestCMapSmallConcurrentOps/test_retrieve (0.23s)
    --- PASS: TestCMapSmallConcurrentOps/test_delete (0.46s)

10,000,000 kv pairs with 64 byte keys and 64 byte values
    --- PASS: TestCMapLargeConcurrentOps/test_insert (7.01s)
    --- PASS: TestCMapLargeConcurrentOps/test_retrieve (3.64s)
    --- PASS: TestCMapLargeConcurrentOps/test_delete (7.41s)

20,000,000 kv pairs with 64 byte keys and 64 byte values
    --- PASS: TestCMapLargeConcurrentOps/test_insert (15.90s)
    --- PASS: TestCMapLargeConcurrentOps/test_retrieve (8.31s)
    --- PASS: TestCMapLargeConcurrentOps/test_delete (17.03s)

1,000,000 kv pairs with 64 byte keys and 64 byte values
    --- PASS: TestParallelReadWrites/test_init_key_val_pairs_in_map (0.68s)
    --- PASS: TestParallelReadWrites/test_write_new_key_vals_in_map (1.14s)
```

comparing to `sync.Map`:

```
3 workers:

1,000,000 kv pairs with 64 byte keys and 64 byte values

  --- PASS: TestSyncMapSmallConcurrentOps/test_insert (0.64s)
  --- PASS: TestSyncMapSmallConcurrentOps/test_retrieve (0.28s)
  --- PASS: TestSyncMapSmallConcurrentOps/test_delete (0.21s)

10,000,000 kv pairs with 64 byte keys and 64 byte values

  --- PASS: TestSyncMapLargeConcurrentOps/test_insert (8.30s)
  --- PASS: TestSyncMapLargeConcurrentOps/test_retrieve (3.33s)
  --- PASS: TestSyncMapLargeConcurrentOps/test_delete (2.13s)

20,000,000 kv pairs with 64 byte keys and 64 byte values
    --- PASS: TestSyncMapLargeConcurrentOps/test_insert (17.04s)
    --- PASS: TestSyncMapLargeConcurrentOps/test_retrieve (9.32s)
    --- PASS: TestSyncMapLargeConcurrentOps/test_delete (3.58s)

1,000,000 kv pairs with 64 byte keys and 64 byte values

  --- PASS: TestParallelSyncMapReadWrites/test_write_new_key_vals_in_map (1.06s)
  --- PASS: TestParallelSyncMapReadWrites/test_init_key_val_pairs_in_map (1.11s)
```
