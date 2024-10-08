// Code generated by goctl. DO NOT EDIT.
package types

type CreateContractReq struct {
	Renter         string `form:"renter"`
	Lessor         string `form:"lessor"`
	ContractRenter string `form:"contractRenter,optional"`
	Room           string `form:"room"`
	CheckIn        int64  `form:"checkIn"`
	Duration       int64  `form:"duration"`
	Purpose        string `form:"purpose"`
	Price          int64  `form:"price"`
	Discount       int64  `form:"discount,optional"`
	Deposit        int64  `form:"deposit"`
	DepositDate    int64  `form:"depositDate"`
}

type CreateContractRes struct {
	Result   Result   `json:"result"`
	Contract Contract `json:"contract"`
}

type UpdateContractReq struct {
	ID             int64  `path:"id"`
	Status         int64  `form:"status"`
	RenterID       int64  `form:"renterID"`
	RenterNumber   string `form:"renterNumber"`
	RenterDate     string `form:"renterDate"`
	RenterName     string `form:"renterName"`
	RenterAddress  string `form:"renterAddress"`
	LessorID       int64  `form:"lessorID"`
	LessorNumber   string `form:"lessorNumber"`
	LessorDate     string `form:"lessorDate"`
	LessorName     string `form:"lessorName"`
	LessorAddress  string `form:"lessorAddress"`
	ContractRenter string `form:"contractRenter"`
	RoomID         int64  `form:"roomID"`
	EIndex         int64  `form:"eIndex"`
	WIndex         int64  `form:"wIndex"`
	CheckIn        int64  `form:"checkIn"`
	Duration       int64  `form:"duration"`
	Purpose        string `form:"purpose"`
	Amount         int64  `form:"amount"`
	Discount       int64  `form:"discount"`
	Deposit        int64  `form:"deposit"`
	DepositDate    int64  `form:"depositDate"`
}

type UpdateContractRes struct {
	Result   Result   `json:"result"`
	Contract Contract `json:"contract"`
}

type GetContractReq struct {
	ID int64 `path:"id"`
}

type GetContractRes struct {
	Result   Result   `json:"result"`
	Contract Contract `json:"contract"`
}

type DeleteContractReq struct {
	ID int64 `path:"id"`
}

type DeleteContractRes struct {
	Result Result `json:"result"`
}

type FilterContractReq struct {
	Search     string `form:"search,optional"`
	CreateFrom int64  `form:"createFrom,optional"`
	CreateTo   int64  `form:"createTo,optional"`
	Status     int64  `form:"status,optional"`
	Limit      int64  `form:"limit"`
	Offset     int64  `form:"offset"`
}

type FilterContractRes struct {
	Result    Result     `json:"result"`
	Contracts []Contract `form:"contracts"`
	Total     int64      `form:"total"`
}

type Result struct {
	Code    int    `json:"code"`    //    Result code: 0 is success. Otherwise, getting an error
	Message string `json:"message"` // Result message: detail response code
}

type User struct {
	UserID      int64  `json:"userID"`
	Phone       string `json:"phone"`
	Role        int64  `json:"role"`
	Status      int64  `json:"status"`
	Address     string `json:"address"`
	FullName    string `json:"fullName"`
	AvatarUrl   string `json:"avatarUrl"`
	Birthday    int64  `json:"birthday"`
	Gender      int64  `json:"gender"`
	CccdNumber  string `json:"cccdNumber"`
	CccdDate    int64  `json:"cccdDate"`
	CccdAddress string `json:"cccdAddress"`
	CreatedAt   int64  `json:"createdAt"`
	UpdatedAt   int64  `json:"updatedAt"`
}

