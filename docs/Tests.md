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
1 worker, 16 shards:

1,000,000 kv pairs with 64 byte keys and 64 byte values

PASS: TestCMapSmallConcurrentOps/test_insert (1.41s)
PASS: TestCMapSmallConcurrentOps/test_retrieve (0.72s)
PASS: TestCMapSmallConcurrentOps/test_delete (1.22s)

10,000,000 kv pairs with 64 byte keys and 64 byte values

PASS: TestCMapLargeConcurrentOps/test_insert (16.84s)
PASS: TestCMapLargeConcurrentOps/test_retrieve (10.39s)
PASS: TestCMapLargeConcurrentOps/test_delete (16.92s)

1,000,000 kv pairs with 64 byte keys and 64 byte values

PASS: TestParallelReadWrites/test_init_key_val_pairs_in_map (1.11s)
PASS: TestParallelReadWrites/test_write_new_key_vals_in_map (1.57s)

2 workers, 16 shards:

1,000,000 kv pairs with 64 byte keys and 64 byte values

PASS: TestCMapSmallConcurrentOps/test_insert (1.00s)
PASS: TestCMapSmallConcurrentOps/test_retrieve (0.46s)
PASS: TestCMapSmallConcurrentOps/test_delete (0.96s)

10,000,000 kv pairs with 64 byte keys and 64 byte values

PASS: TestCMapLargeConcurrentOps/test_insert (12.67s)
PASS: TestCMapLargeConcurrentOps/test_retrieve (5.97s)
PASS: TestCMapLargeConcurrentOps/test_delete (12.51s)

1,000,000 kv pairs with 64 byte keys and 64 byte values

PASS: TestParallelReadWrites/test_init_key_val_pairs_in_map (0.93s)
PASS: TestParallelReadWrites/test_write_new_key_vals_in_map (1.38s)

3 workers, 16 shards:

1,000,000 kv pairs with 64 byte keys and 64 byte values

PASS: TestCMapSmallConcurrentOps/test_insert (0.85s)
PASS: TestCMapSmallConcurrentOps/test_retrieve (0.34s)
PASS: TestCMapSmallConcurrentOps/test_delete (0.84s)

10,000,000 kv pairs with 64 byte keys and 64 byte values

PASS: TestCMapLargeConcurrentOps/test_insert (11.44s)
PASS: TestCMapLargeConcurrentOps/test_retrieve (4.68s)
PASS: TestCMapLargeConcurrentOps/test_delete (11.30s)

1,000,000 kv pairs with 64 byte keys and 64 byte values

PASS: TestParallelReadWrites/test_init_key_val_pairs_in_map (0.76s)
PASS: TestParallelReadWrites/test_write_new_key_vals_in_map (1.05s)

1 worker, 32 shards:

1,000,000 kv pairs with 64 byte keys and 64 byte values

PASS: TestCMapSmallConcurrentOps/test_insert (1.29s)
PASS: TestCMapSmallConcurrentOps/test_retrieve (0.67s)
PASS: TestCMapSmallConcurrentOps/test_delete (1.19s)

10,000,000 kv pairs with 64 byte keys and 64 byte values

PASS: TestCMapLargeConcurrentOps/test_insert (17.11s)
PASS: TestCMapLargeConcurrentOps/test_retrieve (9.66s)
PASS: TestCMapLargeConcurrentOps/test_delete (15.79s)

1,000,000 kv pairs with 64 byte keys and 64 byte values

PASS: TestParallelReadWrites/test_init_key_val_pairs_in_map (0.93s)
PASS: TestParallelReadWrites/test_write_new_key_vals_in_map (1.39s)

2 workers, 32 shards:

1,000,000 kv pairs with 64 byte keys and 64 byte values

PASS: TestCMapSmallConcurrentOps/test_insert (0.92s)
PASS: TestCMapSmallConcurrentOps/test_retrieve (0.44s)
PASS: TestCMapSmallConcurrentOps/test_delete (0.83s)

10,000,000 kv pairs with 64 byte keys and 64 byte values

PASS: TestCMapLargeConcurrentOps/test_insert (12.14s)
PASS: TestCMapLargeConcurrentOps/test_retrieve (6.30s)
PASS: TestCMapLargeConcurrentOps/test_delete (11.68s)

1,000,000 kv pairs with 64 byte keys and 64 byte values

PASS: TestParallelReadWrites/test_init_key_val_pairs_in_map (0.54s)
PASS: TestParallelReadWrites/test_write_new_key_vals_in_map (1.03s)

3 workers, 32 shards:

1,000,000 kv pairs with 64 byte keys and 64 byte values

PASS: TestCMapSmallConcurrentOps/test_insert (0.77s)
PASS: TestCMapSmallConcurrentOps/test_retrieve (0.32s)
PASS: TestCMapSmallConcurrentOps/test_delete (0.69s)

10,000,000 kv pairs with 64 byte keys and 64 byte values

PASS: TestCMapLargeConcurrentOps/test_insert (10.81s)
PASS: TestCMapLargeConcurrentOps/test_retrieve (4.84s)
PASS: TestCMapLargeConcurrentOps/test_delete (9.42s)

1,000,000 kv pairs with 64 byte keys and 64 byte values

PASS: TestParallelReadWrites/test_init_key_val_pairs_in_map (0.76s)
PASS: TestParallelReadWrites/test_write_new_key_vals_in_map (1.24s)
```

comparing to `sync.Map`:

```
3 workers:

1,000,000 kv pairs with 64 byte keys and 64 byte values

PASS: TestSyncMapSmallConcurrentOps/test_insert (0.79s)
PASS: TestSyncMapSmallConcurrentOps/test_retrieve (0.30s)
PASS: TestSyncMapSmallConcurrentOps/test_delete (0.16s)

10,000,000 kv pairs with 64 byte keys and 64 byte values

PASS: TestSyncMapLargeConcurrentOps/test_insert (8.29s)
PASS: TestSyncMapLargeConcurrentOps/test_retrieve (3.35s)
PASS: TestSyncMapLargeConcurrentOps/test_delete (2.18s)

1,000,000 kv pairs with 64 byte keys and 64 byte values

PASS: TestParallelSyncMapReadWrites/test_write_new_key_vals_in_map (1.04s)
PASS: TestParallelSyncMapReadWrites/test_init_key_val_pairs_in_map (1.09s)
```
