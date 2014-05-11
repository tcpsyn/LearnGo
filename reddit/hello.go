/* hello.go */
package main
import (
	"fmt"
	"encoding/json"
	"net/http"
)

type Item struct {
	Title 	string
	URL	string
	Author string
	Comments int `json:"num_comments"`
}



func (i Item) String() string {
	com := ""
	switch i.Comments {
		case 0:
			// nothing
		case 1:
			com = " (1 comment)"
		default: 
			com = fmt.Sprintf(" (%d comments)", i.Comments)
	}
	return fmt.Sprintf("%s%s\n%s", i.Title, com, i.URL)
}

func Get(reddit string) ([]Item, error) {
	url :=fmt.Sprintf("http://reddit.com/r/%s.json", reddit)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, err
	}

	r := new(response)
	err = json.NewDecoder(resp.Body).Decode(r)
	if err != nil {
		return nil, err
	}

	items := make([]Item, len(r.Data.Children))
	for i, child := range r.Data.Children {
		items[i] = child.Data
	}
	return items, nil
}

type response struct {
	Data struct {
		Children []struct {
			Data Item
		}
	}
}

func main() {
	items, err := Get("golang")
	if err != nil {
		fmt.Print(err)
	} else {
		for _, items := range items {
			fmt.Printf("Title: %s\n",items.Title)
			fmt.Printf("Author: %s\n",items.Author)
			fmt.Printf("URL: %s\n",items.URL)
			fmt.Printf("Comments: %d\n",items.Comments)
			fmt.Println("\n")
		}
	}
}
	
