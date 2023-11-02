package main

import "fmt"

type EvalType struct {
	evaluates string
	evaluate  func(tree TreeNode) any
}

func evalTree(node TreeNode) any {
	numberType := EvalType{
		evaluates: "number",
		evaluate: func(tree TreeNode) any {
			num, _ := tree.value.(float64)
			return num
		},
	}
	additionType := EvalType{
		evaluates: "add",
		evaluate: func(tree TreeNode) any {
			left, _ := evalTree(*tree.left).(float64)
			right, _ := evalTree(*tree.right).(float64)
			return left + right
		},
	}
	subtractionType := EvalType{
		evaluates: "sub",
		evaluate: func(tree TreeNode) any {
			left, _ := evalTree(*tree.left).(float64)
			right, _ := evalTree(*tree.right).(float64)
			return left - right
		},
	}
	multiplicationType := EvalType{
		evaluates: "mul",
		evaluate: func(tree TreeNode) any {
			left, _ := evalTree(*tree.left).(float64)
			right, _ := evalTree(*tree.right).(float64)
			return left * right
		},
	}
	divisionType := EvalType{
		evaluates: "div",
		evaluate: func(tree TreeNode) any {
			left, _ := evalTree(*tree.left).(float64)
			right, _ := evalTree(*tree.right).(float64)
			return left / right
		},
	}

	evalTypes := []EvalType{
		numberType,
		additionType, subtractionType, multiplicationType, divisionType,
	}

	for _, evalType := range evalTypes {
		if evalType.evaluates == node.nodeType {
			return evalType.evaluate(node)
		}
	}

	panic(fmt.Sprintf("unsupported tree node type %s\n", node.nodeType))
}
