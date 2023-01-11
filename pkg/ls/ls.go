package ls

// type file struct {
// 	name string
// }

func  IsHiddenFile(name string) bool {
	return name[0] == '.'
}
