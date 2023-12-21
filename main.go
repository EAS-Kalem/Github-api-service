package main

import (
	"net/http"
	"fmt"
	"log"
	"github.com/gin-gonic/gin"
	"encoding/json"
	"encoding/base64"
	"io/ioutil"
	"time"
	"strings"
	) 

type Response3 struct {
	Name        string `json:"name"`
	Path        string `json:"path"`
	Sha         string `json:"sha"`
	Size        int    `json:"size"`
	URL         string `json:"url"`
	HTMLURL     string `json:"html_url"`
	GitURL      string `json:"git_url"`
	DownloadURL string `json:"download_url"`
	Type        string `json:"type"`
	Content     string `json:"content"`
	Encoding    string `json:"encoding"`
	Links       struct {
		Self string `json:"self"`
		Git  string `json:"git"`
		HTML string `json:"html"`
	} `json:"_links"`
}

type Response2 struct {
	ID       int    `json:"id"`
	NodeID   string `json:"node_id"`
	Name     string `json:"name"`
	FullName string `json:"full_name"`
	Private  bool   `json:"private"`
	Owner    struct {
		Login             string `json:"login"`
		ID                int    `json:"id"`
		NodeID            string `json:"node_id"`
		AvatarURL         string `json:"avatar_url"`
		GravatarID        string `json:"gravatar_id"`
		URL               string `json:"url"`
		HTMLURL           string `json:"html_url"`
		FollowersURL      string `json:"followers_url"`
		FollowingURL      string `json:"following_url"`
		GistsURL          string `json:"gists_url"`
		StarredURL        string `json:"starred_url"`
		SubscriptionsURL  string `json:"subscriptions_url"`	
		OrganizationsURL  string `json:"organizations_url"`
		ReposURL          string `json:"repos_url"`
		EventsURL         string `json:"events_url"`
		ReceivedEventsURL string `json:"received_events_url"`
		Type              string `json:"type"`
		SiteAdmin         bool   `json:"site_admin"`
	} `json:"owner"`
	HTMLURL                  string    `json:"html_url"`
	Description              any       `json:"description"`
	Fork                     bool      `json:"fork"`
	URL                      string    `json:"url"`
	ForksURL                 string    `json:"forks_url"`
	KeysURL                  string    `json:"keys_url"`
	CollaboratorsURL         string    `json:"collaborators_url"`
	TeamsURL                 string    `json:"teams_url"`
	HooksURL                 string    `json:"hooks_url"`
	IssueEventsURL           string    `json:"issue_events_url"`
	EventsURL                string    `json:"events_url"`
	AssigneesURL             string    `json:"assignees_url"`
	BranchesURL              string    `json:"branches_url"`
	TagsURL                  string    `json:"tags_url"`
	BlobsURL                 string    `json:"blobs_url"`
	GitTagsURL               string    `json:"git_tags_url"`
	GitRefsURL               string    `json:"git_refs_url"`
	TreesURL                 string    `json:"trees_url"`
	StatusesURL              string    `json:"statuses_url"`
	LanguagesURL             string    `json:"languages_url"`
	StargazersURL            string    `json:"stargazers_url"`
	ContributorsURL          string    `json:"contributors_url"`
	SubscribersURL           string    `json:"subscribers_url"`
	SubscriptionURL          string    `json:"subscription_url"`
	CommitsURL               string    `json:"commits_url"`
	GitCommitsURL            string    `json:"git_commits_url"`
	CommentsURL              string    `json:"comments_url"`
	IssueCommentURL          string    `json:"issue_comment_url"`
	ContentsURL              string    `json:"contents_url"`
	CompareURL               string    `json:"compare_url"`
	MergesURL                string    `json:"merges_url"`
	ArchiveURL               string    `json:"archive_url"`
	DownloadsURL             string    `json:"downloads_url"`
	IssuesURL                string    `json:"issues_url"`
	PullsURL                 string    `json:"pulls_url"`
	MilestonesURL            string    `json:"milestones_url"`
	NotificationsURL         string    `json:"notifications_url"`
	LabelsURL                string    `json:"labels_url"`
	ReleasesURL              string    `json:"releases_url"`
	DeploymentsURL           string    `json:"deployments_url"`
	CreatedAt                time.Time `json:"created_at"`
	UpdatedAt                time.Time `json:"updated_at"`
	PushedAt                 time.Time `json:"pushed_at"`
	GitURL                   string    `json:"git_url"`
	SSHURL                   string    `json:"ssh_url"`
	CloneURL                 string    `json:"clone_url"`
	SvnURL                   string    `json:"svn_url"`
	Homepage                 any       `json:"homepage"`
	Size                     int       `json:"size"`
	StargazersCount          int       `json:"stargazers_count"`
	WatchersCount            int       `json:"watchers_count"`
	Language                 string    `json:"language"`
	HasIssues                bool      `json:"has_issues"`
	HasProjects              bool      `json:"has_projects"`
	HasDownloads             bool      `json:"has_downloads"`
	HasWiki                  bool      `json:"has_wiki"`
	HasPages                 bool      `json:"has_pages"`
	HasDiscussions           bool      `json:"has_discussions"`
	ForksCount               int       `json:"forks_count"`
	MirrorURL                any       `json:"mirror_url"`
	Archived                 bool      `json:"archived"`
	Disabled                 bool      `json:"disabled"`
	OpenIssuesCount          int       `json:"open_issues_count"`
	License                  any       `json:"license"`
	AllowForking             bool      `json:"allow_forking"`
	IsTemplate               bool      `json:"is_template"`
	WebCommitSignoffRequired bool      `json:"web_commit_signoff_required"`
	Topics                   []any     `json:"topics"`
	Visibility               string    `json:"visibility"`
	Forks                    int       `json:"forks"`
	OpenIssues               int       `json:"open_issues"`
	Watchers                 int       `json:"watchers"`
	DefaultBranch            string    `json:"default_branch"`
	TempCloneToken           any       `json:"temp_clone_token"`
	NetworkCount             int       `json:"network_count"`
	SubscribersCount         int       `json:"subscribers_count"`
}

