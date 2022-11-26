package productDto

type ProductRequest struct {
	Nama      string `json:"nama" form:"nama"`
	HargaBeli int    `json:"harga_beli" form:"harga_beli"`
	HargaJual int    `json:"harga_jual" form:"harga_jual"`
	Stok      int    `json:"stok" form:"stok"`
	UserID    int    `json:"user_id" form:"user_id"`
}

type UpdateProductRequest struct {
	Nama      string `json:"nama" form:"nama"`
	HargaBeli int    `json:"harga_beli" form:"harga_beli"`
	HargaJual int    `json:"harga_jual" form:"harga_jual"`
	Stok      int    `json:"stok" form:"stok"`
}
