# Murmur32


## Overview

`Murmur32` is a non-cryptographic hash function that generates `32 bit` values for the given input data.


## Steps


### Initialization

`Murmur32` takes two inputs: the input data and a seed value. The seed value will add randomness to the hash function, where the hash value is initiliazed with the seed value.


### Processing 4-byte Chunks

The input data, a string, is processed in `4-byte` (`32 bit`) chunks, where each chunk is processed individually. The data is interpreted as a sequence of [Little-Endian](https://en.wikipedia.org/wiki/Endianness) `32 bit` unsigned integers, so for each Least Significant Byte comes first in memory. 


### Rotation and Mixing

For each 4-byte chunk, a series of rotations, mixings, and XOR operations are applied.

```
1.) Multiply the chunk by constant1
2.) Rotate the chunk by 15 bits using bitwise OR ('|') and bitwise shift ('<<' and '>>')
3.) Multiply the rotated chunk by constant2
4.) XOR the hash value with the mutated chunk
5.) Rotate the hash value left by 13 bits in the same fashion as 2.)
6.) Multiply the hash by 5 and add constant3
```

here is the rotateRight function in go:
```go
func rotateRight(hash *uint32, chunk uint32) {
	chunk *= constant1
	chunk = (chunk << 15) | (chunk >> 17) // Rotate right by 15
	chunk *= constant2

	*hash ^= chunk
	*hash = (*hash << 13) | (*hash >> 19) // Rotate right by 13
	*hash = *hash * 5 + constant3
}
```


### Handling Remaining Bytes

There may be remaining bytes after processing the chunks.

The following steps may be applied depending on the number of remaining bytes (1, 2, 3).

```
1.) shift the bytes into a 32 bit chunk variable using | and << operations
2.) multiply by constant1
3.) rotate the chunk right by 15 bits using | and << and >>
4.) Multiply the chunk by constant2
5.) XOR the hash value with the mutated chunk
```

here is the handleRemaining function in go:
```go
func handleRemainingBytes(hash *uint32, dataAsBytes []byte) {
	remaining := dataAsBytes[len(dataAsBytes)-len(dataAsBytes) % 4:]
	
	if len(remaining) > 0 {
		var chunk uint32
		
		switch len(remaining) {
			case 3:
				chunk |= uint32(remaining[2]) << 16
				fallthrough
			case 2:
				chunk |= uint32(remaining[1]) << 8
				fallthrough
			case 1:
				chunk |= uint32(remaining[0])
				chunk *= constant1
				chunk = (chunk << 15) | (chunk >> 17) // Rotate right by 15
				chunk *= constant2
				*hash ^= chunk
			}
	}
}
```

the `fallthrough` keyword in the switch allows the case block, if in higher order than the other blocks, to perform the operations on the following case blocks.


### Finalization

After processing the chunks, some final mixing is applied to the hash.

The following steps are applied:

```
1.) XOR the hash with the length of the input data
2.) XOR the hash by the right shifted version of itself by 16 bits
3.) Multiply the hash value by constant4
4.) XOR the hash by the right shifted version of itself by 13 bits
5.) Multiply the hash by constant5
6.) XOR the hash by the right shifted version of itself by 16 bits
```

in go:
```go
hash ^= length
hash ^= hash >> 16
hash *= constant4
hash ^= hash >> 13
hash *= constant5
hash ^= hash >> 16
```


## What are the Constants?

```go
// a prime number that serves as a multiplier during mixing. Distributes bits and improves randomness
const constant1 = 0x85ebca6b 

// a prime number also used for mixing. Enhances distribution of hash value
const constant2 = 0xc2b2ae35

// added to hash after each chunk is mixed in. Contributes to finalization step
const constant3 = 0xe6546b64

// multiplied in the finalization step. Provides additional mixing effect
const constant4 = 0x1b873593

// multiplier in the finalization step. Again, improves hash value distribution
const constant5 = 0x5c4bcea9
```

# Murmur64

`Murmur64` has also been implemented to extend the 32 bit hash to a 64 bit hash. The algorithm is the same, with only the constants changing and the byte chunk size changing (from `4 Bytes` to `8 Bytes`)

## Sources

[Murmur](../pkg/map/Murmur.go)