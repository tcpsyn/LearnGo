/* hello.go */
package main
import (
	"fmt"

)

type Item struct {
	Data struct {
		After    string      `json:"after"`
		Before   interface{} `json:"before"`
		Children []struct {
			Data struct {
				ApprovedBy          interface{} `json:"approved_by"`
				Author              string      `json:"author"`
				AuthorFlairCssClass interface{} `json:"author_flair_css_class"`
				AuthorFlairText     interface{} `json:"author_flair_text"`
				BannedBy            interface{} `json:"banned_by"`
				Clicked             bool        `json:"clicked"`
				Created             float64     `json:"created"`
				CreatedUtc          float64     `json:"created_utc"`
				Distinguished       interface{} `json:"distinguished"`
				Domain              string      `json:"domain"`
				Downs               float64     `json:"downs"`
				Edited              bool        `json:"edited"`
				Gilded              float64     `json:"gilded"`
				Hidden              bool        `json:"hidden"`
				ID                  string      `json:"id"`
				IsSelf              bool        `json:"is_self"`
				Likes               interface{} `json:"likes"`
				LinkFlairCssClass   interface{} `json:"link_flair_css_class"`
				LinkFlairText       interface{} `json:"link_flair_text"`
				Media               interface{} `json:"media"`
				MediaEmbed          struct{}    `json:"media_embed"`
				Name                string      `json:"name"`
				NumComments         float64     `json:"num_comments"`
				NumReports          interface{} `json:"num_reports"`
				Over18              bool        `json:"over_18"`
				Permalink           string      `json:"permalink"`
				Saved               bool        `json:"saved"`
				Score               float64     `json:"score"`
				SecureMedia         interface{} `json:"secure_media"`
				SecureMediaEmbed    struct{}    `json:"secure_media_embed"`
				Selftext            string      `json:"selftext"`
				SelftextHtml        interface{} `json:"selftext_html"`
				Stickied            bool        `json:"stickied"`
				Subreddit           string      `json:"subreddit"`
				SubredditID         string      `json:"subreddit_id"`
				Thumbnail           string      `json:"thumbnail"`
				Title               string      `json:"title"`
				Ups                 float64     `json:"ups"`
				URL                 string      `json:"url"`
				Visited             bool        `json:"visited"`
			} `json:"data"`
			Kind string `json:"kind"`
		} `json:"children"`
		Modhash string `json:"modhash"`
	} `json:"data"`
	Kind string `json:"kind"`
}

func main() {
	fmt.Printf("%+v", Item)
}
	
