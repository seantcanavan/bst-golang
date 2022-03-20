package main

import "testing"

func TestInsertOnEmptyNode(t *testing.T) {
	value := 5
	tree := Node{}
	result := insert(value, &tree)

	if result.value != value {
		t.Error("Failed to insert into an empty node")
	}
}

func TestInsertOnExistingNodeWithIncreasingValues(t *testing.T) {
	value := 5
	firstInsert := 6
	secondInsert := 7

	tree := Node{value: value}

	firstResult := insert(firstInsert, &tree)
	secondResult := insert(secondInsert, &tree)

	if firstResult == nil || firstResult.value != firstInsert {
		t.Error("Failed to insert into existing tree")
	}

	if secondResult == nil || secondResult.value != secondInsert {
		t.Error("Failed to insert increasing value")
	}

	if tree.value != value {
		t.Error("Unexpected tree structure at root level")
	}

	if tree.right.value != firstInsert {
		t.Error("Unexpected tree structure at second level")
	}

	if tree.right.right.value != secondInsert {
		t.Error("Unexpected tree structure at third level")
	}
}

func TestInsertOnExistingNodeWithDecreasingValues(t *testing.T) {
	value := 5
	firstInsert := 4
	secondInsert := 3

	tree := Node{value: value}

	firstResult := insert(firstInsert, &tree)
	secondResult := insert(secondInsert, &tree)

	if firstResult == nil || firstResult.value != firstInsert {
		t.Error("Failed to insert into existing tree")
	}

	if secondResult == nil || secondResult.value != secondInsert {
		t.Error("Failed to insert decreasing value")
	}

	if tree.value != value {
		t.Error("Unexpected tree structure at root level")
	}

	if tree.left.value != firstInsert {
		t.Error("Unexpected tree structure at second level")
	}

	if tree.left.left.value != secondInsert {
		t.Error("Unexpected tree structure at third level")
	}
}
