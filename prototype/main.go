package main

func main() {
	file1 := &file{name: "File1"}
	file2 := &file{name: "File2"}
	file3 := &file{name: "File3"}

	folder1 := &folder{
		name:     "Folder1",
		children: []inode{file1},
	}

	folder2 := &folder{
		name:     "Folder2",
		children: []inode{folder1, file1, file2, file3},
	}

	folder2.print(" ")

	clone := folder2.clone()
	clone.print(" ")

	clone2 := clone.clone()
	clone2.print(" ")

}
