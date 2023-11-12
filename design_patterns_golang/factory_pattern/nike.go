package main

type NikeFactory struct {
}

func (n *NikeFactory) makeShoe() IShoe {
	return &NikeShoe{
		Shoe: Shoe{
			logo: "nike",
			size: 15,
		},
	}
}

func (n *NikeFactory) makeShirt() IShirt {
	return &NikeShirt{
		Shirt: Shirt{
			logo: "nike",
			size: 15,
		},
	}
}
