// Code generated by goctl. DO NOT EDIT.
package types

type NotiInfo struct {
	ID   int64       `json:"id"`
	Name interface{} `json:"name"`
}

type Notification struct {
	NotificationID int64      `json:"id"`
	SenderID       int64      `json:"senderID"`
	ReceiverID     int64      `json:"receiverID"`
	RefID          int64      `json:"refID"`
	RefType        int64      `json:"refType"`
	Unread         int64      `json:"unread"`
	NotiInfos      []NotiInfo `json:"notiInfos"`
	CreatedAt      int64      `json:"createdAt"`
}

type CreateNotificationReq struct {
	Sender      int64  `form:"sender"`
	Receiver    int64  `form:"receiver"`
	RefID       int64  `form:"refID"`
	RefType     int64  `form:"refType"`
	Title       string `form:"title"`
	Description string `form:"description"`
	Priority    int64  `form:"priority"`
	DueDate     int64  `form:"dueDate"`
}

type CreateNotificationRes struct {
	Result       Result       `json:"result"`
	Notification Notification `json:"notification"`
}

type GetListNotificationReq struct {
	Limit  int64 `form:"limit"`
	Offset int64 `form:"offset"`
}

type GetListNotificationRes struct {
	Result        Result         `json:"result"`
	Notifications []Notification `json:"notifications"`
	Total         int            `json:"total"`
}

type MarkReadReq struct {
	ID int64 `path:"id"`
}

type MarkReadRes struct {
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
