package hook

import (
	"encoding/json"
	"path"
	"testing"
	"time"

	"github.com/sirupsen/logrus"
)

func TestLog(t *testing.T) {
	log.Info("测试logrus info")
	log.Warn("测试logrus warn")
	log.Debug("测试logrus debug")
	log.WithFields(logrus.Fields{
		"var": "params",
	}).Warn("测试logrus warn")

}

func TestExecHookBash(t *testing.T) {
	jData := []byte(`{"zen":"Non-blocking is better than blocking.","hook_id":150905429,"hook":{"type":"Repository","id":150905429,"name":"web","active":true,"events":["push"],"config":{"content_type":"json","insecure_ssl":"0","url":"http://39.107.51.167:8080/push"},"updated_at":"2019-10-21T12:40:54Z","created_at":"2019-10-21T12:40:54Z","url":"https://api.github.com/repos/gaopengfei123123/github_hook/hooks/150905429","test_url":"https://api.github.com/repos/gaopengfei123123/github_hook/hooks/150905429/test","ping_url":"https://api.github.com/repos/gaopengfei123123/github_hook/hooks/150905429/pings","last_response":{"code":null,"status":"unused","message":null}},"repository":{"id":216560073,"node_id":"MDEwOlJlcG9zaXRvcnkyMTY1NjAwNzM=","name":"github_hook","full_name":"gaopengfei123123/github_hook","private":true,"owner":{"login":"gaopengfei123123","id":20946369,"node_id":"MDQ6VXNlcjIwOTQ2MzY5","avatar_url":"https://avatars3.githubusercontent.com/u/20946369?v=4","gravatar_id":"","url":"https://api.github.com/users/gaopengfei123123","html_url":"https://github.com/gaopengfei123123","followers_url":"https://api.github.com/users/gaopengfei123123/followers","following_url":"https://api.github.com/users/gaopengfei123123/following{/other_user}","gists_url":"https://api.github.com/users/gaopengfei123123/gists{/gist_id}","starred_url":"https://api.github.com/users/gaopengfei123123/starred{/owner}{/repo}","subscriptions_url":"https://api.github.com/users/gaopengfei123123/subscriptions","organizations_url":"https://api.github.com/users/gaopengfei123123/orgs","repos_url":"https://api.github.com/users/gaopengfei123123/repos","events_url":"https://api.github.com/users/gaopengfei123123/events{/privacy}","received_events_url":"https://api.github.com/users/gaopengfei123123/received_events","type":"User","site_admin":false},"html_url":"https://github.com/gaopengfei123123/github_hook","description":"钩子函数","fork":false,"url":"https://api.github.com/repos/gaopengfei123123/github_hook","forks_url":"https://api.github.com/repos/gaopengfei123123/github_hook/forks","keys_url":"https://api.github.com/repos/gaopengfei123123/github_hook/keys{/key_id}","collaborators_url":"https://api.github.com/repos/gaopengfei123123/github_hook/collaborators{/collaborator}","teams_url":"https://api.github.com/repos/gaopengfei123123/github_hook/teams","hooks_url":"https://api.github.com/repos/gaopengfei123123/github_hook/hooks","issue_events_url":"https://api.github.com/repos/gaopengfei123123/github_hook/issues/events{/number}","events_url":"https://api.github.com/repos/gaopengfei123123/github_hook/events","assignees_url":"https://api.github.com/repos/gaopengfei123123/github_hook/assignees{/user}","branches_url":"https://api.github.com/repos/gaopengfei123123/github_hook/branches{/branch}","tags_url":"https://api.github.com/repos/gaopengfei123123/github_hook/tags","blobs_url":"https://api.github.com/repos/gaopengfei123123/github_hook/git/blobs{/sha}","git_tags_url":"https://api.github.com/repos/gaopengfei123123/github_hook/git/tags{/sha}","git_refs_url":"https://api.github.com/repos/gaopengfei123123/github_hook/git/refs{/sha}","trees_url":"https://api.github.com/repos/gaopengfei123123/github_hook/git/trees{/sha}","statuses_url":"https://api.github.com/repos/gaopengfei123123/github_hook/statuses/{sha}","languages_url":"https://api.github.com/repos/gaopengfei123123/github_hook/languages","stargazers_url":"https://api.github.com/repos/gaopengfei123123/github_hook/stargazers","contributors_url":"https://api.github.com/repos/gaopengfei123123/github_hook/contributors","subscribers_url":"https://api.github.com/repos/gaopengfei123123/github_hook/subscribers","subscription_url":"https://api.github.com/repos/gaopengfei123123/github_hook/subscription","commits_url":"https://api.github.com/repos/gaopengfei123123/github_hook/commits{/sha}","git_commits_url":"https://api.github.com/repos/gaopengfei123123/github_hook/git/commits{/sha}","comments_url":"https://api.github.com/repos/gaopengfei123123/github_hook/comments{/number}","issue_comment_url":"https://api.github.com/repos/gaopengfei123123/github_hook/issues/comments{/number}","contents_url":"https://api.github.com/repos/gaopengfei123123/github_hook/contents/{+path}","compare_url":"https://api.github.com/repos/gaopengfei123123/github_hook/compare/{base}...{head}","merges_url":"https://api.github.com/repos/gaopengfei123123/github_hook/merges","archive_url":"https://api.github.com/repos/gaopengfei123123/github_hook/{archive_format}{/ref}","downloads_url":"https://api.github.com/repos/gaopengfei123123/github_hook/downloads","issues_url":"https://api.github.com/repos/gaopengfei123123/github_hook/issues{/number}","pulls_url":"https://api.github.com/repos/gaopengfei123123/github_hook/pulls{/number}","milestones_url":"https://api.github.com/repos/gaopengfei123123/github_hook/milestones{/number}","notifications_url":"https://api.github.com/repos/gaopengfei123123/github_hook/notifications{?since,all,participating}","labels_url":"https://api.github.com/repos/gaopengfei123123/github_hook/labels{/name}","releases_url":"https://api.github.com/repos/gaopengfei123123/github_hook/releases{/id}","deployments_url":"https://api.github.com/repos/gaopengfei123123/github_hook/deployments","created_at":"2019-10-21T12:16:03Z","updated_at":"2019-10-21T12:36:28Z","pushed_at":"2019-10-21T12:36:26Z","git_url":"git://github.com/gaopengfei123123/github_hook.git","ssh_url":"git@github.com:gaopengfei123123/github_hook.git","clone_url":"https://github.com/gaopengfei123123/github_hook.git","svn_url":"https://github.com/gaopengfei123123/github_hook","homepage":null,"size":0,"stargazers_count":0,"watchers_count":0,"language":"Go","has_issues":true,"has_projects":true,"has_downloads":true,"has_wiki":true,"has_pages":false,"forks_count":0,"mirror_url":null,"archived":false,"disabled":false,"open_issues_count":0,"license":null,"forks":0,"open_issues":0,"watchers":0,"default_branch":"master"},"sender":{"login":"gaopengfei123123","id":20946369,"node_id":"MDQ6VXNlcjIwOTQ2MzY5","avatar_url":"https://avatars3.githubusercontent.com/u/20946369?v=4","gravatar_id":"","url":"https://api.github.com/users/gaopengfei123123","html_url":"https://github.com/gaopengfei123123","followers_url":"https://api.github.com/users/gaopengfei123123/followers","following_url":"https://api.github.com/users/gaopengfei123123/following{/other_user}","gists_url":"https://api.github.com/users/gaopengfei123123/gists{/gist_id}","starred_url":"https://api.github.com/users/gaopengfei123123/starred{/owner}{/repo}","subscriptions_url":"https://api.github.com/users/gaopengfei123123/subscriptions","organizations_url":"https://api.github.com/users/gaopengfei123123/orgs","repos_url":"https://api.github.com/users/gaopengfei123123/repos","events_url":"https://api.github.com/users/gaopengfei123123/events{/privacy}","received_events_url":"https://api.github.com/users/gaopengfei123123/received_events","type":"User","site_admin":false}}`)
	appPath = "."
	params := GithubHook{}
	json.Unmarshal(jData, &params)
	t.Logf("%v ", params)

	cmd, err := selectCMDByHook(params)
	if err != nil {
		t.Error(err)
		return
	}
	err = execBash(cmd)

	if err != nil {
		t.Error(err)
	} else {
		t.Log("pass")
	}

}

