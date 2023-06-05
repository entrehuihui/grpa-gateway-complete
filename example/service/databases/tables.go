package databases

// 用户列表
type Userinfo struct {
	ID         int    `gorm:"primary_key;unique_index;AUTO_INCREMENT:10000"` // 平台唯一ID
	Nicename   string `gorm:"not null;size:100"`                             // 昵称
	Emails     string `gorm:"not null;size:100;index"`                       // 登录账户
	Password   string `gorm:"not null;size:32"`                              // 密码
	Createtime int64  `gorm:"not null"`                                      // 创建时间
	Updatetime int64  `gorm:"not null"`                                      // 更新时间
	Del        int    `gorm:"not null"`                                      // 删除 1删除
	State      int    `gorm:"not null"`                                      // 1未激活 2已激活 3停用
}
