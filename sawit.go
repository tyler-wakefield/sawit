package main
import(
	"html/template"
	"net/http"
	// "net/url"
	 "fmt"
	// #"html/template"
	"math/rand"
	"time"
	"encoding/json"
	"log"
	"io/ioutil"
	// "reflect"
)

type response struct {
	Kind string `json:"kind"`
	Data struct {
		Modhash  string `json:"modhash"`
		Dist     int    `json:"dist"`
		Children []struct {
			Kind string `json:"kind"`
			Data struct {
				ApprovedAtUtc              interface{}   `json:"approved_at_utc"`
				Subreddit                  string        `json:"subreddit"`
				Selftext                   string        `json:"selftext"`
				AuthorFullname             string        `json:"author_fullname"`
				Saved                      bool          `json:"saved"`
				ModReasonTitle             interface{}   `json:"mod_reason_title"`
				Gilded                     int           `json:"gilded"`
				Clicked                    bool          `json:"clicked"`
				Title                      string        `json:"title"`
				LinkFlairRichtext          []interface{} `json:"link_flair_richtext"`
				SubredditNamePrefixed      string        `json:"subreddit_name_prefixed"`
				Hidden                     bool          `json:"hidden"`
				Pwls                       int           `json:"pwls"`
				LinkFlairCSSClass          interface{}   `json:"link_flair_css_class"`
				Downs                      int           `json:"downs"`
				ThumbnailHeight            int           `json:"thumbnail_height"`
				HideScore                  bool          `json:"hide_score"`
				Name                       string        `json:"name"`
				Quarantine                 bool          `json:"quarantine"`
				LinkFlairTextColor         string        `json:"link_flair_text_color"`
				AuthorFlairBackgroundColor interface{}   `json:"author_flair_background_color"`
				SubredditType              string        `json:"subreddit_type"`
				Ups                        int           `json:"ups"`
				TotalAwardsReceived        int           `json:"total_awards_received"`
				MediaEmbed                 struct {
				} `json:"media_embed"`
				ThumbnailWidth        int           `json:"thumbnail_width"`
				AuthorFlairTemplateID interface{}   `json:"author_flair_template_id"`
				IsOriginalContent     bool          `json:"is_original_content"`
				UserReports           []interface{} `json:"user_reports"`
				SecureMedia           interface{}   `json:"secure_media"`
				IsRedditMediaDomain   bool          `json:"is_reddit_media_domain"`
				IsMeta                bool          `json:"is_meta"`
				Category              interface{}   `json:"category"`
				SecureMediaEmbed      struct {
				} `json:"secure_media_embed"`
				LinkFlairText       interface{}   `json:"link_flair_text"`
				CanModPost          bool          `json:"can_mod_post"`
				Score               int           `json:"score"`
				ApprovedBy          interface{}   `json:"approved_by"`
				Thumbnail           string        `json:"thumbnail"`
				Edited              bool          `json:"edited"`
				AuthorFlairCSSClass interface{}   `json:"author_flair_css_class"`
				StewardReports      []interface{} `json:"steward_reports"`
				AuthorFlairRichtext []interface{} `json:"author_flair_richtext"`
				Gildings            struct {
				} `json:"gildings"`
				PostHint          string      `json:"post_hint"`
				ContentCategories interface{} `json:"content_categories"`
				IsSelf            bool        `json:"is_self"`
				ModNote           interface{} `json:"mod_note"`
				Created           float64     `json:"created"`
				LinkFlairType     string      `json:"link_flair_type"`
				Wls               int         `json:"wls"`
				BannedBy          interface{} `json:"banned_by"`
				AuthorFlairType   string      `json:"author_flair_type"`
				Domain            string      `json:"domain"`
				AllowLiveComments bool        `json:"allow_live_comments"`
				SelftextHTML      interface{} `json:"selftext_html"`
				Likes             interface{} `json:"likes"`
				SuggestedSort     interface{} `json:"suggested_sort"`
				BannedAtUtc       interface{} `json:"banned_at_utc"`
				ViewCount         interface{} `json:"view_count"`
				Archived          bool        `json:"archived"`
				NoFollow          bool        `json:"no_follow"`
				IsCrosspostable   bool        `json:"is_crosspostable"`
				Pinned            bool        `json:"pinned"`
				Over18            bool        `json:"over_18"`
				Preview           struct {
					Images []struct {
						Source struct {
							URL    string `json:"url"`
							Width  int    `json:"width"`
							Height int    `json:"height"`
						} `json:"source"`
						Resolutions []struct {
							URL    string `json:"url"`
							Width  int    `json:"width"`
							Height int    `json:"height"`
						} `json:"resolutions"`
						Variants struct {
						} `json:"variants"`
						ID string `json:"id"`
					} `json:"images"`
					Enabled bool `json:"enabled"`
				} `json:"preview"`
				AllAwardings             []interface{} `json:"all_awardings"`
				Awarders                 []interface{} `json:"awarders"`
				MediaOnly                bool          `json:"media_only"`
				CanGild                  bool          `json:"can_gild"`
				Spoiler                  bool          `json:"spoiler"`
				Locked                   bool          `json:"locked"`
				AuthorFlairText          interface{}   `json:"author_flair_text"`
				Visited                  bool          `json:"visited"`
				NumReports               interface{}   `json:"num_reports"`
				Distinguished            interface{}   `json:"distinguished"`
				SubredditID              string        `json:"subreddit_id"`
				ModReasonBy              interface{}   `json:"mod_reason_by"`
				RemovalReason            interface{}   `json:"removal_reason"`
				LinkFlairBackgroundColor string        `json:"link_flair_background_color"`
				ID                       string        `json:"id"`
				IsRobotIndexable         bool          `json:"is_robot_indexable"`
				ReportReasons            interface{}   `json:"report_reasons"`
				Author                   string        `json:"author"`
				DiscussionType           interface{}   `json:"discussion_type"`
				NumComments              int           `json:"num_comments"`
				SendReplies              bool          `json:"send_replies"`
				WhitelistStatus          string        `json:"whitelist_status"`
				ContestMode              bool          `json:"contest_mode"`
				ModReports               []interface{} `json:"mod_reports"`
				AuthorPatreonFlair       bool          `json:"author_patreon_flair"`
				AuthorFlairTextColor     interface{}   `json:"author_flair_text_color"`
				Permalink                string        `json:"permalink"`
				ParentWhitelistStatus    string        `json:"parent_whitelist_status"`
				Stickied                 bool          `json:"stickied"`
				URL                      string        `json:"url"`
				SubredditSubscribers     int           `json:"subreddit_subscribers"`
				CreatedUtc               float64       `json:"created_utc"`
				NumCrossposts            int           `json:"num_crossposts"`
				Media                    interface{}   `json:"media"`
				IsVideo                  bool          `json:"is_video"`
			} `json:"data"`
		} `json:"children"`
		After  string      `json:"after"`
		Before interface{} `json:"before"`
	} `json:"data"`
}

