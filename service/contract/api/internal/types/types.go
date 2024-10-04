// Code generated by goctl. DO NOT EDIT.
package types

type CreateContractReq struct {
	RoomID         int64  `form:"roomID"`
	Description    string `form:"description"`
	ContractUrl    string `form:"contractUrl"`
	Start          int64  `form:"start"`
	End            int64  `form:"end"`
	ContractRenter string `form:"contractRenter"`
	ContractDetail string `form:"contractDetail"`
	Type           int64  `form:"type"`
	Deposit        int64  `form:"deposit"`
	Deadline       int64  `form:"deadline"`
	DepositUrl     string `form:"depositUrl"`
}

type CreateContractRes struct {
	Result   Result   `json:"result"`
	Contract Contract `json:"contract"`
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
	RoomID   int64  `json:"roomID"`
	HouseID  int64  `json:"houseID"`
	Name     string `json:"name"`
	Status   int64  `json:"status"`
	Capacity int64  `json:"capacity"`
	EIndex   int64  `json:"eIndex"`
	WIndex   int64  `json:"wIndex"`
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
	RenterNumber    string           `json:"renterNumber"`
	RenterDate      int64            `json:"renterDate"`
	RenterAddress   string           `json:"renterAddress"`
	RenterName      string           `json:"renterName"`
	LessorID        int64            `json:"lessorID"`
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
	ID         int64 `json:"id"`
	ContractID int64 `json:"contractID"`
	RenterID   int64 `json:"renterID"`
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
