// GameInfo 游戏列表
type GameInfo struct {
ID                   int32  `json:"id"`
CGGameID             int32  `json:"cggameid"`
AppID                int64  `json:"appid"`
OpenAppID            int64  `json:"openappid"`
WxAppID              string `json:"wxappid"`
ScreenDirection      int32  `json:"screenDirection"`
GameType             int32  `json:"gameType"`
DistributeType       int32  `json:"distributeType"`
PkgName              string `json:"pkgname"`
GameName             string `json:"gamename"`
GamePlatformType     int32  `json:"gamePlatformType"` // 游戏平台类型（0:手机游戏，1:PC游戏）
CloudGameTag         string `json:"cloudgametag"`
CloudGameVideoMid    int32  `json:"cloudgamevideomid"`
CloudGameVideoVid    string `json:"cloudgamevideovid"`
CloudGameVideoCover  string `json:"cloudgamevideocover"`
CloudGameActivityUrl string `json:"cloudgameactivityurl"`
NeedPaytoken         int32  `json:"needPaytoken"` // 是否支持云游支付（0不支持，1支持）
/* 以下字段不在dcache中 */
FirstFrameDelay int32 `json:"firstFrameDelay"` // 延迟出现首帧的时间（单位，毫秒）
}

// GameTrainInfo 试玩信息
type GameTrainInfo struct {
ID       int32 `json:"id"`
CGGameID int32 `json:"cggameid"`
TipsList []struct {
ID   int32  `json:"id"`
Tips string `json:"tips"`
} `json:"tipsList"`
TopText      string `json:"topText"`
BottomText   string `json:"bottomText"`
PicURL       string `json:"picurl"`
Provider     string `json:"provider"`
ProviderLogo string `json:"providerLogo"`
MissTips     string `json:"missTips"`
}

// PluginEntranceInfo 插件入口配置
type PluginEntranceInfo struct {
ID                   int32  `json:"id"`
Provider             string `json:"provider"`
CGGameID             int32  `json:"cggameid"`
EntranceID           int32  `json:"entranceid"`
PkgName              string `json:"pkgname"`
EndTime              int64  `json:"endTime"`
BeginTime            int64  `json:"beginTime"` // 毫秒级
CloseTips            string `json:"closeTips"`
IsReal               int32  `json:"isReal"`
IsWhite              bool   `json:"isWhite"`
NoLogin              int32  `json:"nologin"`
MaxTrainTime         int32  `json:"maxTrainTime"`
AntiAddictionType    int32  `json:"antiAddictionType"`
ModScheme            string `json:"modScheme"`
Tag                  string `json:"tag"`
Priority             int32  `json:"cloudgamepriority"`
Icon                 string `json:"icon"`
SupportLoginPlatform int32  `json:"supportLoginPlatform"` // 云游戏支持登录平台
WetestInfo           struct {
Restart      string `json:"restart"`
MaxTrainTime int32  `json:"maxTrainTime"`
WetestGameID string `json:"wetestGameID"`
} `json:"wetestInfo"`
StartInfo struct {
GameID     string `json:"gameid"`
PCGameInfo string `json:"pcgametag"`
} `json:"startInfo"`
MidgameInfo struct {
IsMidgame       bool   `json:"isMidgame"`
MidgameShowText string `json:"midgameShowText"`
} `json:"midgameInfo"`
}

// AuthHostInfo 宿主信息管理
type AuthHostInfo struct {
OpenAppid string `json:"openappid"`
GameName  string `json:"gameName"`
Sign      string `json:"sign"`
}

// GameShareInfo 试玩分享信息
type GameShareInfo struct {
ID           int32  `json:"id"`
CGGameID     int32  `json:"cggameid"`
MainTitle    string `json:"shareTitle"`  // 主标题
SubTitle     string `json:"shareText"`   // 副标题
ShareUrl     string `json:"shareUrl"`    // 分享链接
QrcodeUrl    string `json:"shareSQCode"` // 分享二维码地址
ActivityID   string `json:"sharePicActId"`
IconUrl      string `json:"shareLinkIcon"`
ShareContent string `json:"shareLinkContent"`
}