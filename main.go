package main

import "fmt"

// ut "github.com/OluOwolabi/go-start/src/utils"

func main() {
	// slog.Info("Starting the server...", "version", "1.0.0", "ulid", ut.GenerateUlid())
	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprintf(w, "Hello, World!")
	// })

	// err := http.ListenAndServe(":8080", nil)
	// if err != nil {
	// 	slog.Error("Server failed to start", "error", err)
	// 	log.Fatal(err)
	// }

	root := &TreeNode{data: 1}
	root.left = &TreeNode{data: 2}
	root.right = &TreeNode{data: 3}
	root.left.left = &TreeNode{data: 4}
	root.left.right = &TreeNode{data: 5}
	root.right.left = &TreeNode{data: 6}
	root.right.right = &TreeNode{data: 7}
	root.left.right.left = &TreeNode{data: 8}
	root.left.right.right = &TreeNode{data: 9}
	root.right.right.right = &TreeNode{data: 10}
	root.left.right.left.right = &TreeNode{data: 11}
	root.left.right.right.left = &TreeNode{data: 12}
	// fmt.Println(isValidBST(root))
	// printDfs(root)

	test([]int{1, 2, 3, 4, 5, 6})
}

type TreeNode struct {
	data  int
	left  *TreeNode
	right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	if root == nil {
		return false
	}
	valid := true
	queue := []*TreeNode{root}

	for len(queue) > 0 {

		level := len(queue)
		currentLevelRoot := queue[0].data
		leftVisited, rightVisited := []int{}, []int{}
		for i := 0; i < level; i++ {
			// get the first in the list
			current := queue[0]
			// get all the left and right children
			if current.left != nil {
				leftVisited = append(leftVisited, current.left.data)
				queue = append(queue, current.left)
			}

			if current.right != nil {
				rightVisited = append(rightVisited, current.right.data)
				queue = append(queue, current.right)
			}
			queue = queue[1:]
		}
		for _, i := range leftVisited {
			if i > currentLevelRoot {
				fmt.Printf("the left node %d is greater that node root %d\n", i, currentLevelRoot)
				valid = false
				break
			}
		}
		for _, i := range rightVisited {
			if i < currentLevelRoot {
				fmt.Printf("the right node %d is less that node root %d\n", i, currentLevelRoot)
				valid = false
				break
			}

		}
	}

	return valid
}

func printDfs(root *TreeNode) {

	if root == nil {
		return
	}
	fmt.Print(root.data, ",")
	if root.left != nil {
		printDfs(root.left)
	}
	if root.right != nil {
		printDfs(root.right)
	}
}

func test(arr []int) {
	if len(arr)%2 != 0 {
		fmt.Println("can't do this on uneven len of arraay")
		return
	}
	for j := 2; j < len(arr); j = j + 2 {
		fmt.Printf("elements by two in the array = %v", arr[:j])
	}
}
