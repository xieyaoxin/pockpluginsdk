package repository

import (
	KDHS2 "github.com/xieyaoxin/pockpluginsdk/plugin-sdk/biz/repository/impl/KDHS"
	"github.com/xieyaoxin/pockpluginsdk/plugin-sdk/biz/status"
)

func GetUserRepository() UserRepository {
	switch status.SERVER_NAME {
	case status.KDHS:
		return KDHS2.UserRepositoryImpl4KDHSInstance
	}
	return nil
}

func GetArticleRepository() ArticleRepository {
	switch status.SERVER_NAME {
	case status.KDHS:
		return KDHS2.ArticleRepositoryImpl4KDHSInstance
	}
	return nil
}

func GetMapRepository() MapRepository {
	switch status.SERVER_NAME {
	case status.KDHS:
		return KDHS2.MapRepositoryImpl4KDHSInstance
	}
	return nil
}

func GetPetRepository() PetRepository {
	switch status.SERVER_NAME {
	case status.KDHS:
		return KDHS2.PetRepositoryImpl4KDHS
	}
	return nil
}

func GetBattleRepository() BattleRepository {
	switch status.SERVER_NAME {
	case status.KDHS:
		return KDHS2.BattleRepositoryImplInstance
	}
	return nil
}
