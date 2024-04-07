package main

import (
	"encoding/json"
	
	"fmt"
	"net/http"
	"os"


)

const (
	GithubUrl = "https://api.github.com/notifications"               
)  
var Token = os.Getenv("GITHUB_TOKEN")

type apifunc func(w http.ResponseWriter, r *http.Request) error

func makeApifunc(fn apifunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := fn(w, r); err != nil {
			WriteJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}
	}
}

func WriteJSON(w http.ResponseWriter, code int, v any) error {
	w.Header().Add("content-type", "application/json")
	w.WriteHeader(code)
	return json.NewEncoder(w).Encode(v)
}

func FetchNotificationsRepeated(w http.ResponseWriter, r *http.Request) error {
	// for {
		resp, err := FetchNotificationsFromGithub()
		if err != nil {
			fmt.Errorf("error fetching notifications:", err)
		}
		// time.Sleep(5 * time.Minute)
	// }
	return WriteJSON(w, http.StatusOK, resp)
}

func CreateJsonFiles ()  {
	file, err := os.Create("./git.json")
	if err != nil {
		return 
	}
	defer file.Close()
}

func FetchNotificationsFromGithub() ([]GithubNotification, error) {
	req, err := http.NewRequest("GET", GithubUrl, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}
	req.Header.Add("Authorization", "token "+ Token)
	fmt.Printf("Token: %v\n", Token)
	req.Header.Set("Accept", "application/vnd.github.v3+json")

	fmt.Println("Sending request to", GithubUrl)

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %v", err)
	}
	defer resp.Body.Close()

	fmt.Println("Response status:", resp.Status)

	var Notifications []GithubNotification
	// var htmlURL URL
	err = json.NewDecoder(resp.Body).Decode(&Notifications)
	if err != nil {
		return nil, fmt.Errorf("error decoding response: %v", err)
	}

	// err := json.Unmarshal(resp.body, &htmlURL)
    // if err != nil {
    //     return "", err
    // }
	// return resp.Body.html, nil


	file, err := os.Create("./git.json")
	if err != nil {
		return nil, fmt.Errorf("error creating file: %v", err)
	}
	defer file.Close()

	// creating an additional file to store the data for sending it to the frontend
	CreateJsonFiles()
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	err = encoder.Encode(Notifications)
	if err != nil {
		return nil, fmt.Errorf("error encoding JSON: %v", err)
	}
	// returning the notifications in JSON format as well
	return Notifications, nil
}

// func ExtractURL(url string) (URL,error) {
// 	var htmlURL URL
// 	err := json.Unmarshal(resp.body, &pullRequest)
//     if err != nil {
//         return "", err
//     }

// 	return htmlURL, nil
// }
