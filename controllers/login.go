package controllers

import (
	"yinji/models/bean"
	"time"
	"yinji/service/db"
	"github.com/astaxie/beego/orm"
	"errors"
	"yinji/service"
)

type LoginController struct {
	BeegoController
}

const LOGIN_CONTROLLER_PARAMTER_ACCOUNT = "account"
const LOGIN_CONTROLLER_PARAMTER_PASSWORD = "password"

//对应的 注册接口
func ( self *LoginController ) RegisterByAccount(){

	//获取目标传送而来的信息
	var account = self.GetString( LOGIN_CONTROLLER_PARAMTER_ACCOUNT )

	var password = self.GetString( LOGIN_CONTROLLER_PARAMTER_PASSWORD )

	if account == "" || password == "" {
		//不允许对应的信息为空
		self.FailJson(errors.New(" account and password is not nil "))
		return
	}

	var time = time.Now()

	//根据对方输入的密码 ， 之后进行生成对应的 信息
	var login = bean.Login{}

	login.Acount = account

	login.Password = password

	login.NewEntity( time )

	//为对应的user 设置信息
	var user = bean.User{}

	var userDetails = &bean.UserDetails{}

	user.NewEntity( time )

	//暂时默认 user的初始化名称为 login account 名称
	user.Name = login.Acount

	userDetails.Id = user.Id
	userDetails.ModifyTime = user.CreateTime
	userDetails.Birthday = time

	var ormService = db.GetOrmServiceInstance()

	var userDetailsService = service.GetUserDetailsServiceInstance()

	//下面便开始数据库操作
	var _ , transactionErr = ormService.Transaction(func(o orm.Ormer) (interface{}, error) {

		var loginTemp = bean.Login{}

		loginTemp.Acount = login.Acount

		var readErr = o.Read( &loginTemp , "account")

		if readErr == nil {
			//倘若查找到了 ， 则直接退出
			var err = errors.New(" account is exist ")
			return nil , err
		}

		var _ , insertUserError = o.Insert( &user )

		if insertUserError != nil {
			//错误
			return nil , insertUserError
		}

		//移植对应的user
		login.Id = user.Id

		var _login = &login

		var _ , insertLoginError = o.Insert( _login )

		if insertLoginError != nil {
			return nil , insertLoginError
		}

		var _ , newDetailsErr = userDetailsService.New(o , userDetails)

		if newDetailsErr != nil {
			return nil , newDetailsErr
		}

		return _login, nil
	});

	//最后对结果进行判断
	if transactionErr != nil {
		self.FailJson( transactionErr )
		return
	}

	//倘若生成，将当前 user 调整为 注册账号
	self.StartSession().Set(SESSION_LOCAL_USER_KEY , user)

	//之后返回结果
	self.Json( user )

}

/**
	由于一次性开发容错性高的接口 ，容易造成 高耦合 ，
	因此  ，我个人在这里建议先把登录接口分开，再合并
 */

 const SESSION_LOCAL_USER_KEY = "LOCAL_USER"
 //利用账号密码登录
func ( self *LoginController ) LoginByAccount(){
	//开始对应的输出信息

	//获取对应的账号密码 , 并且进行组装到对应的 login 之中
	var account = self.GetString(LOGIN_CONTROLLER_PARAMTER_ACCOUNT)

	var password = self.GetString(LOGIN_CONTROLLER_PARAMTER_PASSWORD)

	//测试对应的信息是否为空
	if account == "" || password == "" {
		self.FailJson(errors.New("account and password have nil"))
		return
	}

	var login = bean.Login{}

	login.Acount = account

	login.Password = password

	var user = bean.User{}

	//这里 ， 我们设定对应的程序

	//程序进入下一个阶段
	var ormService = db.GetOrmServiceInstance()

	var _ , jdbcErr = ormService.Jdbc(func(o orm.Ormer) (interface{}, error) {

		//这里 之后要加上对应的 验证是否已经登录的信息

		var readLoginErr = o.Read(&login , "account","password")

		if readLoginErr != nil {
			//查询错误 ，
			return nil , readLoginErr
		}

		//再之后 ， 我们便可以获取到对应的 user 的 信息

		user.Id = login.Id;

		var readUserErr = o.Read( &user )

		if readUserErr != nil {
			return nil , readUserErr
		}

		return &user , nil
	})

	//倘若jdbc 出现错误 ， 则进行替换
	if jdbcErr != nil {
		self.FailJson( jdbcErr )
		return
	}

	self.StartSession().Set(SESSION_LOCAL_USER_KEY , user)

	//之后返回结果
	self.Json( user )
}

/**
 修改密码
*/
func ( self *LoginController ) UpdatePwd(){

	var id , getIdErr = self.GetInt64("id")

	if getIdErr != nil {
		self.FailJson( getIdErr )
		return
	}

	var password = self.GetString("password")

	var password_len = len(password)

	if password_len <= 6 {
		self.FailJson(errors.New("密码不符合格式"))
		return
	}

	var ormService = db.GetOrmServiceInstance()
	var loginService = service.GetLoginServiceInstance()

	var login *bean.Login
	var _ , transacErr = ormService.Transaction(func(o orm.Ormer) (interface{}, error) {
		var readErr error

		login , readErr = loginService.ReadLogin( o , id )
		if readErr != nil {
			return nil , readErr
		}

		login.Password = password

		var _ , updateErr = loginService.UpdateLogin( o , login )

		return login , updateErr
	})

	if transacErr != nil {
		self.FailJson( transacErr )
		return
	}

	self.Json( login )
}

//显示当前的 Session 中的 状态
func ( self *LoginController ) CurrentUser(){
	//
	var session = self.StartSession()
	var user = session.Get( SESSION_LOCAL_USER_KEY )

	//是否
	if user == nil {
		self.FailJson(errors.New(" user is nil "))
		return
	}

	self.Json( user )
}

//对应的 Login 页面
func (self *LoginController ) LoginPage(){
	self.Resource("mit/login.html")
}

//对应 注册页面
func (self *LoginController ) RegisterPage(){
	//开始设置对应的信息
	self.Resource("mit/register.html")
}
//对应的政策页面
func (self *LoginController) PolicePage(){
	self.String("hello , world")
}