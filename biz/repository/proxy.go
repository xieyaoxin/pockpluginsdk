package repository

import (
	"github.com/xieyaoxin/plugin-sdk/biz/repository/impl/cqtt"
	"github.com/xieyaoxin/plugin-sdk/biz/status"
)

func GetUserRepository() UserRepository {
	switch status.SERVER_NAME {
	case status.CQTT:
		return cqtt.UserRepositoryImpl4CQTTInstance
	}
	return nil
}

func GetArticleRepository() ArticleRepository {
	switch status.SERVER_NAME {
	case status.CQTT:
		return cqtt.ArticleRepositoryImpl4CQTTInstance
	}
	return nil
}

func GetMapRepository() MapRepository {
	switch status.SERVER_NAME {
	case status.CQTT:
		return cqtt.MapRepositoryImpl4CQTTInstance
	}
	return nil
}

func GetPetRepository() PetRepository {
	switch status.SERVER_NAME {
	case status.CQTT:
		return cqtt.PetRepositoryImpl4CQTT
	}
	return nil
}

func GetBattleRepository() BattleRepository {
	switch status.SERVER_NAME {
	case status.CQTT:
		return cqtt.BattleRepositoryImplInstance
	}
	return nil
}
