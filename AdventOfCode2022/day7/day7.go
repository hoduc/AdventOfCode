// --- Day 7: No Space Left On Device ---
// You can hear birds chirping and raindrops hitting leaves as the expedition proceeds. Occasionally, you can even hear much louder sounds in the distance; how big do the animals get out here, anyway?

// The device the Elves gave you has problems with more than just its communication system. You try to run a system update:

// $ system-update --please --pretty-please-with-sugar-on-top
// Error: No space left on device
// Perhaps you can delete some files to make space for the update?

// You browse around the filesystem to assess the situation and save the resulting terminal output (your puzzle input). For example:

// $ cd /
// $ ls
// dir a
// 14848514 b.txt
// 8504156 c.dat
// dir d
// $ cd a
// $ ls
// dir e
// 29116 f
// 2557 g
// 62596 h.lst
// $ cd e
// $ ls
// 584 i
// $ cd ..
// $ cd ..
// $ cd d
// $ ls
// 4060174 j
// 8033020 d.log
// 5626152 d.ext
// 7214296 k
// The filesystem consists of a tree of files (plain data) and directories (which can contain other directories or files). The outermost directory is called /. You can navigate around the filesystem, moving into or out of directories and listing the contents of the directory you're currently in.

// Within the terminal output, lines that begin with $ are commands you executed, very much like some modern computers:

// cd means change directory. This changes which directory is the current directory, but the specific result depends on the argument:
// cd x moves in one level: it looks in the current directory for the directory named x and makes it the current directory.
// cd .. moves out one level: it finds the directory that contains the current directory, then makes that directory the current directory.
// cd / switches the current directory to the outermost directory, /.
// ls means list. It prints out all of the files and directories immediately contained by the current directory:
// 123 abc means that the current directory contains a file named abc with size 123.
// dir xyz means that the current directory contains a directory named xyz.
// Given the commands and output in the example above, you can determine that the filesystem looks visually like this:

// - / (dir)
//   - a (dir)
//     - e (dir)
//       - i (file, size=584)
//     - f (file, size=29116)
//     - g (file, size=2557)
//     - h.lst (file, size=62596)
//   - b.txt (file, size=14848514)
//   - c.dat (file, size=8504156)
//   - d (dir)
//     - j (file, size=4060174)
//     - d.log (file, size=8033020)
//     - d.ext (file, size=5626152)
//     - k (file, size=7214296)
// Here, there are four directories: / (the outermost directory), a and d (which are in /), and e (which is in a). These directories also contain files of various sizes.

// Since the disk is full, your first step should probably be to find directories that are good candidates for deletion. To do this, you need to determine the total size of each directory. The total size of a directory is the sum of the sizes of the files it contains, directly or indirectly. (Directories themselves do not count as having any intrinsic size.)

// The total sizes of the directories above can be found as follows:

// The total size of directory e is 584 because it contains a single file i of size 584 and no other directories.
// The directory a has total size 94853 because it contains files f (size 29116), g (size 2557), and h.lst (size 62596), plus file i indirectly (a contains e which contains i).
// Directory d has total size 24933642.
// As the outermost directory, / contains every file. Its total size is 48381165, the sum of the size of every file.
// To begin, find all of the directories with a total size of at most 100000, then calculate the sum of their total sizes. In the example above, these directories are a and e; the sum of their total sizes is 95437 (94853 + 584). (As in this example, this process can count files more than once!)

// Find all of the directories with a total size of at most 100000. What is the sum of the total sizes of those directories?

package main

import(
    "fmt"
    _ "embed"
    "strconv"
    "strings"
    "errors"
    "github.com/hoduc/AdventOfCode/AdventOfCode2022/util"
)

type COMMAND = int
const (
    CD COMMAND = iota
    LS
    UNKNOWN
)

type Node struct {
    id string
    sz int
    parent *Node
    children map[string]*Node
}

func isDir(current *Node) bool{
    return len(current.children) > 0
}

func dfs(current *Node, count *int, cap int) {
    if current != nil {
        // fmt.Println(current.id, current.sz, &current, &current.parent)
        if isDir(current) {
            // fmt.Println(current.id, current.sz)
            if current.sz <= cap {
                *count += current.sz
            }
        }
        for _, child := range current.children {
            dfs(child, count, cap)
        }
    }
}

func cd(root, current *Node, dest string) (*Node, error) {
    if current == nil {
        return nil, errors.New(fmt.Sprintf("Cannot operate on nil\n"))
    }
    if dest == ".." {
        // fmt.Println("hello:", current.parent)
        return current.parent, nil
    }
    if dest == "/" {
        return root, nil
    }
    // if current.id == dest {
    //     return current, nil
    // }
    next, ok := current.children[dest]
    if !ok {
        return nil, errors.New(fmt.Sprintf("[%v] does not exist in current\n", dest))
    }
    return next, nil
}

func ls(current, node *Node) {
    // fmt.Println("current:", current.id)
    if _, ok := current.children[node.id]; !ok {
        node.parent = current
        current.children[node.id] = node
        // go all the way back to root to update the sz
        pNode := node
        for pNode.parent != nil {
            pNode.parent.sz += node.sz
            // fmt.Println(pNode.parent.id, pNode.parent.sz)
            pNode = pNode.parent
        }
    }
}

func parseCommand(cmd string) COMMAND {
    if cmd == "cd" {
        return CD
    } else if cmd == "ls" {
        return LS
    }
    return UNKNOWN

}

//go:embed day7.txt
var day7txt string

func current_children(current *Node) string {
    children := ""
    for _, node := range current.children {
        children += node.id + ";"
    }
    return children
}

func browse() int {
    lineNo := 1
    command := UNKNOWN
    root := &Node{id: "/", sz: 0, parent: nil, children: make(map[string]*Node)}
    proot := root
    onLine := func(line string) error {
        lineNo += 1
        if len(line) > 0 {
            // fmt.Println(line)
            commands := strings.Split(line, " ")
            fmt.Println(lineNo, "commands:", commands)
            if commands[0] == "$" {
                command = parseCommand(commands[1])
                if command == UNKNOWN {
                    return errors.New(fmt.Sprintf("Unknown command [%v]\n", commands[1]))
                }
            }

            if command == CD {
                // fmt.Println("before-cd:", proot.id, "{", current_children(proot), "}")
                nr, err := cd(root, proot, commands[2])
                if err != nil {
                    return err
                }
                proot = nr
                // fmt.Println("after-cd:", proot.id, "{", current_children(proot), "}")
            } else if command == LS && commands[0] != "$" {
                _sz, err := strconv.Atoi(commands[0])
                if err != nil {
                    _sz = 0
                }
                ls(proot, &Node{id: commands[1], sz: _sz, parent: nil, children: make(map[string]*Node)})
                // fmt.Println("after:", proot.id, proot.sz, &proot)
            }
            // fmt.Println("current:", proot.id)
        }
        return nil
    }

    if err := util.ReadLinesEmbed(day7txt, onLine); err != nil {
        return -1
    }
    // dfs from root
    fmt.Println("---Finished-reading!!!")
    fmt.Println(root)
    fmt.Println("--------")
    count := 0
    dfs(root, &count, 100000)
    return count
}

func main() {
    fmt.Println("part1:", browse())
}
