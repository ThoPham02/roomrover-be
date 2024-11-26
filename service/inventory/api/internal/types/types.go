// Code generated by goctl. DO NOT EDIT.
package types

type UploadFileHouseReq struct {
	HouseID int64 `form:"houseID,optional"`
}

type UploadFileHouseRes struct {
	Result Result `json:"result"`
	Url    string `json:"url"`
}

type CreateHouseReq struct {
	Name        string `form:"name"`
	Description string `form:"description,optional"`
	Type        int64  `form:"type"`
	Area        int64  `form:"area"`
	Price       int64  `form:"price"`
	BedNum      int    `form:"bedNum,optional"`
	LivingNum   int    `form:"livingNum,optional"`
	Unit        int    `form:"unit,optional"`
	Address     string `form:"address"`
	WardID      int64  `form:"wardID"`
	DistrictID  int64  `form:"districtID"`
	ProvinceID  int64  `form:"provinceID"`
	Albums      string `form:"albums,optional"`
	Rooms       string `form:"rooms,optional"`
	Services    string `form:"services,optional"`
	Option      int    `form:"option,optional"`
}

type CreateHouseRes struct {
	Result Result `json:"result"`
	House  House  `json:"house"`
}

type FilterHouseReq struct {
	Search string `form:"search,optional"`
	Type   int64  `form:"type,optional"`
	Status int64  `form:"status,optional"`
	Limit  int64  `form:"limit"`
	Offset int64  `form:"offset"`
}

type FilterHouseRes struct {
	Result     Result  `json:"result"`
	Total      int64   `json:"total"`
	ListHouses []House `json:"listHouses"`
}

type UpdateHouseStatusReq struct {
	HouseID int64 `path:"id"`
	Status  int64 `form:"status"`
}

type UpdateHouseStatusRes struct {
	Result Result `json:"result"`
}

type UpdateHouseReq struct {
	HouseID     int64  `path:"id"`
	Name        string `form:"name"`
	Description string `form:"description"`
	Type        int64  `form:"type"`
	Area        int64  `form:"area"`
	Price       int64  `form:"price"`
	BedNum      int    `form:"bedNum,optional"`
	LivingNum   int    `form:"livingNum,optional"`
	Unit        int    `form:"unit,optional"`
	Address     string `form:"address"`
	WardID      int64  `form:"wardID"`
	DistrictID  int64  `form:"districtID"`
	ProvinceID  int64  `form:"provinceID"`
	Albums      string `form:"albums,optional"`
	Rooms       string `form:"rooms,optional"`
	Services    string `form:"services,optional"`
	Option      int    `form:"option,optional"`
}

type UpdateHouseRes struct {
	Result Result `json:"result"`
	House  House  `json:"house"`
}

type DeleteHouseReq struct {
	HouseID int64 `path:"id"`
}

type DeleteHouseRes struct {
	Result Result `json:"result"`
}

type FilterRoomReq struct {
	Search string `form:"search,optional"`
	Type   int64  `form:"type,optional"`
	Status int64  `form:"status,optional"`
	Limit  int64  `form:"limit"`
	Offset int64  `form:"offset"`
}

type FilterRoomRes struct {
	Result Result `json:"result"`
	Total  int    `json:"total"`
	Rooms  []Room `json:"rooms"`
}

type UpdateRoomStatusReq struct {
	RoomID int64 `path:"id"`
	Status int64 `form:"status"`
}

type UpdateRoomStatusRes struct {
	Result Result `json:"result"`
}

type GetRoomReq struct {
	ID int64 `path:"id"`
}

type GetRoomRes struct {
	Result   Result   `json:"result"`
	Room     Room     `json:"room"`
	House    House    `json:"house"`
	Contract Contract `json:"contract"`
}

type SearchRoomReq struct {
	Search string `form:"search,optional"`
	Type   int64  `form:"type,optional"`
	Status int64  `form:"status,optional"`
	Limit  int64  `form:"limit"`
	Offset int64  `form:"offset"`
}

type SearchRoomRes struct {
	Result Result `json:"result"`
	Rooms  []Room `json:"rooms"`
	Total  int    `json:"total"`
}

type CreateContactReq struct {
	HouseID  int64 `form:"houseID"`
	LessorID int64 `form:"lessorID"`
	Datetime int64 `form:"datetime"`
}

type CreateContactRes struct {
	Result  Result  `json:"result"`
	Contact Contact `json:"contact"`
}

type DeleteContactReq struct {
	ID int64 `path:"id"`
}

type DeleteContactRes struct {
	Result Result `json:"result"`
}

type FilterContactReq struct {
	From   int64 `form:"from,optional"`
	To     int64 `form:"to,optional"`
	Status int64 `form:"status,optional"`
	Limit  int64 `form:"limit"`
	Offset int64 `form:"offset"`
}

type FilterContactRes struct {
	Result   Result    `json:"result"`
	Contacts []Contact `json:"contacts"`
	Total    int       `json:"total"`
}

type UpdateContactStatusReq struct {
	ID     int64 `path:"id"`
	Status int64 `form:"status"`
}

type UpdateContactStatusRes struct {
	Result Result `json:"result"`
}

type HouseRevenue struct {
	HouseID   int64  `json:"houseID"`
	HouseName string `json:"houseName"`
	Revenue   int64  `json:"revenue"`
}

type GetDashboardReq struct {
}

type GetDashboardRes struct {
	Result          Result         `json:"result"`
	TotalRoom       int            `json:"totalRoom"`
	RentedRoom      int            `json:"rentedRoom"`
	EmptyRoom       int            `json:"emptyRoom"`
	TotalAmount     int64          `json:"totalAmount"`
	CurrentContact  []Contact      `json:"currentContact"`
	ExpiredContract []Contract     `json:"expiredContract"`
	HouseRevenue    []HouseRevenue `json:"houseRevenue"`
}

type SearchHouseReq struct {
	Search     string `form:"search,optional"`
	DistrictID int64  `form:"districtID,optional"`
	ProvinceID int64  `form:"provinceID,optional"`
	WardID     int64  `form:"wardID,optional"`
	Type       int64  `form:"type,optional"`
	PriceFrom  int64  `form:"priceFrom,optional"`
	PriceTo    int64  `form:"priceTo,optional"`
	AreaFrom   int64  `form:"areaFrom,optional"`
	AreaTo     int64  `form:"areaTo,optional"`
	Unit       int64  `form:"unit,optional"`
	Limit      int64  `form:"limit"`
	Offset     int64  `form:"offset"`
}

type SearchHouseRes struct {
	Result Result  `json:"result"`
	Total  int     `json:"total"`
	Houses []House `json:"houses"`
}

type GetHouseReq struct {
	ID int64 `path:"id"`
}

type GetHouseRes struct {
	Result Result `json:"result"`
	House  House  `json:"house"`
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

type RenterContact struct {
	ID          int64  `json:"id"`
	RoomName    string `json:"roomName"`
	Name        string `json:"name"`
	Phone       string `json:"phone"`
	CccdNumber  string `json:"cccdNumber"`
	CccdDate    int64  `json:"cccdDate"`
	CccdAddress string `json:"cccdAddress"`
	Status      int64  `json:"status"`
}
