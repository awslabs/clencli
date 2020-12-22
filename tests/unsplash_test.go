package tests

// func createUnplashCredential() {
// 	aid.DeleteConfigurationsDirectory()
// 	aid.CreateConfigurationsDirectory()

// 	var credentials model.Credentials

// 	var profile model.CredentialProfile
// 	profile.Name = "default"
// 	profile.Enabled = true // enabling profile by default

// 	var credential model.Credential
// 	credential.Name = "unit-testing"
// 	credential.Enabled = true
// 	credential.AccessKey = os.Getenv("UNSPLASH_ACCESS_KEY")
// 	credential.SecretKey = os.Getenv("UNSPLASH_SECRET_KEY")
// 	credential.Provider = "unsplash"

// 	profile.Credentials = append(profile.Credentials, credential)
// 	credentials.Profiles = append(credentials.Profiles, profile)
// 	aid.WriteInterfaceToFile(credentials, aid.GetAppInfo().CredentialsPath)
// }

// func DeleteCredential() {
// 	if aid.CredentialsFileExist() {
// 		aid.DeleteCredentialFile()
// 	}
// }

// func TestUnsplashEmptyWithoutCredentials(t *testing.T) {
// 	aid.DeleteConfigurationsDirectory()
// 	err := tester.ExecuteCommand(controller.UnsplashCmd(), "unsplash")
// 	assert.Contains(t, err.Error(), "unable to read credentials")
// }

// func TestUnsplashEmptyWithCredentials(t *testing.T) {
// 	tester.CreateTestDirectory(t)
// 	createUnplashCredential()

// 	sep := string(os.PathSeparator)
// 	err := tester.ExecuteCommand(controller.UnsplashCmd(), "unsplash")

// 	assert.Nil(t, err)
// 	assert.FileExists(t, "unsplash.yaml")
// 	assert.DirExists(t, "downloads")
// 	assert.DirExists(t, "downloads"+sep+"unsplash")
// 	assert.DirExists(t, "downloads"+sep+"unsplash"+sep+"mountains")

// 	files := helper.ListFiles("downloads" + sep + "unsplash" + sep + "mountains")
// 	assert.GreaterOrEqual(t, len(files), 5)
// }

// func TestUnsplashQuery(t *testing.T) {
// 	tester.CreateTestDirectory(t)
// 	createUnplashCredential()

// 	sep := string(os.PathSeparator)
// 	err := tester.ExecuteCommand(controller.UnsplashCmd(), "unsplash", "--query", "horses")

// 	assert.Nil(t, err)
// 	assert.FileExists(t, "unsplash.yaml")
// 	assert.DirExists(t, "downloads")
// 	assert.DirExists(t, "downloads"+sep+"unsplash")
// 	assert.DirExists(t, "downloads"+sep+"unsplash"+sep+"horses")

// 	files := helper.ListFiles("downloads" + sep + "unsplash" + sep + "horses")
// 	assert.GreaterOrEqual(t, len(files), 5)
// }

// func TestUnsplashWithUntiTestingProfile(t *testing.T) {
// 	// TODO: setup the clencli/credentials before starting the test
// 	pwd, nwd := tester.Setup(t)
// 	createUnplashCredential()
// 	err := tester.ExecuteCommand(controller.UnsplashCmd(), "unsplash", "--query", "horse")
// 	dPath := pwd + "/" + nwd + "/" + "downloads"
// 	assert.Nil(t, err)
// 	assert.FileExists(t, "unsplash.yaml")
// 	assert.DirExists(t, dPath)
// 	assert.DirExists(t, dPath+"/unsplash")
// 	assert.DirExists(t, dPath+"/unsplash/horse")

// 	files := helper.ListFiles(dPath + "/unsplash/horse/")
// 	assert.Greater(t, len(files), 0)

// 	DeleteCredential()
// }