type ImagePage struct {
	Title string
	Image string
}

type ImageHandler struct {
	Image string
}

const UserAgent = "User-Agent: script.sawit.viewer:v0.1 (by /u/thedevopsdojoblog)"

func randomSubreddit() string{
	subreddits := [5]string{"aww","ImaginaryLandscapes","OldSchoolCool","pics","imaginaryCharacters"}
	rand.Seed(time.Now().UnixNano())
	randomNumber := rand.Intn(len(subreddits)-1)

	return (subreddits[randomNumber])
}

func constructURL(subreddit string) string {
	baseUrl := "https://reddit.com"
	
	url := fmt.Sprintf("%s/r/%s/hot.json?limit=1",baseUrl,subreddit)
	return url
}

func (ih *ImageHandler) handler (w http.ResponseWriter,r *http.Request){
	page := ImagePage {Title: "SawIt Image Viewer",Image: ih.Image}
	template, _ := template.ParseFiles("index.html")
	template.Execute(w,page)
}

func getUrlFromJson(jsonbody string) string {
	jsonString := fmt.Sprintf("%v",jsonbody)
	jsonBytes := []byte(jsonString)

	var response response
	err := json.Unmarshal(jsonBytes, &response)
	if err != nil {
		fmt.Println(err)
	}
	
	// return jsonString
	return fmt.Sprintf("%v\n",response.Data.Children[0].Data.URL)
}

func MakeRequest(url string) string{
	req, err := http.NewRequest("GET", url, nil)
    if err != nil {
        fmt.Println(err)
    }
    req.Header.Set("User-Agent", UserAgent)

	resp, err := http.DefaultClient.Do(req)
    if err != nil {
        fmt.Println(err)
    }
    defer resp.Body.Close()
    if resp.StatusCode != http.StatusOK {
        return string(resp.StatusCode)
    }

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error:")
		log.Fatalln(err)
	}

	return (string(body))
}

func main(){
	subreddit := randomSubreddit()
	url := constructURL(subreddit)
	request := MakeRequest(url)
	image := getUrlFromJson(request)
	fmt.Println(subreddit)
	fmt.Println(url)
	fmt.Println(request)
	fmt.Println(image)

	myImageHandler := &ImageHandler{Image: image}
    http.HandleFunc("/", myImageHandler.handler)
    http.ListenAndServe(":8080", nil)
}