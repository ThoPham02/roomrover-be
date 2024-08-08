// Code generated by goctl. DO NOT EDIT.
package types

type Album struct {
	AlbumID   int64  `json:"albumID"`
	HouseID   int64  `json:"houseID"`
	Url       string `json:"url"`
	Type      int64  `json:"type"`
	CreatedAt int64  `json:"createdAt"`
	UpdatedAt int64  `json:"updatedAt"`
	CreatedBy int64  `json:"createdBy"`
	UpdatedBy int64  `json:"updatedBy"`
}

type Contract struct {
	ContractID     int64          `json:"contractID"`
	RenterID       int64          `json:"renterID"`
	RoomID         int64          `json:"roomID"`
	Description    string         `json:"description"`
	ContractUrl    string         `json:"contractUrl"`
	StartDate      int64          `json:"startDate"`
	EndDate        int64          `json:"endDate"`
	Status         int64          `json:"status"`
	Type           int64          `json:"type"`
	Deposit        int64          `json:"deposit"`
	Deadline       int64          `json:"deadline"`
	CreatedAt      int64          `json:"createdAt"`
	UpdatedAt      int64          `json:"updatedAt"`
	CreatedBy      int64          `json:"createdBy"`
	UpdatedBy      int64          `json:"updatedBy"`
	ContractRenter ContractRenter `json:"contractRenter"`
	ContractDetail ContractDetail `json:"contractDetail"`
}

type ContractDetail struct {
	ID         int64 `json:"id"`
	ContractID int64 `json:"contractID"`
	ServiceID  int64 `json:"serviceID"`
	Price      int64 `json:"price"`
	Index      int64 `json:"index"`
	Status     int64 `json:"status"`
	CreatedAt  int64 `json:"createdAt"`
	UpdatedAt  int64 `json:"updatedAt"`
	CreatedBy  int64 `json:"createdBy"`
	UpdatedBy  int64 `json:"updatedBy"`
}

type ContractRenter struct {
	ID         int64 `json:"id"`
	ContractID int64 `json:"contractID"`
	RenterID   int64 `json:"renterID"`
	Type       int64 `json:"type"`
	Status     int64 `json:"status"`
	CreatedAt  int64 `json:"createdAt"`
	UpdatedAt  int64 `json:"updatedAt"`
	CreatedBy  int64 `json:"createdBy"`
	UpdatedBy  int64 `json:"updatedBy"`
}

type CreateHouseReq struct {
	Name        string `form:"name"`
	Description string `form:"description"`
	Type        int64  `form:"type"`
	Area        int64  `form:"area"`
	Price       int64  `form:"price"`
	Address     string `form:"address"`
	WardID      int64  `form:"wardID"`
	DistrictID  int64  `form:"districtID"`
	ProvinceID  int64  `form:"provinceID"`
	Albums      string `form:"albums"`
}

type CreateHouseRes struct {
	Result Result `json:"result"`
	House  House  `json:"house"`
}

type CreateRoomReq struct {
	HouseID     int64  `form:"houseID"`
	Name        string `form:"name"`
	Description string `form:"description"`
	Price       int64  `form:"price"`
	Area        int64  `form:"area"`
}

type CreateRoomRes struct {
	Result Result `json:"result"`
	Room   Room   `json:"room"`
}

type DeleteHouseReq struct {
	HouseID int64 `path:"houseID"`
}

type DeleteHouseRes struct {
	Result Result `json:"result"`
}

type DeleteRoomReq struct {
	RoomID int64 `path:"roomID"`
}

type DeleteRoomRes struct {
	Result Result `json:"result"`
}

type FilterHouseReq struct {
	Area      int64  `form:"area,optional"`
	Price     int64  `form:"price,optional"`
	Longitude string `form:"longitude,optional"`
	Latitude  string `form:"latitude,optional"`
}

type FilterHouseRes struct {
	Result Result  `json:"result"`
	Houses []House `json:"houses"`
}

type GetHouseReq struct {
	HouseID int64 `path:"houseID"`
}

type GetHouseRes struct {
	Result Result `json:"result"`
	House  House  `json:"house"`
}

type GetRoomByHouseReq struct {
	HouseID int64 `path:"houseID"`
}

type GetRoomByHouseRes struct {
	Result Result `json:"result"`
	Rooms  []Room `json:"rooms"`
}

type GetRoomReq struct {
	RoomID int64 `path:"roomID"`
}

type GetRoomRes struct {
	Result Result `json:"result"`
	Room   Room   `json:"room"`
}

