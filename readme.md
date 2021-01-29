# FUNGO

## Why did I create this module

This module was created because I was bored :grin: and decided to replace all basic operations with functions.  

## What has been implemented so far

### Casting interface value into uint, int, float64, bool (:smile:) and string types

Functions that convert the value of an interface{} to a specific type begin with `Cast`, for example `CastFloat()`.  

You can also find out if it is possible to cast the interface{} to this type 
using functions that start with `Castable`, for example `CastableFloat()` 
while the dimension is not important for int and uint, there is support for 
8, 16, 32, 64 bit values and 32 - 64 bit values for float. Will always return a value of the base type.  

### Comparison for base types

All variables are converted to the desired type using the cast package. Сomparison always occurs relative to the first parameter passed to the function.  

$TYPE can be Float, Uint, Int, it can also be String and bool but not all functionality is implemented.  

Functions patterns:

`Less$TYPE(v1, v2 interface{})`  
`LessEqual$TYPE(v1, v2 interface{})`  
`More$TYPE(v1 ,v2 interface{})`  
`MoreEqual$TYPE(v1, v2 interface{})`  
`Equal$TYPE(v1, v2 interface{})`  
`NotEqual$TYPE(v1, v2 interface{})`  

There are also type-independent functions `IsNull(v interface{})` and `NotNull(v interface{})`