func TestSelectCMDByHook(t *testing.T) {
	jData := []byte(`{"zen":"Non-blocking is better than blocking.","hook_id":150905429,"hook":{"type":"Repository","id":150905429,"name":"web","active":true,"events":["push"],"config":{"content_type":"json","insecure_ssl":"0","url":"http://39.107.51.167:8080/push"},"updated_at":"2019-10-21T12:40:54Z","created_at":"2019-10-21T12:40:54Z","url":"https://api.github.com/repos/gaopengfei123123/github_hook/hooks/150905429","test_url":"https://api.github.com/repos/gaopengfei123123/github_hook/hooks/150905429/test","ping_url":"https://api.github.com/repos/gaopengfei123123/github_hook/hooks/150905429/pings","last_response":{"code":null,"status":"unused","message":null}},"repository":{"id":216560073,"node_id":"MDEwOlJlcG9zaXRvcnkyMTY1NjAwNzM=","name":"github_hook","full_name":"gaopengfei123123/github_hook","private":true,"owner":{"login":"gaopengfei123123","id":20946369,"node_id":"MDQ6VXNlcjIwOTQ2MzY5","avatar_url":"https://avatars3.githubusercontent.com/u/20946369?v=4","gravatar_id":"","url":"https://api.github.com/users/gaopengfei123123","html_url":"https://github.com/gaopengfei123123","followers_url":"https://api.github.com/users/gaopengfei123123/followers","following_url":"https://api.github.com/users/gaopengfei123123/following{/other_user}","gists_url":"https://api.github.com/users/gaopengfei123123/gists{/gist_id}","starred_url":"https://api.github.com/users/gaopengfei123123/starred{/owner}{/repo}","subscriptions_url":"https://api.github.com/users/gaopengfei123123/subscriptions","organizations_url":"https://api.github.com/users/gaopengfei123123/orgs","repos_url":"https://api.github.com/users/gaopengfei123123/repos","events_url":"https://api.github.com/users/gaopengfei123123/events{/privacy}","received_events_url":"https://api.github.com/users/gaopengfei123123/received_events","type":"User","site_admin":false},"html_url":"https://github.com/gaopengfei123123/github_hook","description":"钩子函数","fork":false,"url":"https://api.github.com/repos/gaopengfei123123/github_hook","forks_url":"https://api.github.com/repos/gaopengfei123123/github_hook/forks","keys_url":"https://api.github.com/repos/gaopengfei123123/github_hook/keys{/key_id}","collaborators_url":"https://api.github.com/repos/gaopengfei123123/github_hook/collaborators{/collaborator}","teams_url":"https://api.github.com/repos/gaopengfei123123/github_hook/teams","hooks_url":"https://api.github.com/repos/gaopengfei123123/github_hook/hooks","issue_events_url":"https://api.github.com/repos/gaopengfei123123/github_hook/issues/events{/number}","events_url":"https://api.github.com/repos/gaopengfei123123/github_hook/events","assignees_url":"https://api.github.com/repos/gaopengfei123123/github_hook/assignees{/user}","branches_url":"https://api.github.com/repos/gaopengfei123123/github_hook/branches{/branch}","tags_url":"https://api.github.com/repos/gaopengfei123123/github_hook/tags","blobs_url":"https://api.github.com/repos/gaopengfei123123/github_hook/git/blobs{/sha}","git_tags_url":"https://api.github.com/repos/gaopengfei123123/github_hook/git/tags{/sha}","git_refs_url":"https://api.github.com/repos/gaopengfei123123/github_hook/git/refs{/sha}","trees_url":"https://api.github.com/repos/gaopengfei123123/github_hook/git/trees{/sha}","statuses_url":"https://api.github.com/repos/gaopengfei123123/github_hook/statuses/{sha}","languages_url":"https://api.github.com/repos/gaopengfei123123/github_hook/languages","stargazers_url":"https://api.github.com/repos/gaopengfei123123/github_hook/stargazers","contributors_url":"https://api.github.com/repos/gaopengfei123123/github_hook/contributors","subscribers_url":"https://api.github.com/repos/gaopengfei123123/github_hook/subscribers","subscription_url":"https://api.github.com/repos/gaopengfei123123/github_hook/subscription","commits_url":"https://api.github.com/repos/gaopengfei123123/github_hook/commits{/sha}","git_commits_url":"https://api.github.com/repos/gaopengfei123123/github_hook/git/commits{/sha}","comments_url":"https://api.github.com/repos/gaopengfei123123/github_hook/comments{/number}","issue_comment_url":"https://api.github.com/repos/gaopengfei123123/github_hook/issues/comments{/number}","contents_url":"https://api.github.com/repos/gaopengfei123123/github_hook/contents/{+path}","compare_url":"https://api.github.com/repos/gaopengfei123123/github_hook/compare/{base}...{head}","merges_url":"https://api.github.com/repos/gaopengfei123123/github_hook/merges","archive_url":"https://api.github.com/repos/gaopengfei123123/github_hook/{archive_format}{/ref}","downloads_url":"https://api.github.com/repos/gaopengfei123123/github_hook/downloads","issues_url":"https://api.github.com/repos/gaopengfei123123/github_hook/issues{/number}","pulls_url":"https://api.github.com/repos/gaopengfei123123/github_hook/pulls{/number}","milestones_url":"https://api.github.com/repos/gaopengfei123123/github_hook/milestones{/number}","notifications_url":"https://api.github.com/repos/gaopengfei123123/github_hook/notifications{?since,all,participating}","labels_url":"https://api.github.com/repos/gaopengfei123123/github_hook/labels{/name}","releases_url":"https://api.github.com/repos/gaopengfei123123/github_hook/releases{/id}","deployments_url":"https://api.github.com/repos/gaopengfei123123/github_hook/deployments","created_at":"2019-10-21T12:16:03Z","updated_at":"2019-10-21T12:36:28Z","pushed_at":"2019-10-21T12:36:26Z","git_url":"git://github.com/gaopengfei123123/github_hook.git","ssh_url":"git@github.com:gaopengfei123123/github_hook.git","clone_url":"https://github.com/gaopengfei123123/github_hook.git","svn_url":"https://github.com/gaopengfei123123/github_hook","homepage":null,"size":0,"stargazers_count":0,"watchers_count":0,"language":"Go","has_issues":true,"has_projects":true,"has_downloads":true,"has_wiki":true,"has_pages":false,"forks_count":0,"mirror_url":null,"archived":false,"disabled":false,"open_issues_count":0,"license":null,"forks":0,"open_issues":0,"watchers":0,"default_branch":"master"},"sender":{"login":"gaopengfei123123","id":20946369,"node_id":"MDQ6VXNlcjIwOTQ2MzY5","avatar_url":"https://avatars3.githubusercontent.com/u/20946369?v=4","gravatar_id":"","url":"https://api.github.com/users/gaopengfei123123","html_url":"https://github.com/gaopengfei123123","followers_url":"https://api.github.com/users/gaopengfei123123/followers","following_url":"https://api.github.com/users/gaopengfei123123/following{/other_user}","gists_url":"https://api.github.com/users/gaopengfei123123/gists{/gist_id}","starred_url":"https://api.github.com/users/gaopengfei123123/starred{/owner}{/repo}","subscriptions_url":"https://api.github.com/users/gaopengfei123123/subscriptions","organizations_url":"https://api.github.com/users/gaopengfei123123/orgs","repos_url":"https://api.github.com/users/gaopengfei123123/repos","events_url":"https://api.github.com/users/gaopengfei123123/events{/privacy}","received_events_url":"https://api.github.com/users/gaopengfei123123/received_events","type":"User","site_admin":false}}`)
	appPath = "."
	params := GithubHook{}
	json.Unmarshal(jData, &params)
	t.Logf("%v ", params)
	cmd, err := selectCMDByHook(params)

	if err != nil {
		t.Error(err)
	} else {
		t.Log(cmd)
		t.Log("pass")
	}
}