type Response []struct {
	Name        string `json:"name"`
	Path        string `json:"path"`
	Sha         string `json:"sha"`
	Size        int    `json:"size"`
	URL         string `json:"url"`
	HTMLURL     string `json:"html_url"`
	GitURL      string `json:"git_url"`
	DownloadURL string `json:"download_url"`
	Type        string `json:"type"`
	Links       struct {
		Self string `json:"self"`
		Git  string `json:"git"`
		HTML string `json:"html"`
	} `json:"_links"`
}

func main (){
	router := gin.Default()
	repoName := "moto_go"
	user := "kalem"
	fileName := "README.md"
	check := "moto_go"

	//Repo Exists
	router.GET("/api/repository/exists", func (c *gin.Context) {
		resp, err := http.Get("https://api.github.com/repos/EAS-" + user + "/" + repoName)
		if err != nil {
			log.Fatalln(err)
		}
		
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		
		var result2 Response2
		if err := json.Unmarshal(body, &result2); err != nil {   
    	fmt.Println("Can not unmarshal JSON")
		}

		if (result2.Name == repoName){
			fmt.Println("Good")
        } else {
			fmt.Println("Bad")
        }
	})
	
//Repo Contains
	router.GET("/api/repository/contains", func (c *gin.Context) {
		//http req
		resp, err := http.Get("https://api.github.com/repos/EAS-" + user + "/" + repoName + "/contents")
		if err != nil {
			log.Fatalln(err)
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		fmt.Println(body)
		var result Response
		if err := json.Unmarshal(body, &result); err != nil {   
    	fmt.Println("Can not unmarshal JSON")
		}
		for _, rec := range result {
			fmt.Println(rec.Name)
			if (rec.Name == fileName) {
				fmt.Println("Good")
			} else {
				fmt.Println("Bad")
			}
		}
	})

// File Contains
	router.GET("/api/repository/file/contains", func (c *gin.Context) {
		resp, err := http.Get("https://api.github.com/repos/EAS-" + user + "/" + repoName + "/contents/" + fileName)
		if err != nil {
			log.Fatalln(err)
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		fmt.Println(string(body))
		
		var result3 Response3
		if err := json.Unmarshal(body, &result3); err != nil {   
    	fmt.Println("Can not unmarshal JSON")
		}
		

		decoded, err := base64.StdEncoding.DecodeString(result3.Content)
		ans := strings.Contains(string(decoded), check )
		fmt.Println("hello", ans)
    	if err != nil {
        	fmt.Println("Error decoding string:", err)
		} else {
			fmt.Println("hello")
		}
		fmt.Sprintf("%T", check)
		fmt.Println(check)
    	fmt.Println(string(decoded))
	})
	router.Run(":3000")
}


// func PrettyPrint(i interface{}) string {
//     s, _ := json.MarshalIndent(i, "", "\t")
//     return string(s)
// }

