package main

import (
	"fmt"
)

type TreeNode struct {
	nodeType string
	left     *TreeNode
	right    *TreeNode
	value    any
}

func (node TreeNode) Copy() TreeNode {
	newNode := TreeNode{}
	newNode.nodeType = node.nodeType
	newNode.value = node.value

	if node.left != nil {
		leftCopy := node.left.Copy()
		newNode.left = &leftCopy
	}

	if node.right != nil {
		rightCopy := node.right.Copy()
		newNode.right = &rightCopy
	}

	return newNode
}

func (node TreeNode) ToString() string {
	var leftStr string
	if node.left == nil {
		leftStr = "nil"
	} else {
		leftStr = node.left.ToString()
	}

	var rightStr string
	if node.right == nil {
		rightStr = "nil"
	} else {
		rightStr = node.right.ToString()
	}

	valueStr := fmt.Sprintf("%v", node.value)
	if node.value == nil {
		valueStr = "nil"
	}

	return fmt.Sprintf("TreeNode{type: %s, left: %s, right: %s, value: %s}", node.nodeType, leftStr, rightStr, valueStr)
}

func parseTokens(tokens []Token) TreeNode {
	parser := NodeParser{
		index:  0,
		tokens: tokens,
	}
	return parser.parseExpression()
}

type NodeParser struct {
	index  int
	tokens []Token
}

func (parser *NodeParser) parseExpression() TreeNode {
	left := parser.parseTerm()

	for parser.index < len(parser.tokens) {
		if parser.tokens[parser.index].tokenType == AddToken || parser.tokens[parser.index].tokenType == SubToken {
			operator := parser.tokens[parser.index].tokenType.id
			parser.index++
			right := parser.parseTerm()

			left = TreeNode{
				nodeType: operator,
				left:     &left,
				right:    &right,
			}.Copy()
		} else {
			break
		}
	}

	return left
}

func (parser *NodeParser) parseTerm() TreeNode {
	left := parser.parseFactor()

	for parser.index < len(parser.tokens) {
		if parser.tokens[parser.index].tokenType == MulToken || parser.tokens[parser.index].tokenType == DivToken {
			operator := parser.tokens[parser.index].tokenType.id
			parser.index++
			right := parser.parseFactor()

			left = TreeNode{
				nodeType: operator,
				left:     &left,
				right:    &right,
			}.Copy()
		} else {
			break
		}
	}

	return left
}

func (parser *NodeParser) parseFactor() TreeNode {
	if parser.tokens[parser.index].tokenType == NumberToken {
		token := parser.tokens[parser.index]
		parser.index++
		return TreeNode{
			nodeType: "number",
			value:    token.value,
		}.Copy()
	} else if parser.tokens[parser.index].tokenType == ConstantToken {
		token := parser.tokens[parser.index]
		parser.index++
		return TreeNode{
			nodeType: "constant",
			value:    token.value,
		}
	}

	err := fmt.Sprintf("unexpected token %s at index %d", parser.tokens[parser.index].ToString(), parser.index)
	panic(err)
}
