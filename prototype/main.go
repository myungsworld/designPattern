package main

func main() {
	file1 := &file{name: "file1"}
	file2 := &file{name: "file2"}
	file3 := &file{name: "file3"}

	folder1 := &folder{name: "folder1", children: []inode{
		file1, file2,
	}}
	folder2 := &folder{name: "folder2", children: []inode{
		folder1, file3,
	}}

	//folder1.print()
	folder2.print(" ")
}
