# Sort Players
import "sort"

The func sort sorts the data and use Interface as parameter
```
func Sort(data Interface) 
```



#### Create the types
```
type ordenaPorPoints []player
```

#### Create methods for the type
```
func (x ordenaPorPoints) Len() int {return len(x)}
func (x ordenaPorPoints) Less(i, j int) bool {return x[i].Points > x[j].Points}
func (x ordenaPorPoints) Swap(i, j int) {x[i], x[j] = x[j], x[i]}
```

These methods make the types implement the Interface interface.

And the Interface interface {} has the sort function

```
type Interface interface {
    // Len is the number of elements in the collection.
    Len() int
    // Less reports whether the element with
    // index i should sort before the element with index j.
    Less(i, j int) bool
    // Swap swaps the elements with indexes i and j.
    Swap(i, j int)
}
```

#### Sort
```
sort.Sort(ordenaPorPoints(players))
	fmt.Println(players)
```

##### References
[GO pkg](https://golang.org/pkg/)
