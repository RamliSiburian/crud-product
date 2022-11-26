package productDto

type productResponse struct {
	Nama      string `json:"nama" `
	HargaBeli int    `json:"harga_beli" `
	HargaJual int    `json:"harga_jual"`
	Stok      int    `json:"stok" `
}
