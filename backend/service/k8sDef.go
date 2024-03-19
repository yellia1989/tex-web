package service

import "github.com/yellia1989/tex-web/backend/cfg"

func GetDirServiceName() string {
	if cfg.K8s {
		return cfg.App + ".DirServer.DirServiceObj@tcp -h interserver -p 3003 -t 600000"
	} else {
		return cfg.App + ".DirServer.DirServiceObj"
	}
}

func GetBulletServiceName() string {
	if cfg.K8s {
		return cfg.App + ".BulletinServer.BulletinServiceObj@tcp -h interserver -p 3002 -t 600000"
	} else {
		return cfg.App + ".BulletinServer.BulletinServiceObj"
	}
}

func GetMailServiceName() string {
	if cfg.K8s {
		return cfg.App + ".MailServer.MailServiceObj@tcp -h interserver -p 3006 -t 600000"
	} else {
		return cfg.App + ".MailServer.MailServiceObj"
		c
	}
}

func GetMPServiceName() string {
	if cfg.K8s {
		return cfg.App + ".MPServer.MPServiceObj@tcp -h mpserver -p 3001 -t 600000"
	} else {
		return cfg.App + ".MPServer.MPServiceObj"
	}
}

func GetIAPServiceName() string {
	if cfg.K8s {
		return cfg.App + ".IAPServer.IAPServiceObj@tcp -h iapserver -p 3001 -t 600000"
	} else {
		return cfg.App + ".IAPServer.IAPServiceObj"
	}
}

func GetLoginServiceName() string {
	if cfg.K8s {
		return cfg.App + ".LoginServer.LoginServiceObj@tcp -h loginserver -p 3001 -t 600000"
	} else {
		return cfg.App + ".LoginServer.LoginServiceObj"
	}
}

func GetGameServiceName(zoneid string) string {
	if cfg.K8s {
		return cfg.App + ".GameServer.GameServiceObj@tcp -h game" + zoneid + " -p 7001 -t 600000"
	} else {
		return cfg.App + ".GameServer.GameServiceObj%" + cfg.App + ".zone." + zoneid
	}
}