func TestReadConfig(t *testing.T) {
	ReadConfig(configPath)
	t.Log(scriptConf.Get("github_hook"))
	t.Log(scriptConf.Get("nothing"))
}

func TestInit(t *testing.T) {
	initScriptConfig()

	conf, err := scriptConf.Get("github_hook")
	if err != nil {
		t.Error(err)
	}

	cmd, err := conf.EventBash("push")

	t.Logf("cmd: %s, err: %v \n", cmd, err)

}

func TestStop(t *testing.T) {
	Stop()
}

func TestExist(t *testing.T) {
	confFullPath := "/Users/gpf/Documents/www/go_projects/src/hook/example/scripts/config.json"
	res := IsExist(confFullPath)
	t.Logf("exist: %v \n", res)

	dirpath, _ := path.Split(confFullPath)
	res = IsExist(dirpath)
	t.Logf("dir exist: %v \n", res)
}

func TestGithubSecret(t *testing.T) {
	raw := []byte(`{"ref":"refs/heads/master","before":"da8731d29cbf7f52a063717438120dd7872506a3","after":"c6015a150a528283a64a0c4a8c5bbe0af75b65e5","repository":{"id":216560073,"node_id":"MDEwOlJlcG9zaXRvcnkyMTY1NjAwNzM=","name":"hook","full_name":"gaopengfei123123/hook","private":false,"owner":{"name":"gaopengfei123123","email":"5173180@qq.com","login":"gaopengfei123123","id":20946369,"node_id":"MDQ6VXNlcjIwOTQ2MzY5","avatar_url":"https://avatars3.githubusercontent.com/u/20946369?v=4","gravatar_id":"","url":"https://api.github.com/users/gaopengfei123123","html_url":"https://github.com/gaopengfei123123","followers_url":"https://api.github.com/users/gaopengfei123123/followers","following_url":"https://api.github.com/users/gaopengfei123123/following{/other_user}","gists_url":"https://api.github.com/users/gaopengfei123123/gists{/gist_id}","starred_url":"https://api.github.com/users/gaopengfei123123/starred{/owner}{/repo}","subscriptions_url":"https://api.github.com/users/gaopengfei123123/subscriptions","organizations_url":"https://api.github.com/users/gaopengfei123123/orgs","repos_url":"https://api.github.com/users/gaopengfei123123/repos","events_url":"https://api.github.com/users/gaopengfei123123/events{/privacy}","received_events_url":"https://api.github.com/users/gaopengfei123123/received_events","type":"User","site_admin":false},"html_url":"https://github.com/byssh1989/hook","description":"钩子函数","fork":false,"url":"https://github.com/byssh1989/hook","forks_url":"https://api.github.com/repos/gaopengfei123123/hook/forks","keys_url":"https://api.github.com/repos/gaopengfei123123/hook/keys{/key_id}","collaborators_url":"https://api.github.com/repos/gaopengfei123123/hook/collaborators{/collaborator}","teams_url":"https://api.github.com/repos/gaopengfei123123/hook/teams","hooks_url":"https://api.github.com/repos/gaopengfei123123/hook/hooks","issue_events_url":"https://api.github.com/repos/gaopengfei123123/hook/issues/events{/number}","events_url":"https://api.github.com/repos/gaopengfei123123/hook/events","assignees_url":"https://api.github.com/repos/gaopengfei123123/hook/assignees{/user}","branches_url":"https://api.github.com/repos/gaopengfei123123/hook/branches{/branch}","tags_url":"https://api.github.com/repos/gaopengfei123123/hook/tags","blobs_url":"https://api.github.com/repos/gaopengfei123123/hook/git/blobs{/sha}","git_tags_url":"https://api.github.com/repos/gaopengfei123123/hook/git/tags{/sha}","git_refs_url":"https://api.github.com/repos/gaopengfei123123/hook/git/refs{/sha}","trees_url":"https://api.github.com/repos/gaopengfei123123/hook/git/trees{/sha}","statuses_url":"https://api.github.com/repos/gaopengfei123123/hook/statuses/{sha}","languages_url":"https://api.github.com/repos/gaopengfei123123/hook/languages","stargazers_url":"https://api.github.com/repos/gaopengfei123123/hook/stargazers","contributors_url":"https://api.github.com/repos/gaopengfei123123/hook/contributors","subscribers_url":"https://api.github.com/repos/gaopengfei123123/hook/subscribers","subscription_url":"https://api.github.com/repos/gaopengfei123123/hook/subscription","commits_url":"https://api.github.com/repos/gaopengfei123123/hook/commits{/sha}","git_commits_url":"https://api.github.com/repos/gaopengfei123123/hook/git/commits{/sha}","comments_url":"https://api.github.com/repos/gaopengfei123123/hook/comments{/number}","issue_comment_url":"https://api.github.com/repos/gaopengfei123123/hook/issues/comments{/number}","contents_url":"https://api.github.com/repos/gaopengfei123123/hook/contents/{+path}","compare_url":"https://api.github.com/repos/gaopengfei123123/hook/compare/{base}...{head}","merges_url":"https://api.github.com/repos/gaopengfei123123/hook/merges","archive_url":"https://api.github.com/repos/gaopengfei123123/hook/{archive_format}{/ref}","downloads_url":"https://api.github.com/repos/gaopengfei123123/hook/downloads","issues_url":"https://api.github.com/repos/gaopengfei123123/hook/issues{/number}","pulls_url":"https://api.github.com/repos/gaopengfei123123/hook/pulls{/number}","milestones_url":"https://api.github.com/repos/gaopengfei123123/hook/milestones{/number}","notifications_url":"https://api.github.com/repos/gaopengfei123123/hook/notifications{?since,all,participating}","labels_url":"https://api.github.com/repos/gaopengfei123123/hook/labels{/name}","releases_url":"https://api.github.com/repos/gaopengfei123123/hook/releases{/id}","deployments_url":"https://api.github.com/repos/gaopengfei123123/hook/deployments","created_at":1571660163,"updated_at":"2019-11-07T09:57:48Z","pushed_at":1573120882,"git_url":"git://github.com/byssh1989/hook.git","ssh_url":"git@github.com:gaopengfei123123/hook.git","clone_url":"https://github.com/byssh1989/hook.git","svn_url":"https://github.com/byssh1989/hook","homepage":null,"size":12349,"stargazers_count":0,"watchers_count":0,"language":"Go","has_issues":true,"has_projects":true,"has_downloads":true,"has_wiki":false,"has_pages":false,"forks_count":0,"mirror_url":null,"archived":false,"disabled":false,"open_issues_count":0,"license":null,"forks":0,"open_issues":0,"watchers":0,"default_branch":"master","stargazers":0,"master_branch":"master"},"pusher":{"name":"gaopengfei123123","email":"5173180@qq.com"},"sender":{"login":"gaopengfei123123","id":20946369,"node_id":"MDQ6VXNlcjIwOTQ2MzY5","avatar_url":"https://avatars3.githubusercontent.com/u/20946369?v=4","gravatar_id":"","url":"https://api.github.com/users/gaopengfei123123","html_url":"https://github.com/gaopengfei123123","followers_url":"https://api.github.com/users/gaopengfei123123/followers","following_url":"https://api.github.com/users/gaopengfei123123/following{/other_user}","gists_url":"https://api.github.com/users/gaopengfei123123/gists{/gist_id}","starred_url":"https://api.github.com/users/gaopengfei123123/starred{/owner}{/repo}","subscriptions_url":"https://api.github.com/users/gaopengfei123123/subscriptions","organizations_url":"https://api.github.com/users/gaopengfei123123/orgs","repos_url":"https://api.github.com/users/gaopengfei123123/repos","events_url":"https://api.github.com/users/gaopengfei123123/events{/privacy}","received_events_url":"https://api.github.com/users/gaopengfei123123/received_events","type":"User","site_admin":false},"created":false,"deleted":false,"forced":false,"base_ref":null,"compare":"https://github.com/byssh1989/hook/compare/da8731d29cbf...c6015a150a52","commits":[{"id":"c6015a150a528283a64a0c4a8c5bbe0af75b65e5","tree_id":"40f302364f9d54a01240677b6d2fb813072599b5","distinct":true,"message":"测试secret","timestamp":"2019-11-07T18:00:53+08:00","url":"https://github.com/byssh1989/hook/commit/c6015a150a528283a64a0c4a8c5bbe0af75b65e5","author":{"name":"gaopengfei","email":"xxx@qq.com","username":"gaopengfei123123"},"committer":{"name":"gaopengfei","email":"xxx@qq.com","username":"gaopengfei123123"},"added":[],"removed":[],"modified":["README.md"]}],"head_commit":{"id":"c6015a150a528283a64a0c4a8c5bbe0af75b65e5","tree_id":"40f302364f9d54a01240677b6d2fb813072599b5","distinct":true,"message":"测试secret","timestamp":"2019-11-07T18:00:53+08:00","url":"https://github.com/byssh1989/hook/commit/c6015a150a528283a64a0c4a8c5bbe0af75b65e5","author":{"name":"gaopengfei","email":"xxx@qq.com","username":"gaopengfei123123"},"committer":{"name":"gaopengfei","email":"xxx@qq.com","username":"gaopengfei123123"},"added":[],"removed":[],"modified":["README.md"]}}`)
	salt := "123123"
	sign := "sha1=dbe406bb2b6c6356e615f298c2ab004cd6bbcfa9"
	res := checkSecret(raw, salt, sign)
	t.Logf("%v", res)
}

