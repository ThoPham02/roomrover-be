// Code generated by goctl. DO NOT EDIT.
package types

type CreateContractReq struct {
	Renter        string `form:"renter"`
	Lessor        string `form:"lessor"`
	PaymentRenter string `form:"paymentRenter,optional"`
	Room          string `form:"room"`
	CheckIn       int64  `form:"checkIn"`
	Duration      int64  `form:"duration"`
	Purpose       string `form:"purpose"`
	Price         int64  `form:"price"`
	Discount      int64  `form:"discount,optional"`
	Deposit       int64  `form:"deposit"`
	DepositDate   int64  `form:"depositDate"`
}

type CreateContractRes struct {
	Result   Result   `json:"result"`
	Contract Contract `json:"contract"`
}

type UpdateContractReq struct {
	ID            int64  `path:"id"`
	Renter        string `form:"renter"`
	Lessor        string `form:"lessor"`
	PaymentRenter string `form:"paymentRenter,optional"`
	Room          string `form:"room"`
	CheckIn       int64  `form:"checkIn"`
	Duration      int64  `form:"duration"`
	Purpose       string `form:"purpose"`
	Price         int64  `form:"price"`
	Discount      int64  `form:"discount,optional"`
	Deposit       int64  `form:"deposit"`
	DepositDate   int64  `form:"depositDate"`
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

type UpdateContractStatusReq struct {
	ID     int64 `path:"id"`
	Status int64 `form:"status"`
}

type UpdateContractStatusRes struct {
	Result Result `json:"result"`
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
	Contracts []Contract `json:"contracts"`
	Total     int64      `json:"total"`
}

type FilterBillReq struct {
	Search     string `form:"search,optional"`
	CreateFrom int64  `form:"createFrom,optional"`
	CreateTo   int64  `form:"createTo,optional"`
	Status     int64  `form:"status,optional"`
	Limit      int64  `form:"limit"`
	Offset     int64  `form:"offset"`
}

type FilterBillRes struct {
	Result Result `json:"result"`
	Bills  []Bill `json:"bills"`
	Total  int64  `json:"total"`
}

type GetBillDetailReq struct {
	ID int64 `path:"id"`
}

type GetBillDetailRes struct {
	Result Result `json:"result"`
	Bill   Bill   `json:"bill"`
}

type UpdateBillDetailReq struct {
	ID       int64  `path:"id"`
	Amount   int64  `form:"amount"`
	Discount int64  `form:"discount"`
	Note     string `form:"note"`
}

type UpdateBillDetailRes struct {
	Result Result `json:"result"`
	Bill   Bill   `json:"bill"`
}

type CreateBillPayReq struct {
	BillID  int64  `form:"billID"`
	Amount  int64  `form:"amount"`
	PayType int64  `form:"payType"`
	PayDate int64  `form:"payDate"`
	Url     string `form:"url,optional"`
}

type CreateBillPayRes struct {
	Result Result `json:"result"`
	Bill   Bill   `json:"bill"`
}

type DeleteBillPayReq struct {
	ID int64 `path:"id"`
}

type DeleteBillPayRes struct {
	Result Result `json:"result"`
}

type UpdateBillStatusReq struct {
	ID     int64 `path:"id"`
	Status int64 `form:"status"`
}

type UpdateBillStatusRes struct {
	Result Result `json:"result"`
}

type ZaloPaymentReq struct {
	BillID int64 `form:"billID"`
}

type ZaloPaymentRes struct {
	Result   Result `json:"result"`
	OrderUrl string `json:"orderUrl"`
	QrCode   string `json:"qrCode"`
}

type ZaloPaymentCallbackReq struct {
	Mac  string `json:"mac"`
	Data string `json:"data"`
	Type int    `json:"type"`
}

type ZaloPaymentCallbackRes struct {
	ReturnCode    int    `json:"return_code"`
	ReturnMessage string `json:"return_message"`
}

type ConfirmContractReq struct {
	ID       int64  `path:"id"`
	Renters  string `form:"renters,optional"`
	Albums   string `form:"albums,optional"`
	Services string `form:"services,optional"`
}

type ConfirmContractRes struct {
	Result Result `json:"result"`
}

type GetListBillDetailReq struct {
	BillID int64 `path:"billID"`
}

type GetListBillDetailRes struct {
	Result      Result       `json:"result"`
	BillDetails []BillDetail `json:"billDetails"`
}

type UpdateBillDetailsReq struct {
	BillID      int64  `path:"billID"`
	BillDetails string `form:"billDetails"`
}

type UpdateBillDetailsRes struct {
	Result Result `json:"result"`
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
	User        User      `json:"user"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Type        int64     `json:"type"`
	Status      int64     `json:"status"`
	Area        int64     `json:"area"`
	Price       int64     `json:"price"`
	BedNum      int64     `json:"bedNum"`
	LivingNum   int64     `json:"livingNum"`
	Unit        int64     `json:"unit"`
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
	RoomID     int64     `json:"roomID"`
	HouseID    int64     `json:"houseID"`
	Name       string    `json:"name"`
	HouseName  string    `json:"houseName"`
	ProvinceID int64     `json:"provinceID"`
	DistrictID int64     `json:"districtID"`
	WardID     int64     `json:"wardID"`
	Address    string    `json:"address"`
	Area       int64     `json:"area"`
	Price      int64     `json:"price"`
	Type       int64     `json:"type"`
	Status     int64     `json:"status"`
	Capacity   int64     `json:"capacity"`
	EIndex     int64     `json:"eIndex"`
	WIndex     int64     `json:"wIndex"`
	Services   []Service `json:"services"`
}

type Service struct {
	ServiceID int64  `json:"serviceID"`
	HouseID   int64  `json:"houseID"`
	Name      string `json:"name"`
	Price     int64  `json:"price"`
	Unit      int64  `json:"unit"`
}

type Contact struct {
	ID          int64  `json:"id"`
	HouseID     int64  `json:"houseID"`
	HouseName   string `json:"houseName"`
	ProvinceID  int64  `json:"provinceID"`
	DistrictID  int64  `json:"districtID"`
	WardID      int64  `json:"wardID"`
	Address     string `json:"address"`
	RenterID    int64  `json:"renterID"`
	RenterName  string `json:"renterName"`
	RenterPhone string `json:"renterPhone"`
	LessorID    int64  `json:"lessorID"`
	LessorName  string `json:"lessorName"`
	LessorPhone string `json:"lessorPhone"`
	Datetime    int64  `json:"datetime"`
	Status      int64  `json:"status"`
}

type Contract struct {
	ContractID    int64    `json:"contractID"`
	Code          string   `json:"code"`
	Status        int64    `json:"status"`
	Renter        User     `json:"renter"`
	Lessor        User     `json:"lessor"`
	Room          Room     `json:"room"`
	CheckIn       int64    `json:"checkIn"`
	Duration      int64    `json:"duration"`
	Purpose       string   `json:"purpose"`
	Payment       Payment  `json:"payment"`
	ConfirmedImgs []string `json:"confirmedImgs"`
	CreatedAt     int64    `json:"createdAt"`
	UpdatedAt     int64    `json:"updatedAt"`
	CreatedBy     int64    `json:"createdBy"`
	UpdatedBy     int64    `json:"updatedBy"`
}

type PaymentRenter struct {
	ID          int64  `json:"id"`
	PaymentID   int64  `json:"paymentID"`
	RenterID    int64  `json:"renterID"`
	Name        string `json:"name"`
	Phone       string `json:"phone"`
	CccdNumber  string `json:"cccdNumber"`  // so can cuoc
	CccdDate    int64  `json:"cccdDate"`    // ngay cap
	CccdAddress string `json:"cccdAddress"` // noi cap
}

type PaymentDetail struct {
	ID        int64  `json:"id"`
	PaymentID int64  `json:"paymentID"`
	Name      string `json:"name"`
	Price     int64  `json:"price"`
	Type      int64  `json:"type"`
	Index     int64  `json:"index"`
}

type Payment struct {
	PaymentID      int64           `json:"paymentID"`
	ContractID     int64           `json:"contractID"`
	Amount         int64           `json:"amount"`
	Discount       int64           `json:"discount"`
	Deposit        int64           `json:"deposit"`
	DepositDate    int64           `json:"depositDate"`
	NextBill       int64           `json:"nextBill"`
	PaymentRenters []PaymentRenter `json:"paymentRenters"`
	PaymentDetails []PaymentDetail `json:"paymentDetails"`
}

type Bill struct {
	BillID       int64        `json:"billID"`
	Title        string       `json:"title"`
	ContractCode string       `json:"contractCode"`
	RenterID     int64        `json:"renterID"`
	RenterName   string       `json:"renterName"`
	RenterPhone  string       `json:"renterPhone"`
	LessorID     int64        `json:"lessorID"`
	LessorName   string       `json:"lessorName"`
	LessorPhone  string       `json:"lessorPhone"`
	PaymentID    int64        `json:"paymentID"`
	PaymentDate  int64        `json:"paymentDate"`
	Amount       int64        `json:"amount"`
	Remain       int64        `json:"remain"`
	Status       int64        `json:"status"`
	BillDetails  []BillDetail `json:"billDetails"`
	BillPays     []BillPay    `json:"billPays"`
}

type BillDetail struct {
	BillDetailID int64  `json:"billDetailID"`
	BillID       int64  `json:"billID"`
	Name         string `json:"name"`
	Price        int64  `json:"price"`
	Type         int64  `json:"type"`
	OldIndex     int64  `json:"oldIndex"`
	NewIndex     int64  `json:"newIndex"`
	Imgurl       string `json:"imgUrl"`
	Quantity     int64  `json:"quantity"`
}

type BillPay struct {
	BillPayID int64  `json:"billPayID"`
	UserID    int64  `json:"userID"`
	BillID    int64  `json:"billID"`
	Amount    int64  `json:"amount"`
	PayDate   int64  `json:"payDate"`
	Status    int64  `json:"status"`
	Type      int64  `json:"type"`
	Url       string `json:"url"`
	TransId   string `json:"transId"`
}
