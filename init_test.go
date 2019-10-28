package github_hook

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"testing"
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

	params := GithubHook{}
	json.Unmarshal(jData, &params)
	t.Logf("%v ", params)
	err := execHookBash(params)

	if err != nil {
		t.Error(err)
	} else {
		t.Log("pass")
	}

}

func TestReadConfig(t *testing.T) {
	ReadConfig(configPath)
	t.Logf("config: %v \n", configs)
}