type House struct {
	HouseID     int64     `json:"houseID"`
	UserID      int64     `json:"userID"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Type        int64     `json:"type"`
	Area        int64     `json:"area"`
	Price       int64     `json:"price"`
	Status      int64     `json:"status"`
	Address     string    `json:"address"`
	WardID      int64     `json:"wardID"`
	DistrictID  int64     `json:"districtID"`
	ProvinceID  int64     `json:"provinceID"`
	CreatedAt   int64     `json:"createdAt"`
	UpdatedAt   int64     `json:"updatedAt"`
	CreatedBy   int64     `json:"createdBy"`
	UpdatedBy   int64     `json:"updatedBy"`
	Albums      []Album   `json:"albums"`
	Rooms       []Room    `json:"rooms"`
	Services    []Service `json:"services"`
}

type Payment struct {
	ID         int64 `json:"id"`
	ContractID int64 `json:"contractID"`
	Month      int64 `json:"month"`
	Total      int64 `json:"total"`
	Status     int64 `json:"status"`
	CreatedAt  int64 `json:"createdAt"`
	UpdatedAt  int64 `json:"updatedAt"`
	CreatedBy  int64 `json:"createdBy"`
	UpdatedBy  int64 `json:"updatedBy"`
}

type PaymentDetail struct {
	ID        int64  `json:"id"`
	PaymentID int64  `json:"paymentID"`
	Amount    int64  `json:"amount"`
	Type      int64  `json:"type"`
	Utl       string `json:"utl"`
	CreatedAt int64  `json:"createdAt"`
	UpdatedAt int64  `json:"updatedAt"`
	CreatedBy int64  `json:"createdBy"`
	UpdatedBy int64  `json:"updatedBy"`
}

type PaymentUsed struct {
	ID        int64 `json:"id"`
	PaymentID int64 `json:"paymentID"`
	ServiceID int64 `json:"serviceID"`
	Index     int64 `json:"index"`
	Status    int64 `json:"status"`
	CreatedAt int64 `json:"createdAt"`
	UpdatedAt int64 `json:"updatedAt"`
	CreatedBy int64 `json:"createdBy"`
	UpdatedBy int64 `json:"updatedBy"`
}

type Result struct {
	Code    int    `json:"code"`    //    Result code: 0 is success. Otherwise, getting an error
	Message string `json:"message"` // Result message: detail response code
}

type Room struct {
	RoomID    int64  `json:"roomID"`
	HouseID   int64  `json:"houseID"`
	Name      string `json:"name"`
	Status    int64  `json:"status"`
	CreatedAt int64  `json:"createdAt"`
	UpdatedAt int64  `json:"updatedAt"`
	CreatedBy int64  `json:"createdBy"`
	UpdatedBy int64  `json:"updatedBy"`
}

type Service struct {
	ServiceID int64  `json:"serviceID"`
	HouseID   int64  `json:"houseID"`
	Name      string `json:"name"`
	Price     int64  `json:"price"`
	Type      int64  `json:"type"`
	CreatedAt int64  `json:"createdAt"`
	UpdatedAt int64  `json:"updatedAt"`
	CreatedBy int64  `json:"createdBy"`
	UpdatedBy int64  `json:"updatedBy"`
}

type UpdateHouseReq struct {
	HouseID   int64  `form:"houseID"`
	Title     string `form:"title"`
	Price     int64  `form:"price"`
	Area      int64  `form:"area"`
	Address   string `form:"address"`
	Longitude string `form:"longitude"`
	Latitude  string `form:"latitude"`
}

type UpdateHouseRes struct {
	Result Result `json:"result"`
	House  House  `json:"house"`
}

type UpdateRoomReq struct {
	RoomID      int64  `form:"roomID"`
	Name        string `form:"name"`
	Description string `form:"description"`
	Price       int64  `form:"price"`
	Area        int64  `form:"area"`
}

type UpdateRoomRes struct {
	Result Result `json:"result"`
	Room   Room   `json:"room"`
}

type UploadFileHouseReq struct {
	HouseID  int64  `form:"houseID"`
	FileName string `form:"fileName"`
}

type UploadFileHouseRes struct {
	Result Result `json:"result"`
	Url    string `json:"url"`
}

type User struct {
	UserID    int64  `json:"userID"`
	Phone     string `json:"phone"`
	FullName  string `json:"fullName"`
	Birthday  int64  `json:"birthday"`
	AvatarUrl string `json:"avatarUrl"`
	Address   string `json:"address"`
	CreatedAt int64  `json:"createdAt"`
	UpdatedAt int64  `json:"updatedAt"`
}
