# string-cube-puzzle
solves string/snake cube puzzles

One common 4x4x4 example:

```
$ go run scp.go 4 2 4 2 2 2 2 3 2 2 2 2 2 3 2 4 2 3 3 4 2 3 2 2 2 2 2 2 2 2 2 4 2 4 2 4 4 4 3
snake segment lengths: [4 2 4 2 2 2 2 3 2 2 2 2 2 3 2 4 2 3 3 4 2 3 2 2 2 2 2 2 2 2 2 4 2 4 2 4 4 4 3]
&{[true true true true true true true true true true true true true true true true true true true true true true true true true true true true true true true true true true true true true true true true true true true true true true true true true true true true true true true true true true true true true true true true] 4 4 4}
X X X X 
X X X X 
X X X X 
X X X X 

X X X X 
X X X X 
X X X X 
X X X X 

X X X X 
X X X X 
X X X X 
X X X X 

X X X X 
X X X X 
X X X X 
X X X X 

4*+X -> 2*-Y -> 4*-X -> 2*+Z -> 2*+Y -> 2*+X -> 2*-Y -> 3*+X -> 2*+Y -> 2*-X -> 2*+Y -> 2*-X -> 2*-Z -> 3*+X -> 2*+Y -> 4*-X -> 2*-Y -> 3*+Z -> 3*-Y -> 4*+X -> 2*+Y -> 3*-X -> 2*+Y -> 2*+Z -> 2*-Y -> 2*+X -> 2*+Y -> 2*-Z -> 2*+X -> 2*-Z -> 2*+Y -> 4*-X -> 2*+Z -> 4*+X -> 2*+Z -> 4*-X -> 4*-Y -> 4*+X -> 3*+Y -> done
```
