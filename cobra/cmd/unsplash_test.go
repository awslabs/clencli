package cmd

import (
	"os"
	"testing"

	"github.com/awslabs/clencli/cobra/aid"
	"github.com/awslabs/clencli/cobra/controller"
	"github.com/awslabs/clencli/cobra/model"
	"github.com/awslabs/clencli/helper"
	"github.com/awslabs/clencli/tester"
	"github.com/stretchr/testify/assert"
)

func createUnplashCredential() {
	// only setup credentials if non-existent
	if !aid.CredentialsFileExist() {
		var credentials model.Credentials

		var profile model.CredentialProfile
		profile.Name = "unsplash-unit-testing"
		profile.Enabled = true // enabling profile by default

		var credential model.Credential
		credential.Enabled = true
		credential.AccessKey = os.Getenv("UNSPLASH_ACCESS_KEY")
		credential.SecretKey = os.Getenv("UNSPLASH_SECRET_KEY")
		credential.Provider = "unsplash"
		profile.Credentials = append(profile.Credentials, credential)

		credentials.Profiles = append(credentials.Profiles, profile)
		aid.WriteInterfaceToFile(credentials, aid.GetAppInfo().CredentialsPath)
	}
}

func DeleteCredential() {
	if aid.CredentialsFileExist() {
		aid.DeleteCredentialFile()
	}
}

func TestUnsplashEmptyWithoutCredentials(t *testing.T) {
	err := tester.ExecuteCommand(controller.UnsplashCmd(), "unsplash")
	assert.Contains(t, err.Error(), "unable to read credentials")
}

func TestUnsplashEmptyWithEmptyCredentials(t *testing.T) {
	createUnplashCredential()
	err := tester.ExecuteCommand(controller.UnsplashCmd(), "unsplash")
	assert.Contains(t, err.Error(), "no unsplash credential found")
	DeleteCredential()
}

func TestUnsplashEmptyWithCredentials(t *testing.T) {
	createUnplashCredential()
	err := tester.ExecuteCommand(controller.UnsplashCmd(), "unsplash")
	assert.Contains(t, err.Error(), "unable to read credentials")
	// DeleteCredential()
}

func TestUnsplashWithUntiTestingProfile(t *testing.T) {
	// TODO: setup the clencli/credentials before starting the test
	pwd, nwd := tester.Setup(t)
	createUnplashCredential()
	err := tester.ExecuteCommand(controller.UnsplashCmd(), "unsplash", "--profile", "unit-testing")
	dPath := pwd + "/" + nwd + "/" + "downloads"
	assert.Nil(t, err)
	assert.FileExists(t, "unsplash.yaml")
	assert.DirExists(t, dPath)
	assert.DirExists(t, dPath+"/unsplash")
	assert.DirExists(t, dPath+"/unsplash/mountains")

	files := helper.ListFiles(dPath + "/unsplash/mountains/")
	assert.Greater(t, len(files), 0)
}

func TestUnsplashWithQuery(t *testing.T) {
	createUnplashCredential()
	err := tester.ExecuteCommand(controller.UnsplashCmd(), "unsplash", "--query", "horse")
	assert.Contains(t, err.Error(), "invalid argument")
}
