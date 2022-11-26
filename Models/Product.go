package Models

type Product struct {
	ID        int          `json:"id"`
	Image     string       `json:"image" gorm:"type: varchar(255)"`
	Nama      string       `json:"nama" gorm:"type: varchar(255)"`
	HargaBeli int          `json:"harga_beli"`
	HargaJual int          `json:"harga_jual"`
	Stok      int          `json:"stok"`
	UserID    int          `json:"user_id"`
	User      UserResponse `json:"user"`
}
