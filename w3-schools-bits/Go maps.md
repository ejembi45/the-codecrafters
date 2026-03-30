## Go Maps
A map is a collection of key-value pairs. Instead of accessing items by their position like in a slice, you access them by a unique key. Think of it like a dictionary — you look up a word (key) to find its definition (value).

## More On Maps
A map is an unordered and changeable collection that does not allow duplicates.

The length of a map is the number of its elements. You can find it using the len() function.

The default value of a map is nil.

Maps hold references to an underlying hash table.

Go has multiple ways for creating maps.

## Here are the three main ways to create a map in Go
1. Using a Map Literal
The most common way — create and fill the map at the same time.
2. Using make()
Creates an empty map that you can add to later.
3. Using var
Declares a map but leaves it empty for now.

