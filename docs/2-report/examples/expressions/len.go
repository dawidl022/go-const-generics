package expressions

func expressions[T any, N const](arr [N]T) {
	const _ = len(arr) // does not compile
	      _ = len(arr) // compiles - non-constant int type
	
	const _ = N        // does not compile
	      _ = N        // compiles - non-constant int type
}
