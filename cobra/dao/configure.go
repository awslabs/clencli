package dao

import (
	"github.com/awslabs/clencli/cobra/aid"
	"github.com/awslabs/clencli/cobra/model"
	log "github.com/sirupsen/logrus"
)

// Data Access Object

// Optional<T> get(long id);
// List<T> getAll();
// void save(T t);
// void update(T t, String[] params);
// void delete(T t);

// GetCredentials does TODO...
func GetCredentials() model.Credentials {
	creds, err := aid.GetCredentials()
	if err != nil {
		log.Fatalf("Unable to get credentials\n%v", err)
	}
	return creds
}

// SaveCredentials TODO
func SaveCredentials(credentials model.Credentials) error {
	return aid.WriteInterfaceToFile(credentials, aid.GetAppInfo().CredentialsPath)
}

func GetConfigurations() model.Configurations {
	confs, err := aid.GetConfigurations()
	if err != nil {
		log.Fatalf("Unable to get credentials\n%v", err)
	}
	return confs
}

// SaveConfigurations TODO
func SaveConfigurations(configurations model.Configurations) error {
	return aid.WriteInterfaceToFile(configurations, aid.GetAppInfo().ConfigurationsPath)
}

// func createCredentials(){
// 	var credentials model.Credentials
// 	saveCredentials(credentials)
// }
// create create profile
// create crendential

// get credentials
// get profile
// get credential

// update credentials
// update profile
// update credential

// delete credentials
// delete profile
// delete credential
