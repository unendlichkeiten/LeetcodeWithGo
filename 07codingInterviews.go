package leetCode

func CreateBiTreeByPreAndInOrder(preOrder []int, prel, prer int,
																 inOrder []int, inl, inr int)(BiTree, error){
	if prel > prer {
		return nil, nil
	}

	// create binary tree node
	node := &BiTNode{
		data: preOrder[prel],
		lchild: nil,
		rchild: nil,
	}

	// search node value position in inOrder
	// search preOrder left sub tree set and right sub tree set
	// for range bug index
	index := 0
	for index, _ = range inOrder[inl:inr+1] {
		if inOrder[inl+index] == preOrder[prel] {
			break
		}
	}
	index += inl

	node.lchild, _ = CreateBiTreeByPreAndInOrder(preOrder, prel+1, prel+index-inl,
																							 inOrder, inl, index-1)
	node.rchild, _ = CreateBiTreeByPreAndInOrder(preOrder, prel+index-inl+1,
		 																					 prer, inOrder, index+1, inr)

	return node, nil
}