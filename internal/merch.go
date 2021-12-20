package ik_common

type MerchCategory struct {
	Id 					int 	 `json:"id"`
	Title 			string `json:"title"`
	Description	string `json:"description"`
}

type MerchItem struct {
	Id 					int 	 `json:"id"`
	CategoryId 	int 	 `json:"category_id"`
	Title 			string `json:"title"`
	Description	string `json:"description"`
	Quantity 		int 	 `json:"quantity"`
	Size 				int 	 `json:"size"`
	Color 			int		 `json:"color"`
}
