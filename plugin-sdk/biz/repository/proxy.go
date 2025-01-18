package repository

import (
	cqtt2 "github.com/xieyaoxin/pockpluginsdk/plugin-sdk/biz/repository/impl/cqtt"
	"github.com/xieyaoxin/pockpluginsdk/plugin-sdk/biz/status"
)

func GetUserRepository() UserRepository {
	switch status.SERVER_NAME {
	case status.CQTT:
		return cqtt2.UserRepositoryImpl4CQTTInstance
	}
	return nil
}

func GetArticleRepository() ArticleRepository {
	switch status.SERVER_NAME {
	case status.CQTT:
		return cqtt2.ArticleRepositoryImpl4CQTTInstance
	}
	return nil
}

func GetMapRepository() MapRepository {
	switch status.SERVER_NAME {
	case status.CQTT:
		return cqtt2.MapRepositoryImpl4CQTTInstance
	}
	return nil
}

func GetPetRepository() PetRepository {
	switch status.SERVER_NAME {
	case status.CQTT:
		return cqtt2.PetRepositoryImpl4CQTT
	}
	return nil
}

func GetBattleRepository() BattleRepository {
	switch status.SERVER_NAME {
	case status.CQTT:
		return cqtt2.BattleRepositoryImplInstance
	}
	return nil
}
