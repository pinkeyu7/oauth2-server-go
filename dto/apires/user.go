package apires

type User struct {
	Id      int    `xorm:"pk autoincr BIGINT(20)" json:"id"`
	Account string `xorm:"not null default '' comment('account') VARCHAR(64)" json:"account"`
	Phone   string `xorm:"not null default '' comment('phone') VARCHAR(20)" json:"phone"`
	Email   string `xorm:"not null default '' comment('email') VARCHAR(64)" json:"email"`
	Name    string `xorm:"not null default '' comment('name') VARCHAR(64)" json:"name"`
}
