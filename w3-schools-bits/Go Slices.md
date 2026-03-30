## Go Slices
A slice is similar to an array, but with one big difference — its size can grow and shrink. This makes slices much more flexible and they are used far more often than arrays in Go.
In Go, there are several ways to create a slice:

    Using the []datatype{values} format
    Create a slice from an array
    Using the make() function

## Go Access, Change, Append and Copy Slices

## Access Elements of a Slice

You can access a specific slice element by referring to the index number.

In Go, indexes start at 0. That means that [0] is the first element, [1] is the second element, etc.

## Change Elements of a Slice
You can also change a specific slice element by referring to the index number.
## Append Elements To a Slice
You can append elements to the end of a slice using the append()function:
## Append One Slice To Another Slice
To append all the elements of one slice to another slice, use the append()function:
## Change The Length of a Slice
Unlike arrays, it is possible to change the length of a slice.
## Memory Efficiency
 When using slices, Go loads all the underlying elements into the memory.

If the array is large and you need only a few elements, it is better to copy those elements using the copy() function.

The copy() function creates a new underlying array with only the required elements for the slice. This will reduce the memory used for the program. 
