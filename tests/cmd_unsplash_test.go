package tests

import (
	"os"
	"testing"

	"github.com/awslabs/clencli/cobra/aid"
	"github.com/awslabs/clencli/cobra/controller"
	"github.com/awslabs/clencli/cobra/model"
	"github.com/awslabs/clencli/helper"
	"github.com/stretchr/testify/assert"
)

func createUnplashCredential() {
	aid.DeleteConfigurationsDirectory()
	aid.CreateConfigurationsDirectory()

	var credentials model.Credentials

	var profile model.CredentialProfile
	profile.Name = "default"
	profile.Enabled = true // enabling profile by default

	var credential model.Credential
	credential.Name = "unit-testing"
	credential.Enabled = true
	credential.AccessKey = os.Getenv("UNSPLASH_ACCESS_KEY")
	credential.SecretKey = os.Getenv("UNSPLASH_SECRET_KEY")
	credential.Provider = "unsplash"

	profile.Credentials = append(profile.Credentials, credential)
	credentials.Profiles = append(credentials.Profiles, profile)
	aid.WriteInterfaceToFile(credentials, aid.GetAppInfo().CredentialsPath)
}

func DeleteCredential() {
	if aid.CredentialsFileExist() {
		aid.DeleteCredentialFile()
	}
}

func TestUnsplashEmptyWithoutCredentials(t *testing.T) {
	aid.DeleteConfigurationsDirectory()
	args := []string{"unsplash"}
	out, err := executeCommand(t, controller.UnsplashCmd(), args)
	assert.NotNil(t, err)
	assert.Contains(t, out, "")
	assert.Contains(t, err.Error(), "unable to read credentials")
	assert.Contains(t, err.Error(), "unable to read configuration")

}

func TestUnsplashEmptyWithCredentials(t *testing.T) {
	createUnplashCredential()
	defer aid.DeleteConfigurationsDirectory()

	args := []string{"unsplash"}
	_, err := executeCommand(t, controller.UnsplashCmd(), args)

	sep := string(os.PathSeparator)
	dir := t.Name() + sep

	assert.Nil(t, err)
	assert.FileExists(t, dir+sep+"unsplash.yaml")
	assert.DirExists(t, dir+sep+"downloads")
	assert.DirExists(t, dir+sep+"downloads"+sep+"unsplash")
	assert.DirExists(t, dir+sep+"downloads"+sep+"unsplash"+sep+"mountains")

	files := helper.ListFiles(dir + sep + "downloads" + sep + "unsplash" + sep + "mountains")
	assert.GreaterOrEqual(t, len(files), 5)
}

func TestUnsplashQuery(t *testing.T) {
	createUnplashCredential()
	defer aid.DeleteConfigurationsDirectory()

	args := []string{"unsplash", "--query", "horse"}
	_, err := executeCommand(t, controller.UnsplashCmd(), args)

	sep := string(os.PathSeparator)
	dir := t.Name() + sep

	assert.Nil(t, err)
	assert.FileExists(t, dir+sep+"unsplash.yaml")
	assert.DirExists(t, dir+sep+"downloads")
	assert.DirExists(t, dir+sep+"downloads"+sep+"unsplash")
	assert.DirExists(t, dir+sep+"downloads"+sep+"unsplash"+sep+"horse")

	files := helper.ListFiles(dir + sep + "downloads" + sep + "unsplash" + sep + "horse")
	assert.GreaterOrEqual(t, len(files), 5)
}