func TestReplay(t *testing.T) {
	payload := []byte(`{"ref":"refs/heads/master","before":"91162c92dc200064ba2c6fb5b8a6e09542ff2a44","after":"49dde1df4e2bb5dfe3d6bbe1929c1b1388e4c35f","repository":{"id":216560073,"node_id":"MDEwOlJlcG9zaXRvcnkyMTY1NjAwNzM=","name":"hook","full_name":"gaopengfei123123/hook","private":false,"owner":{"name":"gaopengfei123123","email":"5173180@qq.com","login":"gaopengfei123123","id":20946369,"node_id":"MDQ6VXNlcjIwOTQ2MzY5","avatar_url":"https://avatars3.githubusercontent.com/u/20946369?v=4","gravatar_id":"","url":"https://api.github.com/users/gaopengfei123123","html_url":"https://github.com/gaopengfei123123","followers_url":"https://api.github.com/users/gaopengfei123123/followers","following_url":"https://api.github.com/users/gaopengfei123123/following{/other_user}","gists_url":"https://api.github.com/users/gaopengfei123123/gists{/gist_id}","starred_url":"https://api.github.com/users/gaopengfei123123/starred{/owner}{/repo}","subscriptions_url":"https://api.github.com/users/gaopengfei123123/subscriptions","organizations_url":"https://api.github.com/users/gaopengfei123123/orgs","repos_url":"https://api.github.com/users/gaopengfei123123/repos","events_url":"https://api.github.com/users/gaopengfei123123/events{/privacy}","received_events_url":"https://api.github.com/users/gaopengfei123123/received_events","type":"User","site_admin":false},"html_url":"https://github.com/byssh1989/hook","description":"钩子函数","fork":false,"url":"https://github.com/byssh1989/hook","forks_url":"https://api.github.com/repos/gaopengfei123123/hook/forks","keys_url":"https://api.github.com/repos/gaopengfei123123/hook/keys{/key_id}","collaborators_url":"https://api.github.com/repos/gaopengfei123123/hook/collaborators{/collaborator}","teams_url":"https://api.github.com/repos/gaopengfei123123/hook/teams","hooks_url":"https://api.github.com/repos/gaopengfei123123/hook/hooks","issue_events_url":"https://api.github.com/repos/gaopengfei123123/hook/issues/events{/number}","events_url":"https://api.github.com/repos/gaopengfei123123/hook/events","assignees_url":"https://api.github.com/repos/gaopengfei123123/hook/assignees{/user}","branches_url":"https://api.github.com/repos/gaopengfei123123/hook/branches{/branch}","tags_url":"https://api.github.com/repos/gaopengfei123123/hook/tags","blobs_url":"https://api.github.com/repos/gaopengfei123123/hook/git/blobs{/sha}","git_tags_url":"https://api.github.com/repos/gaopengfei123123/hook/git/tags{/sha}","git_refs_url":"https://api.github.com/repos/gaopengfei123123/hook/git/refs{/sha}","trees_url":"https://api.github.com/repos/gaopengfei123123/hook/git/trees{/sha}","statuses_url":"https://api.github.com/repos/gaopengfei123123/hook/statuses/{sha}","languages_url":"https://api.github.com/repos/gaopengfei123123/hook/languages","stargazers_url":"https://api.github.com/repos/gaopengfei123123/hook/stargazers","contributors_url":"https://api.github.com/repos/gaopengfei123123/hook/contributors","subscribers_url":"https://api.github.com/repos/gaopengfei123123/hook/subscribers","subscription_url":"https://api.github.com/repos/gaopengfei123123/hook/subscription","commits_url":"https://api.github.com/repos/gaopengfei123123/hook/commits{/sha}","git_commits_url":"https://api.github.com/repos/gaopengfei123123/hook/git/commits{/sha}","comments_url":"https://api.github.com/repos/gaopengfei123123/hook/comments{/number}","issue_comment_url":"https://api.github.com/repos/gaopengfei123123/hook/issues/comments{/number}","contents_url":"https://api.github.com/repos/gaopengfei123123/hook/contents/{+path}","compare_url":"https://api.github.com/repos/gaopengfei123123/hook/compare/{base}...{head}","merges_url":"https://api.github.com/repos/gaopengfei123123/hook/merges","archive_url":"https://api.github.com/repos/gaopengfei123123/hook/{archive_format}{/ref}","downloads_url":"https://api.github.com/repos/gaopengfei123123/hook/downloads","issues_url":"https://api.github.com/repos/gaopengfei123123/hook/issues{/number}","pulls_url":"https://api.github.com/repos/gaopengfei123123/hook/pulls{/number}","milestones_url":"https://api.github.com/repos/gaopengfei123123/hook/milestones{/number}","notifications_url":"https://api.github.com/repos/gaopengfei123123/hook/notifications{?since,all,participating}","labels_url":"https://api.github.com/repos/gaopengfei123123/hook/labels{/name}","releases_url":"https://api.github.com/repos/gaopengfei123123/hook/releases{/id}","deployments_url":"https://api.github.com/repos/gaopengfei123123/hook/deployments","created_at":1571660163,"updated_at":"2019-11-09T15:56:48Z","pushed_at":1573465288,"git_url":"git://github.com/byssh1989/hook.git","ssh_url":"git@github.com:gaopengfei123123/hook.git","clone_url":"https://github.com/byssh1989/hook.git","svn_url":"https://github.com/byssh1989/hook","homepage":null,"size":12369,"stargazers_count":0,"watchers_count":0,"language":"Go","has_issues":true,"has_projects":true,"has_downloads":true,"has_wiki":false,"has_pages":false,"forks_count":0,"mirror_url":null,"archived":false,"disabled":false,"open_issues_count":0,"license":null,"forks":0,"open_issues":0,"watchers":0,"default_branch":"master","stargazers":0,"master_branch":"master"},"pusher":{"name":"gaopengfei123123","email":"5173180@qq.com"},"sender":{"login":"gaopengfei123123","id":20946369,"node_id":"MDQ6VXNlcjIwOTQ2MzY5","avatar_url":"https://avatars3.githubusercontent.com/u/20946369?v=4","gravatar_id":"","url":"https://api.github.com/users/gaopengfei123123","html_url":"https://github.com/gaopengfei123123","followers_url":"https://api.github.com/users/gaopengfei123123/followers","following_url":"https://api.github.com/users/gaopengfei123123/following{/other_user}","gists_url":"https://api.github.com/users/gaopengfei123123/gists{/gist_id}","starred_url":"https://api.github.com/users/gaopengfei123123/starred{/owner}{/repo}","subscriptions_url":"https://api.github.com/users/gaopengfei123123/subscriptions","organizations_url":"https://api.github.com/users/gaopengfei123123/orgs","repos_url":"https://api.github.com/users/gaopengfei123123/repos","events_url":"https://api.github.com/users/gaopengfei123123/events{/privacy}","received_events_url":"https://api.github.com/users/gaopengfei123123/received_events","type":"User","site_admin":false},"created":false,"deleted":false,"forced":false,"base_ref":null,"compare":"https://github.com/byssh1989/hook/compare/91162c92dc20...49dde1df4e2b","commits":[{"id":"3e46251f2ec03321b1fdb81e544e66664c293122","tree_id":"25647e2eb2e6b3e562b9aaaf7ce1ad3e4b482316","distinct":true,"message":"调整config的数据格式","timestamp":"2019-11-11T17:39:19+08:00","url":"https://github.com/byssh1989/hook/commit/3e46251f2ec03321b1fdb81e544e66664c293122","author":{"name":"gaopengfei","email":"xxx@qq.com","username":"gaopengfei123123"},"committer":{"name":"gaopengfei","email":"xxx@qq.com","username":"gaopengfei123123"},"added":[],"removed":[],"modified":[".gitignore","example/hook.pid","example/scripts/config.json","hook_test.go","http.go","init.go","script_config.go","server.go"]},{"id":"49dde1df4e2bb5dfe3d6bbe1929c1b1388e4c35f","tree_id":"e9fc0b2e22c90b77f59be2d9c055ad9339b4340c","distinct":true,"message":"移除密钥","timestamp":"2019-11-11T17:41:04+08:00","url":"https://github.com/byssh1989/hook/commit/49dde1df4e2bb5dfe3d6bbe1929c1b1388e4c35f","author":{"name":"gaopengfei","email":"xxx@qq.com","username":"gaopengfei123123"},"committer":{"name":"gaopengfei","email":"xxx@qq.com","username":"gaopengfei123123"},"added":[],"removed":[],"modified":["example/scripts/config.json"]}],"head_commit":{"id":"49dde1df4e2bb5dfe3d6bbe1929c1b1388e4c35f","tree_id":"e9fc0b2e22c90b77f59be2d9c055ad9339b4340c","distinct":true,"message":"移除密钥","timestamp":"2019-11-11T17:41:04+08:00","url":"https://github.com/byssh1989/hook/commit/49dde1df4e2bb5dfe3d6bbe1929c1b1388e4c35f","author":{"name":"gaopengfei","email":"xxx@qq.com","username":"gaopengfei123123"},"committer":{"name":"gaopengfei","email":"xxx@qq.com","username":"gaopengfei123123"},"added":[],"removed":[],"modified":["example/scripts/config.json"]}}`)
	params := GithubHook{}
	json.Unmarshal(payload, &params)
	t.Log(params)

	now := time.Now().Unix()

	cha := now - params.Repository.PushedAt

	if cha > 3 {
		t.Logf("timeout: %d \n", cha)
	} else {
		t.Logf("success: %d \n", cha)
	}

	t.Logf("pushed_at: %d \n", params.Repository.PushedAt)
	t.Logf("now: %d \n", now)
}
