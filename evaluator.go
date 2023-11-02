package main

import "fmt"

type EvalType struct {
	evaluates string
	evaluate  func(tree TreeNode) any
}

func evalTree(node TreeNode, customVariables map[string]any) any {
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
			left, _ := evalTree(*tree.left, customVariables).(float64)
			right, _ := evalTree(*tree.right, customVariables).(float64)
			return left + right
		},
	}
	subtractionType := EvalType{
		evaluates: "sub",
		evaluate: func(tree TreeNode) any {
			left, _ := evalTree(*tree.left, customVariables).(float64)
			right, _ := evalTree(*tree.right, customVariables).(float64)
			return left - right
		},
	}
	multiplicationType := EvalType{
		evaluates: "mul",
		evaluate: func(tree TreeNode) any {
			left, _ := evalTree(*tree.left, customVariables).(float64)
			right, _ := evalTree(*tree.right, customVariables).(float64)
			return left * right
		},
	}
	divisionType := EvalType{
		evaluates: "div",
		evaluate: func(tree TreeNode) any {
			left, _ := evalTree(*tree.left, customVariables).(float64)
			right, _ := evalTree(*tree.right, customVariables).(float64)
			return left / right
		},
	}
	constantType := EvalType{
		evaluates: "constant",
		evaluate: func(tree TreeNode) any {
			finalConstants := map[string]any{
				"pi": 3.14159265,
			}

			constName, _ := tree.value.(string)
			if _, ok := finalConstants[constName]; ok {
				return finalConstants[constName]
			}
			if _, ok := customVariables[constName]; ok {
				return customVariables[constName]
			}

			return 0
		},
	}

	evalTypes := []EvalType{
		numberType, constantType,
		additionType, subtractionType, multiplicationType, divisionType,
	}

	for _, evalType := range evalTypes {
		if evalType.evaluates == node.nodeType {
			return evalType.evaluate(node)
		}
	}

	panic(fmt.Sprintf("unsupported tree node type %s\n", node.nodeType))
}
