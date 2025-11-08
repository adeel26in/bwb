package main

import (
	"fmt"

	"github.com/bitwarden/sdk-go"
)

func createproject() {

	var apiURL string = "https://api.bitwarden.com"
	var identityURL string = "https://identity.bitwarden.com"
	var organization_id string
	var project_name string
	var accessToken string

	bitwardenClient, err := sdk.NewBitwardenClient(&apiURL, &identityURL)

	if err != nil {

		fmt.Println(err)
	}

	fmt.Print("Enter your access token to authorize to Bitwarden: ")
	fmt.Scanln(&accessToken)

	err = bitwardenClient.AccessTokenLogin(accessToken, nil)

	if err != nil {

		fmt.Println(err)
	}

	defer bitwardenClient.Close()
	fmt.Print("Enter your organization id: ")
	fmt.Scan(&organization_id)

	fmt.Print("Enter the project name you want to create: ")
	fmt.Scan(&project_name)

	project, err := bitwardenClient.Projects().Create(organization_id, project_name)

	if err != nil {

		fmt.Println(err)

	}
	fmt.Println("Project created successfully:", project)

}

func listprojects() {

	var apiURL string = "https://api.bitwarden.com"
	var identityURL string = "https://identity.bitwarden.com"
	var organization_id string
	var accessToken string

	bitwardenClient, err := sdk.NewBitwardenClient(&apiURL, &identityURL)
	if err != nil {

		fmt.Println(err)
	}

	fmt.Print("Enter your access token to authorize to Bitwarden: ")
	fmt.Scanln(&accessToken)

	err = bitwardenClient.AccessTokenLogin(accessToken, nil)

	if err != nil {

		fmt.Println(err)
	}

	defer bitwardenClient.Close()
	fmt.Print("Enter your organization id: ")
	fmt.Scan(&organization_id)

	projects, err := bitwardenClient.Projects().List(organization_id)

	if err != nil {

		fmt.Println(err)

	}
	fmt.Println("Your projects:", projects)

}

func getprojects() {

	var apiURL string = "https://api.bitwarden.com"
	var identityURL string = "https://identity.bitwarden.com"
	var organization_id string
	var project_id string
	var accessToken string

	bitwardenClient, err := sdk.NewBitwardenClient(&apiURL, &identityURL)
	if err != nil {

		fmt.Println(err)
	}

	fmt.Print("Enter your access token to authorize to Bitwarden: ")
	fmt.Scanln(&accessToken)

	err = bitwardenClient.AccessTokenLogin(accessToken, nil)

	if err != nil {

		fmt.Println(err)
	}

	defer bitwardenClient.Close()
	fmt.Print("Enter your organization id: ")
	fmt.Scan(&organization_id)

	fmt.Println("What project do you want to get?: ")
	fmt.Scan(&project_id)

	project, err := bitwardenClient.Projects().Get(project_id)

	if err != nil {

		fmt.Println(err)

	}
	fmt.Println("Got these projects:", project)

}

func listsecrets() {

	var apiURL string = "https://api.bitwarden.com"
	var identityURL string = "https://identity.bitwarden.com"
	var organization_id string
	var accessToken string

	bitwardenClient, err := sdk.NewBitwardenClient(&apiURL, &identityURL)
	if err != nil {

		fmt.Println(err)
	}

	fmt.Print("Enter your access token to authorize to Bitwarden: ")
	fmt.Scanln(&accessToken)

	err = bitwardenClient.AccessTokenLogin(accessToken, nil)

	if err != nil {

		fmt.Println(err)
	}

	defer bitwardenClient.Close()
	fmt.Print("Enter your organization id: ")
	fmt.Scan(&organization_id)

	secrets, err := bitwardenClient.Secrets().List(organization_id)

	if err != nil {

		fmt.Println(err)

	}
	fmt.Println("List of your secrets", secrets)

}

func getsecret() {

	var apiURL string = "https://api.bitwarden.com"
	var identityURL string = "https://identity.bitwarden.com"
	var organization_id string
	var secret_id string
	var accessToken string

	bitwardenClient, err := sdk.NewBitwardenClient(&apiURL, &identityURL)
	if err != nil {

		fmt.Println(err)
	}

	fmt.Print("Enter your access token to authorize to Bitwarden: ")
	fmt.Scanln(&accessToken)

	err = bitwardenClient.AccessTokenLogin(accessToken, nil)

	if err != nil {

		fmt.Println(err)
	}

	defer bitwardenClient.Close()
	fmt.Print("Enter your organization id: ")
	fmt.Scan(&organization_id)

	fmt.Println("What secret do you want to get?: ")
	fmt.Scan(&secret_id)

	secret, err := bitwardenClient.Secrets().Get(secret_id)

	if err != nil {

		fmt.Println(err)

	}
	fmt.Println("Got these secrets:", secret)

}
