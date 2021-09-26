// 此文件为sdp2go工具自动生成,请不要手动编辑

package rpc

type ErrorCode int32

const (
	Error_Unknown                         = -1
	Error_InvalidParam                    = -2
	Error_Protocol                        = -3
	Error_NoParser                        = -4
	Error_NoAuth                          = -5
	Error_NoAccountId                     = -6
	Error_NoAccountName                   = -7
	Error_SessionExpired                  = -8
	Error_ZoneMaintain                    = -9
	Error_DangerGMLimit                   = -10
	Error_Account_NameExist               = -11
	Error_LoginDatabase                   = -12
	Error_ChannelExist                    = -13
	Error_ChannelNotExist                 = -14
	Error_ChannelAddrError                = -15
	Error_Account_Bound                   = -16
	Error_Account_Bind_Auth               = -17
	Error_Cmd_SendStoreResp               = -18
	Error_APNS_Service                    = -19
	Error_Client_TooOld                   = -20
	Error_InvalidVersion                  = -21
	Error_InvalidZoneId                   = -100
	Error_ZoneNotExist                    = -101
	Error_ZoneHasExist                    = -102
	Error_DirDatabase                     = -103
	Error_DirServer                       = -104
	Error_VersionSmall                    = -105
	Error_VersionBigger                   = -106
	Error_Bulletin_InvalidId              = -121
	Error_Bulletin_NotExist               = -122
	Error_Bulletin_HasExist               = -123
	Error_Bulletin_Database               = -124
	Error_Bulletin_Server                 = -125
	Error_Bulletin_InvalidTime            = -126
	Error_Notice_InvalidId                = -127
	Error_Notice_NotExist                 = -128
	Error_Notice_HasExist                 = -129
	Error_Notice_InvalidTime              = -130
	Error_ConnIdNull                      = -201
	Error_ConnIdExist                     = -202
	Error_ConnNotInManager                = -203
	Error_ZoneIdMismatch                  = -204
	Error_ZonePublishTime                 = -205
	Error_MP_NoProject                    = -301
	Error_MP_NoCDKeyId                    = -302
	Error_MP_CDKeyRunOut                  = -305
	Error_MP_CDKeyGenerating              = -306
	Error_MP_ZoneList                     = -307
	Error_Blob_ShmGet                     = -401
	Error_Blob_ShmSet                     = -402
	Error_Role_InitDb                     = -403
	Error_Role_ShmInit                    = -404
	Error_Mail_InitDb                     = -405
	Error_Mail_ShmInit                    = -406
	Error_FightReport_InitDb              = -407
	Error_FightReport_ShmInit             = -408
	Error_Mail_InvalidId                  = -501
	Error_Mail_Database                   = -502
	Error_Mail_Server                     = -503
	Error_Mail_NotPointZone               = -504
	Error_Google_Service                  = -601
	Error_GameCenter_Service              = -602
	Error_Facebook_Service                = -603
	Error_Fy_Service                      = -604
	Error_Forward_InvalidDstNull          = -701
	Error_Forward_InvalidDstSet           = -702
	Error_Forward_InvalidDivision         = -703
	Error_Forward_InvalidToUid            = -704
	Error_IAP_NoSuchReceipt               = -801
	Error_IAP_DuplicateReceipt            = -802
	Error_IAP_Receipt_Expired             = -803
	Error_IAP_Verify_Fail                 = -804
	Error_IAP_Apple_Verify_Network        = -805
	Error_IAP_Apple_BundleId_Mismatch     = -806
	Error_IAP_Google_PurchaseFail         = -807
	Error_IAP_Google_PurchaseState        = -808
	Error_IAP_Google_PurchaseData         = -809
	Error_IAP_Google_PackageName_Mismatch = -810
	Error_IAP_FY_PurchaseData             = -811
	Error_IAP_FY_Sign                     = -812
	Error_IAP_FY_AMOUNT                   = -813
	Error_IAP_AB_PurchaseData             = -814
	Error_IAP_AB_PurchaseState            = -815
	Error_IAP_AB_APPID_Mismatch           = -816
	Error_IAP_AB_SignTransData            = -817
	Error_IAP_HeePay_AMOUNT               = -818
	Error_IAP_HeePay_CreateOrder          = -819
	Error_IAP_HeePay_PurchaseData         = -820
	Error_IAP_HeePay_Sign                 = -821
	Error_IAP_HeePay_Verify_Network       = -822
	Error_IAP_HeePay_AgentId              = -823
	Error_IAP_HeePay_QueryDismatch        = -824
	Error_IAP_HeePay_HasOrderToPay        = -825
	Error_IAP_HeePayH5_CreateOrder        = -826
	Error_IAP_HeePayH5_PurchaseData       = -827
	Error_IAP_HeePayH5_Sign               = -828
	Error_IAP_HeePayH5_AppId              = -829
	Error_IAP_HeePayH5_MchId              = -830
	Error_IAP_HeePayH5_AMOUNT             = -831
	Error_GCM_Service                     = -1001
	Error_GCM_PushFail                    = -1002
	Error_UPUSH_Service                   = -1003
	Error_Gm_Param                        = 1001
	Error_Gm_NoParser                     = 1002
	Error_Role_Null                       = 1003
	Error_Sys_NoCmdParser                 = 1004
	Error_Gm_Execute                      = 1005
	Error_Cmd_Execute                     = 1006
	Error_TableNoData                     = 1007
	Error_NoBagSpace                      = 1008
	Error_ItemNotEnough                   = 1009
	Error_NoSuchItem                      = 1010
	Error_InvalidItem                     = 1011
	Error_ZoneIdExceedLimit               = 1012
	Error_GUIDAllocExceedLimit            = 1013
	Error_CoinNotEnough                   = 1014
	Error_DiamondNotEnough                = 1015
	Error_KingCoinNotEnough               = 1016
	Error_ItemCannotSell                  = 1017
	Error_ItemCannotUse                   = 1018
	Error_SkillPointNotEnough             = 1019
	Error_RoleNameEmpty                   = 1020
	Error_VipLevelNotEnough               = 1021
	Error_Role_BanLogin                   = 1022
	Error_Role_BanSpeak                   = 1023
	Error_Module_Closed                   = 1024
	Error_Sys_CmdClosed                   = 1025
	Error_TalentNotEnough                 = 1026
	Error_CannotBindMultiple              = 1027
	Error_Role_HasNoFaceFrame             = 1028
	Error_Role_FaceFrameLock              = 1029
	Error_Role_HasNoFace                  = 1030
	Error_Hero_HasBeenAdd                 = 1031
	Error_Hero_NotInList                  = 1032
	Error_Hero_StarTooLarge               = 1033
	Error_Hero_StepTooLarge               = 1034
	Error_Hero_StarBeenMax                = 1035
	Error_Hero_StepBeenMax                = 1036
	Error_Hero_StepUpgrade                = 1037
	Error_Hero_SkillNoLearned             = 1038
	Error_Hero_SkillBeenMax               = 1039
	Error_Hero_SoulCanNotHero             = 1040
	Error_Hero_EquipPosAlready            = 1041
	Error_Hero_EquipInvalidPos            = 1042
	Error_Hero_StepLess                   = 1043
	Error_Hero_LevelNotEnough             = 1044
	Error_Hero_TalentBeenMax              = 1045
	Error_Hero_RuneLevelBeenMax           = 1046
	Error_RuneNotEnough                   = 1047
	Error_Hero_SkillHasLearned            = 1048
	Error_ZoonExpNotEnough                = 1049
	Error_ZoonName_LengthLimit            = 1050
	Error_Stage_Cheat                     = 1055
	Error_Stage_InvalidTeam               = 1056
	Error_Stage_MaxTeam                   = 1057
	Error_Market_NotOpen                  = 1058
	Error_Market_HasSold                  = 1059
	Error_Market_RefreshTime              = 1060
	Error_Market_VipMaxTimes              = 1061
	Error_Arena_NoReport                  = 1062
	Error_Arena_NotOpen                   = 1063
	Error_ArenaCoinNotEnough              = 1067
	Error_Mail_Role_Null                  = 1068
	Error_Mail_NotExist                   = 1069
	Error_Mail_NoAttach                   = 1070
	Error_Mail_LengthLimited              = 1071
	Error_Mail_NoMoreMail                 = 1072
	Error_Mail_MailAttachNotRcv           = 1073
	Error_Mail_RcvUserNotExit             = 1074
	Error_Mail_SendToMyself               = 1075
	Error_Name_RoleNameExist              = 1076
	Error_Name_LengthLimit                = 1077
	Error_HasDirty                        = 1078
	Error_InvalidChar                     = 1079
	Error_Chat_InCD                       = 1080
	Error_Chat_NotOpen                    = 1081
	Error_Chat_LevelLimit                 = 1082
	Error_Chat_NotSelf                    = 1083
	Error_Chat_Channel                    = 1084
	Error_Chat_MessageEmpty               = 1085
	Error_Chat_MessageType                = 1086
	Error_Chat_NotInShieldedList          = 1088
	Error_Chat_WasInShieldedList          = 1089
	Error_Item_BeyondMaxNum               = 1090
	Error_Activity_NotExists              = 1091
	Error_Activity_NotJoin                = 1092
	Error_Activity_Condition              = 1093
	Error_Activity_HideOnFinish           = 1094
	Error_Activity_HasReward              = 1095
	Error_Activity_Close                  = 1096
	Error_Talent_CoinBuy                  = 1097
	Error_Talent_DiamondBuy               = 1098
	Error_Talent_LevelBeenMax             = 1099
	Error_Home_Null                       = 1101
	Error_Home_NotJoin                    = 1104
	Error_Home_HasJoin                    = 1105
	Error_Home_NotMaster                  = 1107
	Error_Home_NotInMemberList            = 1108
	Error_Home_MaxMember                  = 1110
	Error_Home_Exist                      = 1111
	Error_Home_SameName                   = 1114
	Error_Home_HasApply                   = 1115
	Error_Home_MaxApply                   = 1116
	Error_Home_NotInApplyList             = 1118
	Error_Home_HasMembers                 = 1122
	Error_AdvMap_PosHasHome               = 1123
	Error_Home_CannotJoinAnyHome          = 1127
	Error_AdvMap_Null                     = 1128
	Error_AdvMap_Exist                    = 1129
	Error_AdvMap_HasRewarded              = 1130
	Error_AdvMap_PosNoOwner               = 1131
	Error_AdvMap_Fighting                 = 1132
	Error_AdvMap_Refresh                  = 1133
	Error_Chat_SelectNoOne                = 1134
	Error_AdvMap_HeroHasBorrowed          = 1135
	Error_AdvMap_HeroNotSelf              = 1136
	Error_FightReport_Invalid             = 1137
	Error_FightReport_ShareType           = 1138
	Error_FightReport_ShareSelf           = 1139
	Error_FightReport_MessageCharError    = 1140
	Error_CDKey_ZoneList                  = 1141
	Error_MP_NoCDKey                      = 1142
	Error_MP_CDKeyExpired                 = 1143
	Error_MP_InvalidCDKey                 = 1144
	Error_MP_CDKeyExchanged               = 1145
	Error_MP_ExchangeBindLimited          = 1146
	Error_MP_ExchangeNumLimited           = 1147
	Error_MP_InvalidDivision              = 1148
	Error_MP_LimitConfigError             = 1149
	Error_MP_InDelivering                 = 1150
	Error_MP_CDKeyNotOpen                 = 1151
	Error_Hero_GemInvalidSlot             = 1152
	Error_Hero_GemMaxLevel                = 1153
	Error_Gem_NotEnough                   = 1154
	Error_King_NotOpen                    = 1156
	Error_Quest_Reward                    = 1157
	Error_Quest_Step                      = 1158
	Error_Quest_VitalityReward            = 1159
	Error_Achieve_Step                    = 1160
	Error_HeroExp_NotEnough               = 1161
	Error_Arena_WinTimesReward            = 1162
	Error_Arena_WinTimesNotEnough         = 1163
	Error_ChallengeTimes_NotZero          = 1164
	Error_Chat_ShieldListFull             = 1165
	Error_Chat_HasBeenShield              = 1166
	Error_HeroTrialTimesNotEnough         = 1167
	Error_Arena_ChallengeTimesNotEnough   = 1168
	Error_Fight_isFighting                = 1169
	Error_King_ChallengeTimesNotEnough    = 1170
	Error_King_BeAttacking                = 1171
	Error_Recharge_BuyNumBeenMax          = 1172
	Error_Talent_NotOpen                  = 1173
	Error_Trial_NotOpen                   = 1174
	Error_SignIn_HasReward                = 1175
	Error_SignIn_NotReward                = 1176
	Error_SignIn_HasDouble                = 1177
	Error_Vip_HasReward                   = 1178
	Error_Stage_GamesHasReward            = 1179
	Error_Pet_Null                        = 1180
	Error_Stage_VerifyFull                = 1181
	Error_TrialCoinNotEnough              = 1182
	Error_HomeCoinNotEnough               = 1183
	Error_CrowdFund_Complete              = 1184
	Error_Rune_NotExists                  = 1185
	Error_Rune_Lock                       = 1186
	Error_ActBind_BindMe                  = 1187
	Error_AdvWheel_TimesNotEnough         = 1188
	Error_CrowdFund_Recharging            = 1189
	Error_BranchStage_GameOver            = 1190
	Error_AdvMap_Injuring                 = 1191
	Error_Item_OnlyUseOne                 = 1192
	Error_Item_EffectExist                = 1193
	Error_Pet_MaxSlotSize                 = 1194
	Error_Hero_Syncing                    = 1195
	Error_Zoon_CancelMate                 = 1196
	Error_ZoonSlotNotEnough               = 1197
	Error_Obj_NotExists                   = 10000
	Error_Role_Reborn                     = 10001
	Error_Map_RandPos                     = 10002
	Error_Slg_NotOpen                     = 10003
	Error_Map_InvalidPos                  = 10004
	Error_Pos_HasObj                      = 10005
	Error_Pos_HasNoObj                    = 10006
	Error_FoodNotEnough                   = 10007
	Error_WoodNotEnough                   = 10008
	Error_StoneNotEnough                  = 10009
	Error_SilverNotEnough                 = 10010
	Error_Pos_SendTeamAlready             = 10012
	Error_Slg_ReportInvalid               = 10021
	Error_Map_ResNoOwner                  = 10022
	Error_Pos_Forbid                      = 10027
	Error_DungeonTeam_Null                = 10028
	Error_DungeonTeam_NotJoin             = 10029
	Error_DungeonTeam_HasJoin             = 10030
	Error_DungeonTeam_JoinCode            = 10031
	Error_DungeonTeam_FullMember          = 10032
	Error_DungeonTeam_Power               = 10033
	Error_City_DoorBroken                 = 10034
	Error_StorePos_Name                   = 10035
	Error_StorePos_MaxNum                 = 10036
	Error_Wand_Max                        = 10037
	Error_WandNotEnough                   = 10038
	Error_WorkShopLevelNotEnough          = 10039
	Error_SendCacheRsp                    = 10040
	Error_Summon_MaxTimes                 = 10041
	Error_Summon_CD                       = 10042
	Error_Alliance_NotJoin                = 10043
	Error_Alliance_Null                   = 10044
	Error_Alliance_NotMaster              = 10045
	Error_Alliance_HasJoin                = 10046
	Error_Alliance_NotInApplyList         = 10047
	Error_Alliance_CannotJoinAnyAlliance  = 10048
	Error_Alliance_MaxApply               = 10049
	Error_Alliance_HasApply               = 10050
	Error_Alliance_SameName               = 10051
	Error_Alliance_MaxMember              = 10052
	Error_Alliance_NotInMemberList        = 10053
	Error_Alliance_HasMembers             = 10054
	Error_Alliance_Exist                  = 10055
	Error_Alliance_MaxKickTimes           = 10056
	Error_Alliance_PostMaxSize            = 10057
	Error_AllianceCoinNotEnough           = 10058
	Error_AllianceBoss_NotOpen            = 10059
	Error_AllianceBoss_HasReward          = 10060
	Error_NotMyMonster                    = 10061
	Error_Hero_QualityBeenMax             = 10062
	Error_Hero_QualityLess                = 10063
	Error_GemBagSize                      = 10064
	Error_Slg_NotNewbee                   = 10065
	Error_Alliance_NotSendReward          = 10066
	Error_WorkShopNotLevel                = 10067
	Error_Slg_FlagHasActive               = 10068
	Error_Slg_FlagNotSameAlliance         = 10069
	Error_Slg_InFlagNoProtect             = 10070
	Error_Slg_FlagInProtect               = 10071
	Error_Slg_NoCityNoUseGem              = 10072
	Error_Obj_NotInVision                 = 10073
	Error_City_CanotMove                  = 10074
	Error_Map_TeamQuality                 = 10075
	Error_RoleMap_Null                    = 10076
	Error_Flag_AllianceOnlyOne            = 10077
	Error_Alliance_PostNotEnough          = 10078
)

