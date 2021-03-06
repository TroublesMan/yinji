package service

import (
	"github.com/astaxie/beego/orm"
	"yinji/models/base"
	"yinji/models/bean"
	"github.com/astaxie/beego/logs"
	"yinji/service/db"
)

type AudioDashboardService struct {

}

const (
	AUDIO_DASHBOARD_C0LLECTION = 1
	AUDIO_DASHBOARD_FAVORITE = 3
	AUDIO_DASHBOARD_FORWARD = 5
)

//暂时的操作
func ( self *AudioDashboardService ) AddDashboardCountEvent( o orm.Ormer , id int64 , dashboardType int , count int64 ){
	var dashboardBase = &base.DashboardBase{}
	if dashboardType == AUDIO_DASHBOARD_C0LLECTION {
		dashboardBase.CollectionCount = count
	}else if dashboardType == AUDIO_DASHBOARD_FAVORITE {
		dashboardBase.LoveCount = count
	}else if dashboardType == AUDIO_DASHBOARD_FORWARD {
		dashboardBase.ForwardCount = count
	}else if dashboardType == AUDIO_DASHBOARD_C0LLECTION + AUDIO_DASHBOARD_FAVORITE {
		dashboardBase.CollectionCount = count
		dashboardBase.LoveCount = count
	}else if dashboardType == AUDIO_DASHBOARD_C0LLECTION + AUDIO_DASHBOARD_FORWARD {
		dashboardBase.CollectionCount = count
		dashboardBase.ForwardCount = count
	}else if dashboardType == AUDIO_DASHBOARD_FAVORITE + AUDIO_DASHBOARD_FORWARD {
		dashboardBase.LoveCount = count
		dashboardBase.ForwardCount = count
	}else if dashboardType == AUDIO_DASHBOARD_C0LLECTION + AUDIO_DASHBOARD_FAVORITE + AUDIO_DASHBOARD_FORWARD {
		dashboardBase.CollectionCount = count
		dashboardBase.LoveCount = count
		dashboardBase.ForwardCount = count
	}
}


func ( self *AudioDashboardService ) AddCollectionCount( o orm.Ormer , id int64 , count int64 ){
	var dashboradBase = &base.DashboardBase{}
	dashboradBase.CollectionCount = count
	self.AddDashboradCount( o , id , dashboradBase )
}

func ( self *AudioDashboardService ) AddLoveCount( o orm.Ormer , id int64 , count int64 ){
	var dashboradBase = &base.DashboardBase{}
	dashboradBase.LoveCount = count
	self.AddDashboradCount( o , id , dashboradBase )
}

func ( self *AudioDashboardService ) AddForwardCount( o orm.Ormer , id int64 , count int64 ){
	var dashboradBase = &base.DashboardBase{}
	dashboradBase.ForwardCount = count
	self.AddDashboradCount( o , id , dashboradBase )
}

/**
	根据目标的信息
*/
func ( self *AudioDashboardService) AddDashboradCount( o orm.Ormer ,  id int64 ,  base *base.DashboardBase ) ( *bean.AudioDashboard , error){
	//根据对应的 id 来获取数据
	var dashboard , findErr = self.FindById( o , id )

	if findErr != nil {
		return nil , findErr
	}

	dashboard.Add( base )
	dashboard.Refresh()
	var _ , updateErr = o.Update( dashboard )

	return dashboard , updateErr
}

/**
	意义:倘若没有找到 ， 则进行自动更新
	接受的参数 : id 目标audio 音频的声音
 */
func ( self *AudioDashboardService ) FindOrNew( o orm.Ormer , id int64) ( *bean.AudioDashboard , error) {
	var dashboard , findErr = self.FindById(o , id );
	if findErr != nil {
		dashboard , findErr = self.NewDashboard(o,id)
	}

	return dashboard , findErr
}

//根据对应的id 来进行搜索
func ( self *AudioDashboardService) FindById( o orm.Ormer , id int64 ) ( *bean.AudioDashboard , error){
	var dashboard = &bean.AudioDashboard{}
	dashboard.Id = id
	var readErr = o.Read( dashboard )
	return dashboard , readErr
}

//由於主鍵id 策略發生的錯誤
func ( self *AudioDashboardService) NewDashboard( o orm.Ormer  , audioId int64 ) (*bean.AudioDashboard, error) {
	var dashboard = &bean.AudioDashboard{}
	dashboard.New()
	dashboard.Id = audioId
	var _ , insertErr = o.Insert(dashboard)
	return dashboard , insertErr
}

func ( self *AudioDashboardService) NewByAudioId( o orm.Ormer , audioId int64 ) ( *map[int] *bean.AudioDashboard, error ){
	//制定对应的数组 ， 按照数组来进行操作
	var dashboardTypeArray = [3]int{ AUDIO_DASHBOARD_C0LLECTION , AUDIO_DASHBOARD_FAVORITE , AUDIO_DASHBOARD_FORWARD }
	var dashboardMap = make(map[int] *bean.AudioDashboard)
	logs.SetLogger("console")
	logs.Debug("default")
	for dashboardType := range dashboardTypeArray {
		var dashboard , newErr = self.NewDashboard( o ,audioId)

		logs.Debug("dashborad %i",dashboardType)
		//有问题是否会使整个函数失败呢？？ ， 答 ， 不会使整个操作失败
		if newErr != nil {

			logs.Debug(newErr.Error())
			continue
		}
		dashboardMap[dashboardType] = dashboard
	}

	return &dashboardMap , nil
}

/**
	根据对应的 audioId 来搜索对应的长度
*/
func ( self *AudioDashboardService ) SearchDashboard(o orm.Ormer , audioId int64 , limit int) ( *[]*bean.AudioDashboard , error ){
	var array *[]*bean.AudioDashboard
	var qt = o.QueryTable( bean.GetAudioDashboardTableName() )
	// 按照 播放次数（ browseCount ) 来进行倒置输出
	qt = qt.OrderBy("-browseCount").Limit( limit )
	var _ , allErr = qt.All( array )
	return array , allErr
}

/**
	数据库上对应的 修改数据的方法
 */
func ( self *AudioDashboardService ) UpdateDashboardCount( o orm.Ormer , id int64 , dashboardBase *base.DashboardBase) (*bean.AudioDashboard , error){

	/**
		现根据对应的信息 ， 来获取对应的信息
	 */

	 var dashboard , findErr = self.FindOrNew(o , id)

	 if findErr != nil{
	 	//倘若没有则开始添加
	 	return nil , findErr
	 }

	 dashboard.Refresh()
	 dashboard.Add(dashboardBase)

	 var _ , updateErr = o.Update(dashboard)
	 return dashboard , updateErr

}

/**
	真正意义上使用jdbc的方法
 */
func ( self *AudioDashboardService ) AddDashboardCount( id int64 , dashboardBase *base.DashboardBase ){
	var ormService = db.GetOrmServiceInstance()
	ormService.Transaction(func(o orm.Ormer) (interface{}, error) {
		return self.UpdateDashboardCount(o , id , dashboardBase)
	});

}

//对应的数据
var DASHBOARD_SERVICE_INSTANCE *AudioDashboardService = &AudioDashboardService{}

func GetAudioDashboardServiceInstance() *AudioDashboardService {
	return DASHBOARD_SERVICE_INSTANCE
}