type House struct {
	HouseID     int64     `json:"houseID"`
	UserID      int64     `json:"userID"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Type        int64     `json:"type"`
	Status      int64     `json:"status"`
	Area        int64     `json:"area"`
	Price       int64     `json:"price"`
	BedNum      int64     `json:"bedNum"`
	LivingNum   int64     `json:"livingNum"`
	Albums      []string  `json:"albums"`
	Rooms       []Room    `json:"rooms"`
	Services    []Service `json:"services"`
	Address     string    `json:"address"`
	WardID      int64     `json:"wardID"`
	DistrictID  int64     `json:"districtID"`
	ProvinceID  int64     `json:"provinceID"`
	CreatedAt   int64     `json:"createdAt"`
	UpdatedAt   int64     `json:"updatedAt"`
	CreatedBy   int64     `json:"createdBy"`
	UpdatedBy   int64     `json:"updatedBy"`
}

type Album struct {
	AlbumID int64  `json:"albumID"`
	HouseID int64  `json:"houseID"`
	Url     string `json:"url"`
}

type Room struct {
	RoomID    int64  `json:"roomID"`
	HouseID   int64  `json:"houseID"`
	HouseName string `json:"houseName"`
	Area      int64  `json:"area"`
	Price     int64  `json:"price"`
	Type      int64  `json:"type"`
	Name      string `json:"name"`
	Status    int64  `json:"status"`
	Capacity  int64  `json:"capacity"`
	EIndex    int64  `json:"eIndex"`
	WIndex    int64  `json:"wIndex"`
}

type Service struct {
	ServiceID int64  `json:"serviceID"`
	HouseID   int64  `json:"houseID"`
	Name      string `json:"name"`
	Price     int64  `json:"price"`
	Unit      int64  `json:"unit"`
}

type Contract struct {
	ContractID      int64            `json:"contractID"`
	Code            string           `json:"code"`
	Status          int64            `json:"status"`
	RenterID        int64            `json:"renterID"`
	RenterPhone     string           `json:"renterPhone"`
	RenterNumber    string           `json:"renterNumber"`
	RenterDate      int64            `json:"renterDate"`
	RenterAddress   string           `json:"renterAddress"`
	RenterName      string           `json:"renterName"`
	LessorID        int64            `json:"lessorID"`
	LessorPhone     string           `json:"lessorPhone"`
	LessorNumber    string           `json:"lessorNumber"`
	LessorDate      int64            `json:"lessorDate"`
	LessorAddress   string           `json:"lessorAddress"`
	LessorName      string           `json:"lessorName"`
	RoomID          int64            `json:"roomID"`
	CheckIn         int64            `json:"checkIn"`
	Duration        int64            `json:"duration"`
	Purpose         string           `json:"purpose"`
	ContractRenters []ContractRenter `json:"contractRenters"`
	ContractDetails []ContractDetail `json:"contractDetails"`
	Payment         Payment          `json:"payment"`
	CreatedAt       int64            `json:"createdAt"`
	UpdatedAt       int64            `json:"updatedAt"`
	CreatedBy       int64            `json:"createdBy"`
	UpdatedBy       int64            `json:"updatedBy"`
}

type ContractRenter struct {
	ID         int64  `json:"id"`
	ContractID int64  `json:"contractID"`
	RenterID   int64  `json:"renterID"`
	Name       string `json:"name"`
	Phone      string `json:"phone"`
}

type ContractDetail struct {
	ID         int64  `json:"id"`
	ContractID int64  `json:"contractID"`
	Name       string `json:"name"`
	Price      int64  `json:"price"`
	Type       int64  `json:"type"`
}

type Payment struct {
	PaymentID   int64 `json:"paymentID"`
	ContractID  int64 `json:"contractID"`
	Amount      int64 `json:"amount"`
	Discount    int64 `json:"discount"`
	Deposit     int64 `json:"deposit"`
	DepositDate int64 `json:"depositDate"`
	NextBill    int64 `json:"nextBill"`
}

type Bill struct {
	BillID      int64        `json:"billID"`
	PaymentID   int64        `json:"paymentID"`
	PaymentDate int64        `json:"paymentDate"`
	Amount      int64        `json:"amount"`
	Status      int64        `json:"status"`
	BillDetails []BillDetail `json:"billDetails"`
}

type BillDetail struct {
	BillDetailID int64  `json:"billDetailID"`
	BillID       int64  `json:"billID"`
	Name         int64  `json:"name"`
	Price        int64  `json:"price"`
	Type         string `json:"type"`
	Quantity     int64  `json:"quantity"`
}

type BillPay struct {
	BillPayID int64 `json:"billPayID"`
	BillID    int64 `json:"billID"`
	Amount    int64 `json:"amount"`
	PayDate   int64 `json:"payDate"`
	UserID    int64 `json:"userID"`
}
