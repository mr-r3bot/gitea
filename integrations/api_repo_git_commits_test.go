	"code.gitea.io/gitea/models/db"
	user := db.AssertExistsAndLoadBean(t, &models.User{ID: 2}).(*models.User)
	user := db.AssertExistsAndLoadBean(t, &models.User{ID: 2}).(*models.User)
	user := db.AssertExistsAndLoadBean(t, &models.User{ID: 2}).(*models.User)
	user := db.AssertExistsAndLoadBean(t, &models.User{ID: 2}).(*models.User)

func TestDownloadCommitDiffOrPatch(t *testing.T) {
	defer prepareTestEnv(t)()
	user := db.AssertExistsAndLoadBean(t, &models.User{ID: 2}).(*models.User)
	// Login as User2.
	session := loginUser(t, user.Name)
	token := getTokenForLoggedInUser(t, session)

	// Test getting diff
	reqDiff := NewRequestf(t, "GET", "/api/v1/repos/%s/repo16/git/commits/f27c2b2b03dcab38beaf89b0ab4ff61f6de63441.diff?token="+token, user.Name)
	resp := session.MakeRequest(t, reqDiff, http.StatusOK)
	assert.EqualValues(t,
		"commit f27c2b2b03dcab38beaf89b0ab4ff61f6de63441\nAuthor: User2 <user2@example.com>\nDate:   Sun Aug 6 19:55:01 2017 +0200\n\n    good signed commit\n\ndiff --git a/readme.md b/readme.md\nnew file mode 100644\nindex 0000000..458121c\n--- /dev/null\n+++ b/readme.md\n@@ -0,0 +1 @@\n+good sign\n",
		resp.Body.String())

	// Test getting patch
	reqPatch := NewRequestf(t, "GET", "/api/v1/repos/%s/repo16/git/commits/f27c2b2b03dcab38beaf89b0ab4ff61f6de63441.patch?token="+token, user.Name)
	resp = session.MakeRequest(t, reqPatch, http.StatusOK)
	assert.EqualValues(t,
		"From f27c2b2b03dcab38beaf89b0ab4ff61f6de63441 Mon Sep 17 00:00:00 2001\nFrom: User2 <user2@example.com>\nDate: Sun, 6 Aug 2017 19:55:01 +0200\nSubject: [PATCH] good signed commit\n\n---\n readme.md | 1 +\n 1 file changed, 1 insertion(+)\n create mode 100644 readme.md\n\ndiff --git a/readme.md b/readme.md\nnew file mode 100644\nindex 0000000..458121c\n--- /dev/null\n+++ b/readme.md\n@@ -0,0 +1 @@\n+good sign\n",
		resp.Body.String())

}