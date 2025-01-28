package repository

import (
	"github.com/xieyaoxin/pockpluginsdk/plugin-sdk/biz/repository/impl/kdhs"
	"github.com/xieyaoxin/pockpluginsdk/plugin-sdk/biz/status"
)

func GetUserRepository() UserRepository {
	switch status.SERVER_NAME {
	case status.KDHS:
		return kdhs.UserRepositoryImpl4KDHSInstance
	}
	return nil
}

func GetArticleRepository() ArticleRepository {
	switch status.SERVER_NAME {
	case status.KDHS:
		return kdhs.ArticleRepositoryImpl4KDHSInstance
	}
	return nil
}

func GetMapRepository() MapRepository {
	switch status.SERVER_NAME {
	case status.KDHS:
		return kdhs.MapRepositoryImpl4KDHSInstance
	}
	return nil
}

func GetPetRepository() PetRepository {
	switch status.SERVER_NAME {
	case status.KDHS:
		return kdhs.PetRepositoryImpl4KDHS
	}
	return nil
}

func GetBattleRepository() BattleRepository {
	switch status.SERVER_NAME {
	case status.KDHS:
		return kdhs.BattleRepositoryImplInstance
	}
	return nil
}

func GetFusionRepository() FusionRepository {
	switch status.SERVER_NAME {
	case status.KDHS:
		return kdhs.FusionRepositorykdhsImplInstance
	default:
		return nil
	}
}

func GetEquipRepository() EquipRepository {
	switch status.SERVER_NAME {
	case status.KDHS:
		return kdhs.EquipRepositoryImplKDHSInstance
	default:
		return nil
	}
}

func GetTtRepository() TtRepository {
	switch status.SERVER_NAME {
	case status.KDHS:
		return kdhs.TtRepositoryImplKdhsInstance
	default:
		return nil
	}
}
