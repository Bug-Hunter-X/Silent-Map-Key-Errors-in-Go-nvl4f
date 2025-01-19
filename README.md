# Silent Map Key Errors in Go

This example demonstrates a potential pitfall in Go when working with maps: accessing a non-existent key silently returns the zero value for the key's type.  This can be problematic because it doesn't result in an error message, making it harder to debug.

The `bug.go` file shows how accessing a missing key in a map returns 0 without any indication of an error. This could lead to incorrect calculations or unintended behavior, especially in larger codebases.

The `bugSolution.go` file shows an improved approach that uses the comma ok idiom to check for the existence of a key before accessing its value. 

This approach makes the code more robust and prevents subtle bugs caused by silently returning zero values.
