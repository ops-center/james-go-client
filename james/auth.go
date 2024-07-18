package james

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sync"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	v1alpha12 "kmodules.xyz/resource-metadata/apis/identity/v1alpha1"
	"kmodules.xyz/resource-metadata/client/clientset/versioned/typed/identity/v1alpha1"
	ctrl "sigs.k8s.io/controller-runtime"
)

var (
	identityClient *v1alpha1.IdentityV1alpha1Client
	once           sync.Once
)

func getJamesToken() (string, error) {
	{
		if token, ok := os.LookupEnv("INBOX_SERVER_TOKEN"); ok && len(token) > 0 {
			/*log.Println("\n--------------------------------\n" +
			"using james token from env" +
			"\n--------------------------------")*/
			return token, nil
		}
	}

	once.Do(func() {
		if config, err := ctrl.GetConfig(); err != nil {
			log.Fatalf("error fetching config, %v", err)
		} else if identityClient, err = v1alpha1.NewForConfig(config); err != nil {
			log.Fatalf("error building clientset, %v", err)
		}
	})

	tokenRequest := identityClient.InboxTokenRequests()
	tokenResp, err := tokenRequest.Create(context.TODO(), &v1alpha12.InboxTokenRequest{
		TypeMeta: v1.TypeMeta{
			Kind:       "InboxTokenRequest",
			APIVersion: "identity.k8s.appscode.com/v1alpha1",
		},
	}, v1.CreateOptions{})

	if err != nil || tokenResp.Response == nil {
		return "", fmt.Errorf("error creating authtoken %v", err)
	}

	var tokenMap map[string]string
	if jsonErr := json.Unmarshal([]byte(tokenResp.Response.AdminJWTToken), &tokenMap); jsonErr != nil {
		return "", fmt.Errorf("error unmarshalling authtoken %v", jsonErr)
	}

	if token, ok := tokenMap["token"]; ok {
		return token, nil
	}

	return "", fmt.Errorf("key \"token\" missing in authtoken response")
}