func (en ErrorCode) String() string {
	ret := ""
	switch en {
	case Error_Unknown:
		ret = "Error_Unknown"
	case Error_InvalidParam:
		ret = "Error_InvalidParam"
	case Error_Protocol:
		ret = "Error_Protocol"
	case Error_NoParser:
		ret = "Error_NoParser"
	case Error_NoAuth:
		ret = "Error_NoAuth"
	case Error_NoAccountId:
		ret = "Error_NoAccountId"
	case Error_NoAccountName:
		ret = "Error_NoAccountName"
	case Error_SessionExpired:
		ret = "Error_SessionExpired"
	case Error_ZoneMaintain:
		ret = "Error_ZoneMaintain"
	case Error_DangerGMLimit:
		ret = "Error_DangerGMLimit"
	case Error_Account_NameExist:
		ret = "Error_Account_NameExist"
	case Error_LoginDatabase:
		ret = "Error_LoginDatabase"
	case Error_ChannelExist:
		ret = "Error_ChannelExist"
	case Error_ChannelNotExist:
		ret = "Error_ChannelNotExist"
	case Error_ChannelAddrError:
		ret = "Error_ChannelAddrError"
	case Error_Account_Bound:
		ret = "Error_Account_Bound"
	case Error_Account_Bind_Auth:
		ret = "Error_Account_Bind_Auth"
	case Error_Cmd_SendStoreResp:
		ret = "Error_Cmd_SendStoreResp"
	case Error_APNS_Service:
		ret = "Error_APNS_Service"
	case Error_Client_TooOld:
		ret = "Error_Client_TooOld"
	case Error_InvalidVersion:
		ret = "Error_InvalidVersion"
	case Error_InvalidZoneId:
		ret = "Error_InvalidZoneId"
	case Error_ZoneNotExist:
		ret = "Error_ZoneNotExist"
	case Error_ZoneHasExist:
		ret = "Error_ZoneHasExist"
	case Error_DirDatabase:
		ret = "Error_DirDatabase"
	case Error_DirServer:
		ret = "Error_DirServer"
	case Error_VersionSmall:
		ret = "Error_VersionSmall"
	case Error_VersionBigger:
		ret = "Error_VersionBigger"
	case Error_Bulletin_InvalidId:
		ret = "Error_Bulletin_InvalidId"
	case Error_Bulletin_NotExist:
		ret = "Error_Bulletin_NotExist"
	case Error_Bulletin_HasExist:
		ret = "Error_Bulletin_HasExist"
	case Error_Bulletin_Database:
		ret = "Error_Bulletin_Database"
	case Error_Bulletin_Server:
		ret = "Error_Bulletin_Server"
	case Error_Bulletin_InvalidTime:
		ret = "Error_Bulletin_InvalidTime"
	case Error_Notice_InvalidId:
		ret = "Error_Notice_InvalidId"
	case Error_Notice_NotExist:
		ret = "Error_Notice_NotExist"
	case Error_Notice_HasExist:
		ret = "Error_Notice_HasExist"
	case Error_Notice_InvalidTime:
		ret = "Error_Notice_InvalidTime"
	case Error_ConnIdNull:
		ret = "Error_ConnIdNull"
	case Error_ConnIdExist:
		ret = "Error_ConnIdExist"
	case Error_ConnNotInManager:
		ret = "Error_ConnNotInManager"
	case Error_ZoneIdMismatch:
		ret = "Error_ZoneIdMismatch"
	case Error_ZonePublishTime:
		ret = "Error_ZonePublishTime"
	case Error_MP_NoProject:
		ret = "Error_MP_NoProject"
	case Error_MP_NoCDKeyId:
		ret = "Error_MP_NoCDKeyId"
	case Error_MP_CDKeyRunOut:
		ret = "Error_MP_CDKeyRunOut"
	case Error_MP_CDKeyGenerating:
		ret = "Error_MP_CDKeyGenerating"
	case Error_MP_ZoneList:
		ret = "Error_MP_ZoneList"
	case Error_Blob_ShmGet:
		ret = "Error_Blob_ShmGet"
	case Error_Blob_ShmSet:
		ret = "Error_Blob_ShmSet"
	case Error_Role_InitDb:
		ret = "Error_Role_InitDb"
	case Error_Role_ShmInit:
		ret = "Error_Role_ShmInit"
	case Error_Mail_InitDb:
		ret = "Error_Mail_InitDb"
	case Error_Mail_ShmInit:
		ret = "Error_Mail_ShmInit"
	case Error_FightReport_InitDb:
		ret = "Error_FightReport_InitDb"
	case Error_FightReport_ShmInit:
		ret = "Error_FightReport_ShmInit"
	case Error_Mail_InvalidId:
		ret = "Error_Mail_InvalidId"
	case Error_Mail_Database:
		ret = "Error_Mail_Database"
	case Error_Mail_Server:
		ret = "Error_Mail_Server"
	case Error_Mail_NotPointZone:
		ret = "Error_Mail_NotPointZone"
	case Error_Google_Service:
		ret = "Error_Google_Service"
	case Error_GameCenter_Service:
		ret = "Error_GameCenter_Service"
	case Error_Facebook_Service:
		ret = "Error_Facebook_Service"
	case Error_Fy_Service:
		ret = "Error_Fy_Service"
	case Error_Forward_InvalidDstNull:
		ret = "Error_Forward_InvalidDstNull"
	case Error_Forward_InvalidDstSet:
		ret = "Error_Forward_InvalidDstSet"
	case Error_Forward_InvalidDivision:
		ret = "Error_Forward_InvalidDivision"
	case Error_Forward_InvalidToUid:
		ret = "Error_Forward_InvalidToUid"
	case Error_IAP_NoSuchReceipt:
		ret = "Error_IAP_NoSuchReceipt"
	case Error_IAP_DuplicateReceipt:
		ret = "Error_IAP_DuplicateReceipt"
	case Error_IAP_Receipt_Expired:
		ret = "Error_IAP_Receipt_Expired"
	case Error_IAP_Verify_Fail:
		ret = "Error_IAP_Verify_Fail"
	case Error_IAP_Apple_Verify_Network:
		ret = "Error_IAP_Apple_Verify_Network"
	case Error_IAP_Apple_BundleId_Mismatch:
		ret = "Error_IAP_Apple_BundleId_Mismatch"
	case Error_IAP_Google_PurchaseFail:
		ret = "Error_IAP_Google_PurchaseFail"
	case Error_IAP_Google_PurchaseState:
		ret = "Error_IAP_Google_PurchaseState"
	case Error_IAP_Google_PurchaseData:
		ret = "Error_IAP_Google_PurchaseData"
	case Error_IAP_Google_PackageName_Mismatch:
		ret = "Error_IAP_Google_PackageName_Mismatch"
	case Error_IAP_FY_PurchaseData:
		ret = "Error_IAP_FY_PurchaseData"
	case Error_IAP_FY_Sign:
		ret = "Error_IAP_FY_Sign"
	case Error_IAP_FY_AMOUNT:
		ret = "Error_IAP_FY_AMOUNT"
	case Error_IAP_AB_PurchaseData:
		ret = "Error_IAP_AB_PurchaseData"
	case Error_IAP_AB_PurchaseState:
		ret = "Error_IAP_AB_PurchaseState"
	case Error_IAP_AB_APPID_Mismatch:
		ret = "Error_IAP_AB_APPID_Mismatch"
	case Error_IAP_AB_SignTransData:
		ret = "Error_IAP_AB_SignTransData"
	case Error_IAP_HeePay_AMOUNT:
		ret = "Error_IAP_HeePay_AMOUNT"
	case Error_IAP_HeePay_CreateOrder:
		ret = "Error_IAP_HeePay_CreateOrder"
	case Error_IAP_HeePay_PurchaseData:
		ret = "Error_IAP_HeePay_PurchaseData"
	case Error_IAP_HeePay_Sign:
		ret = "Error_IAP_HeePay_Sign"
	case Error_IAP_HeePay_Verify_Network:
		ret = "Error_IAP_HeePay_Verify_Network"
	case Error_IAP_HeePay_AgentId:
		ret = "Error_IAP_HeePay_AgentId"
	case Error_IAP_HeePay_QueryDismatch:
		ret = "Error_IAP_HeePay_QueryDismatch"
	case Error_IAP_HeePay_HasOrderToPay:
		ret = "Error_IAP_HeePay_HasOrderToPay"
	case Error_IAP_HeePayH5_CreateOrder:
		ret = "Error_IAP_HeePayH5_CreateOrder"
	case Error_IAP_HeePayH5_PurchaseData:
		ret = "Error_IAP_HeePayH5_PurchaseData"
	case Error_IAP_HeePayH5_Sign:
		ret = "Error_IAP_HeePayH5_Sign"
	case Error_IAP_HeePayH5_AppId:
		ret = "Error_IAP_HeePayH5_AppId"
	case Error_IAP_HeePayH5_MchId:
		ret = "Error_IAP_HeePayH5_MchId"
	case Error_IAP_HeePayH5_AMOUNT:
		ret = "Error_IAP_HeePayH5_AMOUNT"
	case Error_GCM_Service:
		ret = "Error_GCM_Service"
	case Error_GCM_PushFail:
		ret = "Error_GCM_PushFail"
	case Error_UPUSH_Service:
		ret = "Error_UPUSH_Service"
	case Error_Gm_Param:
		ret = "Error_Gm_Param"
	case Error_Gm_NoParser:
		ret = "Error_Gm_NoParser"
	case Error_Role_Null:
		ret = "Error_Role_Null"
	case Error_Sys_NoCmdParser:
		ret = "Error_Sys_NoCmdParser"
	case Error_Gm_Execute:
		ret = "Error_Gm_Execute"
	case Error_Cmd_Execute:
		ret = "Error_Cmd_Execute"
	case Error_TableNoData:
		ret = "Error_TableNoData"
	case Error_NoBagSpace:
		ret = "Error_NoBagSpace"
	case Error_ItemNotEnough:
		ret = "Error_ItemNotEnough"
	case Error_NoSuchItem:
		ret = "Error_NoSuchItem"
	case Error_InvalidItem:
		ret = "Error_InvalidItem"
	case Error_ZoneIdExceedLimit:
		ret = "Error_ZoneIdExceedLimit"
	case Error_GUIDAllocExceedLimit:
		ret = "Error_GUIDAllocExceedLimit"
	case Error_CoinNotEnough:
		ret = "Error_CoinNotEnough"
	case Error_DiamondNotEnough:
		ret = "Error_DiamondNotEnough"
	case Error_KingCoinNotEnough:
		ret = "Error_KingCoinNotEnough"
	case Error_ItemCannotSell:
		ret = "Error_ItemCannotSell"
	case Error_ItemCannotUse:
		ret = "Error_ItemCannotUse"
	case Error_SkillPointNotEnough:
		ret = "Error_SkillPointNotEnough"
	case Error_RoleNameEmpty:
		ret = "Error_RoleNameEmpty"
	case Error_VipLevelNotEnough:
		ret = "Error_VipLevelNotEnough"
	case Error_Role_BanLogin:
		ret = "Error_Role_BanLogin"
	case Error_Role_BanSpeak:
		ret = "Error_Role_BanSpeak"
	case Error_Module_Closed:
		ret = "Error_Module_Closed"
	case Error_Sys_CmdClosed:
		ret = "Error_Sys_CmdClosed"
	case Error_TalentNotEnough:
		ret = "Error_TalentNotEnough"
	case Error_CannotBindMultiple:
		ret = "Error_CannotBindMultiple"
	case Error_Role_HasNoFaceFrame:
		ret = "Error_Role_HasNoFaceFrame"
	case Error_Role_FaceFrameLock:
		ret = "Error_Role_FaceFrameLock"
	case Error_Role_HasNoFace:
		ret = "Error_Role_HasNoFace"
	case Error_Hero_HasBeenAdd:
		ret = "Error_Hero_HasBeenAdd"
	case Error_Hero_NotInList:
		ret = "Error_Hero_NotInList"
	case Error_Hero_StarTooLarge:
		ret = "Error_Hero_StarTooLarge"
	case Error_Hero_StepTooLarge:
		ret = "Error_Hero_StepTooLarge"
	case Error_Hero_StarBeenMax:
		ret = "Error_Hero_StarBeenMax"
	case Error_Hero_StepBeenMax:
		ret = "Error_Hero_StepBeenMax"
	case Error_Hero_StepUpgrade:
		ret = "Error_Hero_StepUpgrade"
	case Error_Hero_SkillNoLearned:
		ret = "Error_Hero_SkillNoLearned"
	case Error_Hero_SkillBeenMax:
		ret = "Error_Hero_SkillBeenMax"
	case Error_Hero_SoulCanNotHero:
		ret = "Error_Hero_SoulCanNotHero"
	case Error_Hero_EquipPosAlready:
		ret = "Error_Hero_EquipPosAlready"
	case Error_Hero_EquipInvalidPos:
		ret = "Error_Hero_EquipInvalidPos"
	case Error_Hero_StepLess:
		ret = "Error_Hero_StepLess"
	case Error_Hero_LevelNotEnough:
		ret = "Error_Hero_LevelNotEnough"
	case Error_Hero_TalentBeenMax:
		ret = "Error_Hero_TalentBeenMax"
	case Error_Hero_RuneLevelBeenMax:
		ret = "Error_Hero_RuneLevelBeenMax"
	case Error_RuneNotEnough:
		ret = "Error_RuneNotEnough"
	case Error_Hero_SkillHasLearned:
		ret = "Error_Hero_SkillHasLearned"
	case Error_ZoonExpNotEnough:
		ret = "Error_ZoonExpNotEnough"
	case Error_ZoonName_LengthLimit:
		ret = "Error_ZoonName_LengthLimit"
	case Error_Stage_Cheat:
		ret = "Error_Stage_Cheat"
	case Error_Stage_InvalidTeam:
		ret = "Error_Stage_InvalidTeam"
	case Error_Stage_MaxTeam:
		ret = "Error_Stage_MaxTeam"
	case Error_Market_NotOpen:
		ret = "Error_Market_NotOpen"
	case Error_Market_HasSold:
		ret = "Error_Market_HasSold"
	case Error_Market_RefreshTime:
		ret = "Error_Market_RefreshTime"
	case Error_Market_VipMaxTimes:
		ret = "Error_Market_VipMaxTimes"
	case Error_Arena_NoReport:
		ret = "Error_Arena_NoReport"
	case Error_Arena_NotOpen:
		ret = "Error_Arena_NotOpen"
	case Error_ArenaCoinNotEnough:
		ret = "Error_ArenaCoinNotEnough"
	case Error_Mail_Role_Null:
		ret = "Error_Mail_Role_Null"
	case Error_Mail_NotExist:
		ret = "Error_Mail_NotExist"
	case Error_Mail_NoAttach:
		ret = "Error_Mail_NoAttach"
	case Error_Mail_LengthLimited:
		ret = "Error_Mail_LengthLimited"
	case Error_Mail_NoMoreMail:
		ret = "Error_Mail_NoMoreMail"
	case Error_Mail_MailAttachNotRcv:
		ret = "Error_Mail_MailAttachNotRcv"
	case Error_Mail_RcvUserNotExit:
		ret = "Error_Mail_RcvUserNotExit"
	case Error_Mail_SendToMyself:
		ret = "Error_Mail_SendToMyself"
	case Error_Name_RoleNameExist:
		ret = "Error_Name_RoleNameExist"
	case Error_Name_LengthLimit:
		ret = "Error_Name_LengthLimit"
	case Error_HasDirty:
		ret = "Error_HasDirty"
	case Error_InvalidChar:
		ret = "Error_InvalidChar"
	case Error_Chat_InCD:
		ret = "Error_Chat_InCD"
	case Error_Chat_NotOpen:
		ret = "Error_Chat_NotOpen"
	case Error_Chat_LevelLimit:
		ret = "Error_Chat_LevelLimit"
	case Error_Chat_NotSelf:
		ret = "Error_Chat_NotSelf"
	case Error_Chat_Channel:
		ret = "Error_Chat_Channel"
	case Error_Chat_MessageEmpty:
		ret = "Error_Chat_MessageEmpty"
	case Error_Chat_MessageType:
		ret = "Error_Chat_MessageType"
	case Error_Chat_NotInShieldedList:
		ret = "Error_Chat_NotInShieldedList"
	case Error_Chat_WasInShieldedList:
		ret = "Error_Chat_WasInShieldedList"
	case Error_Item_BeyondMaxNum:
		ret = "Error_Item_BeyondMaxNum"
	case Error_Activity_NotExists:
		ret = "Error_Activity_NotExists"
	case Error_Activity_NotJoin:
		ret = "Error_Activity_NotJoin"
	case Error_Activity_Condition:
		ret = "Error_Activity_Condition"
	case Error_Activity_HideOnFinish:
		ret = "Error_Activity_HideOnFinish"
	case Error_Activity_HasReward:
		ret = "Error_Activity_HasReward"
	case Error_Activity_Close:
		ret = "Error_Activity_Close"
	case Error_Talent_CoinBuy:
		ret = "Error_Talent_CoinBuy"
	case Error_Talent_DiamondBuy:
		ret = "Error_Talent_DiamondBuy"
	case Error_Talent_LevelBeenMax:
		ret = "Error_Talent_LevelBeenMax"
	case Error_Home_Null:
		ret = "Error_Home_Null"
	case Error_Home_NotJoin:
		ret = "Error_Home_NotJoin"
	case Error_Home_HasJoin:
		ret = "Error_Home_HasJoin"
	case Error_Home_NotMaster:
		ret = "Error_Home_NotMaster"
	case Error_Home_NotInMemberList:
		ret = "Error_Home_NotInMemberList"
	case Error_Home_MaxMember:
		ret = "Error_Home_MaxMember"
	case Error_Home_Exist:
		ret = "Error_Home_Exist"
	case Error_Home_SameName:
		ret = "Error_Home_SameName"
	case Error_Home_HasApply:
		ret = "Error_Home_HasApply"
	case Error_Home_MaxApply:
		ret = "Error_Home_MaxApply"
	case Error_Home_NotInApplyList:
		ret = "Error_Home_NotInApplyList"
	case Error_Home_HasMembers:
		ret = "Error_Home_HasMembers"
	case Error_AdvMap_PosHasHome:
		ret = "Error_AdvMap_PosHasHome"
	case Error_Home_CannotJoinAnyHome:
		ret = "Error_Home_CannotJoinAnyHome"
	case Error_AdvMap_Null:
		ret = "Error_AdvMap_Null"
	case Error_AdvMap_Exist:
		ret = "Error_AdvMap_Exist"
	case Error_AdvMap_HasRewarded:
		ret = "Error_AdvMap_HasRewarded"
	case Error_AdvMap_PosNoOwner:
		ret = "Error_AdvMap_PosNoOwner"
	case Error_AdvMap_Fighting:
		ret = "Error_AdvMap_Fighting"
	case Error_AdvMap_Refresh:
		ret = "Error_AdvMap_Refresh"
	case Error_Chat_SelectNoOne:
		ret = "Error_Chat_SelectNoOne"
	case Error_AdvMap_HeroHasBorrowed:
		ret = "Error_AdvMap_HeroHasBorrowed"
	case Error_AdvMap_HeroNotSelf:
		ret = "Error_AdvMap_HeroNotSelf"
	case Error_FightReport_Invalid:
		ret = "Error_FightReport_Invalid"
	case Error_FightReport_ShareType:
		ret = "Error_FightReport_ShareType"
	case Error_FightReport_ShareSelf:
		ret = "Error_FightReport_ShareSelf"
	case Error_FightReport_MessageCharError:
		ret = "Error_FightReport_MessageCharError"
	case Error_CDKey_ZoneList:
		ret = "Error_CDKey_ZoneList"
	case Error_MP_NoCDKey:
		ret = "Error_MP_NoCDKey"
	case Error_MP_CDKeyExpired:
		ret = "Error_MP_CDKeyExpired"
	case Error_MP_InvalidCDKey:
		ret = "Error_MP_InvalidCDKey"
	case Error_MP_CDKeyExchanged:
		ret = "Error_MP_CDKeyExchanged"
	case Error_MP_ExchangeBindLimited:
		ret = "Error_MP_ExchangeBindLimited"
	case Error_MP_ExchangeNumLimited:
		ret = "Error_MP_ExchangeNumLimited"
	case Error_MP_InvalidDivision:
		ret = "Error_MP_InvalidDivision"
	case Error_MP_LimitConfigError:
		ret = "Error_MP_LimitConfigError"
	case Error_MP_InDelivering:
		ret = "Error_MP_InDelivering"
	case Error_MP_CDKeyNotOpen:
		ret = "Error_MP_CDKeyNotOpen"
	case Error_Hero_GemInvalidSlot:
		ret = "Error_Hero_GemInvalidSlot"
	case Error_Hero_GemMaxLevel:
		ret = "Error_Hero_GemMaxLevel"
	case Error_Gem_NotEnough:
		ret = "Error_Gem_NotEnough"
	case Error_King_NotOpen:
		ret = "Error_King_NotOpen"
	case Error_Quest_Reward:
		ret = "Error_Quest_Reward"
	case Error_Quest_Step:
		ret = "Error_Quest_Step"
	case Error_Quest_VitalityReward:
		ret = "Error_Quest_VitalityReward"
	case Error_Achieve_Step:
		ret = "Error_Achieve_Step"
	case Error_HeroExp_NotEnough:
		ret = "Error_HeroExp_NotEnough"
	case Error_Arena_WinTimesReward:
		ret = "Error_Arena_WinTimesReward"
	case Error_Arena_WinTimesNotEnough:
		ret = "Error_Arena_WinTimesNotEnough"
	case Error_ChallengeTimes_NotZero:
		ret = "Error_ChallengeTimes_NotZero"
	case Error_Chat_ShieldListFull:
		ret = "Error_Chat_ShieldListFull"
	case Error_Chat_HasBeenShield:
		ret = "Error_Chat_HasBeenShield"
	case Error_HeroTrialTimesNotEnough:
		ret = "Error_HeroTrialTimesNotEnough"
	case Error_Arena_ChallengeTimesNotEnough:
		ret = "Error_Arena_ChallengeTimesNotEnough"
	case Error_Fight_isFighting:
		ret = "Error_Fight_isFighting"
	case Error_King_ChallengeTimesNotEnough:
		ret = "Error_King_ChallengeTimesNotEnough"
	case Error_King_BeAttacking:
		ret = "Error_King_BeAttacking"
	case Error_Recharge_BuyNumBeenMax:
		ret = "Error_Recharge_BuyNumBeenMax"
	case Error_Talent_NotOpen:
		ret = "Error_Talent_NotOpen"
	case Error_Trial_NotOpen:
		ret = "Error_Trial_NotOpen"
	case Error_SignIn_HasReward:
		ret = "Error_SignIn_HasReward"
	case Error_SignIn_NotReward:
		ret = "Error_SignIn_NotReward"
	case Error_SignIn_HasDouble:
		ret = "Error_SignIn_HasDouble"
	case Error_Vip_HasReward:
		ret = "Error_Vip_HasReward"
	case Error_Stage_GamesHasReward:
		ret = "Error_Stage_GamesHasReward"
	case Error_Pet_Null:
		ret = "Error_Pet_Null"
	case Error_Stage_VerifyFull:
		ret = "Error_Stage_VerifyFull"
	case Error_TrialCoinNotEnough:
		ret = "Error_TrialCoinNotEnough"
	case Error_HomeCoinNotEnough:
		ret = "Error_HomeCoinNotEnough"
	case Error_CrowdFund_Complete:
		ret = "Error_CrowdFund_Complete"
	case Error_Rune_NotExists:
		ret = "Error_Rune_NotExists"
	case Error_Rune_Lock:
		ret = "Error_Rune_Lock"
	case Error_ActBind_BindMe:
		ret = "Error_ActBind_BindMe"
	case Error_AdvWheel_TimesNotEnough:
		ret = "Error_AdvWheel_TimesNotEnough"
	case Error_CrowdFund_Recharging:
		ret = "Error_CrowdFund_Recharging"
	case Error_BranchStage_GameOver:
		ret = "Error_BranchStage_GameOver"
	case Error_AdvMap_Injuring:
		ret = "Error_AdvMap_Injuring"
	case Error_Item_OnlyUseOne:
		ret = "Error_Item_OnlyUseOne"
	case Error_Item_EffectExist:
		ret = "Error_Item_EffectExist"
	case Error_Pet_MaxSlotSize:
		ret = "Error_Pet_MaxSlotSize"
	case Error_Hero_Syncing:
		ret = "Error_Hero_Syncing"
	case Error_Zoon_CancelMate:
		ret = "Error_Zoon_CancelMate"
	case Error_ZoonSlotNotEnough:
		ret = "Error_ZoonSlotNotEnough"
	case Error_Obj_NotExists:
		ret = "Error_Obj_NotExists"
	case Error_Role_Reborn:
		ret = "Error_Role_Reborn"
	case Error_Map_RandPos:
		ret = "Error_Map_RandPos"
	case Error_Slg_NotOpen:
		ret = "Error_Slg_NotOpen"
	case Error_Map_InvalidPos:
		ret = "Error_Map_InvalidPos"
	case Error_Pos_HasObj:
		ret = "Error_Pos_HasObj"
	case Error_Pos_HasNoObj:
		ret = "Error_Pos_HasNoObj"
	case Error_FoodNotEnough:
		ret = "Error_FoodNotEnough"
	case Error_WoodNotEnough:
		ret = "Error_WoodNotEnough"
	case Error_StoneNotEnough:
		ret = "Error_StoneNotEnough"
	case Error_SilverNotEnough:
		ret = "Error_SilverNotEnough"
	case Error_Pos_SendTeamAlready:
		ret = "Error_Pos_SendTeamAlready"
	case Error_Slg_ReportInvalid:
		ret = "Error_Slg_ReportInvalid"
	case Error_Map_ResNoOwner:
		ret = "Error_Map_ResNoOwner"
	case Error_Pos_Forbid:
		ret = "Error_Pos_Forbid"
	case Error_DungeonTeam_Null:
		ret = "Error_DungeonTeam_Null"
	case Error_DungeonTeam_NotJoin:
		ret = "Error_DungeonTeam_NotJoin"
	case Error_DungeonTeam_HasJoin:
		ret = "Error_DungeonTeam_HasJoin"
	case Error_DungeonTeam_JoinCode:
		ret = "Error_DungeonTeam_JoinCode"
	case Error_DungeonTeam_FullMember:
		ret = "Error_DungeonTeam_FullMember"
	case Error_DungeonTeam_Power:
		ret = "Error_DungeonTeam_Power"
	case Error_City_DoorBroken:
		ret = "Error_City_DoorBroken"
	case Error_StorePos_Name:
		ret = "Error_StorePos_Name"
	case Error_StorePos_MaxNum:
		ret = "Error_StorePos_MaxNum"
	case Error_Wand_Max:
		ret = "Error_Wand_Max"
	case Error_WandNotEnough:
		ret = "Error_WandNotEnough"
	case Error_WorkShopLevelNotEnough:
		ret = "Error_WorkShopLevelNotEnough"
	case Error_SendCacheRsp:
		ret = "Error_SendCacheRsp"
	case Error_Summon_MaxTimes:
		ret = "Error_Summon_MaxTimes"
	case Error_Summon_CD:
		ret = "Error_Summon_CD"
	case Error_Alliance_NotJoin:
		ret = "Error_Alliance_NotJoin"
	case Error_Alliance_Null:
		ret = "Error_Alliance_Null"
	case Error_Alliance_NotMaster:
		ret = "Error_Alliance_NotMaster"
	case Error_Alliance_HasJoin:
		ret = "Error_Alliance_HasJoin"
	case Error_Alliance_NotInApplyList:
		ret = "Error_Alliance_NotInApplyList"
	case Error_Alliance_CannotJoinAnyAlliance:
		ret = "Error_Alliance_CannotJoinAnyAlliance"
	case Error_Alliance_MaxApply:
		ret = "Error_Alliance_MaxApply"
	case Error_Alliance_HasApply:
		ret = "Error_Alliance_HasApply"
	case Error_Alliance_SameName:
		ret = "Error_Alliance_SameName"
	case Error_Alliance_MaxMember:
		ret = "Error_Alliance_MaxMember"
	case Error_Alliance_NotInMemberList:
		ret = "Error_Alliance_NotInMemberList"
	case Error_Alliance_HasMembers:
		ret = "Error_Alliance_HasMembers"
	case Error_Alliance_Exist:
		ret = "Error_Alliance_Exist"
	case Error_Alliance_MaxKickTimes:
		ret = "Error_Alliance_MaxKickTimes"
	case Error_Alliance_PostMaxSize:
		ret = "Error_Alliance_PostMaxSize"
	case Error_AllianceCoinNotEnough:
		ret = "Error_AllianceCoinNotEnough"
	case Error_AllianceBoss_NotOpen:
		ret = "Error_AllianceBoss_NotOpen"
	case Error_AllianceBoss_HasReward:
		ret = "Error_AllianceBoss_HasReward"
	case Error_NotMyMonster:
		ret = "Error_NotMyMonster"
	case Error_Hero_QualityBeenMax:
		ret = "Error_Hero_QualityBeenMax"
	case Error_Hero_QualityLess:
		ret = "Error_Hero_QualityLess"
	case Error_GemBagSize:
		ret = "Error_GemBagSize"
	case Error_Slg_NotNewbee:
		ret = "Error_Slg_NotNewbee"
	case Error_Alliance_NotSendReward:
		ret = "Error_Alliance_NotSendReward"
	case Error_WorkShopNotLevel:
		ret = "Error_WorkShopNotLevel"
	case Error_Slg_FlagHasActive:
		ret = "Error_Slg_FlagHasActive"
	case Error_Slg_FlagNotSameAlliance:
		ret = "Error_Slg_FlagNotSameAlliance"
	case Error_Slg_InFlagNoProtect:
		ret = "Error_Slg_InFlagNoProtect"
	case Error_Slg_FlagInProtect:
		ret = "Error_Slg_FlagInProtect"
	case Error_Slg_NoCityNoUseGem:
		ret = "Error_Slg_NoCityNoUseGem"
	case Error_Obj_NotInVision:
		ret = "Error_Obj_NotInVision"
	case Error_City_CanotMove:
		ret = "Error_City_CanotMove"
	case Error_Map_TeamQuality:
		ret = "Error_Map_TeamQuality"
	case Error_RoleMap_Null:
		ret = "Error_RoleMap_Null"
	case Error_Flag_AllianceOnlyOne:
		ret = "Error_Flag_AllianceOnlyOne"
	case Error_Alliance_PostNotEnough:
		ret = "Error_Alliance_PostNotEnough"
	}
	return ret
}
